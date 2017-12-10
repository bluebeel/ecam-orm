package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Project struct {
	gorm.Model
	Title       string `gorm:"unique" json:"title"`
	Description string `json:"description"`
	Archived    bool   `json:"archived"`
	Tasks       []Task `json:"tasks"`
	Privated    bool   `json:"private"`
}

func (p *Project) Archive() {
	p.Archived = true
}

func (p *Project) Restore() {
	p.Archived = false
}

func (p *Project) Private() {
	p.Privated = true
}

func (p *Project) Public() {
	p.Privated = false
}

type Task struct {
	gorm.Model
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Priority    string     `json:"priority"`
	Deadline    *time.Time `gorm:"default:null" json:"deadline"`
	Done        bool       `json:"done"`
	ProjectID   uint       `json:"project_id"`
}

func (t *Task) Complete() {
	t.Done = true
}

func (t *Task) Undo() {
	t.Done = false
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Project{}, &Task{})
	return db
}
