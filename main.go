package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/russross/blackfriday"
)

var (
	repeat int
	db     *sql.DB
)

func main() {
	var err error
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	tStr := os.Getenv("REPEAT")
	repeat, err = strconv.Atoi(tStr)
	if err != nil {
		log.Printf("Error converting $REPEAT to an int: %q - Using default\n", err)
		repeat = 5
	}

	// this is opening the db connection
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/mark", func(c *gin.Context) {
		c.String(http.StatusOK, string(blackfriday.MarkdownBasic([]byte("**hi!**"))))
	})

	router.GET("/repeat", repeatHandler)

	router.GET("/db", dbFunc)

	router.GET("/test", getTest)

	router.GET("/templtest", getTemplTest)

	//    -    -    -    -    -    -    -    -    -
	// staff routing block
	router.GET("/staff", getStaff)

	router.GET("/staffedit/:sid", getStaffedit)

	router.POST("/staffupdate", postStaffupdate)

	router.GET("/staffdelete/:sid", getStaffdelete)

	router.GET("/staffcreateform", getStaffcreateform)

	router.POST("/staffcreate", postStaffcreate)

	//    -    -    -    -    -    -    -    -    -
	//tasks routing block
	//    -    -    -    -    -    -    -    -    -
	// staff routing block
	router.GET("/staff", getStaff)

	router.Run(":" + port)
}
