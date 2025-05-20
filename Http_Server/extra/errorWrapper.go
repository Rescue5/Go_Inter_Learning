package extra

import (
	"fmt"
	"runtime"
)

func WrapError(msg string, err error) error {
	_, file, line, ok := runtime.Caller(1)
	location := ""
	if ok {
		location = fmt.Sprintf("[%s:%d]", file, line)
	}
	return fmt.Errorf("%s %s\n→ %v", location, msg, err)
}
