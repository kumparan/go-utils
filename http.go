package utils

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context/ctxhttp"
)

type httpResponse interface {
	IsHTTPResponse()
}

// DoHTTPRequest generic do http request using ctxhttp
func DoHTTPRequest[T httpResponse](ctx context.Context, httpClient *http.Client, httpRequest *http.Request) (respStatusCode int, respBody *T, err error) {
	logger := log.WithFields(log.Fields{
		"ctx":         DumpIncomingContext(ctx),
		"httpRequest": Dump(httpRequest),
	})

	httpResp, err := ctxhttp.Do(ctx, httpClient, httpRequest)
	if err != nil {
		logger.Error(err)
		return respStatusCode, nil, err
	}

	logger = logger.WithField("response header", Dump(httpResp.Header))

	respInBytes, err := io.ReadAll(httpResp.Body)
	if err != nil {
		logger.Error(err)
		return httpResp.StatusCode, nil, err
	}
	defer func() {
		_ = httpResp.Body.Close()
	}()

	if httpResp.StatusCode != http.StatusOK {
		logger.WithField("body", string(respInBytes)).Warn("http status code is not ok")
	}

	var resp T
	err = json.Unmarshal(respInBytes, &resp)
	if err != nil {
		logger.WithField("respInBytes", string(respInBytes)).Error(err)
		return httpResp.StatusCode, nil, err
	}
	return httpResp.StatusCode, &resp, nil
}
