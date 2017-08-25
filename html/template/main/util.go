package main

import (
	"bytes"
	"html/template"
	"log"
	"sync"
	"time"
)

// Util is toolset
type Util struct {
	rwmutex   sync.RWMutex
	templates map[string]string
}

var util = &Util{
	templates: make(map[string]string),
}

func (Util) isInTimeStr(t time.Time, fromDate, untilDate string) (bool, error) {
	validFrom, err := time.ParseInLocation("2006.01.02", fromDate, time.Local)
	if err != nil {
		return false, err
	}
	validUntil, err := time.ParseInLocation("2006.01.02", untilDate, time.Local)
	//log.Println(t, validFrom, validUntil)
	if err != nil {
		return false, err
	}
	if t.Before(validFrom) {
		return false, nil
	}
	if t.After(validUntil.AddDate(0, 0, 1)) {
		return false, nil
	}
	return true, nil
}

func (Util) isInTime(t time.Time, fromDate, untilDate time.Time) bool {
	if t.Before(fromDate) {
		return false
	}
	if t.After(untilDate.AddDate(0, 0, 1)) {
		return false
	}
	return true
}

func (Util) isInArray(list []string, item string) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}
	return false
}

func (Util) getInvoice(in string) string {
	rs := []rune(in)
	if len(rs) > 6 {
		start := len(rs) - 6
		return string(rs[start:])
	}
	return in
}

func (u *Util) GetCouponInfo(name string) string {
	u.rwmutex.RLock()
	if t, ok := u.templates[name]; ok == true {
		u.rwmutex.RUnlock()
		return t
	}
	u.rwmutex.RUnlock()

	u.rwmutex.Lock()
	u.templates[name] = getCouponInfo(name)
	t := u.templates[name]
	u.rwmutex.Unlock()
	return t
}

func getCouponInfo(name string) string {
	path := "./template/" + name + ".html"
	log.Printf("[INFO] getcouponinfo parse path  %s", path)
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		log.Printf("[ERR] getcouponinfo parse path %s:%v", path, err)
		return ""
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, nil)
	if err != nil {
		log.Printf("[ERR] getcouponinfo execute name %s:%v", name, err)
		return ""
	}
	return buf.String()
}
