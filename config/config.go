package config

import (
	"errors"
	"math"
	"strings"
)

type Config struct {
	Silent                bool
	Insecure              bool
	NoColor               bool
	JsonEnvelope          bool
	HashBody              bool
	DiscardBody           bool
	BaseDirectory         string
	DataTemplate          string
	RequestWorkers        int
	ResponseWorkers       int
	SubdirLength          int64
	RequestMethod         string
	ConnectTimeoutSeconds int64
	ThrottlePerSecond     int64
	Retries               int64
	RequestHeaders        []RequestHeader
	RequestFilename       string
}

func New() *Config {
	return &Config{
		RequestMethod:         "GET",
		Insecure:              false,
		Silent:                false,
		NoColor:               false,
		JsonEnvelope:          false,
		HashBody:              false,
		DiscardBody:           false,
		DataTemplate:          "",
		RequestWorkers:        1,
		SubdirLength:          0,
		ConnectTimeoutSeconds: 10,
		ThrottlePerSecond:     math.MaxInt32,
		Retries:               0,
	}
}

type RequestHeader struct {
	Key   string
	Value string
}

func NewRequestHeader(headerString string) (RequestHeader, error) {
	if strings.Contains(headerString, ":") {
		parts := strings.SplitN(headerString, ":", 2)
		return RequestHeader{Key: strings.TrimSpace(parts[0]), Value: strings.TrimSpace(parts[1])}, nil
	}

	return RequestHeader{}, errors.New("Header should be in the format 'Key: value', missing ':' -> " + headerString)
}

func ConvertRequestHeaders(stringHeaders []string) ([]RequestHeader, error) {
	var requestHeaders []RequestHeader

	for _, header := range stringHeaders {
		var requestHeader RequestHeader
		requestHeader, err := NewRequestHeader(header)

		if err != nil {
			return requestHeaders, err
		}

		requestHeaders = append(requestHeaders, requestHeader)
	}

	return requestHeaders, nil
}
