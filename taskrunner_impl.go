package sn_taskrunner

import (
	"context"
	"sync"
)

type taskrunner struct {
	handlers   map[TaskType]HandleFunc
	inputChan  chan Task
	outputChan chan Task
	wait       sync.WaitGroup
	// log 控制是否输出
}

func (t *taskrunner) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			goto EXIT
		case task := <-t.inputChan:
			t.wait.Add(1)
			go func() {
				t.handlers[task.GetType()](ctx, task)
				defer t.wait.Done()
			}()
		}
	}
EXIT:
	close(t.inputChan)
	t.wait.Wait()
	// 所有任务已经退出，关闭输出通道
	close(t.outputChan)
}

// 需要被重写的方法
func (t *taskrunner) BeforeRun(task *Task) {
	return
}

// 需要被重写的方法
func (t *taskrunner) AfterRun(task *Task) {
	return
}

// 注册方法
func (t *taskrunner) AddHandler(handler Handler) {
	for _, tp := range handler.MatchType() {
		if _, ok := t.handlers[tp]; ok {
			panic("重复加载处理方法：" + tp)
		}
		t.handlers[tp] = handler.Handle
	}
}

func (t *taskrunner) GetInputChan() chan<- Task {
	return t.inputChan
}

func (t *taskrunner) GetOutputChan() <-chan Task {
	return t.outputChan
}

// 构造方法
func NewTaskrunner() Taskrunner {
	return &taskrunner{
		handlers:   make(map[TaskType]HandleFunc),
		inputChan:  make(chan Task, 100),
		outputChan: make(chan Task, 100),
	}
}
