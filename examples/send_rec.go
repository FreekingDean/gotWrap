package main

import (
	"gotWrap"
)

func main() {
	go gotWrap.createServer()
	gotWrap.connect()
}