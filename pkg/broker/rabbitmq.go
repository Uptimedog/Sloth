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
func (r *Rabbitmq) SetAddr(addr string) {
	r.Addr = addr
}

// Connect connects to Rabbitmq
func (r *Rabbitmq) Connect() (bool, error) {
	var err error

	r.Conn, err = amqp.Dial(r.Addr)

	if err != nil {
		return false, fmt.Errorf(
			"Failed to connect to RabbitMQ: %s",
			err.Error(),
		)
	}

	return true, nil
}

// Publish publish the message to rabbitmq
func (r *Rabbitmq) Publish(queue string, message string) (bool, error) {
	ch, err := r.Conn.Channel()

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

// Listen listens to published messages
func (r *Rabbitmq) Listen(_ func(message string) error) {

}

// Disconnect closes the connection
func (r *Rabbitmq) Disconnect() {
	r.Conn.Close()
}
