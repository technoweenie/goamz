package s3

import (
	"github.com/technoweenie/goamz/aws"
)

var originalStrategy = attempts

func SetAttemptStrategy(s *aws.AttemptStrategy) {
	if s == nil {
		attempts = originalStrategy
	} else {
		attempts = *s
	}
}

func SetListPartsMax(n int) {
	listPartsMax = n
}

func SetListMultiMax(n int) {
	listMultiMax = n
}
