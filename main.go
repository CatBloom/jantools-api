package main

import (
	"github.com/CatBloom/MahjongMasterApi/db"
	"github.com/CatBloom/MahjongMasterApi/server"
	"github.com/CatBloom/MahjongMasterApi/service"
)

func main() {
	db.Init()
	service.NewLeagueService(db.GetDB())
	defer db.Close()
	server.Init()
}
