package controllers

import(
	// "fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"tutorial2/connections"
	table "tutorial2/databases"
)

// DB is....
var DB = connections.DB


// GetData is....
func GetData(c *gin.Context){
	var datas []table.User
	var perPage int
	var currentPage int
	var search string

	if c.Query("per_page") == "" {
		perPage = 10
	}else{
		perPage,_ = strconv.Atoi(c.Query("per_page"))
	}

	if c.Query("page") == "" {
		currentPage = 1
	}else{
		currentPage, _ = strconv.Atoi(c.Query("page"))
	}

	if c.Query("search") == "" {
		search = ""
	}else{
		search = c.Query("search")
	}

	offset := (currentPage - 1) * perPage

	result := DB.Where("name LIKE ? AND delete_at IS NULL", `%`+ search +`%`).Order("name ASC").Find(&datas).Limit(perPage).Offset(offset)

	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Failed","status": http.StatusOK})
	}else{
		c.JSON(http.StatusOK, gin.H{"data": datas,"status": http.StatusOK})
	}

}

// GetDataByID is....
func GetDataByID(c *gin.Context){
	var datas []table.User
	var id = c.Param("id")
	
	result := DB.Where("id = ? AND delete_at IS NULL", id).Find(&datas).Limit(1)

	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Failed","status": http.StatusOK})
	}else{
		c.JSON(http.StatusOK, gin.H{"data": datas,"status": http.StatusOK})
	}

}

// InData is....
func InData(c *gin.Context){
	datas := table.User{}

	datas.Name = c.PostForm("name")
	datas.Email = c.PostForm("email")
	datas.Age,_ = strconv.Atoi(c.PostForm("age"))
	datas.Birthday,_ = time.Parse("2006-01-02 15:04:05", c.PostForm("birthday"))
	datas.MemberNumber = c.PostForm("member")
	
	result := DB.Create(&datas)

	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Failed","status": http.StatusOK})
	}else{
		c.JSON(http.StatusOK, gin.H{"data": "Success","status": http.StatusOK})
	}
	

}

// UpData is....
func UpData(c *gin.Context){
	id := c.Param("id")
	datas := table.User{}
	
	datas.Name = c.PostForm("name")
	datas.Email = c.PostForm("email")
	datas.Age,_ = strconv.Atoi(c.PostForm("age"))
	datas.Birthday,_ = time.Parse("2006-01-02 15:04:05", c.PostForm("birthday"))
	datas.MemberNumber = c.PostForm("member")

	result := DB.Model(&datas).Where("id = ?", id).Updates(&datas)

	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Failed","status": http.StatusOK})
	}else{
		c.JSON(http.StatusOK, gin.H{"datas": "Success, Changes","status": http.StatusOK})
	}

}

// DelData is Soft Delete
func DelData(c *gin.Context){
	id := c.Param("id")
	
	datas := table.User{}
	datas.DeleteAt = time.Now()

	result := DB.Model(&datas).Where("id = ?", id).Updates(&datas)

	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Failed","status": http.StatusOK})
	}else{
		c.JSON(http.StatusOK, gin.H{"datas": "Success, Delete","status": http.StatusOK})
	}
		
	c.JSON(http.StatusOK, gin.H{"datas": "Success, Delete","status": http.StatusOK})

}
