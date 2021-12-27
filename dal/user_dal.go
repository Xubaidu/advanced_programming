package dal

import (
	"advanced_programming/clients"
	"advanced_programming/common"
	"advanced_programming/constant"
	"advanced_programming/models"
	"log"
)

func CreateUser(user *models.User) (err error) {

	var DB = clients.DB

	// 创建记录
	if err := DB.Create(user).Error; err != nil {
		log.Printf("%+v", err)
		return common.NewError(constant.RecordInsertError, "创建记录失败")
	}

	return nil
}

func GetUsers(filter map[string]interface{}) (users []*models.User, err error) {
	var DB = clients.DB

	// 查询记录
	if err := DB.Where(filter).Find(&users).Error; err != nil {
		log.Printf("%+v", err)
		return nil, common.NewError(constant.RecordNotFound, "没有查询到相关记录")
	}

	// 异常处理
	if len(users) == 0 {
		log.Printf("没有符合条件的查询")
		return nil, common.NewError(constant.DataQueryError, "没有符合条件的查询")
	}

	return users, nil
}

func GetUser(filter map[string]interface{}) (user *models.User, err error) {
	var DB = clients.DB

	user = &models.User{}

	// 查询记录，查不到会返回 ErrRecordNotFound 错误
	if err := DB.Where(filter).Take(user).Error; err != nil {
		log.Printf("%+v", err)
		return nil, common.NewError(constant.RecordNotFound, "没有查询到相关记录")
	}

	return user, nil
}

func GetUserResume(filter map[string]interface{}) (resume string, err error) {
	user, err := GetUser(filter)
	if err != nil {
		return "", err
	}

	if user.Resume == "" {
		return "", common.NewError(constant.DataQueryError, "没有查询到相关记录")
	}

	return user.Resume, nil
}

func UpdateUser(filter, updater map[string]interface{}) (err error) {

	var DB = clients.DB

	// 存在才更新
	if _, err = GetUser(filter); err != nil {
		log.Printf("%+v", err)
		return common.NewError(constant.RecordNotFound, "没有查询到相关记录")
	}

	// 更新
	for k, v := range updater {
		if err := DB.Model(&models.User{}).Where(filter).Update(k, v).Error; err != nil {
			log.Printf("%+v", err)
			return common.NewError(constant.RecordUpdateError, "更新记录失败")
		}
	}

	// 正常返回
	return nil
}

func DeleteUser(filter map[string]interface{}) (err error) {
	var DB = clients.DB

	// 删除
	if err := DB.Where(filter).Delete(&models.User{}).Error; err != nil {
		log.Printf("%+v", err)
		return common.NewError(constant.RecordDeleteError, "删除记录失败")
	}

	return nil
}

func CheckUser(UserID int) (ok bool) {
	filter := map[string]interface{}{
		"id": UserID,
	}
	if _, err := GetUser(filter); err != nil {
		return false
	}
	return true
}
