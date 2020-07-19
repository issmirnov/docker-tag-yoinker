package main

import "net/http"

func init() {
	Ctx.HttpClient = &http.Client{} // FIXME: change back to real thing when done developing.
}
