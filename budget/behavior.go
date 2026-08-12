package budget

import (
	"time"
)

func Budget(parentRemaining, requested time.Duration) time.Duration {
	if parentRemaining <= 0 {
		return 0
	}
	if requested <= 0 || requested > parentRemaining {
		return parentRemaining
	}
	return requested
}
