package app

import (
	"sync"

	"github.com/uptrace/bun"
	"google.golang.org/grpc"
)

type App struct {
	db         *bun.DB
	server     *grpc.Server
	serverOnce sync.Once
}

func NewApp() {}

func (a *App) Start() {

}

func (a *App) Stop() {

}
