package services

import (
	"advanced_programming/clients"
	"advanced_programming/common"
	"advanced_programming/constant"
	"advanced_programming/dal"
	"advanced_programming/models"
	. "advanced_programming/schema"
	"container/heap"
	"log"
)

func AddStarService(req *AddStarRequest) (resp *AddStarResponse, err error) {

	// 开启事务
	DB := clients.DB
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

	if ok := dal.CheckUser(req.UserID); !ok {
		return nil, common.NewError(constant.UserMissingError, "用户不存在")
	}

	// 先查询 blog，获取现在的 star 数
	filter := map[string]interface{}{
		"id": req.BlogID,
	}
	blog, err := dal.GetBlog(filter)
	if err != nil {
		log.Printf("%+v", err)
		tx.Rollback()
		return nil, err
	}

	// 再更新 blog
	updater := map[string]interface{}{
		"likes": blog.Likes + 1,
	}
	if err := dal.UpdateBlog(filter, updater); err != nil {
		log.Printf("%+v", err)
		tx.Rollback()
		return nil, err
	}

	// 落库 blog_liker
	liker := &models.BlogLiker{
		BlogID:   req.BlogID,
		AuthorID: req.UserID,
	}
	if err := dal.CreateBlogLiker(liker); err != nil {
		log.Printf("%+v", err)
		tx.Rollback()
		return nil, err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		log.Printf("%+v", err)
		tx.Rollback()
		return nil, err
	}

	// 返回响应
	resp = &AddStarResponse{
		UserID: req.UserID,
		BlogID: req.BlogID,
	}
	return resp, nil
}

func TopKStarService(k int) (resp *TopKStarResponse, err error) {

	// 先从数据库中读取所有的 blog 记录
	filter := map[string]interface{}{}
	blogs, err := dal.GetAllBlogs(filter)
	if err != nil {
		return nil, err
	}

	// 用小根堆获取点赞数最多的 k 个 blog
	hp := &models.Blogs{}
	heap.Init(hp)
	for _, i := range blogs {
		heap.Push(hp, i)
		if hp.Len() > k {
			heap.Pop(hp)
		}
	}

	// hp 中点赞数少的在堆顶，整理一下
	for i, j := 0, len(*hp)-1; i < j; i, j = i+1, j-1 {
		hp.Swap(i, j)
	}

	// 返回响应
	resp = &TopKStarResponse{
		Blogs: *hp,
	}
	return resp, nil
}
