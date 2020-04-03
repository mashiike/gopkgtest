package subone

import (
	"fmt"

	humanize "github.com/dustin/go-humanize"
)

func Fuga(num int64) string {
	return fmt.Sprintf("%s[fuga]", humanize.Comma(num))
}
