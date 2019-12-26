// Copyright 2019 Silverbackhq. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"time"
)

// Etcd struct
type Etcd struct {
	Connection *clientv3.Client
}

// NewEtcd create a new instance
func NewEtcd() *Etcd {
	return &Etcd{}
}

// Connect connects to etcd
func (e *Etcd) Connect(address []string, dialTimeout int) error {
	var err error

	e.Connection, err = clientv3.New(clientv3.Config{
		DialTimeout: time.Duration(dialTimeout) * time.Second,
		Endpoints:   address,
	})

	if err != nil {
		return err
	}

	return nil
}

// Put store a key value pair
func (e *Etcd) Put(ctx context.Context, key, value string) *clientv3.PutResponse {
	kv := clientv3.NewKV(e.Connection)
	response, _ := kv.Put(ctx, key, value)
	return response
}

// Get get the value of a key
func (e *Etcd) Get(ctx context.Context, key string) *clientv3.GetResponse {
	kv := clientv3.NewKV(e.Connection)
	response, _ := kv.Get(ctx, key)
	return response
}

// Delete delete the value of a key
func (e *Etcd) Delete(ctx context.Context, key string) *clientv3.DeleteResponse {
	kv := clientv3.NewKV(e.Connection)
	response, _ := kv.Delete(ctx, key)
	return response
}

// Disconnect closes etcd connection
func (e *Etcd) Disconnect() {
	e.Connection.Close()
}
