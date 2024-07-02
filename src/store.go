package core

import (
	"sync"
)

type Store struct{
	mutex sync.RWMutex
	chnlMap map[string]chan string
}

func(store *Store) Put(key string, chnl chan string){
	// if !store.mutex.TryLock(){
	// 	// send the request to queue
	// }
	store.mutex.Lock()
	store.chnlMap[key] = chnl
	store.mutex.Unlock()
}

func(store *Store) GetChannel(key string) chan string {
	var chnl chan string
	store.mutex.RLock()
	chnl = store.chnlMap[key]
	store.mutex.RUnlock()
	return chnl
}

var Instance *Store
var once sync.Once

func GetInstance(chnlMap map[string]chan string) *Store{
	once.Do(func(){
		Instance = &Store{
			mutex: sync.RWMutex{},
			chnlMap: chnlMap,
		}
	})
	return Instance
}