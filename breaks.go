package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// breaks
// add log and time back in
//	-    -    -    -    -    -    -
// tasks listing page
func getBreaks(c *gin.Context) {

	var (
		id           int64
		break_name   string
		break_length int64
		//companies_id  int64
		//stores_id     int64
		//dept_id       int64
	)

	rows, err := db.Query("SELECT id, break_name, break_length FROM breaks")
	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error reading breaks: %q", err))
		return
	}

	defer rows.Close()

	results := []Breaks{}
	tRes := Breaks{}
	for rows.Next() {

		rows.Scan(&id, &break_name, &break_length)

		tRes.Id = id
		tRes.BreakName = break_name
		tRes.BreakLength = break_length

		results = append(results, tRes)
	}

	c.HTML(http.StatusOK, "breaks.tmpl.html", gin.H{
		"Breaking": results,
	})
}

//	-    -    -    -    -    -    -
// breaks edit page
func getBreaksedit(c *gin.Context) {

	// the parameter - id of the question
	bid := c.Param("bid")

	var (
		id           int64
		break_name   string
		break_length int64
	)

	err := db.QueryRow("SELECT id, break_name, break_length FROM breaks WHERE id = $1", bid).Scan(&id, &break_name, &break_length)

	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error reading questions: %q", err))
		return
	}

	//defer rows.Close()
	results := []Breaks{}
	tRes := Breaks{}

	tRes.Id = id
	tRes.BreakName = break_name
	tRes.BreakLength = break_length

	results = append(results, tRes)

	c.HTML(http.StatusOK, "breaksedit.tmpl.html", gin.H{
		"Breaking": results,
	})
}

//	-    -    -    -    -    -    -
// breaks update page
func postBreaksupdate(c *gin.Context) {

	// get the current time
	updated_on := time.Now().Local()
	// get all the post vars
	id := c.PostForm("Id")
	break_name := c.PostForm("BreakName")
	break_length := c.PostForm("BreakLength")

	stmt, err := db.Prepare("UPDATE breaks SET break_name = $1, break_length = $2, updated_on = $3 WHERE id = $4")

	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(break_name, break_length, updated_on, id)

	if err != nil {
		log.Fatal(err)
	}

	c.HTML(http.StatusOK, "breaksupdate.tmpl.html", gin.H{
		"break_name": break_name,
		"res":        res,
	})
}

//	-    -    -    -    -    -    -
// breaks delete page row from db
func getBreaksdelete(c *gin.Context) {

	// the parameter - id of the holiday
	hid := c.Param("hid")

	stmt, err := db.Prepare("DELETE FROM holidays WHERE id = $1")

	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(hid)

	if err != nil {
		log.Fatal(err)
	}

	c.HTML(http.StatusOK, "holidaysdelete.tmpl.html", gin.H{
		"holidays": "going to delete ",
		"hid":      hid,
		"res":      res,
	})
}

//	-    -    -    -    -    -    -
// breaks create page just the form
func getBreakscreateform(c *gin.Context) {

	c.HTML(http.StatusOK, "breakscreate.tmpl.html", gin.H{})

}
