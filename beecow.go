package main

import (
	"log"
	"log/syslog"
	"net/http"
)

type beeCow struct {
	Logger *log.Logger
}

func (bc beeCow) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	bc.Logger.Printf("%s -> %s %s", req.RemoteAddr, req.Method, req.URL.Path)
	if req.URL.Path != "/" {
		res.WriteHeader(http.StatusNotFound)
	} else if req.Method != "GET" {
		res.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		res.Write([]byte("<!doctype html><title>moo</title><pre style=margin:2em>(__)\n(<span id=m>oo</span>)\n \\/-------\\\n  ||     | \\\n  ||----||  *\n  ^^    ^^</pre><script>t=setTimeout;j=\"m.innerHTML='\";function b(x){z=Math.random();t(j+\"--';t(\\\"\"+j+\"oo';b(\"+(!x&z<.3)+\")\\\",100)\",(x?100:3e3)*++z)}b()</script>"))
	}
}

func main() {
	logger, err := syslog.NewLogger(syslog.LOG_LOCAL0|syslog.LOG_INFO, 0)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", beeCow{logger})
	err = http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
