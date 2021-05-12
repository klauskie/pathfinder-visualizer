package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"klauskie.com/pathfinder/backend/models"
	"log"
)


// POST /
func HandleRunResults (c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	s := struct {
		Length int `json:"length"`
		Wall []int `json:"wall"`
		Algo int `json:"algo"`
		StartId int `json:"start_id"`
		EndId int `json:"end_id"`
		Rows int `json:"rows"`
		Cols int `json:"cols"`
	}{}
	if err := json.Unmarshal(body, &s); err != nil {
		fmt.Println(err.Error())
	}

	search := models.Create(s.Algo, s.Wall, s.StartId, s.EndId, s.Rows, s.Cols)
	if search == nil {
		log.Println("Invalid algoId...")
		return
	}
	nodes, path := search.Run()

	fmt.Println(s)
	c.JSON(202, gin.H{
		"Data": nodes,
		"Grid": search.GetPathfinder().Grid,
		"Path": path,
		"StartId": search.GetPathfinder().StartId,
		"EndId": search.GetPathfinder().EndId,
		"Walls": s.Wall,
	})
}
