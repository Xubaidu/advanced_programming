package main

import (
	"advanced_programming/handlers/blog"
	"advanced_programming/handlers/common"
	"advanced_programming/handlers/job"
	"advanced_programming/handlers/user"

	"github.com/gin-gonic/gin"
)

func IndexRegister(r *gin.Engine) {
	IndexGroup := r.Group("/index")
	IndexGroup.POST("/register", common.RegisterHandler)
	IndexGroup.POST("/login", common.LoginHandler)
}

func UserRegister(r *gin.Engine) {
	UserGroup := r.Group("/user")
	UserGroup.POST("/upload_resume", user.UploadResumeHandler)
	UserGroup.GET("/download_resume", user.DownloadResumeHandler)
	UserGroup.GET("/profile", user.GetUserInfoHandler)
	UserGroup.GET("/my_applications", user.GetUserApplyHandler)
}

func BlogRegister(r *gin.Engine) {
	BlogGroup := r.Group("/blog")
	BlogGroup.POST("/create_blog", blog.CreateBlogHandler)
	BlogGroup.GET("/show_blog_list", blog.ShowBlogListHandler)
	BlogGroup.GET("/show_blog", blog.ShowBlogHandler)
	BlogGroup.POST("/update_blog", blog.UpdateBlogHandler)
	BlogGroup.DELETE("/delete_blog", blog.DeleteBlogHandler)
	BlogGroup.GET("/top_k", blog.TopKBlogHandler)
	BlogGroup.POST("/add_star", blog.AddStarHandler)
}

func JobRegister(r *gin.Engine) {
	JobGroup := r.Group("/job")
	JobGroup.POST("/campaign_job", job.CampaignJobHandler)
	JobGroup.GET("/show_job_list", job.ShowJobListHandler)
	JobGroup.GET("/show_job", job.ShowJobHandler)
	JobGroup.POST("/update_job", job.UpdateJobHandler)
	JobGroup.DELETE("/delete_job", job.DeleteJobHandler)
	JobGroup.POST("/apply", job.ApplyHandler)
}
