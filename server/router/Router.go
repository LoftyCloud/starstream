/*
 * 管理路由
 */
package router

import (
	"VideoHubGo/controllers/AdminController"
	"VideoHubGo/controllers/ClassController"
	"VideoHubGo/controllers/PhotoController"
	"VideoHubGo/controllers/RelationController"
	"VideoHubGo/controllers/UserController"
	"VideoHubGo/controllers/VideoController"
	"VideoHubGo/controllers/WatchController"
	"VideoHubGo/middlewares/CorsMiddleware"
	"VideoHubGo/middlewares/JwtMiddleware"
	"VideoHubGo/middlewares/NoRouteMiddleware"
	"VideoHubGo/utils/LogUtils"
	"VideoHubGo/utils/UploadUtils"
	"os"

	"net/http"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

/**
 * @Descripttion: 路由管理 - Router Manager
 */
func Router(router *gin.Engine) *gin.Engine {
	r := gin.Default()

	// CORS 中间件 - 配置 CORS
	r.Use(func(c *gin.Context) {
		// 设置 CORS 头部
		c.Header("Access-Control-Allow-Origin", "*")                                            // 允许所有来源（如果只允许某个前端，替换 "*" 为前端地址）
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")             // 支持的方法
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization") // 允许的请求头

		// 允许浏览器发送OPTIONS预检请求并返回204
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		// 继续执行下一个中间件或处理请求
		c.Next()
	})

	//读取配置文件 - Read The Configuration File
	path, err := os.Getwd()
	if err != nil {
		LogUtils.Logger(err.Error())
	}
	config := viper.New()
	config.AddConfigPath(path + "/configs")
	config.SetConfigName("config")
	config.SetConfigType("yaml")

	//尝试进行配置读取 - Try Reading Configuration
	if err := config.ReadInConfig(); err != nil {
		LogUtils.Logger(err.Error())
	}

	router.NoRoute(NoRouteMiddleware.NoRouteHttp)   //设置网页错误返回信息 - Setting Http Error Return Message
	router.NoMethod(NoRouteMiddleware.NoMethodHttp) //设置网页错误返回信息 - Setting Http Error Return Message

	maxSize := config.GetInt("files.maxSize") //设置最大文件上传大小 - Set Max Upload File Size

	router.MaxMultipartMemory = int64(maxSize << 20)

	router.Use(CorsMiddleware.Cors())

	//用户路由 - User Router
	routerList1 := router.Group("/user")
	{
		routerList1.POST("/login", UserController.UserLogin)         //用户登录控制器 - User Login Controller
		routerList1.POST("/register", UserController.UserRegister)   //用户注册控制器 - User Register Controller
		routerList1.POST("/login/admin", AdminController.AdminLogin) //管理员用户登录控制器 - Admin User Login Controller
	}

	//用户信息路由 - User Information Route，需要进行jwt验证，另外作为一个分组
	routerList2 := router.Group("/userinfo").Use(JwtMiddleware.JwtMiddleware()) // JWT中间件 - JWT Middleware
	{
		routerList2.POST("/update/password", UserController.UserUpdatePassword) //用户修改密码控制器 - User Update Password Controller
		routerList2.POST("/upload/avatar", UserController.UploadAvatar)         //用户头像上传控制器 - User Upload Avatar Controller
	}

	//视频分类路由 - Video Class Route
	routerList3 := router.Group("/class").Use(JwtMiddleware.JwtMiddleware()) // JWT中间件 - JWT Middleware
	{
		routerList3.POST("/list", ClassController.GetClassList) //视频控制器 - Video Controller
	}

	//视频路由 - Video Route
	routerList4 := router.Group("/video").Use(JwtMiddleware.JwtMiddleware()) // JWT中间件 - JWT Middleware
	{
		routerList4.POST("/list", VideoController.GetVideoList)                //总视频控制器 - Center Video Controller
		routerList4.POST("/class/list", VideoController.GetVideoClassList)     //视频类型控制器 - Video Class Controller
		routerList4.POST("/search/list", VideoController.SearchVideoClassList) //视频搜索控制器 - Video Search Controller
		routerList4.POST("/upload", VideoController.UploadVideo_StreamFile)    //视频上传 - Video Upload
	}

	//用户收藏路由 - Relation Route
	routerList5 := router.Group("/relation").Use(JwtMiddleware.JwtMiddleware()) // JWT中间件 - JWT Middleware
	{
		routerList5.POST("/list", RelationController.GetRelationList)                //总收藏控制器 - Center Relation Controller
		routerList5.POST("/class/list", RelationController.FindRelationClassList)    //搜藏类型控制器 - Relation Class Controller
		routerList5.POST("/search/list", RelationController.SearchRelationClassList) //收藏搜索控制器 - Relation Search Controller
		routerList5.POST("/add", RelationController.RelationVideo)                   //添加用户收藏 - Add User Relation
		routerList5.POST("/delete", RelationController.RemoveRelation)               //取消用户收藏 - Delete User Relation
		routerList5.POST("/is", RelationController.IsRelation)                       //检查用户是否已经收藏 - Check User Is Already Relation
	}

	//视频详细路由 - Video Detail Route
	routerList6 := router.Group("/watch").Use(JwtMiddleware.JwtMiddleware()) // JWT中间件 - JWT Middleware
	{
		routerList6.POST("find", WatchController.GetVideoDetail) //获取视频详细信息 - Query video details
		routerList6.POST("plus", WatchController.PlusVideoWatch) //增加视频流量 - Increase video traffic
	}

	//文件映射 - Map File
	routerList7 := router.Group("/file")
	{
		routerList7.Static("/avatar", UploadUtils.GetFilePath("user.userAvatar")) //映射头像文件夹 - Map avatar folder
		routerList7.Static("/video", UploadUtils.GetFilePath("video.saveFile"))   //映射视频文件夹 - Map video folder
		routerList7.Static("/cover", UploadUtils.GetFilePath("video.coverFile"))  //映射视频封面文件夹 - Map video cover folder
		routerList7.Static("/photo", UploadUtils.GetFilePath("photo.file"))       //映射图片文件夹 - Map photo folder
		routerList7.GET("/checkcover", VideoController.CheckCoverExists)          //检查视频封面是否存在 - Check if the video cover exists
	}

	//后台路由 - Admin Route
	routerList8 := router.Group("/admin").Use(JwtMiddleware.AdminJwtMiddleware()) // 后台JWT中间件 - Admin JWT Middleware
	{
		routerList8.POST("/edit/video/information", AdminController.EditVideoInformation)           //修改视频信息 - Update video information
		routerList8.POST("/edit/group/video/information", AdminController.BatchVideoCidInformation) //批量修改视频信息 - Batch update video information
		routerList8.POST("/delete/video/information", AdminController.DeleteVideoInformation)       //删除视频信息 - Delete video information
		routerList8.POST("/delete/group/video/information", AdminController.DeleteVideoGroup)       //批量删除视频信息 - Batch delete video information

		routerList8.POST("/edit/user/information", AdminController.EditUserInformation)         //修改用户信息 - Update user information
		routerList8.POST("/delete/group/user/information", AdminController.DeleteUserGroup)     //批量删除用户信息 - Batch delete user information
		routerList8.POST("/delete/user/information", AdminController.DeleteUserInformation)     //删除用户信息 - Delete user information
		routerList8.POST("/authority/user/isadmin", AdminController.AuthorityUserIsadmin)       //修改用户管理权限 - Update user Isadmin
		routerList8.POST("/authority/user/isuploader", AdminController.AuthorityUserIsuploader) //修改用户上传权限 - Update user Isuploader
		routerList8.POST("/upload/user/avatar", AdminController.UploadAvatar)                   //修改用户头像 - Update user avatar

		routerList8.POST("/edit/class/information", AdminController.EditClassInformation)     //修改视频分类信息 - Delete video class information
		routerList8.POST("/delete/class/information", AdminController.DeleteClassInformation) //删除视频分类信息 - Delete video class information
		routerList8.POST("/delete/group/class/information", AdminController.DeleteClassGroup) //批量删除分类信息 - Batch delete class information

		routerList8.POST("/delete/group/photo/information", AdminController.DeletePhotoGroup) //批量删除图片信息 - Batch delete photo information
		routerList8.POST("/delete/photo/information", AdminController.DeletePhotoInformation) //删除图片信息 - Delete photo information

		//数据类操作
		routerList8.POST("/search/video/list", AdminController.GetVideoList)            //获取后台视频信息 - Get admin video information
		routerList8.POST("/search/video/nocid/list", AdminController.GetNoCidVideoList) //搜索未分类的视频 - Seach no cid video data
		routerList8.POST("/search/class/list", AdminController.GetClassList)            //获取后台分类信息 - Get admin class information
		routerList8.POST("/search/user/list", AdminController.GetUserList)              //获取后台用户信息 - Get admin user information

		//服务器信息操作
		routerList8.POST("/get/strong", AdminController.GetDiskStrong)              //获取后台存储空间信息 - Get back-end strong information
		routerList8.POST("/get/recode", AdminController.GetReCodeStatus)            //获取后台二次编码信息 - Get back-end recode information
		routerList8.POST("/change/recode", AdminController.ChangeReCodeStatus)      //修改后台二次编码信息 - Change back-end recode information
		routerList8.POST("/get/temp/file/list", AdminController.GetTempFileList)    //获取后台临时文件列表 - Get back-end temp file list
		routerList8.POST("/get/waste/file/list", AdminController.GetWasteFileList)  //获取后台回收站文件列表 - Get back-end recycle bin file list
		routerList8.POST("/upload/file", AdminController.UploadFileStream)          //后台临时文件列上传 - Back-end upload temp file
		routerList8.POST("/delete/temp/file", AdminController.DeleteTempFile)       //删除临时文件的文件到回收站 - Delete temp file to recycle bin
		routerList8.POST("/delete/waste/file", AdminController.DeleteWasteFile)     //删除回收站文件 - Delete recycle bin file
		routerList8.POST("/recovery/waste/file", AdminController.RecoveryWasteFile) //恢复回收站文件 - Recovery recycle bin file
		routerList8.POST("/save/video", AdminController.SaveVideoToLocal)           //后台存储视频 - Back-end save video
		routerList8.POST("/save/photo", AdminController.SavePhotoToLocal)           //后台存储图片 - Back-end save photo
		routerList8.POST("/created/class", AdminController.CreatedClassInformarion) //创建新视频分类 - Created new video class
		routerList8.POST("/clean/login/logs", AdminController.CleanLoginLogs)       //清空后台登录日志 - Clean back-end login logs
		routerList8.POST("/get/login/logs", AdminController.GetLoginLogs)           //获取管理员登录日志 - Get admin user login logs
	}

	routerList9 := router.Group("/photo").Use(JwtMiddleware.JwtMiddleware()) // JWT中间件 - JWT Middleware
	{
		routerList9.POST("/list", PhotoController.GetPhotoList) //获取图片集 - Get Photo List
	}

	return router
}
