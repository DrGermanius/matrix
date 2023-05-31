package main

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

var (
	errInvalidFileFormatCSV = errors.New("invalid file format, only CSV allowed")
	errNotSquareMatrix      = errors.New("matrix should be square")
	errEmptyRecord          = errors.New("matrix shouldn't be empty")
)

type Handler struct {
	mux *http.ServeMux
}

func NewHandler() *http.ServeMux {
	handler := Handler{}
	mux := http.NewServeMux()
	mux.Handle("/echo", getRecordsMiddleware(handler.Echo))
	mux.Handle("/invert", getRecordsMiddleware(handler.Invert))
	mux.Handle("/multiply", getRecordsMiddleware(handler.Multiply))
	mux.Handle("/flatten", getRecordsMiddleware(handler.Flatten))
	mux.Handle("/sum", getRecordsMiddleware(handler.Sum))
	return mux
}

func (Handler) Echo(w http.ResponseWriter, r *http.Request) {
	records := getRecordsFromCtx(r.Context())
	fmt.Fprint(w, matrixToString(records))
}

func (Handler) Invert(w http.ResponseWriter, r *http.Request) {
	records := getRecordsFromCtx(r.Context())
	records = invertMatrix(records)
	fmt.Fprint(w, matrixToString(records))
}

func (Handler) Flatten(w http.ResponseWriter, r *http.Request) {
	records := getRecordsFromCtx(r.Context())
	fmt.Fprint(w, matrixToFlatString(records))
}

func (Handler) Sum(w http.ResponseWriter, r *http.Request) {
	records := getRecordsFromCtx(r.Context())
	sum, err := sumIntMatrix(records)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Fprint(w, sum)
}

func (Handler) Multiply(w http.ResponseWriter, r *http.Request) {
	records := getRecordsFromCtx(r.Context())
	sum, err := multiplyIntMatrix(records)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Fprint(w, sum)
}

func getRecordsFromCtx(ctx context.Context) [][]string {
	return ctx.Value(recordsKey).([][]string)
}

func getRecordsMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		records, err := readMultipartCsvFile(w, r, multipartFileKey)
		if err != nil {
			// http.Error call inside readMultipartCsvFile
			log.Println(err)
			return
		}
		if !isMatrixSquare(records) {
			http.Error(w, errNotSquareMatrix.Error(), http.StatusBadRequest)
			return
		}
		ctxWithRecords := context.WithValue(r.Context(), "records", records)
		handler.ServeHTTP(w, r.WithContext(ctxWithRecords))
	}
}

func readMultipartCsvFile(w http.ResponseWriter, r *http.Request, key string) ([][]string, error) {
	file, fileheader, err := r.FormFile(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	defer file.Close()

	// check the file extension
	ext := filepath.Ext(fileheader.Filename)
	if ext != csvExtension {
		http.Error(w, errInvalidFileFormatCSV.Error(), http.StatusBadRequest)
		return nil, errInvalidFileFormatCSV
	}

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	if records == nil {
		http.Error(w, errEmptyRecord.Error(), http.StatusBadRequest)
		return nil, errEmptyRecord
	}
	return records, nil
}
