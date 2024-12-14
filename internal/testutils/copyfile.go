package testutils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CopyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	srcInfo, err := srcFile.Stat()
	if err != nil {
		return err
	}

	return os.Chmod(dst, srcInfo.Mode())
}

func PrepareTestCase(tc string, files []string) (string, error) {
	testloc, err := os.MkdirTemp(os.TempDir(), "*")
	if err != nil {
		return "", fmt.Errorf("creating temporary directory to test the created schema: %w", err)
	}
	fmt.Println("using tmp dir:", testloc)

	for _, file := range files {
		src := filepath.Join("testdata", tc, file)
		dst := filepath.Join(testloc, file)
		err = os.MkdirAll(filepath.Dir(dst), 0755)
		if err != nil {
			return "", fmt.Errorf("MkdirAll: %w", err)
		}
		if err := CopyFile(src, dst); err != nil {
			return "", fmt.Errorf("copying %q to %q: %w", file, dst, err)
		}
	}

	return testloc, nil
}
