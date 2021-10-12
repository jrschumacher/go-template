package api

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/jrschumacher/go-template/docs"
	"github.com/jrschumacher/go-template/internal/conf"
	storePkg "github.com/jrschumacher/go-template/internal/store"
	"github.com/mcuadros/go-defaults"
	muxlogrus "github.com/pytimer/mux-logrus"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
)

type ServeParams struct {
	StoreLocation string
	// Default will be a empty string, this is ideal
	Host string `default:""`
	// Default will be
	Port int `default:"8080"`
}

type ErrorResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

func NewServeParams() *ServeParams {
	params := new(ServeParams)
	defaults.SetDefaults(params)
	return params
}

var store storePkg.Store

func SetupStore(storeLocation string) {
	var err error
	store = storePkg.OpenStore(storeLocation)
	if err != nil {
		logrus.WithError(err).Fatal("could not initialize database")
	}
}

func SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Closing store")
		store.Close()
		os.Exit(0)
	}()
}

// @title Go Template API
// @description A go template project
func Serve(params *ServeParams) {
	docs.SwaggerInfo.Version = conf.Version

	SetupStore(params.StoreLocation)

	// Setup our Ctrl+C handler
	SetupCloseHandler()

	// Define server routes
	addr := params.Host + ":" + strconv.Itoa(params.Port)

	// Create new router
	r := mux.NewRouter()
	// Add middleware
	r.Use(
		muxlogrus.NewLogger().Middleware,
		mux.CORSMethodMiddleware(r),
	)

	// Create root handler
	r.HandleFunc("/", rootHandler(r)).Methods("GET")
	// Create health handler
	r.HandleFunc("/health", healthHandler).Methods("GET")
	r.HandleFunc("/version", versionHandler).Methods("GET")
	r.PathPrefix("/swagger").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), //The url pointing to API definition
	))

	// Routes
	r.HandleFunc("/fetch/{id}", fetchHandler).Methods("GET")
	// Used to handle optional queries
	r.HandleFunc("/search", searchHandler).Methods("GET")
	// Define queries will expect all queries
	r.HandleFunc("/search", searchHandler).
		Queries("limit", "{limit}").
		Queries("offset", "{offset}").
		Methods("GET")
	r.HandleFunc("/write", writeHandler).Methods("POST")
	r.HandleFunc("/update/{id}", updateHandler).Methods("PUT")
	r.HandleFunc("/delete/{id}", deleteHandler).Methods("DELETE")

	// Define handler path
	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    addr,
		// // Good practice: enforce timeouts for servers you create!
		WriteTimeout: time.Second * 60,
		ReadTimeout:  time.Second * 60,
	}

	logrus.Infof("Listening on %s", addr)
	logrus.Fatal(srv.ListenAndServe())
}
