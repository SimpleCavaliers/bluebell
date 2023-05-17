/*
*

	@author:XiaoLi
	@data:2022/12/31
	@note: mysql/post 测试

*
*/
package mysql

import (
	"bluebell_mybacken/models"
	"bluebell_mybacken/settings"
	"testing"
	"time"
)

func init() {
	dbCfg := settings.MySQLConfig{
		Host:         "127.0.0.1",
		User:         "root",
		Password:     "10428376",
		DB:           "bluebell",
		Port:         3306,
		MaxOpenConns: 10,
		MaxIdleConns: 10,
	}
	err := Init(&dbCfg)
	if err != nil {
		panic(err)
	}
}

func TestCreatePost(t *testing.T) {
	post := models.Post{
		PostID:      4,
		AuthorId:    1,
		CommunityID: 1,
		Status:      0,
		Title:       "test",
		Content:     "test123",
		CreateTime:  time.Time{},
	}
	err := CreatePost(&post)
	if err != nil {
		t.Fatalf("CreatePost insert record into mysql failed, err:%v\n", err)
	}
	t.Logf("CreatePost insert record into mysql success")
}
