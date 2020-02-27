package main

import (
	messagebus "github.com/vardius/message-bus"
)

var bus = messagebus.New(10)