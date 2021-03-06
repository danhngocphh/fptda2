package convert

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	// "github.com/jinzhu/gorm"
	"github.com/danhngocphh/fptda2/database/models"
	// "github.com/danhngocphh/fptda2/lib/common"
)

// Post type alias
// type Convert = models.Convert

// User type alias
// type User = models.User

// JSON type alias
// type JSON = common.JSON

// get struct

type ReqBody = models.ReqBody
type ResBody = models.ResBody
type ErrBody = models.ErrBody

//get text

func getConvert(c *gin.Context) {
	w := c.Writer
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET, DELETE, PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	fmt.Println(w)
	var reqBody ReqBody
	if err := c.BindJSON(&reqBody); err != nil {
		errBody := &ErrBody{Status: "failed", Message: err.Error()}
		c.JSON(400, errBody)
		return
	}
	body := strings.NewReader(reqBody.Text)
	req, err := http.NewRequest("POST", "https://api.fpt.ai/hmi/tts/v5", body)
	if err != nil {
		errBody := &ErrBody{Status: "failed", Message: err.Error()}
		c.JSON(400, errBody)
		return
	}
	req.Header.Set("api_key", os.Getenv("API_KEY"))
	req.Header.Set("voice", reqBody.Voice)
	req.Header.Set("speed", "0")
	req.Header.Set("format", reqBody.Format)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errBody := &ErrBody{Status: "failed", Message: err.Error()}
		c.JSON(400, errBody)
		return
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	var dat map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &dat); err != nil {
		panic(err)
	}
	resBody := &ResBody{Status: "success", Message: dat["message"].(string), Text: reqBody.Text, Link: dat["async"].(string)}
	c.JSON(200, resBody)
}

// func create(c *gin.Context) {
// 	db := c.MustGet("db").(*gorm.DB)
// 	type RequestBody struct {
// 		Text string `json:"text" binding:"required"`
// 	}
// 	var requestBody RequestBody

// 	if err := c.BindJSON(&requestBody); err != nil {
// 		c.AbortWithStatus(400)
// 		return
// 	}

// 	user := c.MustGet("user").(User)
// 	post := Post{Text: requestBody.Text, User: user}
// 	db.NewRecord(post)
// 	db.Create(&post)
// 	c.JSON(200, post.Serialize())
// }

// func list(c *gin.Context) {
// 	db := c.MustGet("db").(*gorm.DB)

// 	cursor := c.Query("cursor")
// 	recent := c.Query("recent")

// 	var posts []Post

// 	if cursor == "" {
// 		if err := db.Preload("User").Limit(10).Order("id desc").Find(&posts).Error; err != nil {
// 			c.AbortWithStatus(500)
// 			return
// 		}
// 	} else {
// 		condition := "id < ?"
// 		if recent == "1" {
// 			condition = "id > ?"
// 		}
// 		if err := db.Preload("User").Limit(10).Order("id desc").Where(condition, cursor).Find(&posts).Error; err != nil {
// 			c.AbortWithStatus(500)
// 			return
// 		}
// 	}

// 	// length := len(posts)
// 	// serialized := make([]JSON, length, length)
// 	serialized := "dmcs"

// 	// for i := 0; i < length; i++ {
// 	// 	serialized[i] = posts[i].Serialize()
// 	// }

// 	c.JSON(200, serialized)
// }

// func read(c *gin.Context) {
// 	db := c.MustGet("db").(*gorm.DB)
// 	id := c.Param("id")
// 	var post Post

// 	// auto preloads the related model
// 	// http://gorm.io/docs/preload.html#Auto-Preloading
// 	if err := db.Set("gorm:auto_preload", true).Where("id = ?", id).First(&post).Error; err != nil {
// 		c.AbortWithStatus(404)
// 		return
// 	}

// 	c.JSON(200, post.Serialize())
// }

// func remove(c *gin.Context) {
// 	db := c.MustGet("db").(*gorm.DB)
// 	id := c.Param("id")

// 	user := c.MustGet("user").(User)

// 	var post Post
// 	if err := db.Where("id = ?", id).First(&post).Error; err != nil {
// 		c.AbortWithStatus(404)
// 		return
// 	}

// 	if post.UserID != user.ID {
// 		c.AbortWithStatus(403)
// 		return
// 	}

// 	db.Delete(&post)
// 	c.Status(204)
// }

// func update(c *gin.Context) {
// 	db := c.MustGet("db").(*gorm.DB)
// 	id := c.Param("id")

// 	user := c.MustGet("user").(User)

// 	type RequestBody struct {
// 		Text string `json:"text" binding:"required"`
// 	}

// 	var requestBody RequestBody

// 	if err := c.BindJSON(&requestBody); err != nil {
// 		c.AbortWithStatus(400)
// 		return
// 	}

// 	var post Post
// 	if err := db.Preload("User").Where("id = ?", id).First(&post).Error; err != nil {
// 		c.AbortWithStatus(404)
// 		return
// 	}

// 	if post.UserID != user.ID {
// 		c.AbortWithStatus(403)
// 		return
// 	}

// 	post.Text = requestBody.Text
// 	db.Save(&post)
// 	c.JSON(200, post.Serialize())
// }

// func responseWithError(response http.ResponseWriter, statusCode int, msg string) {
// 	responseWithJSON(response, statusCode, map[string]string{
// 		"error": msg,
// 	})
// }

// func responseWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
// 	result, _ := json.Marshal(data)
// 	response.Header().Set("Content-Type", "application/json")
// 	response.WriteHeader(statusCode)
// 	response.Write(result)
// }
