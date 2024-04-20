package main

import (
	"jin-prac/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // 정적 파일 서빙
    r.Static("/static", "./static")

    // HTML 템플릿 파일을 로드하기 위한 설정
    r.LoadHTMLGlob("templates/*.html")

    // 라우팅
    r.GET("/", handlers.IndexHandler)
    r.GET("/shop", handlers.ShopHandler)
    r.GET("/product/:id", handlers.ProductDetailHandler)
    r.POST("/checkout", handlers.CheckoutHandler)

    // 404 페이지 핸들러 설정
    r.NoRoute(handlers.NoRouteHandler)


    r.Run(":8080")

}
