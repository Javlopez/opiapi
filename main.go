package main

import (
	"flag"
	"fmt"
	"github.com/Javlopez/opiapi/app"
	"github.com/Javlopez/opiapi/domain"
	"github.com/Javlopez/opiapi/http/handlers"
	"github.com/Javlopez/opiapi/infrastructure"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

const (
	PORT string = "8005"
)

func main() {

	container := infrastructure.Container{}
	apiApp := &app.OpiAppContext{Container: container}

	if len(os.Args) >= 2 {
		LoadCsvFile(container)
		return
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = PORT
	}

	r := mux.NewRouter()
	r.NewRoute().
		Path("/api/v1/points").
		Handler(app.AppHandler{apiApp, handlers.PointHandler}).
		Methods("GET")

	fmt.Printf("Running api on PORT:%s\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), r))
}

func LoadCsvFile(container infrastructure.Container) {
	csvFile := flag.String("file", "", "Load file from csv to database")
	flag.Parse()

	if *csvFile == "" {
		log.Fatal("Error: you need to provide some csv file: --file=data.csv")
	}

	err := container.DB().AutoMigrate(&domain.Point{})
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	points, err := container.CsvService().ParseCSVtoPoint(*csvFile, true)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	err = container.DBPointRepository().SavePoints(points)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	fmt.Printf("The file %s was procesed successful", *csvFile)
}
