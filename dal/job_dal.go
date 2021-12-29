package dal

import (
	"advanced_programming/clients"
	"advanced_programming/common"
	"advanced_programming/constant"
	"advanced_programming/models"
	"log"
)

func CreateJob(job *models.Job) (err error) {
	var DB = clients.DB

	// 创建记录
	if err := DB.Create(job).Error; err != nil {
		log.Printf("%+v", err)
		return common.NewError(constant.RecordInsertError, "创建记录失败")
	}

	return nil
}

func GetJobs(limit int, filter map[string]interface{}) (jobs []*models.Job, err error) {
	var DB = clients.DB

	// 查询记录
	if err := DB.Limit(limit).Where(filter).Find(&jobs).Error; err != nil {
		log.Printf("%+v", err)
		return nil, common.NewError(constant.RecordNotFound, "没有查询到相关记录")
	}

	// 异常处理
	if len(jobs) == 0 {
		log.Printf("没有符合条件的查询")
		return nil, common.NewError(constant.DataQueryError, "没有符合条件的查询")
	}

	return jobs, nil
}

func GetJob(filter map[string]interface{}) (job *models.Job, err error) {

	var DB = clients.DB

	// 查询记录，查不到会返回 ErrRecordNotFound 错误
	job = &models.Job{}
	if err := DB.Where(filter).Take(job).Error; err != nil {
		log.Printf("%+v", err)
		return nil, common.NewError(constant.RecordNotFound, "没有查询到相关记录")
	}

	return job, nil
}

func UpdateJob(filter, updater map[string]interface{}) (err error) {
	var DB = clients.DB

	// 判断是否存在
	if _, err = GetJob(filter); err != nil {
		log.Printf("%+v", err)
		return common.NewError(constant.RecordNotFound, "没有查询到相关记录")
	}

	// 更新
	for k, v := range updater {
		if err := DB.Model(&models.Job{}).Where(filter).Update(k, v).Error; err != nil {
			log.Printf("%+v", err)
			return common.NewError(constant.RecordUpdateError, "更新记录失败")
		}
	}

	// 正常返回
	return nil
}

func DeleteJob(filter map[string]interface{}) (err error) {
	var DB = clients.DB

	// 删除
	if err := DB.Where(filter).Delete(&models.Job{}).Error; err != nil {
		log.Printf("%+v", err)
		return common.NewError(constant.RecordDeleteError, "删除记录失败")
	}

	return nil
}
