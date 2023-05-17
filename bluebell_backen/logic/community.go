/*
*

	@author:XiaoLi
	@data:2022/12/30
	@note:

*
*/
package logic

import (
	"bluebell_mybacken/dao/mysql"
	"bluebell_mybacken/models"
)

/**
 * @Author XiaoLi
 * @Description //TODO 查询分类社区列表
 * @Date 16:42 2022/12/30
 **/
func GetCommunityList() ([]*models.Community, error) {
	// 查数据库 查找到所有的community 并返回
	return mysql.GetCommunityList()
}

/**
 * @Author XiaoLi
 * @Description //TODO 根据ID查询分类社区详情
 * @Date 17:08 2022/12/30
 **/
func GetCommunityDetailByID(id uint64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityByID(id)
}
