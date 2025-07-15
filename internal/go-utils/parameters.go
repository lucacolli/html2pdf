package otutil

import (
	"log"
	"net/http"
	"strconv"
)

// getPerPage returns the per_page param from the url query.
func GetPerPage(req *http.Request) int {
	perPage := GetQueryInt(req, "per_page")
	if perPage == 0 {
		return 50
	}
	if perPage > 100 {
		return 100
	}
	return perPage
}

func GetLimit(req *http.Request) int {
	limit := GetQueryInt(req, "limit")
	if limit == 0 {
		return 50
	}
	if limit > 1000 {
		return 1000
	}
	return limit
}

func GetQueryInt(req *http.Request, key string) int {
	vals := req.URL.Query()
	val := vals[key]
	if val != nil {
		v, err := strconv.ParseInt(val[0], 10, 0)
		if err != nil {
			log.Println(err)
			return 0
		}
		return int(v)
	}
	return 0
}

func GetQueryString(req *http.Request, key string) string {
	vals := req.URL.Query()
	val := vals[key]
	if val != nil {
		return val[0]
	}
	return ""
}

// FIXME: In case the parameter is false or the parameter is missing the result is the same
func GetQueryBool(req *http.Request, key string) bool {
	vals := req.URL.Query()
	val := vals[key]
	if val != nil {
		if val[0] == "true" {
			return true
		}
	}
	return false
}
