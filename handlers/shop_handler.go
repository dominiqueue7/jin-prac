// handlers/index_handler.go

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShopHandler(c *gin.Context) {
    // 메인 페이지 템플릿을 렌더링합니다.
    c.HTML(http.StatusOK, "shop.html", gin.H{
        "title": "Shop",
    })
}
