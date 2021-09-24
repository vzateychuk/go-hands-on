package work

import (
	"math/rand"
	"strconv"
	"time"
)

/* Это порция работы, "выдаваемая" извне в worker */
type WorkItem struct {
	WorkId string
	Delay  int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewWorkItems(n int) []*WorkItem {

	prefix := time.Now().Format(time.RFC1123) + "-"
	items := make([]*WorkItem, 0, n)
	for i := 0; i < n; i++ {
		wrk := &WorkItem{
			WorkId: prefix + strconv.Itoa(i),
			Delay:  rand.Intn(100),
		}
		items = append(items, wrk)
	}
	return items
}
