package config

import (
	"net/http"
	"sync/atomic"
)

type ApiConfig struct {
    FileserverHits atomic.Int32
    HttpClient *http.Client
    Platform string
    Secretkey string
}