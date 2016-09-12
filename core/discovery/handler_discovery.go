// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package discovery

import (
	"fmt"
	"sync"
	"time"

	pb "github.com/TheThingsNetwork/ttn/api/discovery"
	"github.com/TheThingsNetwork/ttn/core"
	"github.com/pkg/errors"
)

// HandlerCacheTime indicates how long the HandlerDiscovery should cache the services
var HandlerCacheTime = 5 * time.Minute

// HandlerDiscovery is used as a client to discover Handlers
type HandlerDiscovery interface {
	AddAppID(handlerID, appID string) error
	ForAppID(appID string) ([]*pb.Announcement, error)
	Get(id string) (*pb.Announcement, error)
	All() ([]*pb.Announcement, error)
}

type handlerDiscovery struct {
	component       *core.Component
	cache           []*pb.Announcement
	byID            map[string]*pb.Announcement
	byAppID         map[string][]*pb.Announcement
	cacheLock       sync.RWMutex
	cacheValidUntil time.Time
}

// NewHandlerDiscovery returns a new HandlerDiscovery on top of the given gRPC connection
func NewHandlerDiscovery(component *core.Component) HandlerDiscovery {
	return &handlerDiscovery{component: component}
}

func (d *handlerDiscovery) refreshCache() error {
	res, err := d.component.Discovery.GetAll(d.component.GetContext(""), &pb.GetAllRequest{ServiceName: "handler"})
	if err != nil {
		return errors.Wrap(core.FromGRPCError(err), "Failed to refresh handlers from Discovery")
	}
	// TODO: validate response
	d.cacheLock.Lock()
	defer d.cacheLock.Unlock()
	d.cacheValidUntil = time.Now().Add(HandlerCacheTime)
	d.cache = res.Services
	d.updateLookups()
	return nil
}

func (d *handlerDiscovery) update() {
	if time.Now().After(d.cacheValidUntil) {
		d.cacheValidUntil = time.Now().Add(10 * time.Second)
		go d.refreshCache()
	}
}

func (d *handlerDiscovery) updateLookups() {
	d.byID = map[string]*pb.Announcement{}
	d.byAppID = map[string][]*pb.Announcement{}
	for _, service := range d.cache {
		d.byID[service.Id] = service
		for _, meta := range service.Metadata {
			switch meta.Key {
			case pb.Metadata_APP_ID:
				announcedID := string(meta.Value)
				if forID, ok := d.byAppID[announcedID]; ok {
					d.byAppID[announcedID] = append(forID, service)
				} else {
					d.byAppID[announcedID] = []*pb.Announcement{service}
				}
			}
		}
	}
}

func (d *handlerDiscovery) All() (announcements []*pb.Announcement, err error) {
	d.cacheLock.Lock()
	defer d.cacheLock.Unlock()
	d.update()
	announcements = d.cache
	return
}

func (d *handlerDiscovery) Get(id string) (*pb.Announcement, error) {
	d.cacheLock.Lock()
	defer d.cacheLock.Unlock()
	d.update()
	match, ok := d.byID[id]
	if !ok {
		return nil, core.NewErrNotFound(fmt.Sprintf("handler/%s", id))
	}
	return match, nil
}

func (d *handlerDiscovery) ForAppID(appID string) ([]*pb.Announcement, error) {
	d.cacheLock.Lock()
	defer d.cacheLock.Unlock()
	d.update()
	matches, ok := d.byAppID[appID]
	if !ok {
		return []*pb.Announcement{}, nil
	}
	return matches, nil
}

func (d *handlerDiscovery) AddAppID(handlerID, appID string) error {
	d.cacheLock.Lock()
	defer d.cacheLock.Unlock()
	handler, found := d.byID[handlerID]
	if !found {
		return core.NewErrNotFound(fmt.Sprintf("handler/%s", handlerID))
	}
	existing, found := d.byAppID[appID]
	if found && len(existing) > 0 {
		if existing[0].Id == handlerID {
			return nil
		}
		return core.NewErrAlreadyExists(fmt.Sprintf("AppID %s already registered", appID))
	}
	d.byAppID[appID] = []*pb.Announcement{handler}
	return nil
}