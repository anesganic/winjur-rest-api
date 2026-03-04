package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/microsoft/go-mssqldb"
)

func main() {
	InitDB()

	r := gin.Default()

	// Middleware für Token-basierte Systemauthentifizierung
	r.Use(func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		hash := sha256.Sum256([]byte(token))
		hashStr := hex.EncodeToString(hash[:])

		if !validTokenHashes[hashStr] {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		c.Next()
	})

	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })

	// Adressen (Subjekte)
	r.GET("/ansprechpartner", getAnsprechpartner)
	r.GET("/domizil", getDomizil)
	r.GET("/subjekt", getSubjekt)
	r.GET("/subjektstruct", getSubjektStruct)
	r.GET("/subjekt/:id", getSubjektDetail)

	// Projekte (Fall)
	r.GET("/fall", getFall)
	r.GET("/fall/:id", getFallDetail)

	r.Run(":8080")
}
