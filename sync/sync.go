package sync

type Counter struct {
	value int
}

// methods belongs to struct
func (c *Counter) Inc() {
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
