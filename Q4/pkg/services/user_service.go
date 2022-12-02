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
	var response = common.JsonResponse[[]model.User]{Type: "success", Data: users, IsError: false}
	json.NewEncoder(c.Writer).Encode(response)
}

func GetUser(c *gin.Context) {
	var response = common.JsonResponse[model.User]{}
	var user model.User
	userID := c.Param("userid")

	db := persistence.SetupDB()

	err := db.QueryRow("SELECT user_id, username, mail, password_hash FROM public.user WHERE is_deleted=false AND user_id=$1;", userID).Scan(&user.UserID, &user.UserName, &user.Mail, &user.PasswordHash)

	common.CheckError(err)

	response = common.JsonResponse[model.User]{Type: "success", Message: "", Data: user, IsError: false}

	json.NewEncoder(c.Writer).Encode(response)
}

func CreateUser(c *gin.Context) {
	var user model.User

	common.CheckError(json.NewDecoder(c.Request.Body).Decode(&user))

	var response = common.JsonResponse[string]{}

	if !util.IsValidPassword(user.Password) || !util.IsValidUsername(user.Mail) {
		response = common.JsonResponse[string]{Type: "error", Message: "Username must be equal or longer than 7 character. Password must be complex.", IsError: true}
	} else {
		user.PasswordHash = util.GetMD5Hash(user.Password)
		db := persistence.SetupDB()

		var lastInsertID string
		err := db.QueryRow("INSERT INTO public.user (username, password_hash, mail) VALUES($1, $2, $3) returning user_id;", user.UserName, user.PasswordHash, user.Mail).Scan(&lastInsertID)

		common.CheckError(err)

		response = common.JsonResponse[string]{Type: "success", Message: "The user has been inserted successfully!", Data: lastInsertID, IsError: false}
	}

	json.NewEncoder(c.Writer).Encode(response)
}

func UpdateUser(c *gin.Context) {
	var updatedUser model.User

	common.CheckError(json.NewDecoder(c.Request.Body).Decode(&updatedUser))

	var response = common.JsonResponse[string]{}

	updatedUser.PasswordHash = util.GetMD5Hash(updatedUser.Password)
	db := persistence.SetupDB()

	var user model.User

	err := db.QueryRow("SELECT user_id, username, password_hash, mail, is_deleted FROM public.user WHERE user_id =$1; AND is_deleted=false", updatedUser.UserID).Scan(&user.UserID, &user.UserName, &user.PasswordHash, &user.Mail, &user.IsDeleted)
	common.CheckError(err)

	// if !util.IsNullOrEmpty(user.Password) && util.IsValidPassword(user.Password) {
	// 	updatedUser.Password = user.Password
	// }
	if util.IsValidMail(updatedUser.Mail) {
		user.Mail = updatedUser.Mail
	}

	if util.IsValidUsername(updatedUser.UserName) {
		user.UserName = updatedUser.UserName
	}

	var userID string
	err = db.QueryRow("UPDATE public.user SET username = $1, password_hash = $2, mail = $3 WHERE user_id = $4 AND is_deleted=false returning user_id;", user.UserName, user.PasswordHash, user.Mail, user.UserID).Scan(&userID)

	common.CheckError(err)
	response = common.JsonResponse[string]{Type: "success", Message: "The user has been updated successfully!", Data: updatedUser.UserID, IsError: false}

	json.NewEncoder(c.Writer).Encode(response)
}

func DeleteUser(c *gin.Context) {
	var response = common.JsonResponse[model.UserID]{}
	var userID model.UserID

	common.CheckError(json.NewDecoder(c.Request.Body).Decode(&userID))

	db := persistence.SetupDB()

	err := db.QueryRow("UPDATE public.user SET is_deleted=true WHERE user_id = $1 returning user_id;", userID.UserID).Scan(&userID.UserID)

	common.CheckError(err)

	response = common.JsonResponse[model.UserID]{Type: "success", Message: "The user has been deleted successfully!", Data: userID, IsError: false}

	json.NewEncoder(c.Writer).Encode(response)
}
