package convert

import (
	"fmt"
	// "encoding/json"
	"github.com/gin-gonic/gin"
	// "github.com/jinzhu/gorm"
	"github.com/danhngocphh/fptda2/database/models"
	// "github.com/danhngocphh/fptda2/lib/common"
)

// Post type alias
type Convert = models.Convert

// User type alias
type User = models.User

// JSON type alias
// type JSON = common.JSON

//get text

func getConvert(c *gin.Context) {
	// var convert Convert
	type RequestBody struct {
		Text string `json:"text" binding:"required"`
	}
	var requestBody RequestBody

		if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatus(400)
		return
		}
		fmt.Println(requestBody)
		
	
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
