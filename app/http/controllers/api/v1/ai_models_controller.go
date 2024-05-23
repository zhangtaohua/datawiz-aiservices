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
		Name:        request.Name,
		Description: request.Description,
		Type:        request.Type,

		Framework:    request.Framework,
		Algorithm:    request.Algorithm,
		NetStructure: request.NetStructure,
		BackBone:     request.BackBone,

		Accuracy:  request.Accuracy,
		Precision: request.Precision,
		Recall:    request.Recall,
		F1Score:   request.F1Score,
		AUC:       request.AUC,

		InputFeatures: request.InputFeatures,
		OutputLabels:  request.OutputLabels,
		ExecMethod:    request.ExecMethod,

		Size:       request.Size,
		Version:    request.Version,
		Status:     request.Status,
		DeployedAt: request.DeployedAt,
		RetiredAt:  request.RetiredAt,
	}
	aiModelModel.Create()
	if aiModelModel.ID > 0 {
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

	aiModelModel.Name = request.Name
	aiModelModel.Description = request.Description
	aiModelModel.Type = request.Type

	aiModelModel.Framework = request.Framework
	aiModelModel.Algorithm = request.Algorithm
	aiModelModel.NetStructure = request.NetStructure
	aiModelModel.BackBone = request.BackBone

	aiModelModel.Accuracy = request.Accuracy
	aiModelModel.Precision = request.Precision
	aiModelModel.Recall = request.Recall
	aiModelModel.F1Score = request.F1Score
	aiModelModel.AUC = request.AUC

	aiModelModel.InputFeatures = request.InputFeatures
	aiModelModel.OutputLabels = request.OutputLabels
	aiModelModel.ExecMethod = request.ExecMethod

	aiModelModel.Size = request.Size
	aiModelModel.Version = request.Version
	aiModelModel.Status = request.Status
	aiModelModel.DeployedAt = request.DeployedAt
	aiModelModel.RetiredAt = request.RetiredAt

	rowsAffected := aiModelModel.Save()
	if rowsAffected > 0 {
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
