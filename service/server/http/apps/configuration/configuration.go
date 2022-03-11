package configuration

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jahid90-cloud/konfig/service/server/http/apps"
	"github.com/jahid90-cloud/konfig/service/server/http/utils"
)

type configurationApp struct{}

func App() apps.App {
	return &configurationApp{}
}

func (a *configurationApp) Router(server *gin.RouterGroup) *gin.RouterGroup {
	configurationRouter := server.Group("/config")
	{
		configurationRouter.POST("/", a.handleAddNew)
		configurationRouter.GET("/", a.handleGetAll)
		configurationRouter.GET("/:id", a.handleGetOne)
		configurationRouter.PUT("/:id", a.handleUpdateOne)
		configurationRouter.DELETE("/:id", a.handleDeleteOne)
	}

	return configurationRouter
}

func (a *configurationApp) handleGetAll(ctx *gin.Context) {
	utils.GetLoggerFromRequest(ctx).Info("handling get all config request")

	utils.GetLoggerFromRequest(ctx).Info("fetching all configurationa")
	// config, err := service.FindAll()
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	// }

	ctx.JSON(http.StatusOK, gin.H{"data": make([]string, 0)})
}

func (a *configurationApp) handleGetOne(ctx *gin.Context) {
	utils.GetLoggerFromRequest(ctx).Info("handling get one config request")

	id, err := utils.MapParamIdToConfiguration(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	utils.GetLoggerFromRequest(ctx).Info("fetching configuration with id:", id)
	// config, err := service.FindOne(id)	// err == ErrNotFound
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// }

	ctx.JSON(http.StatusOK, gin.H{"data": id})
}

func (a *configurationApp) handleAddNew(ctx *gin.Context) {
	utils.GetLoggerFromRequest(ctx).Info("handling add new config request")

	var request newConfigurationRequest

	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	utils.GetLoggerFromRequest(ctx).Info("fetching configuration with id:", request.Uuid)
	// config, err := service.FindOne(id)
	// if config != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": "already exists"})
	// 	return
	// }

	utils.GetLoggerFromRequest(ctx).Info("creating new configuration:", request)

	ctx.Status(http.StatusCreated)
}

func (a *configurationApp) handleUpdateOne(ctx *gin.Context) {
	utils.GetLoggerFromRequest(ctx).Info("handling update one config request")

	id, err := utils.MapParamIdToConfiguration(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	utils.GetLoggerFromRequest(ctx).Info("fetching configuration with id:", id)
	// config, err := service.FindOne(id)	// err == ErrNotFound
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// }

	var request updateConfigurationRequest

	err = ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	utils.GetLoggerFromRequest(ctx).Info("updating configuration with id:", id)
	// err := service.UpdateOne(request)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	// 	return
	// }

	ctx.JSON(http.StatusNoContent, gin.H{"data": id})
}

func (a *configurationApp) handleDeleteOne(ctx *gin.Context) {
	utils.GetLoggerFromRequest(ctx).Info("handling delete one config request")

	id, err := utils.MapParamIdToConfiguration(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	utils.GetLoggerFromRequest(ctx).Info("fetching configuration with id:", id)
	// err := service.FindOne(id)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }

	utils.GetLoggerFromRequest(ctx).Info("deleting configuration with id:", id)
	// err := service.DeleteOne(id)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	// 	return
	// }

	ctx.JSON(http.StatusNoContent, gin.H{"data": id})
}
