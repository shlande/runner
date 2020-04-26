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
	Status Status                 `json:"status"`
	Type   TaskType               `json:"type"`
	Data   map[string]interface{} `json:"data"`
}

func NewTaskBase(id string, taskType TaskType, data map[string]interface{}) *TaskBase {
	return &TaskBase{Id: id, Type: taskType, Data: data}
}

func (t TaskBase) GetId() string {
	return t.Id
}

func (t TaskBase) GetStatus() Status {
	return t.Status
}

func (t TaskBase) SetStatus(status Status) {
	t.Status = status
}

func (t TaskBase) GetType() TaskType {
	return t.Type
}

func (t TaskBase) GetData() TaskData {
	return t.Data
}
