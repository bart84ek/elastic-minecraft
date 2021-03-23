package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bart84ek/elastic-minecraft-panel/handlers"
	"github.com/gin-gonic/gin"
)

var usersFile = "./mine-panel-users.txt"

func main() {
	accounts, err := loadAccounts()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.LoadHTMLGlob("./public/templates/*")
	router.Static("/shared", "./public/shared/")
	router.StaticFile("/favicon.ico", "./public/assets/favicon.ico")

	authorized := router.Group("/", gin.BasicAuth(accounts))
	authorized.GET("/", handlers.IndexEndpoint())
	authorized.GET("/status", handlers.StatusEndpoint())
	authorized.GET("/start", handlers.StartEndpoint())
	authorized.GET("/stop", handlers.StopEndpoint())

	// Assets
	authorized.Static("/assets", "./public/assets/")

	router.Run("0.0.0.0:8080")
}

func loadAccounts() (gin.Accounts, error) {
	file, err := os.Open(usersFile)
	if err != nil {
		return gin.Accounts{}, fmt.Errorf("error loading cred files %v", err)
	}
	defer file.Close()

	accounts := gin.Accounts{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		creds := strings.Split(scanner.Text(), ":")
		accounts[creds[0]] = creds[1]
	}

	if err := scanner.Err(); err != nil {
		return accounts, err
	}
	return accounts, nil
}
