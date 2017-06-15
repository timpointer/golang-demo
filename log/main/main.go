package main

import "log"
import "os"

func main() {
	debugLog := log.New(os.Stdout, "[Debug]", log.Lshortfile|log.Ldate|log.Ltime)
	debugLog.Print("sdf")
	log.Print("wwg", "ergegr")
	log.Fatal("sdf", "ger", "gdfg")
	log.Print("end wwg", "ergegr")
}
