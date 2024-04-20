package sync

type Counter struct{}

// methods belongs to struct
func (c *Counter) Inc() {}

func (c *Counter) Value() int {
	return 0
}
