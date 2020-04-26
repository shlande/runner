package sn_taskrunner

import "context"

type HandleFunc func(ctx context.Context, task Task) Task

type Handler interface {
	MatchType() []TaskType
	Handle(ctx context.Context, task Task) Task
}

type HandlerBase struct {
	matchType  TaskType
	handleFunc HandleFunc
}

func (h HandlerBase) MatchType() []TaskType {
	return []TaskType{h.matchType}
}

func (h HandlerBase) Handle(ctx context.Context, task Task) Task {
	return h.handleFunc(ctx, task)
}

func NewHandler(taskType TaskType, handleFunc HandleFunc) Handler {
	return HandlerBase{taskType, handleFunc}
}
