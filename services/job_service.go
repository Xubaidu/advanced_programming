package services

import (
	"advanced_programming/clients"
	"advanced_programming/common"
	"advanced_programming/constant"
	"advanced_programming/dal"
	"advanced_programming/models"
	. "advanced_programming/schema"
	"fmt"
	"log"
)

func CampaignJobService(req *CampaignJobRequest) (resp *CampaignJobResponse, err error) {

	// 构造 job 对象
	job := &models.Job{
		UserID:  req.UserID,
		Company: req.Company,
		Base:    req.Base,
		Title:   req.Title,
		Content: req.Content,
	}

	// 创建 job
	if err := dal.CreateJob(job); err != nil {
		return nil, err
	}

	// 返回响应
	resp = &CampaignJobResponse{
		Job: job,
		Msg: "成功发布职位！",
	}
	return resp, nil
}

func ShowJobService(req *GetJobRequest) (resp *GetJobResponse, err error) {

	// 构造过滤器
	filter := map[string]interface{}{
		"id": req.JobID,
	}

	// 查询记录
	job, err := dal.GetJob(filter)
	if err != nil {
		return nil, err
	}

	// 根据 user_id 获取 user_name
	filter = map[string]interface{}{
		"id": job.UserID,
	}
	userName, err := dal.GetUserName(filter)
	if err != nil {
		return nil, err
	}

	// 查询成功后返回响应
	msg := fmt.Sprintf("成功展示职位 %d", job.JobID)
	resp = &GetJobResponse{
		UserName: userName,
		Job:      job,
		Msg:      msg,
	}
	return resp, nil
}

func ShowJobListService(req *GetJobListRequest) (resp *GetJobListResponse, err error) {

	// 查询记录
	jobs, err := dal.GetJobs(req.Limit, nil)
	if err != nil {
		return nil, err
	}

	// 根据 blog 维护的 user_id
	// 获取每一个 blog 的作者名
	var userNames []string
	for _, job := range jobs {
		filter := map[string]interface{}{
			"id": job.UserID,
		}
		if userName, err := dal.GetUserName(filter); err != nil {
			userNames = append(userNames, userName)
		}
	}

	// 返回响应
	msg := fmt.Sprintf("成功展示 %d 条职位", req.Limit)
	resp = &GetJobListResponse{
		UserNames: userNames,
		Jobs:      jobs,
		Msg:       msg,
	}
	return resp, nil
}

func UpdateJobService(req *UpdateJobRequest) (resp *UpdateJobResponse, err error) {

	// 开启事务
	var DB = clients.DB
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 判断事务是否开启
	if err := tx.Error; err != nil {
		log.Printf("%+v", err)
		tx.Rollback()
		return nil, common.NewError(constant.TransactionBeginError, "开启事务失败")
	}

	// 更新 blog
	filter := map[string]interface{}{
		"id": req.JobID,
	}
	updater := map[string]interface{}{
		"title":   req.Title,
		"content": req.Content,
	}
	if err := dal.UpdateJob(filter, updater); err != nil {
		tx.Rollback()
		return nil, err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		log.Printf("%+v", err)
		tx.Rollback()
		return nil, err
	}

	// 构造响应并返回
	msg := fmt.Sprintf("成功更新职位 %d", req.JobID)
	resp = &UpdateJobResponse{
		Msg: msg,
	}
	return resp, nil
}

func DeleteJobService(req *DeleteJobRequest) (resp *DeleteJobResponse, err error) {
	filter := map[string]interface{}{
		"id": req.JobID,
	}
	err = dal.DeleteJob(filter)
	if err != nil {
		return nil, err
	}
	msg := fmt.Sprintf("成功删除职位 %d", req.JobID)
	resp = &DeleteJobResponse{
		Msg: msg,
	}
	return resp, nil
}

func ApplyService(req *ApplyRequest) (resp *ApplyResponse, err error) {

	// 开启事务
	var DB = clients.DB
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 判断事务是否开启
	if err := tx.Error; err != nil {
		log.Printf("%+v", err)
		tx.Rollback()
		return nil, common.NewError(constant.TransactionBeginError, "开启事务失败")
	}

	// 校验岗位信息
	filter := map[string]interface{}{
		"id": req.JobID,
	}
	job, err := dal.GetJob(filter)
	if err != nil {
		tx.Rollback()
		code := err.(*common.CustomError).GetCode()
		return nil, common.NewError(code, "不存在的岗位！")
	}

	// 校验用户信息
	filter = map[string]interface{}{
		"id": req.UserID,
	}
	if _, err = dal.GetUserResume(filter); err != nil {
		tx.Rollback()
		code := err.(*common.CustomError).GetCode()
		return nil, common.NewError(code, "你还没有上传自己的简历！")
	}

	// 校验用户是否已投递
	filter = map[string]interface{}{
		"user_id": req.UserID,
		"job_id":  req.JobID,
	}
	if ok := dal.IsApply(filter); ok {
		tx.Rollback()
		return nil, common.NewError(constant.DuplicateApplyError, "禁止重复投递！")
	}

	// 更新 job 表
	filter = map[string]interface{}{
		"id": req.JobID,
	}
	updater := map[string]interface{}{
		"applicants": job.Applicants + 1,
	}
	if err = dal.UpdateJob(filter, updater); err != nil {
		tx.Rollback()
		return nil, err
	}

	// 给 apply 表插入记录
	newApply := &models.Apply{
		UserID: req.UserID,
		JobID:  req.JobID,
	}
	if err = dal.CreateApply(newApply); err != nil {
		tx.Rollback()
		return nil, err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		log.Printf("%+v", err)
		tx.Rollback()
		return nil, err
	}

	// 构造响应
	msg := fmt.Sprintf("成功投递岗位：%s (%s)，静候佳音！", job.Title, job.Base)
	resp = &ApplyResponse{
		Msg: msg,
	}
	return resp, nil
}
