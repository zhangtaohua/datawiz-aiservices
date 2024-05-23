package v1

import (
	"datawiz-aiservices/app/models/ai_project"
	"datawiz-aiservices/app/requests"
	"datawiz-aiservices/pkg/response"
	"datawiz-aiservices/pkg/translator"

	"github.com/gin-gonic/gin"
)

type AiProjectsController struct {
	BaseAPIController
}

func (ctrl *AiProjectsController) Index(c *gin.Context) {
	aiProjects := ai_project.All()
	response.Data(c, aiProjects)
}

func (ctrl *AiProjectsController) Show(c *gin.Context) {
	aiProjectModel := ai_project.Get(c.Param("id"))
	if aiProjectModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, aiProjectModel)
}

func (ctrl *AiProjectsController) Store(c *gin.Context) {

	request := requests.AiProjectRequest{}
	if ok := requests.Validate(c, &request, requests.AiProjectSave); !ok {
		return
	}

	aiProjectModel := ai_project.AiProject{
		Name:        request.Name,
		Description: request.Description,
		UserID:      "rj-todo",
	}
	aiProjectModel.Create()
	if aiProjectModel.ID > 0 {
		response.Created(c, aiProjectModel)
	} else {
		response.Abort500(c, translator.TransHandler.T("r.cFailed"))
	}
}

func (ctrl *AiProjectsController) Update(c *gin.Context) {

	aiProjectModel := ai_project.Get(c.Param("id"))
	if aiProjectModel.ID == 0 {
		response.Abort404(c)
		return
	}

	request := requests.AiProjectRequest{}
	ok := requests.Validate(c, &request, requests.AiProjectSave)
	if !ok {
		return
	}

	aiProjectModel.Name = request.Name
	aiProjectModel.Description = request.Description
	aiProjectModel.UserID = "rj-todo"

	rowsAffected := aiProjectModel.Save()
	if rowsAffected > 0 {
		response.Data(c, aiProjectModel)
	} else {
		response.Abort500(c, translator.TransHandler.T("r.uFailed"))
	}
}

func (ctrl *AiProjectsController) Delete(c *gin.Context) {

	aiProjectModel := ai_project.Get(c.Param("id"))
	if aiProjectModel.ID == 0 {
		response.Abort404(c)
		return
	}

	rowsAffected := aiProjectModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, translator.TransHandler.T("r.dFailed"))
}
