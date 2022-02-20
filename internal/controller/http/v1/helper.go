package v1

import (
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

func invalidInputMsg(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		jsonMsgKey: "invalid input",
	})
}

func intParam(c *gin.Context, key string) int {
	str := c.Param(key)
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		writeJSONMsg(c, http.StatusBadRequest, "Bad request. Can't parse params.")
	}
	return int(i)

}
