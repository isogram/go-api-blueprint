package controllers

import (
    "net/http"

    "app/commons/response"
    "app/commons/router"
)

var errs [1]Err

type Err struct {
    Key string `json:"key"`
    Value string `json:"value"`
}

func init() {
    // This does not work for routes where the path matches, but the method does not (on HEAD and OPTIONS need to check)
    // https://github.com/julienschmidt/httprouter/issues/13
    var e405 http.HandlerFunc = Error405
    router.Instance().HandleMethodNotAllowed = true
    router.Instance().MethodNotAllowed = e405

    // 404 Page
    var e404 http.HandlerFunc = Error404
    router.Instance().NotFound = e404
}

// Error404 - Page Not Found
func Error404(w http.ResponseWriter, r *http.Request) {
    errs[0].Key = "endpoint"
    errs[0].Value = "Not Found!"

    response.SendError(w, http.StatusNotFound, "Not found!", errs)
}

// Error405 - Method Not Allowed
func Error405(w http.ResponseWriter, r *http.Request) {
    errs[0].Key = "endpoint"
    errs[0].Value = "Method not allowed!"

    response.SendError(w, http.StatusMethodNotAllowed, "Method not allowed!", errs)
}

// Error500 - Internal Server Error
func Error500(w http.ResponseWriter, r *http.Request) {
    errs[0].Key = "endpoint"
    errs[0].Value = "Internal server error!"

    response.SendError(w, http.StatusInternalServerError, "Internal server error!", errs)
}
