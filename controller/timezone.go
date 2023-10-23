package controller

import (
	"encoding/json"
	"fmt"
	"github.com/felixlambertv/nextalent/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetCurrentTime(c *gin.Context) {
	timezone := c.Query("zone")

	_, err := time.LoadLocation(timezone)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: "Invalid timezone",
			Status:  false,
			Data:    nil,
		})
		return
	}

	url := fmt.Sprintf("https://timeapi.io/api/Time/current/zone?timeZone=%s", timezone)
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: "Fail fetch data from time api",
			Status:  false,
			Data:    nil,
		})
		return
	}
	defer resp.Body.Close()

	var timeResponse response.TimezoneResponse
	decoder := json.NewDecoder(resp.Body)
	fmt.Println(decoder)
	if err := decoder.Decode(&timeResponse); err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: "Something went wrong",
			Status:  false,
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusBadRequest, response.BaseResponse{
		Message: "Success get timezone data",
		Status:  true,
		Data:    timeResponse,
	})
}
