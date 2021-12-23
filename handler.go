package main

import (
	"advanced_programming/handlers/blog"
	"advanced_programming/handlers/user"

	"github.com/gin-gonic/gin"
)

func UserRegister(r *gin.Engine) {
	userGroup := r.Group("/user")
	userGroup.POST("/register", user.RegisterHandler)
	/*
		userGroup.POST("/login", user.Login)
		userGroup.POST("/logout", user.Logout)
		userGroup.GET("/get_info", user.GetInfo)
		userGroup.POST("/upload_resume", user.UploadResume)
		userGroup.GET("/download_resume", user.DownloadResume)
	*/
}

func BlogRegister(r *gin.Engine) {
	BlogGroup := r.Group("/blog")
	BlogGroup.POST("/create_blog", blog.CreateBlogHandler)
	BlogGroup.GET("/show_blog_list", blog.ShowBlogListHandler)
	BlogGroup.GET("/show_blog", blog.ShowBlogHandler)
	BlogGroup.POST("/update_blog", blog.UpdateBlogHandler)
	BlogGroup.DELETE("/delete_blog", blog.DeleteBlogHandler)
}
