package service

import (
	db "KlikDigital/config"
	"KlikDigital/models"
	"KlikDigital/repository"
	"KlikDigital/share/times"
	"errors"
	"fmt"
)

func TransactionService(requetData models.TransactionRequestModel)(data models.TransactionResponseModel, err error) {
	// cek user with uuid
	//
	userId, username, err := repository.GetDataUserByUUID(requetData.UUID)
	if err != nil {
		return
	}
	// generate unix time.now

	// cek sessio in redis
	response, err := db.GlobalCache.Do("GET", username)
	if err != nil || response == nil{
		err = errors.New("silahkan login kembali")
		return
	}

	unixTime := times.GenerateUnixTimeNow(times.TimeGmt)

	//generate transaction number
	transactionNumber := fmt.Sprintf("%s%d%s", requetData.Method,unixTime, requetData.UUID[len(requetData.UUID)- 6:] )

	//insert data into db
	err =repository.InsertTransactionData(transactionNumber, userId)
	if err != nil{
		return
	}

	data.TransactionNumber = transactionNumber
	return

}
