package core

type Client interface{
	PUBLISH(chnlName string, message string)
	SUBSCRIBE(chnlName string)
	UNSUBSCRIBE(chnlName string)
	New() chan string
}