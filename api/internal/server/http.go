package server

import (
	"api/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
)

// HttpServer is an interface defining HTTP server methods for handling data-related requests
type HttpServer interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetStatistic(w http.ResponseWriter, r *http.Request)
}

// httpServer is an implementation of the HttpServer interface.
type httpServer struct {
	service serv.DataService
}

// NewHttpServer creates a new HttpServer instance with the provided DataService.
func NewHttpServer(service serv.DataService) HttpServer {
	return &httpServer{service: service}
}

// GetAll handles HTTP requests for retrieving all data.
func (h httpServer) GetAll(w http.ResponseWriter, r *http.Request) {
	var err error

	// Retrieve search text from the query parameter.
	searchText := r.URL.Query().Get("text")

	// Retrieve pagination parameters (from, size).
	from, size, err := h.getPagination(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Call the DataService to get data based on the provided parameters.
	data, err := h.service.GetAllData(searchText, int32(from), int32(size))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set HTTP headers and write the response with the retrieved data.
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(data)
}

// GetStatistic handles HTTP requests for retrieving statistics.
func (h httpServer) GetStatistic(w http.ResponseWriter, _ *http.Request) {
	// Call the DataService to get statistics.
	data, err := h.service.GetStatistic()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set HTTP headers and write the response with the retrieved statistics.
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(data)
}

// getPagination retrieves pagination parameters (from, size) from the request URL.
func (h httpServer) getPagination(r *http.Request) (from int, size int, err error) {
	// Retrieve "from" parameter.
	if fromStr := r.URL.Query().Get("from"); fromStr != "" {
		from, err = strconv.Atoi(fromStr)
		if err != nil {
			return 0, 0, err
		}
	}

	// Retrieve "size" parameter or use default value (10).
	if sizeStr := r.URL.Query().Get("size"); sizeStr != "" {
		size, err = strconv.Atoi(sizeStr)
		if err != nil {
			return 0, 0, err
		}
	} else {
		size = 10
	}

	return from, size, nil
}
