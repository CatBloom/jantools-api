package controllers

import (
	"fmt"
	"strconv"

	"github.com/CatBloom/MahjongMasterApi/services"
	"github.com/gin-gonic/gin"
)

type GameController interface {
	List(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type gameController struct {
	s services.GameService
}

func NewGameController(s services.GameService) GameController {
	return &gameController{s}
}

func (gc gameController) List(c *gin.Context) {
	lid := c.Params.ByName("lid")
	g, err := gc.s.GetGameList(lid)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, g)
	}
}

func (gc gameController) Get(c *gin.Context) {
	id := c.Params.ByName("id")
	g, err := gc.s.GetGame(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, g)
	}
}

func (gc gameController) Create(c *gin.Context) {
	g, err := gc.s.CreateGame(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, g)
	}
}

func (gc gameController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	if err := gc.s.CheckGameAuth(id, c); err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	}
	g, err := gc.s.UpdateGame(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, g)
	}
}

func (gc gameController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	if err := gc.s.CheckGameAuth(id, c); err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	}

	if err := gc.s.DeleteGame(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		resId, _ := strconv.Atoi(id)
		c.JSON(200, gin.H{"id": resId})
	}
}
