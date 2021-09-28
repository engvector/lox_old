package readers

import (
				"testing"
				"log")

func TestMissingFile(t *testing.T) {
	var cf CodeFile
	err := cf.OpenFile("dummy.lox") //This file doesn't exist
	log.Print(err)
}

func TestHelloWorldFile(t *testing.T) {
	var cf CodeFile
	err := cf.OpenFile("testfiles/helloworld.lox")
	if err != nil {
		t.Error("Problem in opening helloworld.lox")
	}
}

func TestReadSingleLine_HelloWorld(t *testing.T) {
	var cf CodeFile
	err := cf.OpenFile("testfiles/helloworld.lox")
	if err != nil { t.Error("Problem opening helloworld.lox") }
	line, _, err := cf.ReadLine()
	if err != nil {
		t.Error(err)
	}
	if line != "print \"Hello, world!\";" {
		t.Error ("Found: ", line)
		t.Error ("Expected: print \"Hello world!\";")
	}
}

func TestReadTwoLines(t *testing.T) {
	var cf CodeFile
	var firstline = "var pi = 22/7;"
	var	secondline = "print pi;"
	err := cf.OpenFile("testfiles/twolines.lox")
	if err != nil { t.Error("Problem opening twolines.lox") }
	line_1, _, err := cf.ReadLine()
	if err != nil {
		t.Error(err)
	}
	if line_1 != firstline {
		t.Error("1:", line_1, "(found); ", firstline, "(expected)")
	}

	line_2, _, err := cf.ReadLine()
	if err != nil {
		t.Error(err)
	}
	if line_2 != secondline {
		t.Error("2:", line_2, "(found); ", secondline, "(expected)")
	}
}
