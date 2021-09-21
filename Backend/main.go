package main

import (
	"FPproject/Backend/env"
	"FPproject/Backend/log"
	"FPproject/Backend/server"
)

func main() {
	log.InitLog()
	env.InitEnv()
	server.InitServer()
}
