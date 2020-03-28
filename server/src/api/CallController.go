package api

import (
	"CallServer/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func (b *BaseController) CreateCalls(c *gin.Context) {
	callArray := &[]model.Call{}
	err := json.NewDecoder(c.Request.Body).Decode(callArray)
	if err != nil {
		log.Printf("Error processing call create json %s", err)
		c.String(http.StatusBadRequest, "Invalid request body format")
	}

	var validationErrors []model.ValidationError

	for idx, _ := range *callArray {
		validationErrors = append(validationErrors, (*callArray)[idx].Validate()...)
		(*callArray)[idx].ComputeDurationAndCost()
	}

	if len(validationErrors) > 0 {
		c.JSON(http.StatusBadRequest, validationErrors)
	} else {
		err := b.Persistence.AddCalls(callArray)
		if err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, "")
		} else {
			c.String(http.StatusOK, "")
		}
	}
}

func (b *BaseController) DeleteCall(c *gin.Context) {
	filterParams, err := b.getFilterParams(c)

	if err != nil {
		c.String(http.StatusBadRequest, "Invalid filter params")
	}

	err = b.Persistence.RemoveCall(filterParams)

	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "")
	} else {
		c.String(http.StatusOK, "")
	}
}

func (b *BaseController) GetAllCalls(c *gin.Context) {
	pageNumStr := c.Query("page")
	pageSizeStr := c.Query("pageSize")

	filterParams,err := b.getFilterParams(c)

	if err != nil {
		c.String(http.StatusBadRequest, "Invalid filter params")
	}

	log.Printf("Get All Calls Request (pageNum: \"%s\", pageSize:\"%s\")", pageNumStr, pageSizeStr)

	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil {
		pageNum = 0
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 0
	}

	calls, err := b.Persistence.GetCalls(filterParams, pageNum, pageSize)

	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "")
	} else {
		c.JSON(http.StatusOK, calls)
	}
}

func (b *BaseController) getFilterParams(c *gin.Context) (map[string]interface{}, error) {
	params := map[string]interface{}{}
	if callerStr := c.Query("caller"); len(callerStr) > 0 {
		params["caller"] = callerStr
	}

	if calleeSrt := c.Query("callee"); len(calleeSrt) > 0 {
		params["callee"] = calleeSrt
	}

	if startTimeStr := c.Query("start_time"); len(startTimeStr) > 0 {
		startTime, err := time.Parse(time.RFC3339, startTimeStr)
		if err != nil {
			log.Printf("Error parsing start date %s", startTimeStr)
			return nil, err
		}
		params["start_time"] = startTime
	}

	if endTimeStr := c.Query("end_time"); len(endTimeStr) > 0 {
		endTime, err := time.Parse(time.RFC3339, endTimeStr)
		if err != nil {
			log.Printf("Error parsing end date %s", endTimeStr)
			return nil, err
		}
		params["end_time"] = endTime
	}

	if inboundStr := c.Query("inbound"); len(inboundStr) > 0 {
		inbound, err := strconv.ParseBool(inboundStr)
		if err != nil {
			log.Printf("Error parsing inbound %s", inboundStr)
			return nil, err
		}
		params["inbound"] = inbound
	}

	if durationStr := c.Query("duration"); len(durationStr) > 0 {
		duration, err := strconv.ParseUint(durationStr, 10, 16)
		if err != nil {
			log.Printf("Error parsing duration %s", durationStr)
			return nil, err
		}
		params["duration"] = uint16(duration)
	}

	if costString := c.Query("cost"); len(costString) > 0 {
		cost, err := strconv.ParseUint(costString, 10, 32)
		if err != nil {
			log.Printf("Error parsing cost %s", costString)
			return nil, err
		}
		params["cost"] = uint32(cost)
	}

	log.Printf("Recieved filter params: %s", params)

	return params, nil
}
