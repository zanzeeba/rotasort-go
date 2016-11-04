package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

//	-    -    -    -    -    -    -
// staff listing page
func getStaff(c *gin.Context) {

	var (
		id        int64
		dept_id   int64
		username  string
		firstname string
		lastname  string
	)

	rows, err := db.Query("SELECT id, username, firstname, lastname, dept_id FROM staff")
	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error reading questions: %q", err))
		return
	}

	defer rows.Close()

	results := []Staff{}
	tRes := Staff{}
	for rows.Next() {
		fmt.Println(".............................")
		fmt.Println(rows)
		rows.Scan(&id, &dept_id, &username, &firstname, &lastname)

		tRes.Id = id
		tRes.Username = username
		tRes.DeptId = dept_id
		tRes.Firstname = firstname
		tRes.Lastname = lastname
		results = append(results, tRes)
	}

	c.HTML(http.StatusOK, "staff.tmpl.html", gin.H{
		"Staffing": results,
	})
}

// after this can all go
func repeatHandler(c *gin.Context) {
	var buffer bytes.Buffer
	for i := 0; i < repeat; i++ {
		buffer.WriteString("Hello from Go!\n")
	}
	c.String(http.StatusOK, buffer.String())
}

func dbFunc(c *gin.Context) {
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS ticks (tick timestamp)"); err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error creating database table: %q", err))
		return
	}

	if _, err := db.Exec("INSERT INTO ticks VALUES (now())"); err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error incrementing tick: %q", err))
		return
	}

	rows, err := db.Query("SELECT tick FROM ticks")
	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error reading ticks: %q", err))
		return
	}

	defer rows.Close()
	for rows.Next() {
		var tick time.Time
		if err := rows.Scan(&tick); err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error scanning ticks: %q", err))
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("Read from DB: %s\n", tick.String()))
	}
}

// to test a get
func getTest(c *gin.Context) {

	var (
		//		task string
		t Task
	//		err  error
	)
	fmt.Println(".............................")
	fmt.Println(c.Query("taskId"))
	fmt.Println(".............................")
	t.ID = c.Query("taskId")
	t.TaskName = c.Query("taskName")

	//
	//	if err := scanParams(c, "task", &clientID); err != nil {
	//		c.AbortWithError(http.StatusBadRequest, err)
	//		return
	//	}
	//
	//	if j, err = dbSelectJobs(clientID); err != nil {
	//		c.AbortWithError(http.StatusServiceUnavailable, err)
	//		return
	//	}
	//
	//	if j == nil {
	//		logInfof("no jobs found for %s", clientID)
	//		c.AbortWithStatus(http.StatusNotFound)
	//		return
	//	}

	c.JSON(http.StatusOK, t)
}

// to test a template invoke
func getTemplTest(c *gin.Context) {
	fmt.Println(".............................")
	fmt.Println("template test")
	fmt.Println(".............................")

	c.HTML(http.StatusOK, "templtest.tmpl.html", gin.H{
		"title": "Main website",
	})
}
