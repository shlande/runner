package sn_taskrunner

type Status int
type TaskType string
type TaskData map[string]interface{}

// 自定义状态码可以在两侧取得
const (
	FAIL      Status = -1
	UNHANDLED Status = 0
	SUCCESS   Status = 1
)

type Task interface {
	// 获取任务ID
	GetId() string
	// 获取任务执行状态
	GetStatus() Status
	// 设置任务状态
	SetStatus(status Status)
	// 获取任务类型
	GetType() TaskType
	// 获取任务数据
	GetData() TaskData
}

type TaskBase struct {
	Id     string                 `json:"id"`
	Status int                    `json:"status"`
	Type   string                 `json:"type"`
	Data   map[string]interface{} `json:"data"`
}
