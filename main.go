package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func main() {
	go api()
	ui()
}

func ui() {
	ui := gin.Default()
	ui.StaticFS("/", http.Dir("./ui"))
	ui.Run(":8080")
}

type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Balance  float32 `json:"balance"`
	Elderly  bool    `json:"elderly"`
	Student  bool    `json:"student"`
	Employee bool    `json:"employee"`
	Passes   []*Pass `json:"passes"`
}

type TransportationType string
type FormOfPayment string

const (
	BUS    TransportationType = "Bus"
	RAIL   TransportationType = "Community Rail"
	SUBWAY TransportationType = "Subway"

	PREPAID FormOfPayment = "Prepaid"
	MONTHLY FormOfPayment = "Monthly"
)

type Pass struct {
	ID      int                `json:"id"`
	Type    TransportationType `json:"type"`
	Payment FormOfPayment      `json:"payment"`
}

func api() {
	api := gin.Default()

	i := 0
	login := make(map[string]int)
	umap := make(map[int]*User)

	// enable CORS
	api.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	api.POST("/login", func(c *gin.Context) {
		uname := c.PostForm("username")
		if uid, ok := login[uname]; ok {
			c.JSON(http.StatusOK, gin.H{
				"message": "Logging in as " + uname + "...",
				"id":      uid,
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "User not found: " + uname,
			})
		}
	})

	api.POST("/register", func(c *gin.Context) {
		uname := c.PostForm("username")
		elderly, _ := strconv.ParseBool(c.PostForm("elderly"))
		student, _ := strconv.ParseBool(c.PostForm("student"))
		employee, _ := strconv.ParseBool(c.PostForm("employee"))

		if _, ok := login[uname]; ok {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Username already registered.",
			})
		} else {
			i += 1
			login[uname] = i
			umap[i] = &User{
				Name:     uname,
				ID:       i,
				Elderly:  elderly,
				Student:  student,
				Employee: employee,
			}
			c.JSON(http.StatusCreated, gin.H{
				"message": "Creating user " + uname,
				"id":      i,
				"details": umap[i],
			})
		}
	})

	api.GET("/user/:uid", func(c *gin.Context) {
		uid, _ := strconv.Atoi(c.Param("uid"))
		c.JSON(http.StatusOK, gin.H{
			"user": umap[uid],
		})
	})

	api.POST("/user/:uid/pass/create", func(c *gin.Context) {
		uid, _ := strconv.Atoi(c.Param("uid"))

		t := TransportationType(c.PostForm("type"))
		p := FormOfPayment(c.PostForm("payment"))

		if user, ok := umap[uid]; ok {
			user.Passes = append(user.Passes, &Pass{
				ID:      len(user.Passes) + 1,
				Type:    t,
				Payment: p,
			})
			c.JSON(http.StatusOK, gin.H{
				"message": "Added a pass to user account #" + strconv.Itoa(uid),
				"user":    user,
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "We couldn't find the requested user.",
			})
		}
	})

	api.Run(":8081")
}
