package v1

import (
	"datawiz-aiservices/pkg/response"

	"github.com/gin-gonic/gin"
)

type HealthsController struct {
	BaseAPIController
}

func (ctrl *HealthsController) Health(c *gin.Context) {
	response.Success(c)
}
