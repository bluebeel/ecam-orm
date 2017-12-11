package handler

import (
	"net/http"

	"log"

	"github.com/bluebeel/ecam-orm/app/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

func GetAllProjects(db *gorm.DB, c *gin.Context) {
	projects := []model.Project{}
	db.Find(&projects)
	c.JSON(http.StatusOK, gin.H{
		"projects": projects,
	})
}

func CreateProject(db *gorm.DB, c *gin.Context) {
	var project model.Project
	if err := c.BindJSON(&project); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	if err := db.Save(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"project": project,
	})
}

func GetProject(db *gorm.DB, c *gin.Context) {
	title := c.Param("title")

	project := getProjectOr404(db, title, c)
	if project == nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"project": project,
	})
}

func UpdateProject(db *gorm.DB, c *gin.Context) {
	title := c.Param("title")

	project := getProjectOr404(db, title, c)
	if project == nil {
		return
	}

	var projectRequest model.Project
	if err := c.BindJSON(&projectRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	copier.Copy(&project, &projectRequest)

	if err := db.Save(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"project": project,
	})
}

func DeleteProject(db *gorm.DB, c *gin.Context) {
	title := c.Param("title")

	project := getProjectOr404(db, title, c)
	if project == nil {
		return
	}
	if err := db.Delete(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "project deleted",
	})
}

func ArchiveProject(db *gorm.DB, c *gin.Context) {
	title := c.Param("title")

	project := getProjectOr404(db, title, c)
	if project == nil {
		return
	}
	project.Archive()
	if err := db.Save(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"project": project,
	})
}

func RestoreProject(db *gorm.DB, c *gin.Context) {
	title := c.Param("title")

	project := getProjectOr404(db, title, c)
	if project == nil {
		return
	}
	project.Restore()
	if err := db.Save(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"project": project,
	})
}

func PublicProject(db *gorm.DB, c *gin.Context) {
	title := c.Param("title")

	project := getProjectOr404(db, title, c)
	if project == nil {
		return
	}
	project.Public()
	if err := db.Save(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"project": project,
	})
}

func PrivateProject(db *gorm.DB, c *gin.Context) {
	title := c.Param("title")

	project := getProjectOr404(db, title, c)
	if project == nil {
		return
	}
	project.Private()
	if err := db.Save(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"project": project,
	})
}

// getProjectOr404 gets a project instance if exists, or respond the 404 error otherwise
func getProjectOr404(db *gorm.DB, title string, c *gin.Context) *model.Project {
	project := model.Project{}
	if err := db.First(&project, model.Project{Title: title}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return nil
	}
	return &project
}
