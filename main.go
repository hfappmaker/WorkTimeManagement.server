package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// // controller
	// http.HandleFunc("/", echoHello)
	// // port
	// http.ListenAndServe(":8000", nil)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		// アクセス許可するオリジン
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		// アクセス許可するHTTPメソッド
		AllowMethods: []string{
			"POST",
			"GET",
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
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"countUp": 5,
		})
	})
	r.Run(":8000") // 0.0.0.0:8080 でサーバーを立てます。
}

// func echoHello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "<h1>Hello World</h1>")
// }
