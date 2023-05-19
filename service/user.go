package service

import (
	"crypto/sha512"
	"encoding/hex"
	"nju-tube/repository"
	"nju-tube/structs"
)

func QueryUserInfo(id int64) (*structs.User, error) {
	dbUser, err := repository.QueryUserById(id)
	if err != nil {
		return nil, err
	}
	if dbUser == nil {
		return nil, nil
	}
	return &structs.User{Id: id, Name: dbUser.Name}, nil
}

func encrypt(rawString string) string {
	secret := sha512.Sum512([]byte(rawString))
	return hex.EncodeToString(secret[:])
}

// UserLogin returns (isSuccessful, msg)
func UserLogin(name string, passwd string) (bool, int64, string) {
	secret := encrypt(passwd)
	user, err := repository.QueryUserByName(name)
	if err != nil {
		return false, -1, "Unknown Error"
	}
	if user == nil {
		return false, -1, "User Not Find"
	}
	if user.Password != secret {
		return false, -1, "Wrong Password"
	}
	return true, user.Id, ""
}

// UserRegister returns (isSuccessful, userId, msg)
func UserRegister(name string, passwd string) (bool, int64, string) {
	user, err := repository.QueryUserByName(name)
	if err != nil {
		return false, -1, "Unknown Error"
	}
	if user != nil {
		return false, -1, "Username Exists"
	}
	secret := encrypt(passwd)
	newUser, err := repository.AddUser(name, secret)
	if err != nil {
		return false, -1, "Unknown Error"
	}
	return true, newUser.Id, ""
}
