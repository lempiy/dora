package utils

import (
	"bytes"
	"compress/bzip2"
	"gopkg.in/h2non/filetype.v1"
	"gopkg.in/h2non/filetype.v1/types"
	"io"
	"os"
	"path/filepath"
)

// UncompressBZ2 - uncompresses *.bz2 files to provided location
func UncompressBZ2(source io.Reader, location string) error {
	body := bzip2.NewReader(source)
	err := copy(location, 0666, body)
	if err != nil {
		return err
	}
	return nil
}

func match(r io.Reader) (io.Reader, types.Type, error) {
	buffer := make([]byte, 512)

	n, err := r.Read(buffer)
	if err != nil && err != io.EOF {
		return nil, types.Unknown, err
	}

	r = io.MultiReader(bytes.NewBuffer(buffer[:n]), r)

	typ, err := filetype.Match(buffer)

	return r, typ, err
}

func copy(path string, mode os.FileMode, src io.Reader) error {
	err := os.MkdirAll(filepath.Dir(path), mode|os.ModeDir|100)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, mode)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, src)
	return err
}
