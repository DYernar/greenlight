package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	quotedJsonValue := strconv.Quote(jsonValue)

	return []byte(quotedJsonValue), nil
}

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	unquotedJson, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		fmt.Println(err)
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(unquotedJson, " ")

	if len(parts) != 2 || parts[1] != "mins" {
		fmt.Println(err)
		return ErrInvalidRuntimeFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		fmt.Println(err)
		return ErrInvalidRuntimeFormat
	}

	*r = Runtime(i)

	return nil

}
