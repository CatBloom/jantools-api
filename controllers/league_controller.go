package controllers

import (
	"fmt"

	"github.com/CatBloom/MahjongMasterApi/services"
	"github.com/gin-gonic/gin"
)

type LeagueController interface {
	Search(c *gin.Context)
	List(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type leagueController struct {
	s services.LeagueService
}

func NewLeagueController(s services.LeagueService) LeagueController {
	return &leagueController{s}
}

func (lc leagueController) Search(c *gin.Context) {
	value := c.Params.ByName("value")
	l, err := lc.s.SearchLeague(value)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, l)
	}
}

func (lc leagueController) List(c *gin.Context) {
	l, err := lc.s.GetLeagueList(c)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, l)
	}
}

func (lc leagueController) Get(c *gin.Context) {
	id := c.Params.ByName("id")
	l, err := lc.s.GetLeague(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, l)
	}
}

func (lc leagueController) Create(c *gin.Context) {
	l, err := lc.s.CreateLeague(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, l)
	}
}

func (lc leagueController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	l, err := lc.s.UpdateLeague(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, l)
	}
}

func (lc leagueController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")

	if err := lc.s.DeleteLeague(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(200, gin.H{"id": id})
	}
}
