package controllers

import (
	"github.com/apot-group/golang-skeleton/x-api/src/helpers"

	"github.com/apot-group/golang-skeleton/x-api/src/apientities"

	"github.com/gin-gonic/gin"
	// "time"
)

// Status Celery By Id
// @Summary      Celery Status
// @Description  show result celery status base on task id
// @Tags         Celery
// @Accept       json
// @Produce      json
// @Param        task_id   path string  true  "Task ID"
// @Success      200  {object}  apientities.CeleryStatusResult
// @Failure      400  {string}  string  "Bad Request"
// @Failure      404  {string}  string  "Not Found"
// @Failure      500  {string}  string  "Internal Server Error"
// @Router       /celery/status/{task_id} [get]
func StatusCelery(c *gin.Context) {
	taskId := c.Param("task_id")
	status := apientities.CeleryStatus{Upload: "SUCCESS", Ml: "SUCCESS", General: "SUCCESS"}
	now := helpers.GetNowTime()
	times := apientities.CeleryTimes{StartUpload: now, EndUpload: now, StartMl: now, EndMl: now}
	result := apientities.CeleryStatusResult{TaskId: taskId, Status: status, Times: times}
	c.JSON(200, result)
}
