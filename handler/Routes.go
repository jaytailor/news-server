package handler

import (
	"net/http"
)

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"GetAllNews",
		"GET",
		"/getallnews",
		GetAllNews,
	},
	Route{
		"PostNews",
		"POST",
		"/postnews",
		PostNews,
	},
}


