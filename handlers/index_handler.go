// handlers/index_handler.go

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
    // 메인 페이지 템플릿을 렌더링합니다.
    c.HTML(http.StatusOK, "index.html", gin.H{
        "title": "EC 사이트",
    })
}
