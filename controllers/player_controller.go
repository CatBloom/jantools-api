package controllers

import (
	"fmt"

	"github.com/CatBloom/MahjongMasterApi/service"
	"github.com/gin-gonic/gin"
)

// Controller is user controlller
type PlayerController struct{}

// Create action: POST /player
func (pc PlayerController) Create(c *gin.Context) {
	var s service.PlayerService
	p, err := s.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

// Show action: GET /player/:id
func (pc PlayerController) Show(c *gin.Context) {
	lid := c.Params.ByName("lid")
	var s service.PlayerService
	p, err := s.GetByLeagueID(lid)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Update action: PUT /player/:id
func (pc PlayerController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.PlayerService
	p, err := s.UpdateByID(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Delete action: DELETE /player/:id
func (pc PlayerController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.PlayerService

	if err := s.DeleteByID(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}
