package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os/exec"
	"path"
	"time"

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
	http.HandleFunc("/execute", db.execute)
	http.HandleFunc("/task", db.task)
	fmt.Println("server start")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type status struct {
	id       string
	progress int
	state    string
	last     string
}

//!-main
type database map[string]*status

func (db database) execute(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	command := query.Get("exe")
	command = path.Join(".", "bin", command)
	log.Printf("[INFO] - Module:[report] -> command execute: %s param: %s", command, query.Encode())

	go func() {
		t := time.Now()
		cmd := exec.Command(command)
		stdout, err := cmd.StdoutPipe()
		stderr, err := cmd.StderrPipe()
		err = cmd.Start()
		var stat = &status{}
		defer func() {
			if err != nil {
				log.Println("[ERR] - Module:[report] -> :", err)
				stat.id = command
				stat.state = command
			} else {
				stat.id = command
				stat.progress = 100
				stat.state = command
			}
			stat.last = time.Since(t).String()
			db[command] = stat
		}()
		if err != nil {
			return
		}
		input := bufio.NewScanner(stdout)
		for input.Scan() {
			stat.state = input.Text()
			db[command] = stat
		}

		errContent, err := ioutil.ReadAll(stderr)
		if err != nil {
			return
		}
		if err = cmd.Wait(); err != nil {
			err = fmt.Errorf(" %v: %s", err, errContent)
			return
		}
	}()
}
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
		db[status.id] = &status
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
