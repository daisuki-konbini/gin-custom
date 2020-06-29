package server

import (
	"gin/config"
)

//Init ...
func Init() {
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	config := config.GetConfig()

	r := SetRouter()
	r.Run(config.GetString("server.port"))
}
