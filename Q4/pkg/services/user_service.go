package services

import (
	common "byFood/Q4/pkg/common"
	model "byFood/Q4/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// common.JsonResponse[model.User]
func GetUsers(c *gin.Context) {
	var user model.User = model.User{UserID: "34", UserName: "Burak", Password: "pa$$w0rd"}
	c.JSON(http.StatusOK, common.JsonResponse[model.User]{Type: "type", Data: user, Message: "Success"})
	//c.IndentedJSON(http.StatusOK, "")
}
