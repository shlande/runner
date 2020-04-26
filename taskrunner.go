package sn_taskrunner

import "context"

type Taskrunner interface {
	// 启动进程
	Run(ctx context.Context)
	// 设置任务处理器
	AddHandler(handler Handler)
	// 获取输入管道
	GetInputChan() chan<- Task
	// 获取输出管道
	GetOutputChan() <-chan Task
}

// 内置处理方法
const (
	BEFORERUN TaskType = "INNER_BEFORERUN"
	AFTERRUN  TaskType = "INNER_AFTERRUN"
)
