// Reads a file from a local path and pushes out one line at a time.
package readers

import ("os"
				"bufio"
				"errors")

type CodeFile struct {
	lines []string
	lineNumber int
}

func (cf *CodeFile) OpenFile(path string) error {
	file, err := os.Open(path)
	defer file.Close() // Close file before function exit
	cf.lineNumber = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cf.lines = append(cf.lines, scanner.Text())
	}

	return err
}

// Read next line
func (f *CodeFile) ReadLine() (string, int, error) {
	if f.lineNumber >= len(f.lines) {
		return "", -1, errors.New("No more lines left")
	}
	line := f.lines[f.lineNumber]
	f.lineNumber++
	return line, f.lineNumber, nil
}
