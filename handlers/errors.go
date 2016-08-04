package handlers

import (
	"encoding/json"
	"net/http"

	"code.cloudfoundry.org/routing-api"
	"code.cloudfoundry.org/routing-api/metrics"
	"github.com/pivotal-golang/lager"
)

func handleProcessRequestError(w http.ResponseWriter, procErr error, log lager.Logger) {
	log.Error("error", procErr)
	err := routing_api.NewError(routing_api.ProcessRequestError, "Cannot process request: "+procErr.Error())
	retErr := marshalRoutingApiError(err, log)

	w.WriteHeader(http.StatusBadRequest)
	_, writeErr := w.Write(retErr)
	log.Error("error writing to request", writeErr)
}

func handleNotFoundError(w http.ResponseWriter, err error, log lager.Logger) {
	log.Error("error", err)
	retErr := marshalRoutingApiError(routing_api.NewError(routing_api.ResourceNotFoundError, err.Error()), log)

	w.WriteHeader(http.StatusNotFound)
	_, writeErr := w.Write(retErr)
	log.Error("error writing to request", writeErr)
}

func handleApiError(w http.ResponseWriter, apiErr *routing_api.Error, log lager.Logger) {
	log.Error("error", apiErr)
	retErr := marshalRoutingApiError(*apiErr, log)

	w.WriteHeader(http.StatusBadRequest)
	_, writeErr := w.Write(retErr)
	log.Error("error writing to request", writeErr)
}

func handleDBCommunicationError(w http.ResponseWriter, err error, log lager.Logger) {
	log.Error("error", err)
	retErr := marshalRoutingApiError(routing_api.NewError(routing_api.DBCommunicationError, err.Error()), log)

	w.WriteHeader(http.StatusInternalServerError)
	_, writeErr := w.Write(retErr)
	log.Error("error writing to request", writeErr)
}

func handleUnauthorizedError(w http.ResponseWriter, err error, log lager.Logger) {
	log.Error("error", err)

	retErr := marshalRoutingApiError(routing_api.NewError(routing_api.UnauthorizedError, err.Error()), log)
	metrics.IncrementTokenError()

	w.WriteHeader(http.StatusUnauthorized)
	_, writeErr := w.Write(retErr)
	log.Error("error writing to request", writeErr)
}

func handleDBConflictError(w http.ResponseWriter, err error, log lager.Logger) {
	log.Error("error", err)
	retErr := marshalRoutingApiError(routing_api.NewError(routing_api.DBConflictError, err.Error()), log)

	w.WriteHeader(http.StatusConflict)
	_, writeErr := w.Write(retErr)
	log.Error("error writing to request", writeErr)
}

func marshalRoutingApiError(err routing_api.Error, log lager.Logger) []byte {
	retErr, jsonErr := json.Marshal(err)
	if jsonErr != nil {
		log.Error("could-not-marshal-json", jsonErr)
	}

	return retErr
}
