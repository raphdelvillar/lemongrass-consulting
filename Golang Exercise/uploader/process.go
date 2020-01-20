package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

// Zip --
type Zip struct {
	filename string
	writer   *zip.Writer
}

func (z Zip) compress() error {
	file, err := z.fileExist()

	defer file.Close()

	if err != nil {
		return err
	}

	wr, err := z.createWriter()

	if err != nil {
		return err
	}

	err = z.copy(wr, file)

	if err != nil {
		return err
	}

	return nil
}

func (z Zip) fileExist() (*os.File, error) {
	file, err := os.Open(z.filename)
	if err != nil {
		return nil, fmt.Errorf("Failed to open %s:%s", z.filename, err)
	}

	return file, nil
}

func (z Zip) createWriter() (io.Writer, error) {
	wr, err := z.writer.Create(z.filename)

	if err != nil {
		return nil, fmt.Errorf("Failed to create entry for %s in zip file:%s", z.filename, err)
	}

	return wr, nil
}

func (z Zip) copy(wr io.Writer, file *os.File) error {
	if _, err := io.Copy(wr, file); err != nil {
		return fmt.Errorf("Failed to write %s to zip %s", z.filename, err)
	}

	return nil
}
