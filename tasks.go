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
