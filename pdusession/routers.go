// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package pdusession

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	logger_util "github.com/omec-project/util/logger"
	"github.com/omec-project/smf/logger"
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
	router := logger_util.NewGinWithLogrus(logger.GinLog)
	AddService(router)
	return router
}

func AddService(engine *gin.Engine) *gin.RouterGroup {
	group := engine.Group("/nsmf-pdusession/v1")

	for _, route := range routes {
		switch route.Method {
		case "GET":
			group.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			group.POST(route.Pattern, route.HandlerFunc)
		case "PUT":
			group.PUT(route.Pattern, route.HandlerFunc)
		case "DELETE":
			group.DELETE(route.Pattern, route.HandlerFunc)
		}
	}
	return group
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
		"ReleasePduSession",
		strings.ToUpper("Post"),
		"/pdu-sessions/:pduSessionRef/release",
		ReleasePduSession,
	},

	{
		"UpdatePduSession",
		strings.ToUpper("Post"),
		"/pdu-sessions/:pduSessionRef/modify",
		UpdatePduSession,
	},

	{
		"ReleaseSmContext",
		strings.ToUpper("Post"),
		"/sm-contexts/:smContextRef/release",
		HTTPReleaseSmContext,
	},

	{
		"RetrieveSmContext",
		strings.ToUpper("Post"),
		"/sm-contexts/:smContextRef/retrieve",
		RetrieveSmContext,
	},

	{
		"UpdateSmContext",
		strings.ToUpper("Post"),
		"/sm-contexts/:smContextRef/modify",
		HTTPUpdateSmContext,
	},

	{
		"PostPduSessions",
		strings.ToUpper("Post"),
		"/pdu-sessions",
		PostPduSessions,
	},

	{
		"PostSmContexts",
		strings.ToUpper("Post"),
		"/sm-contexts",
		HTTPPostSmContexts,
	},
}
