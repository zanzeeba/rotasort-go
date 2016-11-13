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
		id            int64
		holiday_name  string
		holiday_start time.Time
		holiday_end   time.Time
		//companies_id  int64
		//stores_id     int64
		//dept_id       int64
	)

	rows, err := db.Query("SELECT id, holiday_name, holiday_start, holiday_end FROM holidays")
	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error reading questions: %q", err))
		return
	}

	defer rows.Close()

	results := []Breaks{}
	tRes := Breaks{}
	for rows.Next() {

		rows.Scan(&id, &holiday_name, &holiday_start, &holiday_end)

		tRes.Id = id
		//tRes.BreakName = holiday_name
		//tRes.BreakLength = holiday_start

		results = append(results, tRes)
	}

	c.HTML(http.StatusOK, "holidays.tmpl.html", gin.H{
		"Holidaying": results,
	})
}

//	-    -    -    -    -    -    -
// holidays edit page
func getBreaksedit(c *gin.Context) {

	// the parameter - id of the question
	hid := c.Param("hid")

	var (
		id            int64
		holiday_name  string
		holiday_start time.Time
		holiday_end   time.Time
	)

	err := db.QueryRow("SELECT id, holiday_name, holiday_start, holiday_end FROM holidays WHERE id = $1", hid).Scan(&id, &holiday_name, &holiday_start, &holiday_end)

	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error reading questions: %q", err))
		return
	}

	//defer rows.Close()
	results := []Breaks{}
	tRes := Breaks{}

	tRes.Id = id

	//tRes.BreakName = holiday_name
	//tRes.BreakLength = holiday_start

	results = append(results, tRes)

	c.HTML(http.StatusOK, "holidaysedit.tmpl.html", gin.H{
		"Holidaying": results,
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
	res, err := stmt.Exec(holiday_name, holiday_start, holiday_end, updated_on, id)

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
