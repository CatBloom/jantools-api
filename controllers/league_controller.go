package controllers

import (
	"fmt"

	"github.com/CatBloom/MahjongMasterApi/service"
	"github.com/gin-gonic/gin"
)

// Controller is league controlller
type LeagueController struct{}

// Index action: GET /league
func (pc LeagueController) Index(c *gin.Context) {
	var s service.LeagueService
	p, err := s.GetAll()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Create action: POST /player
func (pc LeagueController) Create(c *gin.Context) {
	var s service.LeagueService
	p, err := s.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

// Show action: GET /league/:uid
func (pc LeagueController) Show(c *gin.Context) {
	uid := c.Params.ByName("uid")
	var s service.LeagueService
	p, err := s.GetByUID(uid)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Update action: PUT /league/:id
func (pc LeagueController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.LeagueService
	p, err := s.UpdateByID(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Delete action: DELETE /league/:id
func (pc LeagueController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.LeagueService

	if err := s.DeleteByID(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}
