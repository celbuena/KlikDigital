package Controller

import (
	"KlikDigital/models"
	"KlikDigital/service"
	"KlikDigital/share/response"
	"encoding/json"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request)  {
	var err error
	var requestData models.RegisterRequestModel
	err = json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = requestData.Validate()
	if err != nil{
		response.SendErrorResponse(w, err.Error(), http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// call register service
 	responsData, err := service.RegisterService(requestData)
 	if err != nil{
		response.SendErrorResponse(w, err.Error(), http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	response.SendSuccessResponse(w, responsData, http.StatusText(http.StatusOK))
}

func Login(w http.ResponseWriter, r *http.Request){
	var err error
	var requestData models.LoginRequestModel

	err = json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil{
		response.SendErrorResponse(w, err.Error(), http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	data, err := service.LoginService(requestData)
	if err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return

	}
	response.SendSuccessResponse(w, data, http.StatusText(http.StatusOK))


}


