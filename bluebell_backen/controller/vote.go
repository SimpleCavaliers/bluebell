package controller

import (
	"bluebell_mybacken/logic"
	"bluebell_mybacken/models"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

/**
 * @Author XiaoLi
 * @Description //TODO 投票
 * @Date 10:53 2023/1/2
 **/
// VoteHandler 投票
// @Summary 投票
// @Description 投票
// @Tags 投票业务接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /vote [POST]
func VoteHandler(c *gin.Context) {
	// 参数校验,给哪个文章投什么票
	vote := new(models.VoteDataForm)
	if err := c.ShouldBindJSON(&vote); err != nil {
		errs, ok := err.(validator.ValidationErrors) // 类型断言
		if !ok {
			zap.L().Error("VoteHandler with invalid params", zap.Error(err))
			ResponseError(c, CodeInvalidParams)
			return
		}
		errdata := removeTopStruct(errs.Translate(trans)) // 翻译并去除掉错误提示中的结构体标识
		ResponseErrorWithMsg(c, CodeInvalidParams, errdata)
		return
	}
	// 获取当前请求用户的id
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNotLogin)
		return
	}
	// 具体投票的业务逻辑
	if err := logic.VoteForPost(userID, vote); err != nil {
		zap.L().Error("logic.VoteForPost() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
