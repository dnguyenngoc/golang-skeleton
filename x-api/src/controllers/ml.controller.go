/**
 * @author: Duy Nguyen
 * @email: duynguyenngoc@hotmail.com
 * @create: 2022-1-1
 */

package controllers

import (
	"fmt"
	"net/http"

	"github.com/apot-group/golang-skeleton/x-api/src/apientities"
	"github.com/apot-group/golang-skeleton/x-api/src/helpers"
	"github.com/apot-group/golang-skeleton/x-api/src/tasks"
	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
	// "github.com/apot-group/golang-skeleton/x-api/src/loggers"
)

// Face detection
// @Summary      upload image and detection
// @Description  upload image and get face detection
// @Tags         ML-Face
// @Accept       multipart/form-data
// @Produce      json
// @Param        file formData file true  "select file"
// @Success      200  {object}  apientities.MlFace
// @Failure      400  {string}  string  "Bad Request"
// @Failure      404  {string}  string  "Not Found"
// @Failure      500  {string}  string  "Internal Server Error"
// @Router       /ml/face/detection [post]
func FaceDetection(c *gin.Context) {
	id := uuid.NewV4()
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename := header.Filename
	path, err := helpers.UploadImage(filename, file)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	result := apientities.MlFace{Status: "SUCCESS", FileUrl: "http://localhost:8080" + path, FilePath: path, TaskId: id.String()}
	tasks.SentTask(
		"",
		"x-ml",
		"x-ml.face_detection",
		[]string{},
		map[string]interface{}{"path": path, "task_id": id.String()})
	c.JSON(http.StatusOK, result)
}

// Face recognition
// @Summary      upload image and recognition
// @Description  upload image and get face recognition
// @Tags         ML-Face
// @Accept       multipart/form-data
// @Produce      json
// @Param        file formData file true  "select file"
// @Success      200  {object}  apientities.MlFace
// @Failure      400  {string}  string  "Bad Request"
// @Failure      404  {string}  string  "Not Found"
// @Failure      500  {string}  string  "Internal Server Error"
// @Router       /ml/face/recognition [post]
func FaceRecognition(c *gin.Context) {
	id := uuid.NewV4()
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename := header.Filename
	path, err := helpers.UploadImage(filename, file)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	result := apientities.MlFace{Status: "SUCCESS", FileUrl: "http://localhost:8080" + path, FilePath: path, TaskId: id.String()}
	tasks.SentTask(
		"",
		"x-ml",
		"x-ml.face_recognition",
		[]string{},
		map[string]interface{}{"path": path, "task_id": id.String()})
	c.JSON(http.StatusOK, result)
}
