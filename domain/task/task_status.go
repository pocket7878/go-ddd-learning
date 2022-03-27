package task

type TaskStatus int

const (
	Undone TaskStatus = iota + 1
	Done
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
