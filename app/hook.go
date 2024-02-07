package app

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

var onStart appHooks

func OnStart(name string, fn HookFunc) {
	onStart.Add(newHook(name, fn))
}

type HookFunc func(ctx context.Context, a *App) error

type appHooks struct {
	mu    sync.Mutex
	hooks []appHook
}

type appHook struct {
	name string
	fn   HookFunc
}

func (h appHook) run(ctx context.Context, a *App) error {
	const timeout = 30 * time.Second

	done := make(chan struct{})
	errc := make(chan error)

	go func() {
		start := time.Now()
		if err := h.fn(ctx, a); err != nil {
			errc <- err
			return
		}
		if d := time.Since(start); d > time.Second {
			fmt.Printf("hook=%q took %s\n", h.name, d)
		}
		close(done)
	}()
	select {
	case <-done:
		return nil
	case err := <-errc:
		return err
	case <-time.After(timeout):
		return fmt.Errorf("hook=%q timed out afrer %s", h.name, timeout)
	}
}

func newHook(name string, fn HookFunc) appHook {
	return appHook{name: name, fn: fn}
}

func (hs *appHooks) Add(hook appHook) {
	hs.mu.Lock()
	defer hs.mu.Unlock()
	hs.hooks = append(hs.hooks, hook)
}

func (hs *appHooks) Run(ctx context.Context, a *App) error {

	hs.mu.Lock()
	defer hs.mu.Unlock()

	g := new(errgroup.Group)

	for _, h := range hs.hooks {
		h := h
		g.Go(func() error {
			err := h.run(ctx, a)
			if err != nil {
				fmt.Printf("hook=%q failed: %s\n", h.name, err)
			}
			return err
		},
		)

	}
	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}
