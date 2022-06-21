package service

import (
	"fmt"

	"github.com/CatBloom/MahjongMasterApi/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LeagueService struct {
	db *gorm.DB
}

func NewLeagueService(db *gorm.DB) *LeagueService {
	return &LeagueService{db}
}

// league is alias of league struct
type League models.League

// GetAll is get all leagues
func (s LeagueService) GetAll() ([]League, error) {
	// db := db.GetDB()
	var u []League

	if err := s.db.Find(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

// CreateModel is create league model
func (s LeagueService) CreateModel(c *gin.Context) (League, error) {
	// db := db.GetDB()
	var u League

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	if err := s.db.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

func (s LeagueService) GetByUID(uid string) ([]League, error) {

	var u []League
	if err := s.db.Where("uid = ?", uid).Find(&u).Error; err != nil {
		return nil, err
	}
	fmt.Println(u)

	return u, nil
}

// UpdateByID is update a league
func (s LeagueService) UpdateByID(id string, c *gin.Context) (League, error) {
	var u League

	if err := s.db.Where("league_id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	s.db.Save(&u)

	return u, nil
}

// DeleteByID is delete a league
func (s LeagueService) DeleteByID(id string) error {
	// db := db.GetDB()
	var u League

	if err := s.db.Where("league_id = ?", id).Delete(&u).Error; err != nil {
		return err
	}

	return nil
}
