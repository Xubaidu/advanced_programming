package dal

import (
	"advanced_programming/clients"
	"advanced_programming/common"
	"advanced_programming/constant"
	"advanced_programming/models"
	"log"
)

func CreateApply(apply *models.Apply) (err error) {
	var DB = clients.DB

	// 创建记录
	if err := DB.Create(apply).Error; err != nil {
		log.Printf("%+v", err)
		return common.NewError(constant.RecordInsertError, "创建记录失败")
	}

	return nil
}

func GetApply(filter map[string]interface{}) (apply *models.Apply, err error) {
	var DB = clients.DB

	// 查询记录，查不到会返回 ErrRecordNotFound 错误
	apply = &models.Apply{}
	if err := DB.Where(filter).Take(apply).Error; err != nil {
		log.Printf("%+v", err)
		return nil, common.NewError(constant.RecordNotFound, "没有查询到相关记录")
	}

	return apply, nil
}

func IsApply(filter map[string]interface{}) (ok bool) {
	_, err := GetApply(filter)
	return err == nil
}

func GetUserApply(applierID int) (jobs []*models.Job, err error) {
	var DB = clients.DB
	DB.Table("job").Joins("join apply on job.id = apply.job_id").
		Where("apply.user_id = ?", applierID).Find(&jobs)
	if err = DB.Error; err != nil {
		return nil, err
	}
	return jobs, nil
}
