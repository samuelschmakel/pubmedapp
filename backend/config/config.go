package config

import "sync/atomic"

type ApiConfig struct {
    FileserverHits atomic.Int32
    Platform string
    Secretkey string
}