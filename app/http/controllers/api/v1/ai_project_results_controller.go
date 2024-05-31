package v1

import (
	"datawiz-aiservices/app/models/ai_project_result"
	"datawiz-aiservices/app/models/translation"
	"datawiz-aiservices/app/requests"
	"datawiz-aiservices/pkg/app"
	"datawiz-aiservices/pkg/response"
	"datawiz-aiservices/pkg/translator"

	"github.com/gin-gonic/gin"
)

type AiProjectResultsController struct {
	BaseAPIController
}

func getTransAiProjectResult(aiProjectResultModel *ai_project_result.AiProjectResult) {
	namekey := aiProjectResultModel.Name
	desckey := aiProjectResultModel.Description

	tranV := translation.GetTs([]string{namekey, desckey}, app.Language)

	aiProjectResultModel.Name = tranV[namekey]
	aiProjectResultModel.Description = tranV[desckey]
}

func (ctrl *AiProjectResultsController) Index(c *gin.Context) {
	aiProjectResults := ai_project_result.All()
	length := len(aiProjectResults)
	for i := 0; i < length; i++ {
		getTransAiProjectResult(&aiProjectResults[i])
	}
	response.Data(c, aiProjectResults)
}

func (ctrl *AiProjectResultsController) Show(c *gin.Context) {
	id := c.Param("id")
	aiProjectResultModel := ai_project_result.Get(id)
	if aiProjectResultModel.ID == 0 {
		aiProjectResultModels := ai_project_result.GetByUUID(id)
		length := len(aiProjectResultModels)
		for i := 0; i < length; i++ {
			getTransAiProjectResult(&aiProjectResultModels[i])
		}
		response.Data(c, aiProjectResultModels)
	} else {
		getTransAiProjectResult(&aiProjectResultModel)
		response.Data(c, aiProjectResultModel)
	}
}

func (ctrl *AiProjectResultsController) Store(c *gin.Context) {

	request := requests.AiProjectResultRequest{}
	if ok := requests.Validate(c, &request, requests.AiProjectResultSave); !ok {
		return
	}

	aiProjectResultModel := ai_project_result.AiProjectResult{
		Name:        "",
		Description: "",

		Input:  request.Input,
		Output: request.Output,

		Status: request.Status,

		UserID:        request.UserID,
		AiModelUUID:   request.AiModelUUID,
		AiProjectUUID: request.AiProjectUUID,
	}

	err := aiProjectResultModel.CreateTx(&request)

	if err == nil {
		aiProjectResultModel.Name = request.Name
		aiProjectResultModel.Description = request.Description
		response.Created(c, aiProjectResultModel)
	} else {
		response.Abort500(c, translator.TransHandler.T("r.cFailed"))
	}
}

func (ctrl *AiProjectResultsController) Update(c *gin.Context) {

	aiProjectResultModel := ai_project_result.Get(c.Param("id"))
	if aiProjectResultModel.ID == 0 {
		response.Abort404(c)
		return
	}

	request := requests.AiProjectResultRequest{}
	ok := requests.Validate(c, &request, requests.AiProjectResultSave)
	if !ok {
		return
	}

	aiProjectResultModel.Input = request.Input
	aiProjectResultModel.Output = request.Output

	aiProjectResultModel.Status = request.Status

	// aiProjectResultModel.UserID = request.UserID
	// aiProjectResultModel.AiModelUUID = request.AiModelUUID
	// aiProjectResultModel.AiProjectUUID = request.AiProjectUUID

	err := aiProjectResultModel.SaveTx(&request, false)
	if err == nil {
		aiProjectResultModel.Name = request.Name
		aiProjectResultModel.Description = request.Description
		response.Data(c, aiProjectResultModel)
	} else {
		response.Abort500(c, translator.TransHandler.T("r.uFailed"))
	}
}

func (ctrl *AiProjectResultsController) Delete(c *gin.Context) {

	aiProjectResultModel := ai_project_result.Get(c.Param("id"))
	if aiProjectResultModel.ID == 0 {
		response.Abort404(c)
		return
	}

	rowsAffected := aiProjectResultModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, translator.TransHandler.T("r.dFailed"))
}
