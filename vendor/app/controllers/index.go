package controllers

import (
    "net/http"

    "app/commons/response"
    "app/commons/router"
)

func init() {
    // Main page
    router.Get("/", Index)
}

func Index(w http.ResponseWriter, r *http.Request) {
    response.SendOK(w, http.StatusOK, "Hi there!", [0]int{} ,struct{}{})
}
