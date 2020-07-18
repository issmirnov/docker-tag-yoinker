package main

import "net/http"

func init() {
	Context.HttpClient = &http.Client{} // FIXME: change back to real thing when done developing.
}
