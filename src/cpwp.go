package main

import (
	"fmt"
	"sync"
	"common"
	"runtime"
	"net/http"
	"crypto/tls"
	"http/router"
	"http/router/api"
	"http/middleware"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	defer common.Logger.Print("[Server] Stop")

	// set core
	runtime.GOMAXPROCS(runtime.NumCPU())

	// log
	common.InitalizeLogger()
	common.Logger.Printf("[Server Initialization] Number of cores to use : %v", runtime.GOMAXPROCS(0))
	waitSignal := &sync.WaitGroup{}

	// http
	waitSignal.Add(1)
	go func(){
		/******************************************************************
		 ROUTER INITIALIZATION
		 ******************************************************************/
		r := mux.NewRouter()

		// Index - Redirect To Prod
		r.Handle("/", middleware.SetMiddleware(http.HandlerFunc(router.GetIndex))).Methods("GET")
		// Favicon
		r.Handle("/favicon.ico", middleware.SetMiddleware(http.HandlerFunc(router.GetFavicon))).Methods("GET")
		// Static
		r.PathPrefix("/static/").Handler(middleware.SetMiddleware(http.StripPrefix("/static/", http.FileServer(http.Dir(fmt.Sprintf("%v/static/", common.BasePath)))))).Methods("GET")
		// Angular Prod
		r.PathPrefix("/prod/").Handler(middleware.SetMiddleware(http.StripPrefix("/prod/", http.FileServer(http.Dir(fmt.Sprintf("%v/angular/prod/", common.BasePath)))))).Methods("GET")

		/******************************************************************
		 API
		 ******************************************************************/
		// Status
		r.Handle("/api/status", middleware.SetMiddleware(http.HandlerFunc(api.GetStatus))).Methods("GET")

		/******************************************************************
		 ERROR
		 ******************************************************************/
		// Not Found
		r.NotFoundHandler = http.HandlerFunc(router.NotFound)

		/******************************************************************
		 SERVE by STAGING
		 ******************************************************************/
		common.Logger.Printf("[Server STAGING] %v", common.Staging)
		if common.Staging == "real"{
			// SSL/TLS
			certManager := autocert.Manager{
				Prompt: autocert.AcceptTOS,
				Cache:  autocert.DirCache("var/cert"),
			}
			server := &http.Server{
				Addr: fmt.Sprintf(":%v", common.HttpsPort),
				Handler: r,
				TLSConfig: &tls.Config{
					GetCertificate: certManager.GetCertificate,
				},
			}
			go http.ListenAndServe(fmt.Sprintf(":%v", common.HttpPort), certManager.HTTPHandler(nil))
			common.Logger.Fatal(server.ListenAndServeTLS("", ""))
		} else {
			server := &http.Server{
				Addr: fmt.Sprintf(":%v", common.HttpPort),
				Handler: r,
			}
			common.Logger.Fatal(server.ListenAndServe())
		}
		waitSignal.Done()
	}()
	waitSignal.Wait()
}