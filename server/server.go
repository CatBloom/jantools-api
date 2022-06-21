package server

import (
	"time"

	"github.com/CatBloom/MahjongMasterApi/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Init is initialize server
func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
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
		ctrl := controllers.LeagueController{}
		// l.GET("", ctrl.Index)
		l.GET("/:uid", ctrl.Show)
		l.POST("", ctrl.Create)
		l.PUT("/:id", ctrl.Update)
		l.DELETE("/:id", ctrl.Delete)
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
