package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

var from string
var to string
var limit int64
var offset int64

func openFile(path string) (*os.File, error) {
	file, err := os.Open(from)
	if err != nil {
		if os.IsNotExist(err) {
			return file, fmt.Errorf("file %s was not found", path)
		}
		return file, err
	}
	return file, nil
}

func writeFile(to string, reader io.Reader) error {
	target, err := os.Create(to)

	if err != nil {
		return fmt.Errorf("cannot write the content to provided destination: %s", err)
	}
	defer target.Close()

	writer := bufio.NewWriter(target)

	if _, err := io.Copy(writer, reader); err != nil {
		return fmt.Errorf("error occurred during writing to the file: %s", err)
	}

	writer.Flush()

	return nil
}

func checkParams() error {
	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file destination")
	flag.Int64Var(&limit, "limit", 0, "limit in input file")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")

	flag.Parse()
	if from == "" {
		return errors.New("required parameter 'from' is not provided")
	}
	if to == "" {
		return errors.New("required parameter 'from' is not provided")
	}
	return nil
}

func main() {
	// Check required variables
	err := checkParams()
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	// Open the file
	file, err := openFile(from)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error occurred during opening the file: %s", err)
		os.Exit(1)
	}
	defer file.Close()

	// Copy the whole file, when parameter is not provided
	if limit == 0 {
		info, err := file.Stat()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error, while getting the file info: %s", err)
			os.Exit(1)
		}
		limit = info.Size()
	}

	// Write the file
	err = writeFile(to, io.NewSectionReader(file, offset, limit))
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	fmt.Println("File was copied successfully")
}
