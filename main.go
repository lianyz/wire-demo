/*
@Time : 2022/3/19 17:09
@Author : lianyz
@Description :
*/

package main

import "fmt"

func main() {
	event := InitializeEvent("hello world")
	event.Start()
}

type Message struct {
	msg string
}

type Greeter struct {
	Message Message
}

func (g Greeter) Greet() Message {
	return g.Message
}

type Event struct {
	Greeter Greeter
}

// NewMessage Message的构造函数
func NewMessage(msg string) Message {
	return Message{msg: msg}
}

func NewGreeter(message Message) Greeter {
	return Greeter{Message: message}
}

func NewEvent(greeter Greeter) Event {
	return Event{Greeter: greeter}
}

func (e Event) Start() {
	msg := e.Greeter.Greet().msg
	fmt.Println(msg)
}
