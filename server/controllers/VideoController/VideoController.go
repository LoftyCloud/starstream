/*
 * @Descripttion: 视频控制层 - Video Controller
 */
package VideoController

import (
	"VideoHubGo/caches/VideoCache"
	"VideoHubGo/models/VideoModel"
	"VideoHubGo/services/VideoServices"
	"VideoHubGo/utils/JsonUtils"
	"VideoHubGo/utils/LogUtils"
	"VideoHubGo/utils/UploadUtils"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
)

// 使用SingleFlight防止Redis缓存击穿
var sfg singleflight.Group

/**
 * @Descripttion: 获取视频数据 - Get Video List Data
 * @Param: 分页 - page (int)
 * @Param: 数据条数 - size (int)
 * @Return: Json
 */
func GetVideoList(ctx *gin.Context) {
	requestBody := VideoModel.VideoRequest{}
	err := ctx.BindJSON(&requestBody)
	if err != nil {
		ctx.JSON(http.StatusOK, JsonUtils.JsonResult(600, "参数错误 - Parameter error", ""))
		return
	}

	page := requestBody.Page
	size := requestBody.Size
	if page < 1 {
		page = 1
	}
	if size < 20 {
		size = 20
	}
	if size > 40 {
		size = 20
	}
	offset := size * (page - 1)

	// 直接从数据库获取视频总数和列表
	count := VideoServices.GetCountVideoList()             // 查询数据库中的视频总数
	videoData := VideoServices.FindVideoList(size, offset) // 获取视频列表

	rData := map[string]interface{}{"list": videoData, "count": count}

	// 返回查询结果
	ctx.JSON(http.StatusOK, JsonUtils.JsonResult(200, "200", rData))
}

// func GetVideoList(ctx *gin.Context) {
// 	requestBody := VideoModel.VideoRequest{}
// 	err := ctx.BindJSON(&requestBody)
// 	if err != nil {
// 		ctx.JSON(http.StatusOK, JsonUtils.JsonResult(600, "参数错误 - Parameter error", ""))
// 		return
// 	}
// 	page := requestBody.Page
// 	size := requestBody.Size
// 	if page < 1 {
// 		page = 1
// 	}
// 	if size < 20 {
// 		size = 20
// 	}
// 	if size > 40 {
// 		size = 20
// 	}
// 	offset := size * (page - 1)
// 	retData, err, _ := sfg.Do("getvideolist", func() (interface{}, error) {
// 		var videoData = VideoCache.VideoGetListCache(page, size)
// 		count := VideoCache.VideoGetCount() // 读取缓存中视频总数
// 		redcount := VideoCache.GetReidsVideoListCount()

// 		if videoData == nil && len(videoData) < 20 { // 缓存数据小于20条
// 			count = VideoServices.GetCountVideoList() // 查询数据库中视频总数
// 			VideoCache.VideoSaveCountList(count)      // 保存视频总数到缓存
// 			if redcount < count {
// 				//查询数据库数据 - Find Sql Data
// 				videoData = VideoServices.FindVideoList(size, offset)
// 				//缓存到Redis里 - Cache Redis
// 				VideoCache.VideoWriteListCache(videoData)
// 			}
// 		}
// 		rData := map[string]interface{}{"list": videoData, "count": count}
// 		return rData, nil
// 	})
// 	if err != nil {
// 		LogUtils.Logger("[GIN ERROR]在处理getvideolist缓存击穿中异常：" + err.Error())
// 		ctx.JSON(http.StatusOK, JsonUtils.JsonResult(201, "未知的错误-击穿", ""))
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, JsonUtils.JsonResult(200, "200", retData))
// }

/**
 * @Descripttion: 获取定义分类的视频数据 - Get video data defining classification
 * @Param: cid (int)
 * @Param: 分页 - page (int)
 * @Param: 数据条数 - size (int)
 * @Return: Json
 */
func GetVideoClassList(ctx *gin.Context) {
	requestBody := VideoModel.VideoRequestClass{}
	err := ctx.BindJSON(&requestBody)
	if err != nil {
		ctx.JSON(http.StatusOK, JsonUtils.JsonResult(600, "参数错误 - Parameter error", ""))
		return
	}
	cid := requestBody.Cid
	page := requestBody.Page
	size := requestBody.Size
	if page < 1 {
		page = 1
	}
	if size < 20 {
		size = 20
	}
	if size > 40 {
		size = 20
	}
	if cid < 0 {
		cid = 0
	}
	offset := size * (page - 1)

	count, err, _ := sfg.Do("getclasscount", func() (interface{}, error) {
		count := VideoCache.VideoGetClassCount(cid)
		if count == 0 {
			tempCount := VideoServices.GetCountVideoClassList(cid)
			VideoCache.VideoSaveClassCountList(cid, tempCount)
			count = tempCount
		}
		return count, nil
	})
	if err != nil {
		LogUtils.Logger("[GIN ERROR]在处理getclasscount缓存击穿中异常：" + err.Error())
		ctx.JSON(http.StatusOK, JsonUtils.JsonResult(201, "未知的错误-击穿", ""))
		return
	}

	videoData := VideoServices.FindVideoInClass(cid, size, offset)
	rData := map[string]interface{}{"list": videoData, "count": count}
	ctx.JSON(http.StatusOK, JsonUtils.JsonResult(200, "200", rData))
}

/**
 * @Descripttion: 搜索视频 - Search Video
 * @Param: 分类ID - cid (int)
 * @Param: 搜索关键字 - key (string)
 * @Param: 分页 - page (int)
 * @Param: 数据条数 - size (int)
 * @Return: Json
 */
func SearchVideoClassList(ctx *gin.Context) {
	requestBody := VideoModel.VideoRequestSearch{}
	err := ctx.BindJSON(&requestBody)
	if err != nil {
		ctx.JSON(http.StatusOK, JsonUtils.JsonResult(600, "参数错误 - Parameter error", ""))
		return
	}
	cid := requestBody.Cid
	key := requestBody.Key
	page := requestBody.Page
	size := requestBody.Size
	if page < 1 {
		page = 1
	}
	if size < 20 {
		size = 20
	}
	if size > 40 {
		size = 20
	}
	if cid < 0 {
		cid = 0
	}
	offset := size * (page - 1)
	videoData, count := VideoServices.SearchVideoList_Class(cid, key, size, offset)

	rData := map[string]interface{}{"list": videoData, "count": count}
	ctx.JSON(http.StatusOK, JsonUtils.JsonResult(200, "200", rData))
}

/**
 * @Descripttion: 检查视频封面是否存在 - Check if the video cover exists
 * @Param: vid (int)
 */
func CheckCoverExists(ctx *gin.Context) {
	vid := ctx.Param("vid")
	coverPath := fmt.Sprintf("/path/to/cover/%s.jpg", vid)

	if _, err := os.Stat(coverPath); err == nil {
		ctx.JSON(http.StatusOK, gin.H{"exists": true})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"exists": false})
	}
}

/**
 * @Descripttion: 视频上传 - Video Upload
 * @Param: File
 * @Return: Json
 */
func UploadVideo_StreamFile(ctx *gin.Context) {
	// fmt.Println("用户正在上传视频")
	file, header, _ := ctx.Request.FormFile("file")
	detail := ctx.PostForm("detail")
	cid := ctx.PostForm("cid")

	recid, err := strconv.Atoi(cid)
	uid := ctx.MustGet("uid").(int)

	if err != nil {
		ctx.JSON(http.StatusOK, JsonUtils.JsonResult(602, "参数错误 - CID Base Parameter error", ""))
		return
	}
	if len(detail) < 1 {
		ctx.JSON(http.StatusOK, JsonUtils.JsonResult(601, "detail基本参数错误 - Detail Base Parameter error", ""))
		return
	}

	// 视频文件保存地址
	savePath := UploadUtils.GetFilePath("video.saveFile")

	fileSuffix := path.Ext(header.Filename) // 获取文件后缀
	if fileSuffix != ".mp4" {
		ctx.JSON(http.StatusOK, JsonUtils.JsonResult(201, "文件类型不符合要求 - Document type does not meet requirements", ""))
		return
	}
	// 保存视频文件
	saveVid := VideoServices.UploadVideo(uid, detail, recid, "")
	fmt.Println("saveVid:", saveVid)

	save, err := os.OpenFile(savePath+strconv.Itoa(saveVid)+fileSuffix, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		ctx.JSON(http.StatusOK, JsonUtils.JsonResult(202, "读取文件流失败 - Failed to read file stream", ""))
		return
	}
	// 保存视频流
	for {
		buf := make([]byte, 1024*2)
		read, err := file.Read(buf)
		if err != nil && err != io.EOF {
			ctx.JSON(http.StatusOK, JsonUtils.JsonResult(203, "视频上传出现异常 - Abnormal video uploading", ""))
			return
		}
		if read == 0 {
			break
		}
		_, err = save.Write(buf)
		if err != nil {
			ctx.JSON(http.StatusOK, JsonUtils.JsonResult(204, "视频存储过程出现异常 - Exception in video stored procedure", ""))
			return
		}
	}

	// save.Close()
	// 使用FFmprg提取第一帧作为缩略图
	// 封面保存地址
	// thumbnailPath := savePath + strconv.Itoa(saveVid) + ".jpg"
	thumbnailPath := UploadUtils.GetFilePath("video.coverFile") + strconv.Itoa(saveVid) + ".jpg"
	// fmt.Println(thumbnailPath)
	// if _, err := os.Stat(thumbnailPath); err == nil {
	// 	err = os.Remove(thumbnailPath) // 删除已存在的封面文件
	// 	if err != nil {
	// 		ctx.JSON(http.StatusOK, JsonUtils.JsonResult(205, "无法删除旧的缩略图 - Failed to delete old thumbnail", ""))
	// 		return
	// 	}
	// }

	cmd := exec.Command("ffmpeg", "-i", savePath+strconv.Itoa(saveVid)+fileSuffix, "-y", "-f", "image2", "-t", "0.001", thumbnailPath)
	err = cmd.Run()
	if err != nil {
		ctx.JSON(http.StatusOK, JsonUtils.JsonResult(206, "提取缩略图失败 - Failed to extract thumbnail", ""))
		return
	}

	rData := map[string]interface{}{"vid": saveVid}
	ctx.JSON(http.StatusOK, JsonUtils.JsonResult(200, "视频上传成功 - Video upload successed", rData))
}
