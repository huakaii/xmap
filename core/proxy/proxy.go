package proxy

import (
	"flag"
	"fmt"
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
	"net/http/httputil"
)

func StartProxy() {
	verbose := flag.Bool("v", false, "should every proxy request be logged to stdout")
	flag.Parse()
	_ = SetCA()
	proxy := goproxy.NewProxyHttpServer()
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
	proxy.Verbose = *verbose

	proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {

			buf, _ := httputil.DumpRequest(r, false)
			// TODO send data to scanner
			fmt.Println(string(buf))
			return r, nil
		})

	log.Fatal(http.ListenAndServe(":8888", proxy))

}
