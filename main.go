package main

import (
	"github.com/CatBloom/MahjongMasterApi/controllers"
	"github.com/CatBloom/MahjongMasterApi/db"
	"github.com/CatBloom/MahjongMasterApi/server"
	"github.com/CatBloom/MahjongMasterApi/service"
)

func main() {
	db.Init()
	defer db.Close()

	leagueService := service.NewLeagueService(db.GetDB())
	leagueController := controllers.NewLeagueController(*leagueService)

	serve := server.NewServer(*leagueController)
	serve.Init()
}
