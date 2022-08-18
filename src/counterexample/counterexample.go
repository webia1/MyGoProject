package counterexample

import (
	"fmt"
	"time"
)

type Counter struct {
	no      int
	updated time.Time
}

func (c *Counter) Inc() {
	c.no++
	c.updated = time.Now()
}

func (c Counter) Log() string {
	return fmt.Sprintf("No: %d, updated: %v", c.no, c.updated)
}
