// Copyright 2015 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
This file contains tests which verify that the scenarios described
in the raft paper (https://ramcloud.stanford.edu/raft.pdf) are
handled by the raft implementation correctly. Each test focuses on
several sentences written in the paper. This could help us to prevent
most implementation bugs.
Each test is composed of three parts: init, test and check.
Init part uses simple and understandable way to simulate the init state.
Test part uses Step function to generate the scenario. Check part checks
outgoing messages and state.
*/

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
	router.LoadHTMLGlob("templates/skills/*.tmpl.html")
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

	router.GET("/tasks", getTasks)

	router.GET("/tasksedit/:tid", getTasksedit)

	router.POST("/tasksupdate", postTasksupdate)

	router.GET("/tasksdelete/:tid", getTasksdelete)

	router.GET("/taskscreateform", getTaskscreateform)

	router.POST("/taskscreate", postTaskscreate)

	//    -    -    -    -    -    -    -    -    -
	//skills routing block

	router.GET("/skills", getSkills)
	//    -    -    -    -    -    -    -    -    -

	router.Run(":" + port)
}
