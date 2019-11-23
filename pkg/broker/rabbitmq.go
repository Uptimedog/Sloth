// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package broker

import (
	"fmt"
	"github.com/streadway/amqp"
)

// rabbitmq struct
type rabbitmq struct {
	Conn *amqp.Connection
	Addr string
}

// NewRabbitmq create a new instance
func NewRabbitmq(addr string) *rabbitmq {
	return &rabbitmq{
		Addr: addr,
	}
}

// Connect connects to rabbitmq
func (c *rabbitmq) Connect() (bool, error) {
	var err error

	c.Conn, err = amqp.Dial(c.Addr)

	if err != nil {
		return false, fmt.Errorf(
			"Failed to connect to RabbitMQ: %s",
			err.Error(),
		)
	}

	defer c.Conn.Close()

	return true, nil
}

// Disconnect closes the connection
func (c *rabbitmq) Disconnect() {
	c.Conn.Close()
}
