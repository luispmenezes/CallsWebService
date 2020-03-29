package api

import (
	"CallServer/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func (b *BaseController) GetCallMetadata(c *gin.Context) {
	var startTime time.Time
	var endTime time.Time
	var err error
	var metadata model.CallMetadata

	if startTimeStr := c.Query("startTime"); len(startTimeStr) > 0 {
		startTime, err = time.Parse(time.RFC3339, startTimeStr)
		if err != nil {
			log.Printf("Error parsing start time %s", startTimeStr)
			c.String(http.StatusBadRequest, "Invalid start time format, use RFC3339")
			return
		}
	} else {
		startTime = time.Now().Add(time.Duration(-1) * time.Hour)
	}

	if endTimeStr := c.Query("endTime"); len(endTimeStr) > 0 {
		endTime, err = time.Parse(time.RFC3339, endTimeStr)
		if err != nil {
			log.Printf("Error parsing end time %s", endTimeStr)
			c.String(http.StatusBadRequest, "Invalid end time format, use RFC3339")
			return
		}
	} else {
		endTime = time.Now()
	}

	metadata, err = b.Persistence.GetMetadata(startTime, endTime)

	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "")
	} else {
		c.JSON(http.StatusOK, metadata)
	}
}
