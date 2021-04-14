package dao

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
)

type CmyTopic struct {
	ID    int    `gorm:"primary_key" json:"id"`
	Name  string `json:"name,omitempty"`
	IsTop int64  `json:"is_top"`
}

func (dao *Dao) GetTopic(id int) (CmyTopic, error) {
	var query CmyTopic

	err := dao.db.Table("community_topic").Where("id=?", id).First(&query).Error
	if err != nil {
		return query, err
	}

	return query, nil
}

func (dao *Dao) GetTopicList() ([]*CmyTopic, error) {
	var query []*CmyTopic

	err := dao.db.Table("community_topic").Where("deleted_at is null").Limit(20).Offset(0).Find(&query).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil, err
	}

	return query, nil
}

func (dao *Dao) GetFirstRds() string {

	c := dao.rds.Get()
	defer c.Close()

	s, err := redis.String(c.Do("GET", "firstrds"))
	if err == redis.ErrNil {
		return "empty key"
	}

	if err != nil {
		return err.Error()
	}

	return s
}
