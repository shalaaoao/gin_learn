package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

func (h Handle) HelloHandler(c *gin.Context) {

	fmt.Println(h.cfg)

	success(c, gin.H{
		"": "hello world",
	})

	//fail(c, "报错了")
}

func (h Handle) FirstTopic(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))

	s, err := h.handleDao.GetTopic(id)
	if err != nil {

		if err == gorm.ErrRecordNotFound {
			fail(c, "找不到记录")
			return
		}

		fail(c, err.Error())
		return
	}
	
	success(c, s)
}

func (h Handle) TopicList(c *gin.Context) {
	s, err := h.handleDao.GetTopicList()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, s)
}

func (h Handle) FirstRds(c *gin.Context) {

	s := h.handleDao.GetFirstRds()

	c.JSON(http.StatusOK, s)

}
