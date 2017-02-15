package controllers

import (
    "app/routes/middleware/pprofhandler"
    "app/commons/router"
)

func init() {
    // Enable Pprof
    router.Get("/debug/pprof/*pprof", pprofhandler.Handler)
}
