package main

import (
	"log"
	"log/syslog"
	"net/http"
	"time"
)

type beeCow struct {
	Logger   *log.Logger
	Location *time.Location
}

func (bc beeCow) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	bc.Logger.Printf("%s -> %s %s", req.RemoteAddr, req.Method, req.URL.Path)
	if req.URL.Path != "/" {
		res.WriteHeader(http.StatusNotFound)
	} else if req.Method != "GET" {
		res.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		hour := time.Now().In(bc.Location).Hour()
		if hour < 22 && hour > 5 {
			res.Write([]byte("<!doctype html><title>moo</title><pre style=margin:2em>(__)\n(<span id=m>oo</span>)\n \\/-------\\\n  ||     | \\\n  ||----||  *\n  ^^    ^^</pre><script>t=setTimeout;j=\"m.innerHTML='\";function b(x){z=Math.random();t(j+\"--';t(\\\"\"+j+\"oo';b(\"+(!x&z<.3)+\")\\\",100)\",(x?100:3e3)*++z)}b()</script>"))
		} else {
			res.Write([]byte("<!doctype html><title>moo</title><pre style=margin:2em>    <span id=z>ZzZ</span>\n(__)\n(--)-----\\\n \\/     ||\n  \\\\---//*\n   ^^ ^^</pre><script>j='z.innerHTML';setInterval(j+'='+j+'[1]+'+j+'[0]+'+j+'[1]',500)</script>"))
		}
	}
}

func main() {
	logger, err := syslog.NewLogger(syslog.LOG_LOCAL0|syslog.LOG_INFO, 0)
	if err != nil {
		log.Fatal(err)
	}
	location, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", beeCow{logger, location})
	err = http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
