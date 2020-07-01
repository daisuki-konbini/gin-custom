package main

import (
	"flag"
	"gin/server"

	"github.com/gin-gonic/gin"
)

var env = flag.String("env", "debug", "input the env debug|test|release")

func main() {
	flag.Parse()
	gin.SetMode(*env)

	server.Init()
}
