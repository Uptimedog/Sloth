// Copyright 2019 Silverbackhq. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package broker

import (
	"fmt"
	"github.com/streadway/amqp"
)

// Rabbitmq struct
type Rabbitmq struct {
	Conn *amqp.Connection
	Addr string
}

// NewRabbitmq create a new instance
func NewRabbitmq() *Rabbitmq {
	return &Rabbitmq{}
}

// SetAddr sets rabbitmq address
func (c *Rabbitmq) SetAddr(addr string) {
	c.Addr = addr
}

// Connect connects to Rabbitmq
func (c *Rabbitmq) Connect() (bool, error) {
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
func (c *Rabbitmq) Disconnect() {
	c.Conn.Close()
}
