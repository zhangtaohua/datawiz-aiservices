package v1

import (
	"datawiz-aiservices/app/models/ai_project"
	"datawiz-aiservices/app/models/translation"
	"datawiz-aiservices/app/requests"
	"datawiz-aiservices/pkg/app"
	"datawiz-aiservices/pkg/response"
	"datawiz-aiservices/pkg/translator"

	"github.com/gin-gonic/gin"
)

type AiProjectsController struct {
	BaseAPIController
}

func getTransAiProject(aiProject *ai_project.AiProject) {
	namekey := aiProject.Name
	desckey := aiProject.Description
	tranV := translation.GetTs([]string{namekey, desckey}, app.Language)

	aiProject.Name = tranV[namekey]
	aiProject.Description = tranV[desckey]
}

func (ctrl *AiProjectsController) Index(c *gin.Context) {
	aiProjects := ai_project.All()
	length := len(aiProjects)
	for i := 0; i < length; i++ {
		getTransAiProject(&aiProjects[i])
	}
	response.Data(c, aiProjects)
}

func (ctrl *AiProjectsController) Show(c *gin.Context) {
	aiProjectModel := ai_project.Get(c.Param("id"))
	if aiProjectModel.ID == 0 {
		response.Abort404(c)
		return
	}
	getTransAiProject(&aiProjectModel)
	response.Data(c, aiProjectModel)
}

// func (ctrl *AiProjectsController) Store(c *gin.Context) {

// 	request := requests.AiProjectRequest{}
// 	if ok := requests.Validate(c, &request, requests.AiProjectSave); !ok {
// 		return
// 	}

// 	// 生成UUID作为 translated_id
// 	nameUUID := helpers.UUID()
// 	descUUID := helpers.UUID()

// 	aiProjectModel := ai_project.AiProject{
// 		Name:        nameUUID,
// 		Description: descUUID,
// 		UserID:      "rj-todo",
// 	}
// 	aiProjectModel.Create()

// 	nameTranslationModel := translation.Translation{
// 		TranslationId:  nameUUID,
// 		Language:       request.Language,
// 		TranslatedText: request.Name,
// 	}
// 	nameTranslationModel.Create()

// 	descTranslationModel := translation.Translation{
// 		TranslationId:  descUUID,
// 		Language:       request.Language,
// 		TranslatedText: request.Description,
// 	}
// 	descTranslationModel.Create()

// 	if aiProjectModel.ID > 0 && nameTranslationModel.ID > 0 && descTranslationModel.ID > 0 {
// 		aiProjectModel.Name = request.Name
// 		aiProjectModel.Description = request.Description
// 		response.Created(c, aiProjectModel)
// 	} else {
// 		response.Abort500(c, translator.TransHandler.T("r.cFailed"))
// 	}
// }

func (ctrl *AiProjectsController) Store(c *gin.Context) {

	request := requests.AiProjectRequest{}
	if ok := requests.Validate(c, &request, requests.AiProjectSave); !ok {
		return
	}

	aiProjectModel := ai_project.AiProject{
		Name:        "",
		Description: "",
		UserID:      "rj-todo",
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

	// 实际更新的应该是翻译表中
	nameKey := aiProjectModel.Name
	descKey := aiProjectModel.Description

	var nameRowsAffected int64 = 0
	var descRowsAffected int64 = 0

	// todo 这里应该学习如何用事物处理。
	nameTranslationModel := translation.GetByTidLang(nameKey, request.Language)
	if nameTranslationModel.ID == 0 {
		// 没有对应语言的翻译 就创建
		nameTranslationModel = translation.Translation{
			TranslationId:  nameKey,
			Language:       request.Language,
			TranslatedText: request.Name,
		}
		nameTranslationModel.Create()

		if nameTranslationModel.ID > 0 {
			nameRowsAffected = 1
		}
	} else {
		nameTranslationModel.TranslatedText = request.Name
		nameRowsAffected = nameTranslationModel.Save()
	}

	descTranslationModel := translation.GetByTidLang(descKey, request.Language)
	if descTranslationModel.ID == 0 {
		descTranslationModel = translation.Translation{
			TranslationId:  descKey,
			Language:       request.Language,
			TranslatedText: request.Description,
		}
		descTranslationModel.Create()
		if descTranslationModel.ID > 0 {
			descRowsAffected = 1
		}
	} else {
		descTranslationModel.TranslatedText = request.Description
		descRowsAffected = descTranslationModel.Save()
	}

	if nameRowsAffected > 0 && descRowsAffected > 0 {
		aiProjectModel.Name = request.Name
		aiProjectModel.Description = request.Description
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

	// todo 删除翻译表中 对应的翻译

	response.Abort500(c, translator.TransHandler.T("r.dFailed"))
}
