package struct_metadata

import (
	"reflect"
	"sync"
)

var storeMutex = sync.RWMutex{}
var store = make(map[string]any)

func Metadata[TKey any, TData any](data TData) {
	storeMutex.Lock()
	defer storeMutex.Unlock()
	var key TKey
	store[reflect.TypeOf(key).String()] = &data
}

func GetMetadata[TKey any, TData any]() *TData {
	storeMutex.RLock()
	defer storeMutex.RUnlock()
	var key TKey
	instance, ok := store[reflect.TypeOf(key).String()].(*TData)
	if !ok {
		return nil
	}
	return instance
}
