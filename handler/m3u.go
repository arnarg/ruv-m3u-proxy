package handler

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/arnarg/ruv-m3u-proxy/logging"
)

type M3uHandler struct {
	Prefix string
}

type TmplParams struct {
	Prefix string
}

func NewM3uHandler(prefix string) M3uHandler {
	return M3uHandler{Prefix: prefix}
}

func (h M3uHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logging.LogRequest(r)
	w.Header().Set("Content-Type", "audio/x-mpegurl")

	m3u := `#EXTM3U
#EXTINF:-1 tvg-id="" tvg-name="RUV" tvg-logo="{{.Prefix}}/static/ruv.png" group-title="",RUV
{{.Prefix}}/ruv.m3u8
#EXTINF:-1 tvg-id="" tvg-name="RUV 2" tvg-logo="{{.Prefix}}/static/ruv2.png" group-title="",RUV 2
{{.Prefix}}/ruv2.m3u8`

	tmplParams := TmplParams{}
	if h.Prefix != "" {
		tmplParams.Prefix = h.Prefix
	} else {
		tmplParams.Prefix = fmt.Sprintf("http://%s", r.Host)
	}
	tmpl, err := template.New("m3u").Parse(m3u)
	if err != nil {
		w.WriteHeader(500)
	}

	err = tmpl.Execute(w, tmplParams)
	if err != nil {
		w.WriteHeader(500)
	}
}
