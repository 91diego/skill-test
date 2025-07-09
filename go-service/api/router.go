package api

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sevengit-wq/skill-test/config"
	"github.com/sevengit-wq/skill-test/internal/handlers"
	"github.com/sevengit-wq/skill-test/internal/services"
	"github.com/sevengit-wq/skill-test/pkg/constants"
	"github.com/sevengit-wq/skill-test/pkg/utils"
	/*"github.com/sevengit-wq/skill-test/internal/handlers"
	repositories "github.com/sevengit-wq/skill-test/internal/repository"
	"github.com/sevengit-wq/skill-test/internal/services"
	"github.com/sevengit-wq/skill-test/constants"
	"github.com/sevengit-wq/skill-test/pkg/utils"*/)

type Router struct {
	engine   *gin.Engine
	handler  handlers.Handler
	apiGroup *gin.RouterGroup
}

func NewRouter(core *config.CoreResources) *Router {
	service := services.NewService(utils.NewLogrusLogger(), core.ApiService)
	handler := handlers.NewHandler(service)
	gin.SetMode(gin.DebugMode)
	engine := gin.Default()
	return &Router{
		engine:   engine,
		handler:  handler,
		apiGroup: engine.Group(constants.API),
	}
}

func (router *Router) SetupRoutes(core *config.CoreResources) *gin.Engine {
	router.engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet, http.MethodPost, http.MethodPut,
			http.MethodPatch, http.MethodDelete, http.MethodHead, http.MethodOptions,
		},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	router.setupStudentsRoutes()
	return router.engine
}

func (router *Router) setupStudentsRoutes() {
	students := router.apiGroup.Group(fmt.Sprintf("%s/students/:id/report", constants.APIVersion))
	students.Use()
	{
		students.GET("", router.handler.GetReportByStudentID)
	}
}
