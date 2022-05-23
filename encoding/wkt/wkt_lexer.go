package wkt

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"unicode"
)

type tokenType int

// const ...
const (
	// Separator
	LeftParen tokenType = iota
	RightParen
	Comma
	EqualSign
	Semicolon

	// Keyword
	Empty
	Z
	M
	ZM

	Srid

	// Geometry type
	PointEnum
	Linestring
	PolygonEnum
	Multipoint
	MultilineString
	MultiPolygonEnum
	GeometryCollection

	// Values
	Float
	Int

	EOF
)

// eof is used to simplify treatment of file end
const eof = rune(0)

// Token ...
type Token struct {
	ttype  tokenType
	lexeme string
	pos    int
}

// Lexer ...
type Lexer struct {
	reader *bufio.Reader

	pos int
}

// NewLexer ...
func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		reader: bufio.NewReader(reader),
	}
}

// getToken add a parsed token to the token list
func (l *Lexer) getToken(ttype tokenType, lexeme string) Token {
	t := Token{ttype, lexeme, l.pos}
	l.pos += len(lexeme)
	return t
}

func (l *Lexer) read() rune {
	ch, _, err := l.reader.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (l *Lexer) unread() {
	_ = l.reader.UnreadRune()
}

// Peek ...
func (l *Lexer) Peek() rune {
	ch, _, err := l.reader.ReadRune()
	if err != nil {
		ch = eof
	}
	_ = l.reader.UnreadRune()
	return ch
}

// scanToLowerWord scan a word and returns its value in lower letters
func (l *Lexer) scanToLowerWord(r rune) string {
	var buf bytes.Buffer
	buf.WriteRune(unicode.ToLower(r))
	r = l.read()
	for unicode.IsLetter(r) {
		buf.WriteRune(unicode.ToLower(r))
		r = l.read()
	}
	l.unread()
	return buf.String()
}

//TODO
// // scanInt scan a string representing a int
// func (l *Lexer) scanInt(r rune) string {
// 	var buf bytes.Buffer
// 	buf.WriteRune(r)
// 	r = l.read()
// 	for beginFloat(r) {
// 		buf.WriteRune(r)
// 		r = l.read()
// 	}
// 	l.unread()
// 	return buf.String()
// }

// scanFloat scan a string representing a float
func (l *Lexer) scanFloat(r rune) string {
	var buf bytes.Buffer
	buf.WriteRune(r)
	r = l.read()
	for isFloatRune(r) {
		buf.WriteRune(r)
		r = l.read()
	}
	l.unread()
	return buf.String()
}

// scanToken scans the next lexeme
// return false is eof is reached true otherwise
// error is non nil only in case of unexpected character or word
func (l *Lexer) scanToken() (Token, error) {
	r := l.read()
	switch {
	case unicode.IsSpace(r):
		l.pos++
		return l.scanToken()
	case r == '(':
		return l.getToken(LeftParen, "("), nil
	case r == ')':
		return l.getToken(RightParen, ")"), nil
	case r == ',':
		return l.getToken(Comma, ","), nil
	case r == '=':
		return l.getToken(EqualSign, "="), nil
	case unicode.IsLetter(r):
		w := l.scanToLowerWord(r)
		switch w {
		case "empty":
			return l.getToken(Empty, "empty"), nil
		case "z":
			return l.getToken(Z, "z"), nil
		case "m":
			return l.getToken(M, "m"), nil
		case "zm":
			return l.getToken(ZM, "zm"), nil
		case "point":
			return l.getToken(PointEnum, "point"), nil
		case "linestring":
			return l.getToken(Linestring, "linestring"), nil
		case "polygon":
			return l.getToken(PolygonEnum, "polygon"), nil
		case "multipoint":
			return l.getToken(Multipoint, "multipoint"), nil
		case "multilinestring":
			return l.getToken(MultilineString, "multilinestring"), nil
		case "multipolygon":
			return l.getToken(MultiPolygonEnum, "multipolygon"), nil
		case "geometrycollection":
			return l.getToken(GeometryCollection, "geometrycollection"), nil
		case "srid":
			return l.getToken(Srid, "srid"), nil
		default:
			return Token{}, fmt.Errorf("Unexpected word %s on character %d", w, l.pos)
		}
	case beginFloat(r):
		w := l.scanFloat(r)
		return l.getToken(Float, w), nil
	case r == eof:
		return l.getToken(EOF, ""), nil
	default:
		return Token{}, fmt.Errorf("Unexpected rune %s on character %d", string(r), l.pos)
	}
}

//TODO
// func beginInt(r rune) bool {
// 	return unicode.IsNumber(r)
// }

func beginFloat(r rune) bool {
	return r == '-' || r == '.' || unicode.IsNumber(r)
}

func isFloatRune(r rune) bool {
	return beginFloat(r) || r == 'e'
}
