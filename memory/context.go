package memory

import (
	"arena"
	"context"
)

func NewContext(ctx context.Context) (context.Context, func()) {
	a := arena.NewArena()
	context.AfterFunc(ctx, a.Free)

	ctx, cancel := context.WithCancel(ctx)

	return &arenaCtx{alloc: a, Context: ctx}, cancel
}

type arenaCtx struct {
	alloc *arena.Arena
	context.Context
}

func NewWithContext[T any](ctx context.Context) *T {
	return arena.New[T](ctx.(*arenaCtx).alloc)
}
