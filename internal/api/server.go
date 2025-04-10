package api

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/recrsn/mail-sink/internal/email"
)

type Server struct {
	store    *email.Store
	smtpPort int
	httpPort int
	uiAssets embed.FS
}

func NewServer(store *email.Store, smtpPort, httpPort int, uiAssets embed.FS) *Server {
	// Set Gin mode based on environment
	if os.Getenv("GIN_MODE") != "" {
		gin.SetMode(os.Getenv("GIN_MODE"))
	}

	return &Server{
		store:    store,
		smtpPort: smtpPort,
		httpPort: httpPort,
		uiAssets: uiAssets,
	}
}

func (s *Server) Start() error {
	r := gin.Default()

	fs, err := static.EmbedFolder(s.uiAssets, "ui/dist")
	if err != nil {
		return err
	}

	r.Use(static.Serve("/", fs))
	r.NoRoute(static.Serve("/", fs))

	// API routes
	api := r.Group("/api")
	{
		api.GET("/emails", func(c *gin.Context) {
			c.JSON(http.StatusOK, s.store.GetAll())
		})

		api.GET("/emails/:id", func(c *gin.Context) {
			id := c.Param("id")
			email := s.store.GetByID(id)
			if email == nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Email not found"})
				return
			}
			c.JSON(http.StatusOK, email)
		})

		api.DELETE("/emails", func(c *gin.Context) {
			s.store.Clear()
			c.Status(http.StatusNoContent)
		})
	}

	// Server info API
	api.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"smtpPort": s.smtpPort,
			"httpPort": s.httpPort,
		})
	})

	log.Printf("Starting HTTP server on :%d", s.httpPort)
	log.Printf("Web UI available at http://localhost:%d", s.httpPort)
	return r.Run(fmt.Sprintf(":%d", s.httpPort))
}
