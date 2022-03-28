package task

type TaskStatus string

const (
	Undone TaskStatus = "undone"
	Done              = "done"
)

func (s *TaskStatus) String() string {
	switch *s {
	case Undone:
		return "undone"
	case Done:
		return "done"
	default:
		panic("Unknown TaskStatus")
	}
}
