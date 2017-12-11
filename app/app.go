package app

import (
	"time"

	"github.com/bluebeel/orm/app/handler"
	"github.com/bluebeel/orm/app/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// App has router and db instances
type App struct {
	Router *gin.Engine
	DB     *gorm.DB
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize() {
	//db, err := gorm.Open("sqlite3", "orm.db")
	db, err := gorm.Open("postgres", "host=crypto.cbnwm8jvjnvp.eu-west-1.rds.amazonaws.com user=bluebeel dbname=crypto password=crypto2132 sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = gin.Default()
	a.Router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))
	//gin.SetMode(gin.ReleaseMode)
	a.setRouters()
}

// setRouters sets the all required routers
// setRouters sets the all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/projects", a.GetAllProjects)
	a.Post("/projects", a.CreateProject)
	a.Get("/projects/:title", a.GetProject)
	a.Put("/projects/:title", a.UpdateProject)
	a.Delete("/projects/:title", a.DeleteProject)
	a.Put("/projects/:title/archive", a.ArchiveProject)
	a.Delete("/projects/:title/archive", a.RestoreProject)
	a.Put("/projects/:title/private", a.PrivateProject)
	a.Delete("/projects/:title/private", a.PublicProject)

	// Routing for handling the tasks
	a.Get("/projects/:title/tasks", a.GetAllTasks)
	a.Post("/projects/:title/tasks", a.CreateTask)
	a.Get("/projects/:title/tasks/:id", a.GetTask)
	a.Put("/projects/:title/tasks/:id", a.UpdateTask)
	a.Delete("/projects/:title/tasks/:id", a.DeleteTask)
	a.Put("/projects/:title/tasks/:id/complete", a.CompleteTask)
	a.Delete("/projects/:title/tasks/:id/complete", a.UndoTask)
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(c *gin.Context)) {
	//a.Router.Handle("GET", path, authMiddleware, f)
	a.Router.Handle("GET", path, f)
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(c *gin.Context)) {
	//a.Router.Handle("POST", path, authMiddleware, f)
	a.Router.Handle("POST", path, f)
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(c *gin.Context)) {
	//a.Router.Handle("PUT", path, authMiddleware, f)
	a.Router.Handle("PUT", path, f)
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(c *gin.Context)) {
	//a.Router.Handle("DELETE", path, authMiddleware, f)
	a.Router.Handle("DELETE", path, f)
}

/*
** Projects Handlers
 */
func (a *App) GetAllProjects(c *gin.Context) {
	handler.GetAllProjects(a.DB, c)
}

func (a *App) CreateProject(c *gin.Context) {
	handler.CreateProject(a.DB, c)
}

func (a *App) GetProject(c *gin.Context) {
	handler.GetProject(a.DB, c)
}

func (a *App) UpdateProject(c *gin.Context) {
	handler.UpdateProject(a.DB, c)
}

func (a *App) DeleteProject(c *gin.Context) {
	handler.DeleteProject(a.DB, c)
}

func (a *App) ArchiveProject(c *gin.Context) {
	handler.ArchiveProject(a.DB, c)
}

func (a *App) RestoreProject(c *gin.Context) {
	handler.RestoreProject(a.DB, c)
}

func (a *App) PrivateProject(c *gin.Context) {
	handler.PrivateProject(a.DB, c)
}

func (a *App) PublicProject(c *gin.Context) {
	handler.PublicProject(a.DB, c)
}

/*
** Tasks Handlers
 */
func (a *App) GetAllTasks(c *gin.Context) {
	handler.GetAllTasks(a.DB, c)
}

func (a *App) CreateTask(c *gin.Context) {
	handler.CreateTask(a.DB, c)
}

func (a *App) GetTask(c *gin.Context) {
	handler.GetTask(a.DB, c)
}

func (a *App) UpdateTask(c *gin.Context) {
	handler.UpdateTask(a.DB, c)
}

func (a *App) DeleteTask(c *gin.Context) {
	handler.DeleteTask(a.DB, c)
}

func (a *App) CompleteTask(c *gin.Context) {
	handler.CompleteTask(a.DB, c)
}

func (a *App) UndoTask(c *gin.Context) {
	handler.UndoTask(a.DB, c)
}

// Run the app on it's router
func (a *App) Run(host string) {
	a.Router.Run(host)
}
