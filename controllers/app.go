package controller

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	Users "github.com/abhinavramkumar/go-rss/models"
	"github.com/abhinavramkumar/go-rss/structs"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

func PingController(ctx *gin.Context) {
	message := "pong"
	ctx.String(http.StatusOK, message)
}

func LoginController(ctx *gin.Context) {
	type RequestStruct struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var request RequestStruct
	err := ctx.BindJSON(&request)
	if err != nil {
		log.Println(err)
	}

	var userID string
	userID, err = Users.CheckIfUserExists(request.Email)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":  err.Error(),
			"status": false,
		})
		return
	}

	if userID == "" {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":  "user does not exist",
			"status": false,
		})
		return
	}

	var isPasswordVerified bool
	if isPasswordVerified, err = Users.VerifyPassword(request.Password, userID); isPasswordVerified {
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":  err.Error(),
				"status": false,
			})

			return
		}
		expiration := time.Now().Add(time.Hour * 24).Unix()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"exp":     expiration,
			"user_id": userID,
			"scope":   "user",
			"email":   request.Email,
		})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString([]byte("secretkey"))
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":  err.Error(),
				"status": false,
			})
			return
		}

		ctx.SetCookie("token", tokenString, int(expiration), "/", "127.0.0.1", false, false)
		ctx.JSON(http.StatusOK, gin.H{
			"status": "true",
		})
	} else {
		ctx.JSON(http.StatusConflict, gin.H{
			"status": "false",
			"error":  "username/password does not match",
		})
	}
}

func CreateUserController(ctx *gin.Context) {
	var request structs.UserStruct
	err := ctx.BindJSON(&request)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "false",
			"error":  err,
		})
		return
	}

	v := validator.New()
	err = v.Struct(request)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "false",
			"error":  err,
		})
		return
	}

	var user structs.UserStruct
	user, err = Users.InsertUser(request)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "false",
			"error":  err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":   user,
		"status": true,
	})
}

func GetSuggestedFeeds(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}

func GetFeed(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}

func UploadOPML(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	f, _ := file.Open()
	// Read the XML data from the file.
	data, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Parse the XML data.
	xmlData := structs.Opml{}
	err = xml.Unmarshal(data, &xmlData)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("\n %+v", xmlData)

	ctx.JSON(http.StatusBadRequest, gin.H{
		"status": "false",
		"error":  "",
	})
}
