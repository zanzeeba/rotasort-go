package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// tasks
// add log and time back in
//	-    -    -    -    -    -    -
// tasks listing page
func getHolidays(c *gin.Context) {

	var (
		id            int64
		holiday_start time.Time
		holiday_end   time.Time
		//companies_id  int64
		//stores_id     int64
		//dept_id       int64
	)

	rows, err := db.Query("SELECT id, holiday_start, holiday_end FROM holidays")
	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error reading questions: %q", err))
		return
	}

	defer rows.Close()

	results := []Holidays{}
	tRes := Holidays{}
	for rows.Next() {

		rows.Scan(&id, &holiday_start, &holiday_end)

		tRes.Id = id
		tRes.HolidayStart = holiday_start
		tRes.HolidayEnd = holiday_end
		results = append(results, tRes)
	}

	c.HTML(http.StatusOK, "tasks.tmpl.html", gin.H{
		"Holidaying": results,
	})
}
