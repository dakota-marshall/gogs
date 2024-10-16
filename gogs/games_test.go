package gogs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetGameById(t *testing.T) {
	assert := assert.New(t)
	var gameId int = 67679508
	game, err := ogsServer.GetGameById(gameId)
	assert.NoError(err)

	assert.Equal(gameId, game.ID, game.Gamedata.GameID)
	assert.NotEmpty(game.Name)
	assert.NotEmpty(game.Players.White.ID, game.Players.Black.ID)
	assert.NotEmpty(game.Gamedata.Clock.CurrentPlayer)
}
func TestGetGameByIdBadId(t *testing.T) {
	assert := assert.New(t)
	var gameId int = 111111111
	game, err := ogsServer.GetGameById(gameId)
	assert.Error(err)

	assert.Empty(game.ID, game.Name)
}
func TestGetGameByIdDifferentGameType(t *testing.T) {
	assert := assert.New(t)
	var gameId int = 56399513
	game, err := ogsServer.GetGameById(gameId)
	assert.NoError(err)

	assert.Equal(gameId, game.ID, game.Gamedata.GameID)
	assert.NotEmpty(game.Name)
	assert.NotEmpty(game.Players.White.ID, game.Players.Black.ID)
	assert.NotEmpty(game.Gamedata.Clock.CurrentPlayer)
}
func TestGetGameByIdRengo(t *testing.T) {
	assert := assert.New(t)
	var gameId int = 68251525
	game, err := ogsServer.GetGameById(gameId)
	assert.NoError(err)

	assert.Equal(gameId, game.ID, game.Gamedata.GameID)
	assert.NotEmpty(game.Name)
	assert.NotEmpty(game.Players.White.ID, game.Players.Black.ID)
	assert.NotEmpty(game.Gamedata.Clock.CurrentPlayer)
	assert.True(game.Gamedata.Rengo)
	assert.NotEmpty(game.Gamedata.RengoTeams.Black, game.Gamedata.RengoTeams.Black)
}
