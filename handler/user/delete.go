package user

import (
	"strconv"
	
	"haishenming/go-api/model"
	"haishenming/go-api/pkg/errno"
	. "haishenming/go-api/handler"
	
	"github.com/gin-gonic/gin"
)

// Delete delete an user by the user identifier.
func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	
	SendResponse(c, nil, nil)
}
