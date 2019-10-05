ruv-m3u-proxy:
	go get github.com/GeertJohan/go.rice/rice
	CGO_ENABLED=0 go build -o ruv-m3u-proxy
	`go env GOPATH`/bin/rice append --exec ruv-m3u-proxy

