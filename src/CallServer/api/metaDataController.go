package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (b *BaseController) GetCallMetadata(c *gin.Context) {
	metadata, err := b.Persistence.GetMetadata()

	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "")
	} else {
		c.JSON(http.StatusOK, metadata)
	}
}
