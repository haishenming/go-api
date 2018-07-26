package user

import (
	"fmt"
	
	. "haishenming/go-api/handler"
	"haishenming/go-api/pkg/errno"
	
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct {
	Username string `json:"username"`
}

// Create creates a new user account.
func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	
	admin2 := c.Param("username")
	log.Infof("URL username: %s", admin2)
	
	desc := c.Query("desc")
	log.Infof("URL key param desc: %s", desc)
	
	contentType := c.GetHeader("Content-Type")
	log.Infof("Header Content-Type: %s", contentType)
	
	log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)
	if r.Username == "" {
		SendResponse(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")), nil)
		return
	}
	
	if r.Password == "" {
		SendResponse(c, fmt.Errorf("password is empty"), nil)
	}
	
	rsp := CreateResponse{
		Username: r.Username,
	}
	
	// Show the user information.
	SendResponse(c, nil, rsp)
}