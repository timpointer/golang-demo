package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"

	"github.com/robfig/cron"
)

var cronrunner = cron.New()
var CRONs = []*CRONItem{
	{
		Schedule:   "@every 5s",
		Command:    "./bin/bin",
		Parameters: []string{"1", "2"},
	},
	{
		Schedule:   "@every 2m",
		Command:    "Task for every 2 minute",
		Parameters: []string{"3", "4"},
	},
}

var msg = make(chan status)

func main() {
	db := database{}
	setupCorns()
	go db.collectionStatus(msg)
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/entries", db.entries)
	http.HandleFunc("/task", db.task)
	fmt.Println("server start")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type status struct {
	id       string
	progress int
	state    string
}

//!-main
type database map[string]status

func (db database) entries(w http.ResponseWriter, req *http.Request) {
	entries := cronrunner.Entries()
	data := struct {
		Items []*cron.Entry
	}{
		entries,
	}
	err := entryList.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	data := struct {
		Items database
	}{
		db,
	}
	err := reportList.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
}

func (db database) task(w http.ResponseWriter, req *http.Request) {
	values := req.URL.Query()
	id := values.Get("id")
	if id == "" {
		http.Error(w, "bad query", http.StatusBadRequest)
		return
	}
	switch rand.Intn(5) {
	case 0:
		task(id, msg, 100)
	case 1, 2, 3, 4:
		task(id, msg, 10)
	default:
		log.Fatal("out bound")
	}
}
func (db database) collectionStatus(in <-chan status) {
	for {
		status := <-in
		db[status.id] = status
	}
}
func (db database) printStatus() {
	for key, value := range db {
		var pg string
		for i := 0; i <= value.progress; i++ {
			pg += "*"
		}
		fmt.Printf("task %s ,progress %s\n", key, pg)
	}
}

var reportList = template.Must(template.New("reportlist").Parse(`
<html>
<body>
<h1>Status</h1>
<table>
<tr style='text-align: left'>
  <th>Key</th>
  <th>State</th>
</tr>
{{range $i, $v := .Items}}
<tr>
  <td>{{$i}}</td>
  <td>{{$v}}</td>
</tr>
{{end}}
</table>
</body>
</html>
`))

var entryList = template.Must(template.New("entrylist").Parse(`
<html>
<body>
<h1>Status</h1>
<table>
<tr style='text-align: left'>
  <th>Key</th>
  <th>Next</th>
  <th>Prev</th>
  <th>Job</th>
</tr>
{{range  .Items}}
<tr>
  <td>{{.Schedule}}</td>
  <td>{{.Next}}</td>
  <td>{{.Prev}}</td>
  <td>{{.Job}}</td>
</tr>
{{end}}
</table>
</body>
</html>
`))
