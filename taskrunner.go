package sn_taskrunner

import "context"

type Taskrunner interface {
	// 启动进程
	Run(ctx context.Context)
	// 任务开始前的钩子函数
	BeforeRun(task *Task)
	// 任务完成后的钩子函数
	AfterRun(task *Task)
	// 设置任务处理器
	SetHandler(handler Handler)
	// 获取输入管道
	GetInputChan() chan<- Task
	// 获取输出管道
	GetOutputChan() <-chan Task
}
