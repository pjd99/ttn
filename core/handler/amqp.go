// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package handler

import (
	"github.com/apex/log"

	"github.com/TheThingsNetwork/ttn/amqp"
	"github.com/TheThingsNetwork/ttn/core/types"
)

func (h *handler) HandleAMQP(username, password, host, exchange, downlinkQueue string) error {
	h.amqpClient = amqp.NewClient(h.Ctx, username, password, host)

	err := h.amqpClient.Connect()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			h.amqpClient.Disconnect()
		}
	}()

	h.amqpUp = make(chan *types.UplinkMessage)

	subscriber := h.amqpClient.NewSubscriber(h.amqpExchange, downlinkQueue, downlinkQueue != "", downlinkQueue == "")
	err = subscriber.Open()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			subscriber.Close()
		}
	}()
	err = subscriber.SubscribeDownlink(func(_ amqp.Subscriber, _, _ string, req types.DownlinkMessage) {
		h.EnqueueDownlink(&req)
	})
	if err != nil {
		return err
	}

	ctx := h.Ctx.WithField("Protocol", "AMQP")

	go func() {
		publisher := h.amqpClient.NewPublisher(h.amqpExchange)
		err := publisher.Open()
		if err != nil {
			ctx.WithError(err).Error("Could not open publisher channel")
			return
		}
		defer publisher.Close()

		for up := range h.amqpUp {
			ctx.WithFields(log.Fields{
				"DevID": up.DevID,
				"AppID": up.AppID,
			}).Debug("Publish Uplink")
			err := publisher.PublishUplink(*up)
			if err != nil {
				ctx.WithError(err).Warn("Could not publish Uplink")
			}
		}
	}()

	return nil
}
