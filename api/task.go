package api

import (
	"gin+gorm_study/pkg/utils"
	"gin+gorm_study/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func CreateTask(c *gin.Context) {
	var createTask service.CreateTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createTask); err != nil {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	} else {
		res := createTask.Create(claim.Id)
		c.JSON(200, res)
	}
}

func ShowTask(c *gin.Context) {
	var showtask service.ShowTaskService
	if err := c.ShouldBind(&showtask); err != nil {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	} else {
		res := showtask.Show(c.Param("id"))
		c.JSON(200, res)
	}
}

func ListTask(c *gin.Context) {
	var listtask service.ListTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listtask); err != nil {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	} else {
		res := listtask.List(claim.Id)
		c.JSON(200, res)
	}
}

func UpdateTask(c *gin.Context) {
	var updatetask service.UpdateTaskService
	if err := c.ShouldBind(&updatetask); err != nil {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	} else {
		res := updatetask.Update(c.Param("id"))
		c.JSON(200, res)
	}
}

func SearchTask(c *gin.Context) {
	var searchtask service.SearchTaskService
	claim,_:=utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&searchtask); err != nil {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	} else {
		res := searchtask.Search(claim.Id)
		c.JSON(200, res)
	}
}

func DeleteTask(c *gin.Context)  {
	var deletetask service.DeleteTaskService
	if err := c.ShouldBind(&deletetask); err != nil {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	} else {
		res := deletetask.Delete(c.Param("id"))
		c.JSON(200, res)
	}
}
