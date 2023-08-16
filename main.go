package main

import "github.com/gin-gonic/gin"

func main() {
	go handleBroadcast()
	router := gin.Default()

	router.GET("/message", handleWebSocket)

	router.Run(":3000")
}
