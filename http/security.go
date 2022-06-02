package http

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ragoncsa/todo/authz"
	"google.golang.org/api/idtoken"
)

type AuthEnforcement int

const (
	mandatory AuthEnforcement = iota
	optional
)

func authHandler(clientId string, enforcing AuthEnforcement) gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request

		// OPTIONS is used only for doing CORS preflight check
		// Also OpenAPI spec can be accessed without authentication
		if c.Request.Method == "OPTIONS" || c.Request.URL.Path == "/swagger/doc.json" {
			return
		}
		authH, ok := c.Request.Header["Authorization"]
		var user string
		if !ok && enforcing == mandatory {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "no Authorization header in request")
			return
		} else if !ok && enforcing == optional {
			user = "default-test-user@"
			if callerId := c.Request.Header.Get("CallerId"); callerId != "" {
				user = callerId
			}
		} else { // Authorization header had a value
			token := strings.Replace(authH[0], "Bearer ", "", 1)
			payload, err := idtoken.Validate(context.Background(), token, clientId)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid token")
				return
			}
			user = payload.Claims["email"].(string)
		}
		c.Set("user", user) // only for debugging
		path := strings.Split(strings.Trim(c.Request.URL.Path, "/"), "/")
		c.Set("decisionRequest", &authz.DecisionRequest{
			Method: c.Request.Method,
			Path:   path,
			User:   user,
		})
		c.Next()
		// after request
	}
}
