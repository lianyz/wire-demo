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

type Message interface {
	Print()
}

type Greeter interface {
	Greet()
}

type MessageImpl struct {
	msg string
}

type GreeterImpl struct {
	Message Message
}

func (m MessageImpl) Print() {
	fmt.Println(m.msg)
}

func (g GreeterImpl) Greet() {
	g.Message.Print()
}

type Event struct {
	Greeter Greeter
}

// NewMessage Message的构造函数
func NewMessage(msg string) Message {
	return MessageImpl{msg: msg}
}

func NewGreeter(message Message) Greeter {
	return GreeterImpl{Message: message}
}

func NewEvent(greeter Greeter) Event {
	return Event{Greeter: greeter}
}

func (e Event) Start() {
	e.Greeter.Greet()
}
