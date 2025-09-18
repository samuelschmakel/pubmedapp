package config

import (
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

type ApiConfig struct {
    FileserverHits atomic.Int32
    HttpClient *http.Client
    PythonClient *http.Client
    Platform string
    Secretkey string
    PythonBaseURL string
}

func CreateConfig() *ApiConfig {
    return &ApiConfig{
        FileserverHits: atomic.Int32{},
        HttpClient: &http.Client{
            Timeout: 10 * time.Second,
        },
        PythonClient: &http.Client{
            Timeout: 10 * time.Second,
        },
        PythonBaseURL: os.Getenv("PYTHON_BASE_URL"),
    }
}