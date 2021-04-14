package dao

import (
	"fmt"
	"gggin/config"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"time"
)

import _ "github.com/go-sql-driver/mysql"

type Dao struct {
	db  *gorm.DB
	rds *redis.Pool
}

var daoInstance *Dao

func NewDao(cfg *config.Config) *Dao {
	dao := new(Dao)

	// db
	dsnr := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Mysql.User, cfg.Mysql.Psw, cfg.Mysql.Host.Read, cfg.Mysql.Port, cfg.Mysql.DbName)
	var err error
	dao.db, err = gorm.Open("mysql", dsnr)
	if err != nil {
		panic(err.Error())
	}

	// redis
	dao.rds = NewRedisPool(&cfg.RedisDefault)

	daoInstance = dao

	return daoInstance
}

func GetDao() *Dao {
	return daoInstance
}

func (dao *Dao) Close() {
	if dao.db != nil {
		dao.db.Close()
	}
}

// NewRedisPool new redis pool
func NewRedisPool(cfg *config.RedisCfg) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", cfg.Addr)
			if err != nil {
				return nil, err
			}
			if len(cfg.Psw) > 0 {
				if _, err := c.Do("AUTH", cfg.Psw); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
