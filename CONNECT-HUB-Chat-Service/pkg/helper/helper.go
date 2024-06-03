package helper

import (
	"errors"
	"strconv"
)

func Pagination(limit, offset string) (string, error) {
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		return "", err
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return "", err
	}

	if limitInt < 1 || offsetInt < 1 {
		return "", errors.New("pagination value must be positive")
	}

	return strconv.Itoa((offsetInt * limitInt) - limitInt), nil
}
