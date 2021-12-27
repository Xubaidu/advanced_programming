package main

import (
	"advanced_programming/handlers/blog"
	"advanced_programming/handlers/common"
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
	UserGroup.GET("/get_user_info", user.GetUserInfoHandler)
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
