package apis

import (
	"build-web-application-with-golang/gin_web/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func AddOrderApi(c *gin.Context) {
	id := c.Request.FormValue("id")
	code := c.Request.FormValue("code")
	p := Order{id, code}

	ra, err := p.AddOrder()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
