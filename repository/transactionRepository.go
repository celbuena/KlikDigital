package repository

import (
	db "KlikDigital/config"
	"KlikDigital/share/times"
)

func GetDataUserByUUID(uuid string)(userId int, username string, err error){
	query := `SELECT id, nama FROM users.user WHERE uuid = $1`

	err = db.GlobalDbSQL.QueryRow(query, uuid).Scan(&userId, &username)

	return
}

func InsertTransactionData(TransactionNumber string, userId int)(err error){
	now := times.Now(times.TimeGmt, times.TimeFormat)
	query := `INSERT INTO transactions.transactions(user_id, nomor_transaction, created_date) VALUES ($1, $2, $3)`

	_, err = db.GlobalDbSQL.Exec(query, userId, TransactionNumber, now)

	return
}
