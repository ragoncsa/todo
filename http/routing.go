package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ragoncsa/todo/config"
	"google.golang.org/api/idtoken"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// TaskServiceHTTPHandlers defines all the handlers the TaskService needs. It's
// possible to register routes for a different implementation (like a mock).
type TaskServiceHTTPHandlers interface {
	GetTask(c *gin.Context)
	GetTasks(c *gin.Context)
	CreateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	DeleteTasks(c *gin.Context)
}

type Server struct {
	router *gin.Engine
	port   int
}

type AuthEnforcement int

const (
	mandatory AuthEnforcement = iota
	optional
)

func googleAuth(clientId string, enforcing AuthEnforcement) gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request

		// OPTIONS is used only for doing CORS preflight check
		// Also OpenAPI spec can be accessed without authentication
		if c.Request.Method != "OPTIONS" && c.Request.URL.Path != "/swagger/doc.json" {
			authH, ok := c.Request.Header["Authorization"]
			if !ok {
				if enforcing == mandatory {
					c.AbortWithStatusJSON(http.StatusUnauthorized, "no Authorization header in request")
				}
				return
			}
			token := strings.Replace(authH[0], "Bearer ", "", 1)

			payload, err := idtoken.Validate(context.Background(), token, clientId)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid token")
				return
			}
			c.Set("email", payload.Claims["email"])

		}
		c.Next()
		// after request
	}
}

func InitServer(conf *config.Config) *Server {
	server := &Server{
		router: gin.Default(),
		port:   conf.Server.Port,
	}
	server.router.Use(cors.New(cors.Config{
		AllowOrigins: conf.Frontend.Endpoints,
		AllowMethods: []string{"GET", "POST", "DELETE"},
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type", "CallerId"},
		MaxAge:       12 * time.Hour,
	}))
	if conf.Authn.DevMode {
		server.router.Use(googleAuth(conf.Authn.ClientId, optional))
	} else {
		server.router.Use(googleAuth(conf.Authn.ClientId, mandatory))
	}
	return server
}

func (s Server) RegisterRoutes(t TaskServiceHTTPHandlers) {
	s.router.GET("/tasks/", t.GetTasks)
	s.router.GET("/tasks/:taskid", t.GetTask)
	s.router.POST("/tasks/", t.CreateTask)
	s.router.DELETE("/tasks/:taskid", t.DeleteTask)
	s.router.DELETE("/tasks/", t.DeleteTasks)

	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))

}

func (s Server) Start() {
	fmt.Printf("Server at %d\n", s.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.router))
}
