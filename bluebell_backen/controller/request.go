package controller

/**
	对请求做一些处理
**/
import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	ContextUserIDKey = "userID"
)

var (
	ErrorUserNotLogin = errors.New("当前用户未登录")
)

/**
 * @Author XiaoLi
 * @Description //TODO 获取当前登录用户ID
 * @Date 22:00 2022/12/26
 **/
// getCurrentUserID 获取当前登录用户ID
// @Summary 获取当前登录用户ID
// @Description 获取当前登录用户ID
// @Tags 公共接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query userID  path    int     true        "_userID"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
func getCurrentUserID(c *gin.Context) (userID uint64, err error) {
	_userID, ok := c.Get(ContextUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = _userID.(uint64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

/**
 * @Author XiaoLi
 * @Description //TODO 分页参数
 * @Date 23:41 2022/12/30
 **/
// getPageInfo 分页参数
// @Summary 分页参数
// @Description 分页参数
// @Tags 公共接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
func getPageInfo(c *gin.Context) (int64, int64) {
	pageStr := c.Query("page")
	SizeStr := c.Query("size")

	var (
		page int64 // 第几页 页数
		size int64 // 每页几条数据
		err  error
	)
	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(SizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return page, size
}
