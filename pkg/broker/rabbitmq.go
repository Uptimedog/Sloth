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

// Publish publish the message to rabbitmq
func (c *Rabbitmq) Publish(queue string, message string) (bool, error) {
	ch, err := c.Conn.Channel()

	if err != nil {
		return false, fmt.Errorf(
			"Failed to open a channel: %s",
			err.Error(),
		)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		return false, fmt.Errorf(
			"Failed to declare a queue: %s",
			err.Error(),
		)
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)

	if err != nil {
		return false, fmt.Errorf(
			"Failed to publish a message: %s",
			err.Error(),
		)
	}

	return true, nil
}

// Disconnect closes the connection
func (c *Rabbitmq) Disconnect() {
	c.Conn.Close()
}
