package core

import (
	"fmt"
	"sync"
)


type CreateClient struct{
	chnl chan string
	unsub chan bool
	store *Store
}

func(client *CreateClient) New() chan string {
	client.chnl = make(chan string)
	return client.chnl
}


func(client *CreateClient) PUBLISH (chnlName string, message string){
	chnl:= client.store.GetChannel(chnlName)
	chnl <- message
}

func(client *CreateClient) SUBSCRIBE (chnlName string){
	client.store.Put(chnlName, client.chnl)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
	loop:
		for{
			select {
			case msg := <- client.chnl:
				fmt.Println(msg)
				fmt.Println("Here1")
				continue loop;
			case msg := <-client.unsub:
				fmt.Println(msg)
				fmt.Println("Here2")
				break loop;
			default:
				fmt.Println("Here3")
				continue loop;
			}
		}
		wg.Done()
	}()
	wg.Wait()
}

func(client *CreateClient) UNSUBSCRIBE (chnlName string){
	client.unsub <- true
}

func GetClient(store *Store) *CreateClient{
	return &CreateClient{
		chnl: make(chan string),
		unsub: make(chan bool),
		store: store,
	}
}