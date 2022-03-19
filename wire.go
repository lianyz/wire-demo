//go:build wireinject

/*
@Time : 2022/3/19 17:15
@Author : lianyz
@Description :
*/

package main

import "github.com/google/wire"

// InitializeEvent 声明injector的函数签名
func InitializeEvent(msg string) Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)

	// 返回值没有实际意义，只需符合函数签名即可
	return Event{}
}
