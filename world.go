package m2

import (
	"fmt"

	"github.com/devlights/m1"
)

func World() string {
	return fmt.Sprintf("%s world", m1.Hello())
}
