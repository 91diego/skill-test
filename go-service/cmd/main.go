package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sevengit-wq/skill-test/api"
	"github.com/sevengit-wq/skill-test/config"
)

var (
	routes *gin.Engine
	core   *config.CoreResources
)

func init() {
	var err error
	core, err = config.SetupCoreResources()
	if err != nil {
		log.Fatalf("failed to setup core resources: %v", err)
	}
}

func main() {
	core.Log.Info("starting sutdents microservice")
	routes := api.NewRouter(core)
	err := routes.SetupRoutes(core).Run(fmt.Sprintf(":%v", core.Port))
	if err != nil {
		log.Fatalln(fmt.Sprintf("failed to listen and serve on port %v", core.Port) +
			err.Error())
	}
}
