package main

import "net/http"

type Paths struct {
	path    string
	process func(http.ResponseWriter, *http.Request)
}

func handle_paths(paths_objs []*Paths) {
	for _, v := range paths_objs {
		http.HandleFunc(v.path, v.process)
	}
}
