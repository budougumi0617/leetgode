package leetgode

import "fmt"

func buildPath(id, slug string) string {
	format := "tmp/%s.%s.go"
	return fmt.Sprintf(format, id, slug)
}
