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
// holidays update page
func postBreaksupdate(c *gin.Context) {

	// get the current time
	updated_on := time.Now().Local()
	// get all the post vars
	id := c.PostForm("Id")
	holiday_name := c.PostForm("BreakName")
	holiday_start := c.PostForm("BreakLength")

	stmt, err := db.Prepare("UPDATE holidays SET holiday_name = $1, holiday_start = $2, holiday_end = $3, updated_on = $4 WHERE id = $5")

	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(holiday_name, holiday_start, updated_on, id)

	if err != nil {
		log.Fatal(err)
	}

	c.HTML(http.StatusOK, "holidaysupdate.tmpl.html", gin.H{
		"holiday_name": holiday_name,
		"res":          res,
	})
}

//	-    -    -    -    -    -    -
// holidays delete page row from db
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
