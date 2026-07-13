package worker

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Result struct {
	Line    string
	LineNum int
	Path    string
}

type Results struct {
	Inner []Result
}

func FindInPath(path string, find string) (r *Results, err error) {
	if path == "" {
		return nil, nil
	}
	file, err := os.Open(path)
	if err != nil {
		err = fmt.Errorf("Unable to open file: %v", err)
		return
	}

	results := Results{
		make([]Result, 0),
	}
	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		ok := strings.Contains(scanner.Text(), find)
		if ok {
			res := Result{
				Line:    scanner.Text(),
				LineNum: lineNum,
				Path:    path,
			}
			results.Inner = append(results.Inner, res)
		}
		lineNum++
	}

	if len(results.Inner) == 0 {
		return nil, nil
	}

	return &results, nil
}
