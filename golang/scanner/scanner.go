package scanner

import "errors"

type TokenType int32
const (
	// Single-character tokens.
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
  COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR
	// One or two character tokens.
  BANG
	BANG_EQUAL
  EQUAL
	EQUAL_EQUAL
  GREATER
	GREATER_EQUAL
  LESS
	LESS_EQUAL
  // Literals.
  IDENTIFIER
	STRING
	NUMBER
  // Keywords.
  AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
  PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE
	// End of file
  EOF
)

type Token struct {
	Type TokenType
	Lexeme string
	Literal string
}

type Scanner struct {
	tokens []Token
}

func (s *Scanner) addToken(typ TokenType) {
	s.tokens = append(s.tokens,
										Token {
											Type : typ,
										})
}

func (s *Scanner) Scan(code string) ([]Token, error) {
	for index, character := range code {
		switch character {
			case '(':
				s.addToken(LEFT_PAREN)
      case ')':
				s.addToken(RIGHT_PAREN)
      case '{':
				s.addToken(LEFT_BRACE)
      case '}':
				s.addToken(RIGHT_BRACE)
      case ',':
				s.addToken(COMMA)
      case '.':
				s.addToken(DOT)
      case '-':
				s.addToken(MINUS)
      case '+':
				s.addToken(PLUS)
      case ';':
				s.addToken(SEMICOLON)
      case '*':
				s.addToken(STAR)
			case ' ':
				continue
			default:
				return s.tokens,
							 errors.New("unexpected character - " + code[index:index+1])
		}
	}
	return s.tokens, nil
}

