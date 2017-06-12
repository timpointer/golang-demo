package main

import (
	"fmt"
	"time"

	ttime "github.com/timpointer/golang-demo/time"
)

func main() {
	// Date
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())

	// Duration
	t0 := time.Now()
	time.Sleep(time.Second)
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))

	// Month
	_, month, day := time.Now().Date()
	if month == time.November && day == 10 {
		fmt.Println("Happy Go day!")
	}

	// Parse
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ = time.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)")
	fmt.Println(t)

	const shortForm = "2006-Jan-02"
	t, _ = time.Parse(shortForm, "2013-Feb-03")
	fmt.Println(t)
	year, week := t.ISOWeek()
	fmt.Printf("year %d week %d\n", year, week)

	t = ttime.FirstDayOfISOWeek(2017, 2, time.UTC)
	fmt.Println(t)

	unittime := t.Unix()
	fmt.Printf("unitime %d\n", unittime)

	t = t.AddDate(0, 0, 7)
	fmt.Println(t)

	unittime = t.Unix()
	fmt.Printf("unitime %d\n", unittime)
}
