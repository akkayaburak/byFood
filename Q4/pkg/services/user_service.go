package services

import (
	common "byFood/Q4/pkg/common"
	model "byFood/Q4/pkg/models"
	persistence "byFood/Q4/pkg/persistence"
	util "byFood/Q4/pkg/utils"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

	db := persistence.SetupDB()
	rows, err := db.Query("SELECT user_id, username, mail, password_hash FROM public.user WHERE is_deleted=false;")

	common.CheckError(err)

	var users []model.User

	for rows.Next() {
		var userID string
		var userName string
		var mail string
		var passwordHash string

		err = rows.Scan(&userID, &userName, &mail, &passwordHash)

		common.CheckError(err)

		users = append(users, model.User{UserID: userID, UserName: userName, Mail: mail, PasswordHash: passwordHash})
	}
	var response = common.JsonResponse[[]model.User]{Type: "success", Data: users}
	json.NewEncoder(c.Writer).Encode(response)
}

func CreateUser(c *gin.Context) {
	var user model.User

	common.CheckError(json.NewDecoder(c.Request.Body).Decode(&user))

	var response = common.JsonResponse[string]{}

	if !util.IsValidPassword(user.Password) || !util.IsValidUsername(user.Mail) {
		response = common.JsonResponse[string]{Type: "error", Message: "Username must be equal or longer than 7 character. Password must be complex."}
	} else {
		user.PasswordHash = util.GetMD5Hash(user.Password)
		db := persistence.SetupDB()

		var lastInsertID string
		err := db.QueryRow("INSERT INTO public.user (username, password_hash, mail) VALUES($1, $2, $3) returning user_id;", user.UserName, user.PasswordHash, user.Mail).Scan(&lastInsertID)

		common.CheckError(err)

		response = common.JsonResponse[string]{Type: "success", Message: "The user has been inserted successfully!", Data: lastInsertID}
	}

	json.NewEncoder(c.Writer).Encode(response)
}
