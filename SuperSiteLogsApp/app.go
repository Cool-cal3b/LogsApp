package main

import (
	"context"
)

type App struct {
	ctx context.Context
}

// startup is called at application startup
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}
