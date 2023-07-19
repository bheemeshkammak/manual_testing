package controllers

import (
	"github.com/bheemeshkammak/manual_testing/manual_testing/pkg/rest/server/models"
	"github.com/bheemeshkammak/manual_testing/manual_testing/pkg/rest/server/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type ManController struct {
	manService *services.ManService
}

func NewManController() (*ManController, error) {
	manService, err := services.NewManService()
	if err != nil {
		return nil, err
	}
	return &ManController{
		manService: manService,
	}, nil
}

func (manController *ManController) CreateMan(context *gin.Context) {
	// validate input
	var input models.Man
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger man creation
	if _, err := manController.manService.CreateMan(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Man created successfully"})
}

func (manController *ManController) UpdateMan(context *gin.Context) {
	// validate input
	var input models.Man
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger man update
	if _, err := manController.manService.UpdateMan(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Man updated successfully"})
}

func (manController *ManController) FetchMan(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger man fetching
	man, err := manController.manService.GetMan(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, man)
}

func (manController *ManController) DeleteMan(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger man deletion
	if err := manController.manService.DeleteMan(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Man deleted successfully",
	})
}

func (manController *ManController) ListMen(context *gin.Context) {
	// trigger all men fetching
	men, err := manController.manService.ListMen()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, men)
}

func (*ManController) PatchMan(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*ManController) OptionsMan(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*ManController) HeadMan(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
