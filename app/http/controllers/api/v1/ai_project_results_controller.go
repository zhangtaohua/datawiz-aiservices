package v1

import (
	"bytes"
	"datawiz-aiservices/app/models/ai_model"
	"datawiz-aiservices/app/models/ai_project_result"
	"datawiz-aiservices/app/requests"
	"datawiz-aiservices/pkg/config"
	"datawiz-aiservices/pkg/helpers"
	"datawiz-aiservices/pkg/response"
	"datawiz-aiservices/pkg/translator"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type AiProjectResultsController struct {
	BaseAPIController
}

func notifyAiProcess(result_uuid, ai_model_code string) bool {
	aiProcessBaseUrl := config.Get("app.ai_process_base_url")
	aiProcessUrl := aiProcessBaseUrl + "/ais/api/v1/process"
	fmt.Println("通知处理", aiProcessBaseUrl, aiProcessUrl)
	method := "POST"
	reqData := map[string]string{
		"uuid":          result_uuid,
		"ai_model_code": ai_model_code,
	}

	jsonData, err := json.Marshal(reqData)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return false
	}

	req, err := http.NewRequest(method, aiProcessUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return false
	}

	req.Header.Add("Content-Type", "application/json")

	// sent req
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return false
	}
	defer resp.Body.Close()

	// read Body
	if resp.StatusCode == http.StatusOK {
		// fmt.Println("Response Status:", resp.Status)
		// fmt.Println("Response Body:", string(body))

		// body, err := io.ReadAll(resp.Body)
		// if err != nil {
		//   fmt.Println("Error reading response:", err)
		//   return false
		// }

		// var result map[string]interface{}
		// err = json.Unmarshal(body, &result)
		// if err != nil {
		//   fmt.Println("Error parsing Body JSON response:", err)
		//   return false
		// }
		return true
	} else {
		return false
	}
}

func (ctrl *AiProjectResultsController) Index(c *gin.Context) {
	aiProjectResults := ai_project_result.All()
	response.Data(c, aiProjectResults)
}

func (ctrl *AiProjectResultsController) Show(c *gin.Context) {
	id := c.Param("id")
	if helpers.IsUUID(id) {
		aiProjectResultModels := ai_project_result.GetByUUID(id)
		response.Data(c, aiProjectResultModels)
	} else {
		aiProjectResultModel := ai_project_result.Get(id, false)
		if aiProjectResultModel.ID == 0 {
			response.Abort404(c)
			return
		}
		response.Data(c, aiProjectResultModel)
	}
}

func (ctrl *AiProjectResultsController) Store(c *gin.Context) {

	request := requests.AiProjectResultRequest{}
	if ok := requests.Validate(c, &request, requests.AiProjectResultSave); !ok {
		return
	}

	aiModel := ai_model.GetBy("uuid", string(request.AiModelUUID))
	if aiModel.ID == 0 {
		response.Abort404(c)
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

	// authHeader := c.Request.Header.Get("Authorization")
	// if authHeader == "" {
	// 	return "", ErrHeaderEmpty
	// }

	err := aiProjectResultModel.CreateTx(&request)

	if err == nil {
		aiProjectResultModel.Name = request.Name
		aiProjectResultModel.Description = request.Description
		// 发送请求到Python 进行AI处理
		notifyAiProcess(cast.ToString(aiProjectResultModel.UUID), string(aiModel.Code))
		response.Created(c, aiProjectResultModel)
	} else {
		response.Abort500(c, translator.TransHandler.T("r.cFailed"))
	}
}

func (ctrl *AiProjectResultsController) Update(c *gin.Context) {

	aiProjectResultModel := ai_project_result.Get(c.Param("id"), true)
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

func (ctrl *AiProjectResultsController) Patch(c *gin.Context) {

	aiProjectResultModel := ai_project_result.Get(c.Param("id"), true)
	if aiProjectResultModel.ID == 0 {
		response.Abort404(c)
		return
	}

	request := requests.AiProjectResultRequest{}
	ok := requests.Validate(c, &request, requests.AiProjectResultSave)
	if !ok {
		return
	}

	if !helpers.Empty(request.Input) {
		aiProjectResultModel.Input = request.Input
	}

	if !helpers.Empty(request.Output) {
		aiProjectResultModel.Output = request.Output
	}

	if !helpers.Empty(request.Status) {
		aiProjectResultModel.Status = request.Status
	}

	// aiProjectResultModel.UserID = request.UserID
	// aiProjectResultModel.AiModelUUID = request.AiModelUUID
	// aiProjectResultModel.AiProjectUUID = request.AiProjectUUID

	err := aiProjectResultModel.SaveTx(&request, true)
	if err == nil {
		if !helpers.Empty(request.Name) {
			aiProjectResultModel.Name = request.Name
		}
		if !helpers.Empty(request.Description) {
			aiProjectResultModel.Description = request.Description
		}

		response.Data(c, aiProjectResultModel)
	} else {
		response.Abort500(c, translator.TransHandler.T("r.uFailed"))
	}
}

func (ctrl *AiProjectResultsController) Restart(c *gin.Context) {

	aiProjectResultModel := ai_project_result.Get(c.Param("id"), true)
	if aiProjectResultModel.ID == 0 {
		response.Abort404(c)
		return
	} else {
		aiModel := ai_model.GetBy("uuid", string(aiProjectResultModel.AiModelUUID))
		if aiModel.ID == 0 {
			response.Abort404(c)
			return
		} else {
			notifyAiProcess(cast.ToString(aiProjectResultModel.UUID), string(aiModel.Code))
			response.Success(c)
		}
	}
}

func (ctrl *AiProjectResultsController) Delete(c *gin.Context) {

	aiProjectResultModel := ai_project_result.Get(c.Param("id"), true)
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
