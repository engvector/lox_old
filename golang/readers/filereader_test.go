package filereader

import (
				"testing"
				"log")

func TestMissingFile(t *testing.T) {
	// NOTE: This file shouldn't exist

	err := OpenFile("dummy.lox")
	log.Print(err)
}
