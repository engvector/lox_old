// Reads a file from a local path and pushes out one line at a time.
package filereader

import "os"

type CodeFile struct {
	lines string
	//fileHandle *File
	lineNumber int
}

func OpenFile(path string) error {
	_, err := os.Open(path)

	return err
}
/*
func (CodeFile *f) ReadNextLine() string {
	return ""
}

func (CodeFile *f) ReadPreviousLine() string {
	return ""
}*/
