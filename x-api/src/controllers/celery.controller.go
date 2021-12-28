package controllers

import (
	_ "net/http"

	"github.com/apot-group/golang-skeleton/x-api/src/apientities"
	"github.com/gin-gonic/gin"
)

// Status Celery By Id
// @Summary      Celery Status
// @Description  show result celery status base on task id
// @Tags         Celery
// @Accept       json
// @Produce      json
// @Param        task_id   path string  true  "Task ID"
// @Success      200 {object}  apientities.CeleryStatus
// @Failure      400  {string}  string  "Bad Request"
// @Failure      404  {string}  string  "Not Found"
// @Failure      500  {string}  string  "Internal Server Error"
// @Router       /celery/status/{task_id} [get]
func StatusCelery(c *gin.Context) {
	taskId := c.Param("task_id")
	var result = apientities.CeleryStatus{TaskId: taskId}
	c.JSON(200, result)
	// return
}
