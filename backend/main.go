package main

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"klauskie.com/pathfinder/backend/models"
	"klauskie.com/pathfinder/backend/src"
)

// GOOS=linux go build -o pathLambda

func main() {
	lambda.Start(HandlerApiGateway)
}

func HandlerApiGateway(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	input := unMarshalInputEvent([]byte(request.Body))

	output, err := Run(input)
	if err != nil {
		return lResponse(err.Error(), 404), err
	}

	outputJson, err := json.Marshal(&output)
	if err != nil {
		return lResponse("Failed marshaling output to json", 500), err
	}

	return lResponse(string(outputJson), 200), nil
}

func lResponse(body string, statusCode int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{Body: body, StatusCode: statusCode}
}

func Run(event InputEvent) (OutputEvent, error) {
	search := models.Create(event.Algo, event.Wall, event.StartId, event.EndId, event.Rows, event.Cols)
	if search == nil {
		return OutputEvent{}, errors.New("No algorithm found with provided algoId")
	}
	nodes, path := search.Run()

	output := OutputEvent {
		Data:    nodes,
		Grid:    search.GetPathfinder().Grid,
		Path:    path,
		StartId: search.GetPathfinder().StartId,
		EndId:   search.GetPathfinder().EndId,
		Walls:   event.Wall,
	}

	return output, nil
}

func unMarshalInputEvent(body []byte) InputEvent {
	var input InputEvent
	json.Unmarshal(body, &input)
	return input
}

type InputEvent struct {
	Wall []int `json:"wall"`
	Algo int `json:"algo"`
	StartId int `json:"start_id"`
	EndId int `json:"end_id"`
	Rows int `json:"rows"`
	Cols int `json:"cols"`
}

type OutputEvent struct {
	Data []src.Node `json:"Data"`
	Grid [][]*src.Node `json:"Grid"`
	Path []src.Node `json:"Path"`
	StartId int `json:"StartId"`
	EndId int `json:"EndId"`
	Walls []int `json:"Walls"`
}
