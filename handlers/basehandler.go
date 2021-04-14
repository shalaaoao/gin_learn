package handlers

import (
	"gggin/config"
	"gggin/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handle struct {
	A         int
	handleDao *dao.Dao
	cfg       config.Config
}

func NewHandle(cfg config.Config) *Handle {
	handle := new(Handle)
	handle.A = 99
	handle.handleDao = dao.NewDao(&cfg)
	handle.cfg = cfg
	return handle
}

type ResponseSt struct {
	ErrCode int         `json:"errcode"`
	ErrMsg  string      `json:"errmsg"`
	Data    interface{} `json:"data,omitempty"`
}

func success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseSt{
		ErrCode: 0,
		ErrMsg:  "",
		Data:    data,
	})
}

func fail(c *gin.Context, errMsg string) {
	c.JSON(http.StatusBadRequest, &ResponseSt{
		ErrCode: 1,
		ErrMsg:  errMsg,
	})
}
