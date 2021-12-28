package utils

import (
	"strconv"

	"github.com/chinmay2197/stripe-integration-api/logging"
)

func ConvertStringToInt64(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		logging.Logger.Error("string to int64 conversion failed Error:", err)
		return 0, err
	}
	return i, nil
}
