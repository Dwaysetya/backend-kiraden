package controllers

import (
	"backend/database"
	"backend/models"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)


func CreateProduct(c *gin.Context){
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 10<<20)
	title := c.PostForm("title")
	description := c.PostForm("description")
	priceStr := c.PostForm("price")

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image file requaired"})
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".webp" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only JPG, JPEG, PNG, WEBP allowed"})
		return
	}

	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	path := filepath.Join("uploads", filename)
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failde to save image"})
		return
	}

	price, _ := strconv.ParseFloat(priceStr, 64)
	product := models.Products{Title: title, Description: description, Price: price, Image: filename}
	if err := database.DB.Create(&product).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": product})
}

func GetProduct(c *gin.Context){
	var product []models.Products
	database.DB.Find(&product)
	c.JSON(http.StatusOK, gin.H{"success": true, "data": product})
}