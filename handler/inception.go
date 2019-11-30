package handler

import (
	"fmt"
	"net/http"
	"rds/goinception/v1/service"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type Response struct {
	Code    string      `json:"status"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

type incResponse struct {
	TaskId string `json:"task_id"`
}

// @Summary inception 预检或者执行
// @Description 预检或者执行
// @Tags inception
// @Accept  json
// @Produce  json
// @Param inception
// @Success 200 {object} user.CreateResponse "{"status":0,"msg":"OK","data":[]}"
// @Router /user [post]
func Inception(c *gin.Context) {
	var r service.CreateInception
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    "404",
			Message: err.Error(),
			Data:    []string{},
		})
	}
	// or error handling
	var err error
	u1 := uuid.Must(uuid.NewV4(), err)
	fmt.Printf("UUIDv4: %s\n", u1)

	ml, err := service.InceptionTest(r)
	ml = ml
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    "404",
			Message: err.Error(),
			Data:    []string{},
		})
	}
	var resp incResponse
	resp.TaskId = fmt.Sprintf("%s", u1)
	c.JSON(http.StatusOK, Response{
		Code:    "0000",
		Message: "成功",
		Data:    resp,
	})

}
