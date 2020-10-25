package service

import (
	db "KlikDigital/config"
	"KlikDigital/models"
	"KlikDigital/repository"
	"KlikDigital/utils"
	"errors"
	guuid "github.com/google/uuid"
)

func RegisterService(requstData models.RegisterRequestModel) (data models.RegisterResponseModel, err error) {
// cek username apakah exist or not
	isExist, err := repository.IsExistUsername(requstData.Username)
	if err != nil{
		return
	}

	if isExist {
		err = errors.New("usernama sudah terdaftar")
		return
	}
   // encrypt pasword
   passwordEncrypt, err := utils.HashPassword(requstData.Password)
   if err != nil{
   	return
   }

   // Generate uuid
   uuidString := generateUUID()


   //insert user data into database
   err = repository.InsertDatauser(requstData, passwordEncrypt, uuidString)
   if err != nil{
   	return
   }
   data.Username = requstData.Username
   data.Email = requstData.Email
   data.PhoneNumber = requstData.PhoneNumber
   data.UUID = uuidString

   return
}

func generateUUID() string{
	id := guuid.New()

	return id.String()
}

func LoginService(requestData models.LoginRequestModel) (data models.LoginResponseModel, err error){
	uuid, passwordEncrypt, err := repository.GetUUIDAndPasswordByUsernameAndPassword(requestData.Username)
	if err != nil || uuid == ""{
		err = errors.New("username atau password salah")
		return
	}
	// check password
	isPasswordMatch := utils.CheckPasswordHash(requestData.Password, passwordEncrypt)
	if !isPasswordMatch {
		err = errors.New("username atau password salah")
		return
	}

 	// save into redis
	response, err := saveSessionIntoRedis(uuid, requestData.Username)
	if err != nil || response != "OK"{
		return
	}

	data.UUID = uuid

	return
}

func saveSessionIntoRedis(uuid string, username string) (response interface{}, err error)  {
	response, err = db.GlobalCache.Do("SETEX", username,"3600" ,uuid)

	return
}