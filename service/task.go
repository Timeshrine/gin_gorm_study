package service

import (
	"gin+gorm_study/model"
	"gin+gorm_study/serializer"
	"time"
)

type CreateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` //0未做，1已做
}

type ShowTaskService struct {
}

type ListTaskService struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}

type UpdateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` //0未做，1已做
}

type SearchTaskService struct {
	Info     string `json:"info" form:"info"`
	PageNum  int    `json:"page_num" form:"page_num"`
	PageSize int    `json:"page_size" form:"page_size"`
}

type DeleteTaskService struct {
}

// 新增备忘录
func (service *CreateTaskService) Create(id uint) serializer.Response {
	var user model.User
	code := 200
	model.DB.First(&user, id)
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Status:    0,
		Content:   service.Content,
		StartTime: time.Now().Unix(),
		EndTime:   0,
	}
	err := model.DB.Create(&task).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "创建失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "创建成功",
	}
}

// 展示一条备忘录
func (service *ShowTaskService) Show(tid string) serializer.Response {
	var task model.Task
	code := 200
	err := model.DB.First(&task, tid).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "查询失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildTask(task),
	}
}

// 列表返回用户所有备忘录
func (service *ListTaskService) List(uid uint) serializer.Response {
	var tasks []model.Task
	count := 0
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).Count(&count).
		Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageNum).Find(&tasks)
	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(count))
}

func (service *UpdateTaskService) Update(tid string) serializer.Response {
	var task model.Task
	model.DB.First(&task, tid)
	task.Content = service.Content
	task.Status = service.Status
	task.Title = service.Title
	model.DB.Save(&task)
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildTask(task),
		Msg:    "更新完成",
	}
}

func (service *SearchTaskService) Search(uid uint) serializer.Response {
	var tasks []model.Task
	count := 0
	if service.PageSize == 0 {
		service.PageSize = 10
	}
	model.DB.Model(&model.Task{}).Preload("User").Where("uid=?").
		Where("title LIKE ? OR content LIKE ?", "%"+service.Info+"%").
		Count(&count).Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&tasks)
	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(count))
}

func (service *DeleteTaskService)Delete(tid string)serializer.Response  {
	var task model.Task
	err:=model.DB.Delete(&task,tid).Error
	if err!=nil{
		return serializer.Response{
			Status: 500,
			Msg: "删除失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg: "删除成功",
	}
}