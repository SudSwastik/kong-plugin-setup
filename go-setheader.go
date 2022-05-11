package main

import (
	"fmt"
	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
)

type Config struct {
	Message string
}

func New() interface{} {
	return &Config{}
}


func (config Config) Access(kong *pdk.PDK) {

	message := config.Message

	if message == "" {
		message = "hello"
	}
	kong.Response.SetHeader("go-config-message", fmt.Sprintf("%s", message))

}

func main() {
	Version := "1.1"
	Priority := 1000
	_ = server.StartServer(New, Version, Priority)
}