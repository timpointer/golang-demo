package main

import (
	"fmt"
	"time"

	ttime "github.com/timpointer/golang-demo/time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Format("Jan 2, 2006 at 3:04pm (MST)"))
	timestamp := time.Unix(0, 0)
	fmt.Println("timestamp", timestamp)
	fmt.Println("timestamp is zero", timestamp.IsZero())
	timestamp2 := time.Time{}
	fmt.Println("timestamp2", timestamp2)
	fmt.Println("timestamp2 is zero", timestamp2.IsZero())
	date := time.Date(2016, 1, 0, 0, 0, 0, 0, time.Local)
	fmt.Println("date", date)
	Date := "20170203"
	Time := "170203"
	datetime := Date + Time
	t, _ := time.Parse("20060102150405", datetime)
	fmt.Println("timestamp3", t)
	fmt.Println("timestamp3", t.Format("2006年1月2日 15:04"))

	oldtime := time.Time{}
	fmt.Println("time.Time{}", oldtime)
	fmt.Println("timestamp4", time.Now().Format("20060102"))
	fmt.Println("timestamp4", time.Now().Add(-time.Hour*8).Format("20060102"))
	// Date
	t = time.Date(2009, time.November, 10, 23, 3, 6, 2, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())
	tu := time.Unix(t.Unix(), 0)
	fmt.Printf("unix %v", tu)
	//

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
	fmt.Println("**********************")
	// Parse
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ = time.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)")
	fmt.Println(t)

	const shortForm = "2006-Jan-02"
	t, _ = time.Parse(shortForm, "2013-Feb-03")
	fmt.Println(t)
	year, week := t.ISOWeek()
	fmt.Printf("year %d week %d\n", year, week)

	fmt.Println("*****************************")
	// shortForm
	const shortForm2 = "20060102"
	t, err := time.Parse(shortForm2, "20133203")
	if err != nil {
		fmt.Printf("error:%v\n", err)
	}
	fmt.Println("shortform2:", t)

	t = ttime.FirstDayOfISOWeek(2017, 2, time.UTC)
	fmt.Println(t)

	unittime := t.Unix()
	fmt.Printf("unitime %d\n", unittime)

	t = t.AddDate(0, 0, 3322)
	fmt.Println(t)

	unittime = t.Unix()
	fmt.Printf("unitime %d\n", unittime)
}
