package handler

import (
	"net/http"
	"strconv"

	"github.com/bluebeel/orm/app/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

func GetAllTasks(db *gorm.DB, c *gin.Context) {
	projectTitle := c.Param("title")
	project := getProjectOr404(db, projectTitle, c)
	if project == nil {
		return
	}

	tasks := []model.Task{}
	if err := db.Model(&project).Related(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}

func CreateTask(db *gorm.DB, c *gin.Context) {
	projectTitle := c.Param("title")
	project := getProjectOr404(db, projectTitle, c)
	if project == nil {
		return
	}

	task := model.Task{ProjectID: project.ID}

	var taskRequest model.Task
	if err := c.BindJSON(&taskRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	copier.Copy(&task, &taskRequest)
	if err := db.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

func GetTask(db *gorm.DB, c *gin.Context) {
	projectTitle := c.Param("title")
	project := getProjectOr404(db, projectTitle, c)
	if project == nil {
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	task := getTaskOr404(db, id, c)
	if task == nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

func UpdateTask(db *gorm.DB, c *gin.Context) {
	projectTitle := c.Param("title")
	project := getProjectOr404(db, projectTitle, c)
	if project == nil {
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	task := getTaskOr404(db, id, c)
	if task == nil {
		return
	}

	var taskRequest model.Task
	if err := c.BindJSON(&taskRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	copier.Copy(&task, &taskRequest)
	if err := db.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

func DeleteTask(db *gorm.DB, c *gin.Context) {
	projectTitle := c.Param("title")
	project := getProjectOr404(db, projectTitle, c)
	if project == nil {
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	task := getTaskOr404(db, id, c)
	if task == nil {
		return
	}

	if err := db.Delete(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "task deleted",
	})
}

func CompleteTask(db *gorm.DB, c *gin.Context) {
	projectTitle := c.Param("title")
	project := getProjectOr404(db, projectTitle, c)
	if project == nil {
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	task := getTaskOr404(db, id, c)
	if task == nil {
		return
	}

	task.Complete()
	if err := db.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

func UndoTask(db *gorm.DB, c *gin.Context) {
	title := c.Param("title")
	project := getProjectOr404(db, title, c)
	if project == nil {
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	task := getTaskOr404(db, id, c)
	if task == nil {
		return
	}

	task.Undo()
	if err := db.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

// getTaskOr404 gets a task instance if exists, or respond the 404 error otherwise
func getTaskOr404(db *gorm.DB, id int, c *gin.Context) *model.Task {
	task := model.Task{}
	if err := db.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return nil
	}
	return &task
}
