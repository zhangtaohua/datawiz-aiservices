package v1

import (
	"datawiz-aiservices/app/models/ai_model"
	"datawiz-aiservices/app/requests"
	"datawiz-aiservices/pkg/response"
	"datawiz-aiservices/pkg/translator"

	"github.com/gin-gonic/gin"
)

type AiModelsController struct {
	BaseAPIController
}

func (ctrl *AiModelsController) Index(c *gin.Context) {
	aiModels := ai_model.All()
	response.Data(c, aiModels)
}

func (ctrl *AiModelsController) Show(c *gin.Context) {
	aiModelModel := ai_model.Get(c.Param("id"))
	if aiModelModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, aiModelModel)
}

func (ctrl *AiModelsController) Store(c *gin.Context) {

	request := requests.AiModelRequest{}
	if ok := requests.Validate(c, &request, requests.AiModelSave); !ok {
		return
	}

	aiModelModel := ai_model.AiModel{
		Name:        "",
		Description: "",
		Code:        request.Code,
		Type:        request.Type,
		Category:    request.Category,
		Icon:        request.Icon,
		Cover:       request.Cover,

		Framework:    request.Framework,
		Algorithm:    request.Algorithm,
		NetStructure: request.NetStructure,
		BackBone:     request.BackBone,

		Accuracy:  request.Accuracy,
		Precision: request.Precision,
		Recall:    request.Recall,
		F1Score:   request.F1Score,
		AUC:       request.AUC,

		InputFeatures:   "",
		OutputLabels:    "",
		InputParameters: request.InputParameters,
		ExecMethod:      request.ExecMethod,
		OutputFormatter: request.OutputFormatter,

		Size:       request.Size,
		Version:    request.Version,
		Status:     request.Status,
		DeployedAt: request.DeployedAt,
		RetiredAt:  request.RetiredAt,
	}

	err := aiModelModel.CreateTx(&request)
	if err == nil {
		aiModelModel.Name = request.Name
		aiModelModel.Description = request.Description
		aiModelModel.InputFeatures = request.InputFeatures
		aiModelModel.OutputLabels = request.OutputLabels
		response.Created(c, aiModelModel)
	} else {
		response.Abort500(c, translator.TransHandler.T("r.cFailed"))
	}
}

func (ctrl *AiModelsController) Update(c *gin.Context) {

	aiModelModel := ai_model.Get(c.Param("id"))
	if aiModelModel.ID == 0 {
		response.Abort404(c)
		return
	}

	request := requests.AiModelRequest{}
	ok := requests.Validate(c, &request, requests.AiModelSave)
	if !ok {
		return
	}

	// aiModelModel.Name = request.Name
	// aiModelModel.Description = request.Description
	aiModelModel.Code = request.Code
	aiModelModel.Type = request.Type
	aiModelModel.Category = request.Category
	aiModelModel.Icon = request.Icon
	aiModelModel.Cover = request.Cover

	aiModelModel.Framework = request.Framework
	aiModelModel.Algorithm = request.Algorithm
	aiModelModel.NetStructure = request.NetStructure
	aiModelModel.BackBone = request.BackBone

	aiModelModel.Accuracy = request.Accuracy
	aiModelModel.Precision = request.Precision
	aiModelModel.Recall = request.Recall
	aiModelModel.F1Score = request.F1Score
	aiModelModel.AUC = request.AUC

	// aiModelModel.InputFeatures = request.InputFeatures
	// aiModelModel.OutputLabels = request.OutputLabels
	aiModelModel.InputParameters = request.InputParameters
	aiModelModel.ExecMethod = request.ExecMethod
	aiModelModel.OutputFormatter = request.OutputFormatter

	aiModelModel.Size = request.Size
	aiModelModel.Version = request.Version
	aiModelModel.Status = request.Status
	aiModelModel.DeployedAt = request.DeployedAt
	aiModelModel.RetiredAt = request.RetiredAt

	err := aiModelModel.SaveTx(&request)
	if err == nil {
		aiModelModel.Name = request.Name
		aiModelModel.Description = request.Description
		aiModelModel.InputFeatures = request.InputFeatures
		aiModelModel.OutputLabels = request.OutputLabels
		response.Data(c, aiModelModel)
	} else {
		response.Abort500(c, translator.TransHandler.T("r.uFailed"))
	}
}

func (ctrl *AiModelsController) Delete(c *gin.Context) {

	aiModelModel := ai_model.Get(c.Param("id"))
	if aiModelModel.ID == 0 {
		response.Abort404(c)
		return
	}

	rowsAffected := aiModelModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, translator.TransHandler.T("r.dFailed"))
}
