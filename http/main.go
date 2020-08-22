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
	router.PATCH("/urls/:id", handler.PatchUrls)
	router.POST("/urls/:id/activate", handler.ActivateUrls)
	router.POST("/urls/:id/deactivate", handler.DeactivateUrls)
	router.DELETE("/urls/:id", handler.DeleteUrls)

	router.Run()

}
