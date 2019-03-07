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
	"strings"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := gin.Default()
	for _, route := range routes {
		switch route.Method {
		case "GET":
			router.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			router.POST(route.Pattern, route.HandlerFunc)
		case "PUT":
			router.PUT(route.Pattern, route.HandlerFunc)
		case "DELETE":
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		"GET",
		"/",
		Index,
	},

	{
		"CreateComment",
		strings.ToUpper("Post"),
		"/posts/:postId/comments",
		CreateComment,
	},

	{
		"CreatePost",
		strings.ToUpper("Post"),
		"/posts",
		CreatePost,
	},

	{
		"DeleteComment",
		strings.ToUpper("Delete"),
		"/posts/:postId/comments/:commentId",
		DeleteComment,
	},

	{
		"DeletePost",
		strings.ToUpper("Delete"),
		"/posts/:postId",
		DeletePost,
	},

	{
		"GetComments",
		strings.ToUpper("Get"),
		"/posts/:postId/comments",
		GetComments,
	},

	{
		"GetPost",
		strings.ToUpper("Get"),
		"/posts/:postId",
		GetPost,
	},

	{
		"GetPosts",
		strings.ToUpper("Get"),
		"/posts",
		GetPosts,
	},

	{
		"UpdatePost",
		strings.ToUpper("Put"),
		"/posts/:postId",
		UpdatePost,
	},
}