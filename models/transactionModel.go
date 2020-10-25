package models

type TransactionRequestModel struct {
	Method string `json:"method"`
	UUID string `json:"uuid"`

}

type TransactionResponseModel struct {
	TransactionNumber string `json:"transaction_number"`
}
