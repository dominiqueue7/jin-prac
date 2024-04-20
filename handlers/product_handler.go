// handlers/product_handler.go

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProductDetailHandler(c *gin.Context) {
    // URL 파라미터에서 상품 ID를 가져옵니다.
    id := c.Param("id")

    // 해당 상품의 상세 정보를 데이터베이스에서 조회합니다.
    // 여기서는 간단히 상품 ID만 출력하도록 예시로 작성합니다.
    c.HTML(http.StatusOK, "product_detail.html", gin.H{
        "title": "상품 상세 페이지",
        "productId": id,
    })
}
