package scanner

import ("testing"
				"log")

func TestSingleCharSymbolLexer(t *testing.T) {
	symbols := "{}(),.-+;/*"
	var scanner Scanner
	tokens, err := scanner.Scan(symbols);
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < len(tokens); i++ {
		if string(symbols[i]) != tokens[i].String {
			t.Error("Expected:", string(symbols[i]), "Found:", tokens[i])
		}
	}
	log.Print(tokens)
}

func TestSingleOrDoubleCharSymbolLexer(t *testing.T) {
	symbols := "!!====>>=<<="
	expectedResult := []string{"!","!=","==","=",">",">=","<","<="}
	var scanner Scanner
	tokens, err := scanner.Scan(symbols);
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < len(tokens); i++ {
		if expectedResult[i] != tokens[i].String {
			t.Error("Expected:", expectedResult[i], "Found:", tokens[i])
		}
	}
	log.Print(tokens)
}

func TestAllSymbols(t *testing.T) {
	text := "{ !(!=) == (>) = * }"
	expectedResult := []string {"{","!","(","!=",")","==",
															"(",">",")","=","*","}"}
	var scanner Scanner
	tokens, err := scanner.Scan(text);
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < len(tokens); i++ {
		if expectedResult[i] != tokens[i].String {
			t.Error("Expected:", expectedResult[i], "Found:", tokens[i])
		}
	}
	log.Print(tokens)
}

func TestKeywordDetection(t *testing.T) {
	text := "pri print \"ABC\";" // This is an invalid lox syntax.
	expectedResult := []string {"pri", "print", "ABC", ";"}
	var scanner Scanner
	tokens, err := scanner.Scan(text)
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < len(tokens); i++ {
		if expectedResult[i] != tokens[i].String {
			t.Error("Expected:", expectedResult[i], "Found:", tokens[i])
		}
	}
	log.Print(tokens)
}

func TestQuotedText(t *testing.T) {
	var scanner Scanner
	expectedTokens := []string{"print", "Hello world!", ";", "// Comment"}
	tokens, err := scanner.Scan("print \"Hello world!\"; // Comment")
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < len(tokens); i++ {
		if expectedTokens[i] != tokens[i].String {
			t.Error("Found:",tokens[i]," Expected:", expectedTokens[i])
		}
	}
	log.Print(tokens)
}
