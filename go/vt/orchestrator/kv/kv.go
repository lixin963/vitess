/*
   Copyright 2017 Shlomi Noach, GitHub Inc.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package kv

import (
	"fmt"
	"sync"
)

type KeyValuePair struct {
	Key   string
	Value string
}

func NewKVPair(key string, value string) *KeyValuePair {
	return &KeyValuePair{Key: key, Value: value}
}

func (kvPair *KeyValuePair) String() string {
	return fmt.Sprintf("%s:%s", kvPair.Key, kvPair.Value)
}

type KeyValueStore interface {
	PutKeyValue(key string, value string) (err error)
	GetKeyValue(key string) (value string, found bool, err error)
	DistributePairs(kvPairs [](*KeyValuePair)) (err error)
}

var kvMutex sync.Mutex
var kvInitOnce sync.Once
var kvStores []KeyValueStore

// InitKVStores initializes the KV stores (duh), once in the lifetime of this app.
// Configuration reload does not affect a running instance.
func InitKVStores() {
	kvMutex.Lock()
	defer kvMutex.Unlock()

	kvInitOnce.Do(func() {
		kvStores = []KeyValueStore{
			NewInternalKVStore(),
			NewConsulStore(),
			NewZkStore(),
		}
	})
}

func getKVStores() (stores []KeyValueStore) {
	kvMutex.Lock()
	defer kvMutex.Unlock()

	stores = kvStores
	return stores
}

func GetValue(key string) (value string, found bool, err error) {
	for _, store := range getKVStores() {
		// It's really only the first (internal) that matters here
		return store.GetKeyValue(key)
	}
	return value, found, err
}

func PutValue(key string, value string) (err error) {
	for _, store := range getKVStores() {
		if err := store.PutKeyValue(key, value); err != nil {
			return err
		}
	}
	return nil
}

func PutKVPair(kvPair *KeyValuePair) (err error) {
	if kvPair == nil {
		return nil
	}
	return PutValue(kvPair.Key, kvPair.Value)
}

func DistributePairs(kvPairs [](*KeyValuePair)) (err error) {
	for _, store := range getKVStores() {
		if err := store.DistributePairs(kvPairs); err != nil {
			return err
		}
	}
	return nil
}
