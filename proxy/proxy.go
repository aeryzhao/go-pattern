package main

import "fmt"

type server interface {
	handleRequest(string, string) (int, string)
}

type nginx struct {
	app         *application
	maxAllowed  int
	rateLimiter map[string]int
}

func newNginxServer() *nginx {
	return &nginx{
		app:         &application{},
		maxAllowed:  2,
		rateLimiter: make(map[string]int),
	}
}

func (n *nginx) handleRequest(url, method string) (int, string) {
	if !n.checkRateLimiting(url) {
		return 403, "Not Allowed"
	}
	return n.app.handleRequest(url, method)
}

func (n *nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.maxAllowed {
		return false
	}
	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}

type application struct{}

func (a *application) handleRequest(url, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "Ok"
	}
	if url == "/create/user" && method == "POST" {
		return 201, "User Created"
	}
	return 404, "Not Found"
}

func main() {
	server := newNginxServer()
	appStatusURL := "/app/status"
	createuserURL := "/create/user"
	urls := []string{appStatusURL, createuserURL, appStatusURL, createuserURL, appStatusURL, createuserURL}
	for _, url := range urls {
		code, message := server.handleRequest(url, "GET")
		println(fmt.Sprintf("Response for URL %s with code %d and message %s", url, code, message))
	}
}
