package Controller

import (
	"KlikDigital/models"
	"KlikDigital/service"
	"KlikDigital/share/response"
	"encoding/json"
	"net/http"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request){
	var err error

	var requestData models.TransactionRequestModel

	err = json.NewDecoder(r.Body).Decode(&requestData)

	if err != nil{
		response.SendErrorResponse(w, err.Error(), http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	data, err := service.TransactionService(requestData)
	if err != nil {
		if err.Error() == "silahkan login kembali"{
			response.SendErrorResponse(w, err.Error(), http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		response.SendErrorResponse(w, err.Error(), http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response.SendSuccessResponse(w, data, http.StatusText(http.StatusOK))
}