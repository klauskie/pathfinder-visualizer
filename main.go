package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"klauskie.com/pathfinder/api"
	"klauskie.com/pathfinder/models"
	"log"
)


func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")
	r.Static("/static", "./static")

	r.GET("/", api.GetScene)
	r.POST("/", api.HandleRunResults)


	r.Run(":8080")
}

func m() {
	search := models.Create(1, []int{}, 510, 524)
	if search == nil {
		log.Println("Invalid algoId...")
		return
	}
	nodes, path := search.Run()

	fmt.Println(nodes)
	fmt.Println(path)
}


