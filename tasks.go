package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
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
	tid := c.Param("tid")

	var (
		id           int64
		task_name    string
		task_type    string
		weighting    string
		time_when    time.Time
		time_offset  int64
		time_float   bool
		time_length  int64
		no_of_jobs   int64
		time_min     int64
		time_max     int64
		day_of_week  string
		active       bool
		companies_id int64
		stores_id    int64
		dept_id      int64
	)

	err := db.QueryRow("SELECT id, task_name, task_type, weighting, time_when, time_offset, time_float, time_length, no_of_jobs, time_min, time_max, day_of_week, active, companies_id, stores_id, dept_id FROM tasks WHERE id = $1", tid).Scan(&id, &task_name, &task_type, &weighting, &time_when, &time_offset, &time_float, &time_length, &no_of_jobs, &time_min, &time_max, &day_of_week, &active, &companies_id, &stores_id, &dept_id)

	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error reading questions: %q", err))
		return
	}

	//defer rows.Close()
	results := []Tasks{}
	tRes := Tasks{}

	tRes.Id = id
	tRes.TaskName = task_name
	tRes.TaskType = task_type
	tRes.Weighting = weighting
	tRes.TimeWhen = time_when
	tRes.Time_Offset = time_offset
	tRes.TimeFloat = time_float
	tRes.TimeLength = time_length
	tRes.NoOfJobs = no_of_jobs
	tRes.TimeMin = time_min
	tRes.TimeMax = time_max
	tRes.DayOfWeek = day_of_week
	tRes.Active = active
	tRes.CompaniesId = companies_id
	tRes.StoresId = stores_id
	tRes.DeptId = dept_id

	results = append(results, tRes)

	c.HTML(http.StatusOK, "tasksedit.tmpl.html", gin.H{
		"Tasking": results,
	})
}

//	-    -    -    -    -    -    -
// tasks update page
func postTasksupdate(c *gin.Context) {

	// get the current time
	updated_on := time.Now().Local()
	// get all the post vars
	id := c.PostForm("Id")
	task_name := c.PostForm("TaskName")
	task_type := c.PostForm("TaskType")
	weighting := c.PostForm("Weighting")
	time_when := c.PostForm("TimeWhen")
	time_offset := c.PostForm("Time_Offset")
	time_float := c.PostForm("TimeFloat")
	time_length := c.PostForm("TimeLength")
	no_of_jobs := c.PostForm("NoOfJobs")
	time_min := c.PostForm("TimeMin")
	time_max := c.PostForm("TimeMax")
	day_of_week := c.PostForm("DayOfWeek")
	active := c.PostForm("Active")
	companies_id := c.PostForm("CompaniesId")
	stores_id := c.PostForm("StoresId")
	dept_id := c.PostForm("DeptId")

	stmt, err := db.Prepare("UPDATE staff SET task_name = $1 ,task_type = $2, weighting = $3, time_when = $4, time_offset = $5, time_float = $6, time_length = $7, no_of_jobs = $8, time_min = $9, time_max = $10, day_of_week = $11, active = $12, companies_id = $13, stores_id = $14, dept_id  = $15, updated_on = $16 WHERE id = $17")
	fmt.Println("........................")
	fmt.Println(stmt)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(task_name, task_type, weighting, time_when, time_offset, time_float, time_length, no_of_jobs, time_min, time_max, day_of_week, active, companies_id, stores_id, dept_id, updated_on, id)

	if err != nil {
		log.Fatal(err)
	}

	c.HTML(http.StatusOK, "tasksupdate.tmpl.html", gin.H{
		"taskname": task_name,
		"res":      res,
	})

}
