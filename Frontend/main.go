package main

import (
	"FPproject/Frontend/env"
	"FPproject/Frontend/log"
	"FPproject/Frontend/server"
	"fmt"
)

func main() {
	log.InitLog()
	env.InitEnv()
	fmt.Println("Frontend Client")
	server.Run()
}
