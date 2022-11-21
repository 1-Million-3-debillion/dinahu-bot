package tools

import (
	"fmt"
	"os"
	"sort"
)

const failGetFiles string = "GetFiles() failed: %v"

func GetFiles(dir string) ([]string, error) {
	files, err := os.Open(dir)
	if err != nil {
		return nil, fmt.Errorf(failGetFiles, err)
	}
	defer func() {
		_ = files.Close()
	}()

	list, err := files.Readdirnames(0)
	if err != nil {
		return nil, fmt.Errorf(failGetFiles, err)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})

	return list, nil
}
