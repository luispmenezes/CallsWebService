package api

import (
	"CallServer/persistence"
	"github.com/gin-gonic/gin"
)

type BaseController struct {
	Engine      *gin.Engine
	Persistence *persistence.Manager
}

func NewBaseController(persistence *persistence.Manager) *BaseController {
	return &BaseController{
		Engine:      gin.Default(),
		Persistence: persistence,
	}
}

func (b *BaseController) initializeRoutes(){
	b.Engine.PUT("/call", b.CreateCalls)
	b.Engine.DELETE("/call", b.DeleteCall)
	b.Engine.GET("/call", b.GetAllCalls)
	b.Engine.GET("/callMetadata", b.GetCallMetadata)
}

func (b *BaseController) Start(port string) error {
	b.initializeRoutes()
	return b.Engine.Run(":"+port)
}
