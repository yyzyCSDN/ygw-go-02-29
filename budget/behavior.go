package budget

import (
	"time"
)

func Budget(parentRemaining, requested time.Duration) time.Duration {
	if requested > 0 {
		return requested
	}
	return parentRemaining
}
