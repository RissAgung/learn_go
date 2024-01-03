package category

import (
	"db_api/models/category"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	data, err := category.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}
	c.JSON(http.StatusOK, gin.H{"data": data, "message": "Success Get Data Categories"})
}

func GetOne(c *gin.Context) {

	id := c.Param("id")

	data, err := category.GetOne(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}
	c.JSON(http.StatusOK, gin.H{"data": data, "message": "Success Get Data Categories"})
}

func Store(c *gin.Context) {

	var requestBody struct {
		Name string `json:"name" binding:"required"`
	}

	// parsing JSON from body request into variable requestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	insertedID, err := category.Store(requestBody.Name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": insertedID, "message": "Success Insert Data"})
}

func Update(c *gin.Context) {

	// get id from parameter
	id := c.Param("id")

	var requestBody struct {
		Name string `json:"name" binding:"required"`
	}

	// parsing JSON from body request into variable requestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := category.Update(requestBody.Name, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id, "message": "Success Update Data"})
}

func Destroy(c *gin.Context) {
	// get id from parameter
	id := c.Param("id")

	if err := category.Destroy(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Success Delete Data"})
}
