package context

import (
	"context"
	"sync"
)

type WaitStopContext struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	wg         sync.WaitGroup
}

func NewWaitStopContext() *WaitStopContext {
	ctx, cancel := context.WithCancel(context.Background())
	return &WaitStopContext{
		ctx:        ctx,
		cancelFunc: cancel,
	}
}

func (c *WaitStopContext) Cancel() {
	c.cancelFunc()
	c.wg.Wait()
}

func (c *WaitStopContext) Add(delta int) {
	c.wg.Add(delta)
}

func (c *WaitStopContext) Finish() {
	c.wg.Done()
}

func (c *WaitStopContext) Done() <-chan struct{} {
	return c.ctx.Done()
}

func (c *WaitStopContext) Value(key interface{}) interface{} {
	return c.ctx.Value(key)
}

func WithValue(parent *WaitStopContext, key interface{}, value interface{}) {
	parent.ctx = context.WithValue(parent.ctx, key, value)
}

const (
	TempCancelContextKey = "temp_cancel_func"
)

func WithTempCancel(parent *WaitStopContext, cancelFunc context.CancelFunc) {
	parent.ctx = context.WithValue(parent.ctx, TempCancelContextKey, cancelFunc)
}

func (c *WaitStopContext) TempCancel() context.CancelFunc {
	return c.ctx.Value(TempCancelContextKey).(context.CancelFunc)
}
