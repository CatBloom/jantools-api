package firebase

import (
	"context"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
)

var (
	app *firebase.App
	err error
)

func Init() {
	app, err = firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
}

//middlewear
func APIAuthWrap() func(*gin.Context) {
	return func(c *gin.Context) {
		//特定URL以外のGETは認証確認しない
		if c.Request.Method == "GET" && c.Request.RequestURI != "/api/v1/league/list" {
			return
		}

		bearer := c.GetHeader("Authorization")
		if bearer == "" {
			c.String(http.StatusBadRequest, "Authorizationヘッダが設定されていません")
			c.Abort()
			return
		}
		idToken := strings.TrimPrefix(bearer, "Bearer ")
		if idToken == bearer {
			c.String(http.StatusBadRequest, "Authorization: Bearer ヘッダが設定されていません")
			c.Abort()
			return
		}

		ctx := context.Background()
		client, err := app.Auth(ctx)
		if err != nil {
			log.Print(err)
		}
		token, err := client.VerifyIDToken(ctx, idToken)
		if err != nil {
			log.Print(err)
		}

		//contextにuidを追加
		c.Set("fb_uid", token.UID)
		//contextにdisplayNameを追加
		c.Set("fb_name", token.Claims["name"])
	}

}
