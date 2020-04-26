package sn_taskrunner

import (
	"context"
	"fmt"
	"testing"
	"time"
)

var t_taskrunner Taskrunner

func TestMain(t *testing.T) {
	t.Run("whole", testAddTask)
}
func testAddTask(t *testing.T) {
	t_taskrunner = NewTaskrunner()
	inputTask := []Task{
		NewTaskBase("1", "task1", nil),
		NewTaskBase("2", "task2", nil),
		NewTaskBase("3", "task3", nil),
	}
	inputHander := []Handler{
		NewHandler("task1", func(ctx context.Context, task Task) Task {
			return task
		}),
		NewHandler("task2", func(ctx context.Context, task Task) Task {
			return task
		}),
		NewHandler("task3", func(ctx context.Context, task Task) Task {
			return task
		}),
	}
	for i, _ := range inputHander {
		t_taskrunner.AddHandler(inputHander[i])
	}
	//for tp,handler := range t_taskrunner.handlers {
	//	for _,ha := range inputHander {
	//		if ha.MatchType()[0] == tp {
	//			if &handler == &ha.(HandlerBase).handleFunc {
	//				t.Error("Task Run error")
	//			}
	//			fmt.Println(task.GetId(),"success")
	//			continue
	//		}
	//		t.Error("No task find!")
	//	}
	//}
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	go t_taskrunner.Run(ctx)
	for i, _ := range inputTask {
		fmt.Println("input ", i)
		t_taskrunner.GetInputChan() <- inputTask[i]
	}
	for task := range t_taskrunner.GetOutputChan() {
		for _, ta := range inputTask {
			if task.GetId() == ta.GetId() {
				if task != ta {
					t.Error("Task Run error")
				}
				fmt.Println(task.GetId(), "success")
				goto NXT
			}
			continue
		}
		t.Error("No task find:", task)
	NXT:
	}
}
