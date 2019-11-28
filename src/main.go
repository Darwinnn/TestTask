package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/darwinnn/TestTask/src/controllers"
	"github.com/darwinnn/TestTask/src/db"
	"github.com/darwinnn/TestTask/src/restapi"
	"github.com/darwinnn/TestTask/src/restapi/operations"
	"github.com/darwinnn/TestTask/src/workers"
	"github.com/go-openapi/loads"
)

var (
	dbConnString string
	tMin         int
	tOddNum      int
)

func init() {
	var err error
	if dbConnString = os.Getenv("DB_CONN_STRING"); dbConnString == "" {
		log.Fatal("env DB_CONN_STRING is missing")
	}
	if tMinVal := os.Getenv("T_MINUTES"); tMinVal != "" {
		tMin, err = strconv.Atoi(tMinVal)
	} else if err != nil || tMinVal == "" {
		log.Fatal("env T_MINUTES is missing")
	}
	if tOddNumVal := os.Getenv("T_ODD_NUM"); tOddNumVal != "" {
		tOddNum, err = strconv.Atoi(tOddNumVal)
	} else if err != nil || tOddNumVal == "" {
		log.Fatal("env T_ODD_NUM is missing")
	}
}

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "2.0")
	if err != nil {
		log.Fatalf("can't get swagger spec: %v", err)
	}
	dbh, err := db.Init(dbConnString)
	if err != nil {
		log.Fatalf("can't init db: %v", err)
	}
	api := operations.NewTestTaskAPI(swaggerSpec)
	server := restapi.NewServer(api)
	server.Port = 8080
	controllers.Init(api, dbh)
	server.ConfigureAPI()

	canceller := workers.Init(dbh, log.Printf, time.Duration(tMin)*time.Minute, tOddNum)
	go canceller.Work()
	if err := server.Serve(); err != nil {
		log.Fatalf("Server stopped: %v", err)
	}
}
