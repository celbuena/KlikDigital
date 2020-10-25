package repository

import (
	db "KlikDigital/config"
	"KlikDigital/models"
	"KlikDigital/share/times"
)

func IsExistUsername(username string )(isExist bool, err error)   {
	query := `SELECT EXISTS(SELECT 1 FROM users.user WHERE nama = $1)`
	err = db.GlobalDbSQL.QueryRow(query, username).Scan(&isExist)
	return

}
func InsertDatauser (requestData models.RegisterRequestModel, passwordEncrypt, uuid string) (err error) {
	now := times.Now(times.TimeGmt, times.TimeFormat)

	query := `INSERT INTO users.user(created_date, nama, email, phone_number, password, uuid) VALUES ($1, $2, $3, $4, $5 , $6)`
	_, err = db.GlobalDbSQL.Exec(query, now, requestData.Username, requestData.Email, requestData.PhoneNumber, passwordEncrypt, uuid)
	return
}

func GetUUIDAndPasswordByUsernameAndPassword(username string)(uuid, password string, err error){
	query := `SELECT uuid, password  FROM users.user WHERE nama = $1`
	err = db.GlobalDbSQL.QueryRow(query, username).Scan(&uuid, &password)
	return
}
