package core


func Pub(){
	client:= GetClient(Instance)
	client.PUBLISH("hello", "message from manik")
}