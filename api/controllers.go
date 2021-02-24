package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"klauskie.com/pathfinder/models"
	"klauskie.com/pathfinder/src"
	"log"
	"net/http"
)

// GET /
func GetScene (c *gin.Context) {
	p := src.CreatePathFinder([]int{}, -1, -1)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Grid": p.Grid,
		"StartId": p.GetStartNode(),
		"EndId": p.GetEndNode(),
		"Rows": 30,
		"Cols": 50,
	})
}

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
	}{}
	if err := json.Unmarshal(body, &s); err != nil {
		fmt.Println(err.Error())
	}

	search := models.Create(s.Algo, s.Wall, s.StartId, s.EndId)
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
