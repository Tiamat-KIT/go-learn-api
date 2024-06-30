package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// https://qiita.com/mugi111/items/9063e7c6d9e86164d6c3#verceljson%E3%81%AE%E4%BD%9C%E6%88%90

type Paths struct {
	path    string
	process func(http.ResponseWriter, *http.Request)
}

func handle_paths(paths_objs []*Paths) {
	for _, v := range paths_objs {
		http.HandleFunc(v.path, v.process)
	}
}

func main() {
	log.Print("starting server...")
	paths_objs := []*Paths{
		{
			"/",
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Hello World!\n")
			},
		}
		{
			"/hello",
			func(w http.ResponseWriter, r *http.Request) {
				name := os.Getenv("NAME")
				if name == "" {
					name = "World"
				}
				fmt.Fprintf(w, "Hello %s!\n", name)
			},
		},
		/* {
			"/employee/new",
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Process!\n")
			},
		},
		{
			"/employee/delete",
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Process!\n")
			},
		},
		{
			"/employee/edit",
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Process!\n")
			},
		},
		{
			"/employee/read",
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Process!\n")
			},
		},
		{
			"/shift/new",
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Process!\n")
			},
		},
		{
			"/shift/delete",
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Process!\n")
			},
		},
		{
			"/shift/edit",
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Process!\n")
			},
		},
		{
			"/shift/read",
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Process!\n")
			},
		}, */
		{
			"/view/request",
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "Method:", r.Method)
				// fmt.Fprintln(w, "URL:", r.URL.String())
				fmt.Fprintln(w, "Proto:", r.Proto)
				fmt.Fprintln(w, "Header:", r.Header)
				fmt.Fprintln(w, "Body:", r.Body)
			},
		},
	}

	handle_paths(paths_objs)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err, "\n")
	}
}
