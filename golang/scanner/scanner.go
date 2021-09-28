package scanner

import "strings"

type Scanner struct {
	lexemes []string
}

func (s *Scanner) Scan(code string) ([]string, error) {
	// Valid single character symbols.
	singleCharSymbols := "{}(),.-+;/*"
	// Single/Double character symbols are codified as two strings. First 
	// string contains the first symbol and the second string contains the
	// matching second symbol.
	singleOrDoubleCharSymbols := []string{"!=><","===="}
	keywords := []string{
			"and", "class", "else", "false", "fun", "for",
			"if", "nil", "or", "print", "return", "super",
			"this", "true", "var", "while"}
	var literal = ""

	// First Pass: Break string into lexemes
	for i := 0; i < len(code); i++ {
		// First make sure symbols inside quotes are not detected.
		if code[i] == '"' {
			aString := ""
			for i++; code[i] != '"'; i++ {
				aString = aString + string(code[i])
			}
			// If we were processing a literal before, finish it.
			if len(literal) > 0 {
				s.lexemes = append(s.lexemes, literal)
				literal = ""
			}
			s.lexemes = append(s.lexemes, aString)
			continue // Move on
		}
		// Start with single character symbols
		if strings.IndexByte(singleCharSymbols, code[i]) != -1 {
			// If we were processing a literal before, finish it.
			if len(literal) > 0 {
				s.lexemes = append(s.lexemes, literal)
				literal = ""
			}
			s.lexemes = append(s.lexemes, string(code[i]))
			continue // Move on
		}
		// Check for single/double character symbols
		if strings.IndexByte(singleOrDoubleCharSymbols[0], code[i]) != -1 {
			if strings.IndexByte(singleOrDoubleCharSymbols[1], code[i+1]) != -1 {
				// If we were processing a literal before, finish it.
				if len(literal) > 0 {
					s.lexemes = append(s.lexemes, literal)
					literal = ""
				}
				// If both characters matched
				s.lexemes = append(s.lexemes, string(code[i]) + string(code[i+1]))
				i++
			} else {
				// If we were processing a literal before, finish it.
				if len(literal) > 0 {
					s.lexemes = append(s.lexemes, literal)
					literal = ""
				}
				// If we only matched a single character
				s.lexemes = append(s.lexemes, string(code[i]))
			}
			continue // Move on
		}
		// Look for keywords
		kwd := findKeyword(code, i, keywords)
		if (kwd != "") {
			// If we were processing a literal before, finish it.
			if len(literal) > 0 {
				s.lexemes = append(s.lexemes, literal)
				literal = ""
			}
			s.lexemes = append(s.lexemes, kwd)
			i += len(kwd) - 1 // Fast-forward index to end of detected keyword
			continue // Move on
		}

		// If none of the above searches yield anything, it is a literal
		// Process the string till you find the next valid symbol.
		if (code[i] != ' ') {
			literal = literal + string(code[i]);
		}
	}
	return s.lexemes, nil
}

// Look for keywords starting at supplied index.
// Returns the detected keyword.
func findKeyword(input string, startIdx int, keywords []string) string {
	for k := 0; k < len(keywords); k++ {
		if matchKeyword(input[startIdx:], keywords[k]) {
			return keywords[k]
		}
	}
	return ""
}

// We just use naive matching for now.
// Algorithms like KMP should give better performance.
// TODO: Can I create a KMP lookup table across keywords?
func matchKeyword(input string, keyword string) bool {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(keyword); j++ {
			if input[i+j] != keyword[j] {
				return false
			}
		}
		return true
	}
	return false
}
