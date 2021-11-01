package main

import (
	"FPproject/Frontend/env"
	"FPproject/Frontend/log"
	"FPproject/Frontend/server"
)

func main() {
	log.InitLog()
	env.InitEnv()
	server.Run()
}
