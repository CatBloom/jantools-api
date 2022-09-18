package controllers

import (
	"fmt"

	"github.com/CatBloom/MahjongMasterApi/models/response"
	"github.com/CatBloom/MahjongMasterApi/services"
	"github.com/gin-gonic/gin"
)

type ResultController interface {
	GetPlayerResults(c *gin.Context)
	GetLeagueResults(c *gin.Context)
	GetPlayerPie(c *gin.Context)
	GetPlayerLine(c *gin.Context)
}

type resultController struct {
	s services.ResultService
}

func NewResultController(s services.ResultService) ResultController {
	return &resultController{s}
}

func (pc resultController) GetPlayerResults(c *gin.Context) {
	pid := c.Params.ByName("pid")
	result, err := pc.s.GetPlayerResults(pid)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	a, err := pc.s.GetPlayerAgg(pid)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	res := response.PlayerAggResponse{}

	res.Player = result
	res.TotalGame = a.TotalGame
	res.TotalPoint = a.TotalPoint
	res.TotalCalcPoint = a.TotalCalcPoint
	res.AverageRank = a.AverageRank

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, res)
	}
}

func (pc resultController) GetLeagueResults(c *gin.Context) {
	lid := c.Params.ByName("lid")
	result, err := pc.s.GetLeagueResults(lid)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, result)
	}
}

func (pc resultController) GetPlayerPie(c *gin.Context) {
	pid := c.Params.ByName("pid")
	result, err := pc.s.GetPlayerPie(pid)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, result)
	}
}

func (pc resultController) GetPlayerLine(c *gin.Context) {
	pid := c.Params.ByName("pid")
	result, err := pc.s.GetPlayerLine(pid)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, result)
	}
}
