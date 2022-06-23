package server

import (
	"time"

	"github.com/CatBloom/MahjongMasterApi/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	league controllers.LeagueController
	// todo player
}

func NewServer(
	league controllers.LeagueController,
	// todo player
) Server {
	return Server{
		league: league,
	}
}

// Init is initialize server
func (s Server) Init() {
	r := s.router()
	r.Run()
}

func (s Server) router() *gin.Engine {
	r := gin.Default()

	// Cors
	r.Use(cors.New(cors.Config{
		// アクセス許可するオリジン
		AllowOrigins: []string{
			"http://localhost:4200",
		},
		// アクセス許可するHTTPメソッド
		AllowMethods: []string{
			"POST",
			"GET",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		// 許可するHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Content-Type",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: false,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

	l := r.Group("/league")
	{
		// l.GET("", ctrl.Index)
		l.GET("/:uid", s.league.Show)
		l.POST("", s.league.Create)
		l.PUT("/:id", s.league.Update)
		l.DELETE("/:id", s.league.Delete)
	}

	p := r.Group("/player")
	{
		ctrl := controllers.PlayerController{}
		p.GET("/:lid", ctrl.Show)
		p.POST("", ctrl.Create)
		p.PUT("/:id", ctrl.Update)
		p.DELETE("/:id", ctrl.Delete)
	}

	return r
}
