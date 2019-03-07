/*
 * 匿名掲示板API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	l "github.com/po3rin/go-playground/try-go-openapi/server/logger"
)

// CreateComment -
func CreateComment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// CreatePost -
func CreatePost(c *gin.Context) {
	var p Post
	err := c.BindJSON(&p)
	if err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, p)
}

// DeleteComment -
func DeleteComment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeletePost -
func DeletePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetComments -
func GetComments(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetPost -
func GetPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetPosts -
func GetPosts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdatePost -
func UpdatePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}