package scanner

import ("testing"
				"log")

func TestScanner(t *testing.T) {
	var scanner Scanner
	tokens, err := scanner.Scan("print \"Hello world!\";")
	if err != nil {
		t.Error(err)
	}

	log.Print(tokens)
}
