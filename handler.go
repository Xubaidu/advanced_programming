package main

import (
	"advanced_programming/handlers/blog"

	"github.com/gin-gonic/gin"
)

/*
func UserRegister(r *gin.Engine) {
	userGroup := r.Group("/user")
	userGroup.POST("/user/register", user.Register)
	userGroup.POST("/user/login", user.Login)
	userGroup.POST("/user/logout", user.Logout)
	userGroup.GET("/user/get_info", user.GetInfo)
	userGroup.POST("/user/upload_resume", user.UploadResume)
	userGroup.GET("/user/download_resume", user.DownloadResume)
}
*/

func BlogRegister(r *gin.Engine) {
	BlogGroup := r.Group("/blog")
	BlogGroup.GET("blog/show_blog_list", blog.ShowBlogList)
	BlogGroup.GET("blog/show_blog", blog.ShowBlog)
	BlogGroup.POST("blog/create_blog", blog.CreateBlog)
	BlogGroup.POST("blog/update_blog", blog.UpdateBlog)
	BlogGroup.DELETE("blog/delete_blog", blog.DeleteBlog)
}