package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const jsonMsgKey = "message"

func writeJSONMsg(c *gin.Context, code int, msg interface{}) {
	c.JSON(code, gin.H{
		jsonMsgKey: msg,
	})
}

func writeJSONOkMsg(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusOK, gin.H{
		jsonMsgKey: msg,
	})
}

func invalidInput(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		jsonMsgKey: "invalid input",
	})
}
func invalidInputMsg(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{
		jsonMsgKey: msg,
	})
}
func notFoundMsg(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusNotFound, gin.H{
		jsonMsgKey: msg,
	})
}

func serverErrorMsg(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		jsonMsgKey: "error occurred",
	})
}

func intParam(c *gin.Context, key string) (int, error) {
	str := c.Param(key)
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return -1, errors.New("can't parse int param")
	}
	return int(i), nil

}
