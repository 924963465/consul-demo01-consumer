package main

import (
	"fmt"
	"vscode/consul-demo01-consumer/proto"

	"github.com/micro/go-micro/v2"
	_ "github.com/micro/go-plugins/registry/consul/v2"
	"golang.org/x/net/context"
)

func main() {
	service := micro.NewService(
		micro.Name("client.service"),
	)

	service.Init()

	greeter := proto.NewGreeterService("greeter", service.Client())

	rsp, err := greeter.HelloTest(context.TODO(), &proto.SayRequest{
		Name: "Hello test",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rsp.GetGreeting)
}
