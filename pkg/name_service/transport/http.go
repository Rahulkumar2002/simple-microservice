package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/Rahulkumar2002/simple-microservice/pkg/name_service/endpoints"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPHandler(ep endpoints.Set) http.Handler {
	m := http.NewServeMux()

	m.Handle("/name", httptransport.NewServer(
		ep.GiveNameEndpoint,
		decodeHTTPGiveNameRequest,
		encodeResponse,
	))

	return m
}

func decodeHTTPGiveNameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	logger.Log("Decoding", "RequestBody")
	var req endpoints.GiveNameRequest
	if r.ContentLength == 0 {
		logger.Log("Request with no body")
		return req, nil
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	if e, ok := response.(error); ok && e != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": e.Error(),
		})
	}
	logger.Log("Response", "Encoded")
	return json.NewEncoder(w).Encode(response)
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
