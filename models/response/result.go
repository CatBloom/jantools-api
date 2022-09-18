package response

import (
	"time"

	"github.com/CatBloom/MahjongMasterApi/models"
)

//result用のresponse構造体
type PlayerAggResponse struct {
	models.Player
	TotalGame      uint    `json:"totalGame"`
	TotalPoint     int32   `json:"totalPoint"`
	TotalCalcPoint float64 `json:"totalCalcPoint"`
	AverageRank    float64 `json:"averageRank"`
}

type GameResponse struct {
	ID        uint             `json:"id" gorm:"primary_key"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
	DeletedAt *time.Time       `json:"deletedAt"`
	Results   []ResultResponse `json:"results"`
	Players   []models.Player  `json:"players" gorm:"many2many:games_players"`
}

type ResultResponse struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	PlayerId  uint       `json:"playerId"`
	Name      string     `json:"playerName"`
	Rank      uint       `json:"rank"`
	Point     int        `json:"point"`
	CalcPoint float64    `json:"calcPoint"`
	GameID    uint       `json:"gameId"`
}

type LeagueResultResponce struct {
	PlayerId       uint    `json:"playerId"`
	Rank           string  `json:"rank"`
	PlayerName     string  `json:"name"`
	TotalGame      uint    `json:"totalGame"`
	TotalPoint     int     `json:"totalPoint"`
	TotalCalcPoint float64 `json:"totalCalcPoint"`
	AverageRank    float64 `json:"averageRank"`
	LeagueId       string  `json:"leagueId"`
}

type PieResponse struct {
	PlayerId  uint `json:"playerId"`
	Rank      uint `json:"rank"`
	CountRank uint `json:"countRank"`
}

type LineResponse struct {
	PlayerId  uint      `json:"playerId"`
	Rank      uint      `json:"rank"`
	CreatedAt time.Time `json:"createdAt"`
}
