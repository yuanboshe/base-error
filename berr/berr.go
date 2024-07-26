package berr

import (
	"errors"
)

// ErrInitFailed if InitAddr not called, SetErr will return nil
var ErrInitFailed = errors.New("BaseErr init failed, pleas call InitAddr(t *T) first")

// BaseErr used as base struct
type BaseErr[T any] struct {
	err error
	t   *T
}

// Err get the error
func (p *BaseErr[T]) Err() error {
	return p.err
}

// SetErr must call InitAddr once before use SetErr
//
//	otherwise panic arrived:
//	"panic: runtime error: invalid memory address or nil pointer dereference"
func (p *BaseErr[T]) SetErr(err error) *T {
	if p.t == nil {
		p.err = ErrInitFailed
		return nil
	}

	p.err = err

	return p.t
}

// InitAddr must be call once, set the child struct address
//
//	otherwise panic arrived:
//	"panic: runtime error: invalid memory address or nil pointer dereference"
func (p *BaseErr[T]) InitAddr(t *T) *T {
	p.t = t
	return p.t
}
