package filesystem

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func BasePath(path string) string {
	_, b, _, _ := runtime.Caller(0)
	Root := filepath.Join(filepath.Dir(b), "../../..")

	return fmt.Sprintf("%s%c%s", Root, os.PathSeparator, path)
}
