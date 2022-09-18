package main

import (
	"github.com/CatBloom/MahjongMasterApi/controllers"
	"github.com/CatBloom/MahjongMasterApi/db"
	"github.com/CatBloom/MahjongMasterApi/firebase"
	"github.com/CatBloom/MahjongMasterApi/logger"
	"github.com/CatBloom/MahjongMasterApi/server"
	"github.com/CatBloom/MahjongMasterApi/services"
)

func main() {
	db.Init()
	defer db.Close()

	logger.LoggerInit()
	defer logger.LoggerClose()

	firebase.Init()

	leagueService := services.NewLeagueService(db.GetDB())
	leagueController := controllers.NewLeagueController(leagueService)

	playerService := services.NewPlayerService(db.GetDB())
	playerController := controllers.NewPlayerController(playerService)

	gameService := services.NewGameService(db.GetDB())
	gameController := controllers.NewGameController(gameService)

	resultService := services.NewResultService(db.GetDB())
	resultController := controllers.NewResultController(resultService)

	serve := server.NewServer(leagueController, playerController, gameController, resultController)
	serve.Init()
}
