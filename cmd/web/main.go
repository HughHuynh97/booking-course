package main

import (
	"github.com/HughHuynh97/booking-course/pkg/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	_ = http.ListenAndServe(":8080", nil)
}
