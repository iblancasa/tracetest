package trigger

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/kubeshop/tracetest/server/model"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

func HTTP() Triggerer {
	return &httpTriggerer{
		traceProvider: traceProvider(),
	}
}

type httpTriggerer struct {
	traceProvider *sdktrace.TracerProvider
}

func (te *httpTriggerer) Trigger(_ context.Context, test model.Test, tid trace.TraceID, sid trace.SpanID) (Response, error) {

	response := Response{
		Result: model.TriggerResult{
			Type: te.Type(),
		},
	}

	trigger := test.ServiceUnderTest
	if trigger.Type != model.TriggerTypeHTTP {
		return response, fmt.Errorf(`trigger type "%s" not supported by HTTP triggerer`, trigger.Type)
	}

	client := http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport,
			otelhttp.WithTracerProvider(te.traceProvider),
			otelhttp.WithPropagators(propagators()),
		),
	}

	var tf trace.TraceFlags
	sc := trace.NewSpanContext(trace.SpanContextConfig{
		TraceID:    tid,
		SpanID:     sid,
		TraceFlags: tf.WithSampled(true),
		TraceState: trace.TraceState{},
		Remote:     true,
	})

	var req *http.Request
	tReq := trigger.HTTP
	var body io.Reader
	if tReq.Body != "" {
		body = bytes.NewBufferString(tReq.Body)
	}
	req, err := http.NewRequest(strings.ToUpper(string(tReq.Method)), tReq.URL, body)
	if err != nil {
		return response, err
	}
	for _, h := range tReq.Headers {
		req.Header.Set(h.Key, h.Value)
	}

	tReq.Authenticate(req)

	resp, err := client.Do(req.WithContext(trace.ContextWithSpanContext(context.Background(), sc)))
	if err != nil {
		return response, err
	}

	mapped := mapResp(resp)
	response.Result.HTTP = &mapped
	response.SpanAttributes = map[string]string{
		"tracetest.run.trigger.http.response_code": strconv.Itoa(resp.StatusCode),
	}

	return response, nil
}

func (t *httpTriggerer) Type() model.TriggerType {
	return model.TriggerTypeHTTP
}

func mapResp(resp *http.Response) model.HTTPResponse {
	var mappedHeaders []model.HTTPHeader
	for key, headers := range resp.Header {
		for _, val := range headers {
			val := model.HTTPHeader{
				Key:   key,
				Value: val,
			}
			mappedHeaders = append(mappedHeaders, val)
		}
	}
	var body string
	if b, err := io.ReadAll(resp.Body); err == nil {
		body = string(b)
	} else {
		fmt.Println(err)
	}

	return model.HTTPResponse{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Headers:    mappedHeaders,
		Body:       body,
	}
}
