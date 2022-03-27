package task

type TaskStatus string

const (
	Undone TaskStatus = "undone"
	Done              = "done"
)

func (s *TaskStatus) String() string {
	switch *s {
	case Undone:
		return "Undone"
	case Done:
		return "Done"
	default:
		panic("Unknown TaskStatus")
	}
}
