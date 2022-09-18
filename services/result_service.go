package services

import (
	"github.com/CatBloom/MahjongMasterApi/models"
	"github.com/CatBloom/MahjongMasterApi/models/response"
	"gorm.io/gorm"
)

type ResultService interface {
	GetPlayerResults(pid string) (models.Player, error)
	GetLeagueResults(lid string) ([]response.LeagueResultResponce, error)
	GetPlayerAgg(pid string) (response.PlayerAggResponse, error)
	GetPlayerPie(pid string) ([]response.PieResponse, error)
	GetPlayerLine(pid string) ([]response.LineResponse, error)
}

type resultService struct {
	db *gorm.DB
}

func NewResultService(db *gorm.DB) ResultService {
	return &resultService{db}
}

//player毎の成績
func (s resultService) GetPlayerResults(pid string) (models.Player, error) {
	var p models.Player

	if err := s.db.Preload("Games.Results", func(db *gorm.DB) *gorm.DB {
		return db.Select("*").Order("Results.rank ASC").Joins("left join players on players.id = results.player_id")
	}).Preload("Games", func(db *gorm.DB) *gorm.DB {
		return db.Order("Games.created_at DESC")
	}).Where("id = ?", pid).Find(&p).Error; err != nil {
		return p, err
	}

	return p, nil
}

//league毎の成績
func (s resultService) GetLeagueResults(lid string) ([]response.LeagueResultResponce, error) {
	var results []response.LeagueResultResponce

	if err := s.db.
		Table("results").
		Select("RANK() OVER(ORDER BY SUM(results.calc_point) DESC )rank,players.id AS player_Id,players.name AS player_name ,COUNT(*) AS total_game,SUM(results.point) AS total_point ,SUM(results.calc_point) AS total_calc_point ,Round(avg(results.rank),2) AS average_rank ,league_id").
		Joins("LEFT JOIN players ON players.id = results.player_id").
		Group("players.id").
		Order("total_calc_point desc,total_game desc").
		Where("league_id = ?", lid).
		Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}

//playerの集計
func (s resultService) GetPlayerAgg(pid string) (response.PlayerAggResponse, error) {
	var results response.PlayerAggResponse

	if err := s.db.
		Table("results").
		Select("players.id AS player_Id,players.name AS player_name ,COUNT(*) AS total_game,SUM(results.point) AS total_point ,SUM(results.calc_point) AS total_calc_point,Round(avg(results.rank),2) AS average_rank").
		Joins("LEFT JOIN players ON players.id = results.player_id").
		Group("players.id").
		Where("player_id = ?", pid).
		Find(&results).Error; err != nil {
		return results, err
	}

	return results, nil
}

//円グラフ用
func (s resultService) GetPlayerPie(pid string) ([]response.PieResponse, error) {
	var pie []response.PieResponse

	if err := s.db.
		Table("results").
		Select("rank, COUNT(rank) AS count_rank ,player_id").
		Group("rank,player_id").
		Where("player_id = ?", pid).
		Find(&pie).Error; err != nil {
		return pie, err
	}

	return pie, nil
}

//折れ線グラフ用
func (s resultService) GetPlayerLine(pid string) ([]response.LineResponse, error) {
	var line []response.LineResponse

	subQuery1 := s.db.
		Table("results").
		Select("created_at,rank").
		Limit(10).
		Where("player_id = ?", pid).
		Order("created_at DESC")

	if err := s.db.
		Table("(?) as line", subQuery1).
		Order("created_at ASC").
		Find(&line).Error; err != nil {
		return line, err
	}

	return line, nil
}
