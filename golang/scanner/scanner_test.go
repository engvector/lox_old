package scanner

import ("testing"
				"log")

func TestScanner(t *testing.T) {
	tokens, err := Scan("print \"Hello world!\";")
	if err != nil {
		t.Error(err)
	}

	log.Print(tokens)
}
