package v1

import (
	"datawiz-aiservices/app/models/ai_project"
	"datawiz-aiservices/app/requests"
	"datawiz-aiservices/pkg/auth"
	"datawiz-aiservices/pkg/helpers"
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

func (ctrl *AiProjectsController) PaginateIndex(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	whereFields := requests.MakeWhereFields(request, true)

	data, pager := ai_project.Paginate(c, nil, whereFields, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (ctrl *AiProjectsController) Show(c *gin.Context) {
	aiProjectModel := ai_project.Get(c.Param("id"), false)
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

	userId := auth.CurrentUID(c)

	aiProjectModel := ai_project.AiProject{
		Name:        "",
		Description: "",
		Cover:       request.Cover,
		UserID:      userId,
	}
	err := aiProjectModel.CreateTx(&request)

	if err != nil {
		response.Abort500(c, translator.TransHandler.T("r.cFailed"))
	} else {
		aiProjectModel.Name = request.Name
		aiProjectModel.Description = request.Description
		response.Created(c, aiProjectModel)
	}
}

func (ctrl *AiProjectsController) Update(c *gin.Context) {

	aiProjectModel := ai_project.Get(c.Param("id"), true)
	if aiProjectModel.ID == 0 {
		response.Abort404(c)
		return
	}

	request := requests.AiProjectRequest{}
	ok := requests.Validate(c, &request, requests.AiProjectSave)
	if !ok {
		return
	}

	aiProjectModel.Cover = request.Cover
	err := aiProjectModel.SaveTx(&request, false)
	if err == nil {
		aiProjectModel.Name = request.Name
		aiProjectModel.Description = request.Description
		response.Data(c, aiProjectModel)
	} else {
		response.Abort500(c, translator.TransHandler.T("r.uFailed"))
	}
}

func (ctrl *AiProjectsController) Patch(c *gin.Context) {

	aiProjectModel := ai_project.Get(c.Param("id"), true)
	if aiProjectModel.ID == 0 {
		response.Abort404(c)
		return
	}

	request := requests.AiProjectRequest{}
	ok := requests.Validate(c, &request, requests.AiProjectUpdate)
	if !ok {
		return
	}

	err := aiProjectModel.SaveTx(&request, true)
	if err == nil {
		if !helpers.Empty(request.Name) {
			aiProjectModel.Name = request.Name
		}
		if !helpers.Empty(request.Description) {
			aiProjectModel.Description = request.Description
		}
		response.Data(c, aiProjectModel)
	} else {
		response.Abort500(c, translator.TransHandler.T("r.uFailed"))
	}
}

func (ctrl *AiProjectsController) Delete(c *gin.Context) {

	aiProjectModel := ai_project.Get(c.Param("id"), true)
	if aiProjectModel.ID == 0 {
		response.Abort404(c)
		return
	}

	rowsAffected := aiProjectModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	// todo 删除翻译表中 对应的翻译

	response.Abort500(c, translator.TransHandler.T("r.dFailed"))
}
