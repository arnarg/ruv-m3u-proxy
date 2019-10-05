package main

import (
	"flag"
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/arnarg/ruv-m3u-proxy/handler"
)

func main() {
	prefix := flag.String("prefix", "", "Prefix string to put in urls in index.m3u")
	flag.Parse()

	// Handle static files
	http.Handle("/static/", staticFileServer("/static/"))

	http.Handle("/index.m3u", handler.NewM3uHandler(*prefix))

	http.Handle("/ruv.m3u8", handler.NewRuvHandler("https://geo.spilari.ruv.is/channel/ruv"))
	http.Handle("/ruv2.m3u8", handler.NewRuvHandler("https://geo.spilari.ruv.is/channel/ruv2"))

	http.ListenAndServe(":8080", nil)
}

func staticFileServer(prefix string) http.Handler {
	box := rice.MustFindBox("static")
	return http.StripPrefix(prefix, http.FileServer(box.HTTPBox()))
}
