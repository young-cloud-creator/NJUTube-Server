package repository

import "gorm.io/gorm"

func QueryUserById(id int64) (*DBUser, error) {
	var user DBUser
	err := database.Model(&user).Where("id = ?", id).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func QueryUserByName(name string) (*DBUser, error) {
	var user DBUser
	err := database.Model(&user).Where("name = ?", name).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func AddUser(name string, passwd string) (*DBUser, error) {
	user := DBUser{Name: name, Password: passwd}
	err := database.Model(&user).Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
