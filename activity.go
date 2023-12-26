package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/carlmjohnson/requests"
)

func TryDownload(ctx context.Context, url string, dir string) (string, error) {
	slices := strings.SplitAfter(url, "=")
	if len(slices) == 1 {
		slices = strings.SplitAfter(url, "/")
	}
	filename := slices[len(slices)-1]
	destination := filepath.Join(dir, filename)
	err := requests.URL(url).ToFile(destination).Fetch(ctx)

	if err != nil {
		return "", err
	}

	stat, err := os.Stat(destination)
	if err != nil {
		log.Fatal(err)
	}

	size := fmt.Sprintf("file is %d bytes\n", stat.Size())
	return size, nil
}
