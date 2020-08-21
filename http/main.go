package main

import (
	"Uptime-monitor/http/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	//http.HandleFunc("/urls/", handler.PostUrls)

	//log.Fatal(http.ListenAndServe(":8080", nil))
	router.POST("/urls/", handler.PostUrls)
	router.GET("/urls/:id", handler.GetUrls)
	router.PATCH("/urls/:id", )
	router.POST("/urls/:id/activate", activateUrls)
	router.POST("/urls/:id/deactivate", deactivateUrls)
	router.DELETE("/urls/:id", deleteUrls)

	router.Run()

}
