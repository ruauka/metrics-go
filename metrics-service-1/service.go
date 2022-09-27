package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/goccy/go-json"
	"github.com/gorilla/handlers"
	"github.com/julienschmidt/httprouter"

	"metrics-service-1/internal/logic"
	"metrics-service-1/internal/request"
	"metrics-service-1/internal/response"
	"metrics-service-1/middleware"
)

// Execute - основная функция скрипта.
func Execute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Парсинг входящего JSON
	req := &request.Request{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	// Валидация объекта структуры Request.
	if err := request.ValidateStruct(req); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	// создание словаря данных.
	data := logic.NewData(req)
	// вызов методов логики.
	data.LocalCount()
	data.ResultCount()
	// создание объекта ответа.
	resp := response.NewResponse(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp) //nolint:errcheck
}

func main() {
	// Create metrics
	metrics := middleware.NewMetrics()
	// Create middlewares for lambda handler
	metricsMiddleware := middleware.NewMetricsMiddleware(metrics)

	router := httprouter.New()
	router.POST("/execute", metricsMiddleware(Execute))

	// Get prometheus metrics
	router.GET("/metrics", middleware.NewMetricsHandler(metrics))

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	log.Println("Starting server...")

	log.Fatal(http.ListenAndServe(":8000", loggedRouter))
}
