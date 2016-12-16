package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

//	-    -    -    -    -    -    -
// shifts listing page
func getShifts(c *gin.Context) {

	var (
		id   int64
		name string
		//companies_id  int64
		//stores_id     int64
		//dept_id       int64
	)

	rows, err := db.Query("SELECT id, name FROM shifts")
	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error reading breaks: %q", err))
		return
	}

	defer rows.Close()

	results := []Shifts{}
	tRes := Shifts{}
	for rows.Next() {
		rows.Scan(&id, &name)
		tRes.Id = id
		tRes.Name = name
		results = append(results, tRes)
	}

	c.HTML(http.StatusOK, "shifts.tmpl.html", gin.H{
		"Shifting": results,
	})
}

//	-    -    -    -    -    -    -
// shiftpattern listing page
func getShiftPatterns(c *gin.Context) {

	var (
		id            int64
		sp_name       string
		sp_start_date time.Time
		sp_end_date   time.Time
		roll_over     bool
	)

	rows, err := db.Query("SELECT id, sp_name, sp_start_date, sp_end_date, roll_over FROM shift_patterns")
	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error reading breaks: %q", err))
		return
	}

	defer rows.Close()

	results := []ShiftPatterns{}
	tRes := ShiftPatterns{}
	for rows.Next() {

		rows.Scan(&id, &sp_name, &sp_start_date, &sp_end_date, &roll_over)

		tRes.Id = id
		tRes.SpName = sp_name
		tRes.SpStartDate = sp_start_date
		tRes.SpEndDate = sp_end_date
		tRes.RollOver = roll_over

		results = append(results, tRes)
	}

	c.HTML(http.StatusOK, "shiftpatterns.tmpl.html", gin.H{
		"ShiftPatterning": results,
	})
}
