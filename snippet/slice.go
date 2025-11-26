package snippet

import (
	"log/slog"
)

func DealSliceByStep() {
	tarList := make([]int, 1000)

	step := 3
	total := len(tarList)
	for indexNow := 0; indexNow < total; {
		next := indexNow + step
		if next > total {
			next = total
		}
		slog.Info("current progress", "from", indexNow, "to", next, "total", total)
		DoList(tarList[indexNow:next])
		indexNow = next
	}
}

func DoList(ins []int) {
	for k := range ins {
		ins[k] = 5
	}
}
