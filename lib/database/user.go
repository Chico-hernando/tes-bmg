package database

import (
	"bmg/configs"
	"bmg/models/user"
)

func GetUser() (dataresult []user.User, err error) {
	err = configs.DB.Find(&dataresult).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetUserByName(name string) (dataresult []user.UserResponse, err error) {
	err = configs.DB.Model(&user.User{}).Find(&dataresult, "name", name).Error
	if err != nil {
		return nil, err
	}
	return
}

func CheckUserByReffCode (reffCode string, data user.User) (status bool, err error) {
	var count int64
	err = configs.DB.Model(&data).Select("id").Where("refferal_code_input = ?", reffCode).Count(&count).Error
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, err
	}

	return false, err
}