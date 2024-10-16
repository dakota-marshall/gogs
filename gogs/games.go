package gogs

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Gets a game by ID from the REST API and returns a Game object
func (server *Server) GetGameById(gameId int) (Game, error) {

	var game Game

	result, err := server.NewAPIRequest("GET", "/games/"+strconv.Itoa(gameId), nil)
	if err != nil {
		return Game{}, err
	}
	if result.Code != 200 {
		return Game{}, fmt.Errorf("Got non 200 response from API: %d", result.Code)
	}

	jsonStr, marshalErr := json.Marshal(result.Results)
	if marshalErr != nil {
		return Game{}, err
	}

	if unmarshalErr := json.Unmarshal(jsonStr, &game); unmarshalErr != nil {
		return Game{}, unmarshalErr
	}

	return game, nil

}
