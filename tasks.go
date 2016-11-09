package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
)

// tasks
// add log and time back in
//	-    -    -    -    -    -    -
// tasks listing page
func getTasks(c *gin.Context) {

	var (
		id        int64
		task_name string
		task_type string
		weighting string
	)

	rows, err := db.Query("SELECT id, task_name, task_type, weighting FROM tasks")
	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error reading questions: %q", err))
		return
	}

	defer rows.Close()

	results := []Tasks{}
	tRes := Tasks{}
	for rows.Next() {

		rows.Scan(&id, &task_name, &task_type, &weighting)

		tRes.Id = id
		tRes.TaskName = task_name
		tRes.TaskType = task_type
		tRes.Weighting = weighting
		results = append(results, tRes)
	}

	c.HTML(http.StatusOK, "tasks.tmpl.html", gin.H{
		"Tasking": results,
	})
}

//	-    -    -    -    -    -    -
// tasks edit page
func getTasksedit(c *gin.Context) {

	// the parameter - id of the question
	sid := c.Param("tid")

	var (
		id           int64
		task_name    string
		task_type    string
		weighting    string
		time_when    time.Time
		time_offset  int64
		time_float   time.Time
		time_length  int64
		no_of_jobs   int64
		time_min     string
		time_max     string
		day_of_week  string
		active       bool
		companies_id int64
		stores_id    int64
		dept_id      int64
	)

	err := db.QueryRow("SELECT d,task_name,task_type,weighting,time_when,time_offset,time_float,time_length,no_of_jobs,time_min,time_max,day_of_week,active,companies_id,stores_id,dept_id FROM tasks WHERE id = $1", tid).Scan(&id, &task_name, &task_type, &weighting, &time_when, &time_offset, &time_float, &time_length, &no_of_jobs, &time_min, &time_max, &day_of_week, &active, &companies_id, &stores_id, &dept_id)

	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error reading questions: %q", err))
		return
	}

	//defer rows.Close()
	results := []Tasking{}
	tRes := Tasking{}

	tRes.Id = int64
	tRes.TaskName = string
	tRes.TaskType = string
	tRes.Weightin = string
	tRes.TimeWhen = time.Time
	tRes.Time_Offset = int64
	tRes.TimeFloat = time.Time
	tRes.TimeLength = int64
	tRes.NoOfJobs = int64
	tRes.TimeMin = string
	tRes.TimeMax = string
	tRes.DayOfWeek = string
	tRes.Active = bool
	tRes.CompaniesId = int64
	tRes.StoresId = int64
	tRes.DeptId = int64

	results = append(results, tRes)

	c.HTML(http.StatusOK, "tasksedit.tmpl.html", gin.H{
		"Tasking": results,
	})
}
