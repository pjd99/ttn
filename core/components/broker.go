// Copyright © 2015 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package components

import (
	"fmt"

	"github.com/TheThingsNetwork/ttn/core"
	"github.com/TheThingsNetwork/ttn/lorawan"
	"github.com/TheThingsNetwork/ttn/utils/log"
)

type Broker struct {
	log.Logger
	db brokerStorage
}

func NewBroker(loggers ...log.Logger) (*Broker, error) {
	localDB, err := NewBrokerStorage()

	if err != nil {
		return nil, err
	}

	return &Broker{
		Logger: log.MultiLogger{Loggers: loggers},
		db:     localDB,
	}, nil
}

func (b *Broker) HandleUp(p core.Packet, an core.AckNacker, adapter core.Adapter) error {
	// 1. Lookup for entries for the associated device
	devAddr, err := p.DevAddr()
	if err != nil {
		an.Nack()
		return ErrInvalidPacket
	}
	entries, err := b.db.lookup(devAddr)
	switch err {
	case nil:
	case ErrDeviceNotFound:
		return an.Nack()
	default:
		an.Nack()
		return err
	}

	// 2. Several handler might be associated to the same device, we distinguish them using MIC
	// check. Only one should verify the MIC check.
	var handler *core.Recipient
	for _, entry := range entries {
		ok, err := p.Payload.ValidateMIC(entry.NwsKey)
		if err != nil {
			continue
		}
		if ok {
			handler = &core.Recipient{
				Id:      entry.Id,
				Address: entry.Url,
			}
			b.LogEntry(log.DebugLevel, "Associated device with handler", log.Meta{"devAddr": devAddr, "handler": handler})
			break
		}
	}
	if handler == nil {
		b.LogEntry(log.WarnLevel, "Could not find handler for device", log.Meta{"devAddr": devAddr})
		return an.Nack()
	}

	// 3. If one was found, we forward the packet and wait for the response
	response, err := adapter.Send(p, *handler)
	if err != nil {
		an.Nack()
		return err
	}
	return an.Ack(response)
}

func (b *Broker) HandleDown(p core.Packet, an core.AckNacker, a core.Adapter) error {
	return fmt.Errorf("Not Implemented")
}

func (b *Broker) Register(r core.Registration, an core.AckNacker) error {
	id, okId := r.Recipient.Id.(string)
	url, okUrl := r.Recipient.Address.(string)
	nwsKey, okNwsKey := r.Options.(lorawan.AES128Key)

	if !(okId && okUrl && okNwsKey) {
		an.Nack()
		return ErrInvalidRegistration
	}

	entry := brokerEntry{Id: id, Url: url, NwsKey: nwsKey}
	if err := b.db.store(r.DevAddr, entry); err != nil {
		an.Nack()
		return err
	}
	return an.Ack()
}