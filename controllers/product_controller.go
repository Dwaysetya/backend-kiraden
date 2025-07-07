package controllers

import (
	"backend/database"
	"backend/models"
	"fmt"
	"net/http"
	"os"
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
		fmt.Println("DB Error:", err)
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

func UpdateProduct(c *gin.Context){
	id := c.Param("id")
	var product models.Products
	if err := database.DB.First(&product, id).Error; err != nil{
		fmt.Println("Produk tidak ditemukan:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	title := c.PostForm("title")
	description := c.PostForm("description")
	priceStr := c.PostForm("price")
	price, _ := strconv.ParseFloat(priceStr, 64)

	if file, err := c.FormFile("image"); err == nil{
		ext := strings.ToLower(filepath.Ext(file.Filename))
		filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
		path := filepath.Join("Uploads", filename)
		if err := c.SaveUploadedFile(file, path); err == nil{
			os.Remove(filepath.Join("uploads", product.Image)) // delete old
			product.Image = filename
		}
	}
	product.Title = title
	product.Description = description
	product.Price = price
	database.DB.Save(&product)
	c.JSON(http.StatusOK, gin.H{"success": true, "data": product})
}

func DeleteProduct(c *gin.Context){
	id := c.Param("id")
	var product models.Products
	if err := database.DB.First(&product, id).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	os.Remove(filepath.Join("uploads", product.Image))
	database.DB.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"success": true , "message": "Product deleted"})
}