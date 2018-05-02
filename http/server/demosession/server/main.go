package main

import (
	"net/http"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type cust struct {
	firstname string
	lastname  string
	mobile    string
}

var sessionMap map[string]cust
var sessionCleanList [][]string
var sessionMutex *sync.Mutex

const (
	CleanLen = 4
	CleanGap = 5
)

func init() {
	sessionMap = make(map[string]cust)
	sessionMutex = new(sync.Mutex)
	sessionCleanList = [][]string{}
	for index := 0; index < CleanLen; index++ {
		sessionCleanList = append(sessionCleanList, []string{})
	}
}

func main() {
	go func() {
		for true {
			time.Sleep(time.Second * CleanGap)
			for k, i := range sessionCleanList {
				log.WithFields(log.Fields{
					"index": k,
					"item":  len(i),
				}).Info("start")
			}
			log.WithFields(log.Fields{
				"sessionMap": len(sessionMap),
			}).Info("start")
			sessionMutex.Lock()
			for _, session := range sessionCleanList[len(sessionCleanList)-1] {
				delete(sessionMap, session)
			}

			for index := CleanLen - 1; index > 0; index-- {
				sessionCleanList[index] = sessionCleanList[index-1]
			}
			sessionCleanList[0] = []string{}
			sessionMutex.Unlock()
			for k, i := range sessionCleanList {
				log.WithFields(log.Fields{
					"index": k,
					"item":  len(i),
				}).Info("end")
			}
			log.WithFields(log.Fields{
				"sessionMap": len(sessionMap),
			}).Info("end")
		}
	}()

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {

		for index := 0; index < 1000; index++ {
			session, err := generateRandomSession()
			if err != nil {
				log.WithFields(log.Fields{
					"err": err,
				}).Error("generateRandomSession")
			}
			sessionMutex.Lock()
			sessionMap[session] = cust{"asdfs", "fweg", "12123124"}
			sessionCleanList[0] = append(sessionCleanList[0], session)
			sessionMutex.Unlock()
		}

		log.WithFields(log.Fields{}).Info("bar")
		http.Error(w, "ok", http.StatusOK)
		return
	})

	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
