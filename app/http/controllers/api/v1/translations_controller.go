package v1

import (
	"datawiz-aiservices/app/models/translation"
	"datawiz-aiservices/app/requests"
	"datawiz-aiservices/pkg/response"
	"datawiz-aiservices/pkg/translator"

	"github.com/gin-gonic/gin"
)

type TranslationsController struct {
	BaseAPIController
}

func (ctrl *TranslationsController) Index(c *gin.Context) {
	translations := translation.All()
	response.Data(c, translations)
}

func (ctrl *TranslationsController) Show(c *gin.Context) {
	translationModel := translation.Get(c.Param("id"))
	if translationModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, translationModel)
}

func (ctrl *TranslationsController) Store(c *gin.Context) {

	request := requests.TranslationRequest{}
	if ok := requests.Validate(c, &request, requests.TranslationSave); !ok {
		return
	}

	translationModel := translation.Translation{
		TranslationId:  request.TranslationId,
		Language:       request.Language,
		TranslatedText: request.TranslatedText,
	}
	translationModel.Create()
	if translationModel.ID > 0 {
		response.Created(c, translationModel)
	} else {
		response.Abort500(c, translator.TransHandler.T("r.cFailed"))
	}
}

func (ctrl *TranslationsController) Update(c *gin.Context) {

	translationModel := translation.Get(c.Param("id"))
	if translationModel.ID == 0 {
		response.Abort404(c)
		return
	}

	// if ok := policies.CanModifyTranslation(c, translationModel); !ok {
	// 	response.Abort403(c)
	// 	return
	// }

	request := requests.TranslationRequest{}
	ok := requests.Validate(c, &request, requests.TranslationSave)
	if !ok {
		return
	}

	translationModel.TranslationId = request.TranslationId
	translationModel.Language = request.Language
	translationModel.TranslatedText = request.TranslatedText

	rowsAffected := translationModel.Save()
	if rowsAffected > 0 {
		response.Data(c, translationModel)
	} else {
		response.Abort500(c, translator.TransHandler.T("r.uFailed"))
	}
}

func (ctrl *TranslationsController) Delete(c *gin.Context) {

	translationModel := translation.Get(c.Param("id"))
	if translationModel.ID == 0 {
		response.Abort404(c)
		return
	}

	// if ok := policies.CanModifyTranslation(c, translationModel); !ok {
	// 	response.Abort403(c)
	// 	return
	// }

	rowsAffected := translationModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, translator.TransHandler.T("r.dFailed"))
}
