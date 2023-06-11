/* This file contains the 'lexer' object which turns characters from the
 * 'source.Source' object into tokens that represent the symbols and values
 * found in the file. These tokens will next be used by the parser to creat an ast
 */
package lexer

import (
	"turtlego/message"
	"turtlego/source"
	"turtlego/tokens"
)

type Lexer struct {
	src        *source.Source
	CurTok     tokens.Token
	singleChar map[string]byte
	doubleChar map[string]byte
	keywords   map[string]byte
}

func New(src *source.Source) *Lexer {
	l := &Lexer{}
	l.src = src
	return l
}

func (l *Lexer) SetMaps(singleChr, doubleChr, keyword map[string]byte) {
	l.singleChar = singleChr
	l.doubleChar = doubleChr
	l.keywords = keyword
	l.MoveUp() //Fill in 'l.CurTok' value to be used by parser
}

func (l *Lexer) MoveUp() {
	l.skipWhitespace()

	focusString := l.curString() + l.nextString()
	alias, ok := l.doubleChar[focusString]

	if ok {
		l.setTok(focusString, alias)
		l.src.MoveUp()
		l.src.MoveUp()
		return
	} else if focusString == "(#" {
		l.skipComments()
		l.MoveUp()
		return
	}

	focusString = l.curString()
	alias, ok = l.singleChar[focusString]

	if ok {
		l.setTok(focusString, alias)
		l.src.MoveUp()
		return
	} else if focusString == "#" {
		l.skipComments()
		l.MoveUp()
		return
	}

	if isLetter(l.src.CurChar) {
		id := l.buildIdent()

		alias, ok = l.keywords[id]
		if ok {
			l.setTok(id, alias)
		} else {
			l.setTok(id, tokens.IDENT)
		}
		return

	} else if isDigit(l.src.CurChar) {
		num, isflt := l.buildNum()
		if isflt {
			l.setTok(num, tokens.FLT)
		} else {
			l.setTok(num, tokens.INT)
		}
		return

	} else if l.src.CurChar == '"' {
		str := l.buildStr()
		l.setTok(str, tokens.STR)
		return

	} else if l.src.CurChar == '\'' {
		chr := l.buildChar()
		l.setTok(chr, tokens.CHR)
		return

	} else {
		message.Error(l.src.Filename, "Lexer",
			"Could not lex character this value '"+string(l.src.CurChar)+"'", l.src.LineNo, l.src.ColumnNo)
	}

}

func (l *Lexer) Src() *source.Source {
	return l.src
}
