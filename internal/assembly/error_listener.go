package assembly

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

// ErrorListener обрабатывает ошибки парсера
type ErrorListener struct {
	*antlr.DefaultErrorListener
	errors []string
}

// NewErrorListener создает новый error listener
func NewErrorListener() *ErrorListener {
	return &ErrorListener{
		errors: make([]string, 0),
	}
}

// SyntaxError обрабатывает синтаксические ошибки
func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, 
	line, column int, msg string, e antlr.RecognitionException) {
	errorMsg := fmt.Sprintf("line %d:%d %s", line, column, msg)
	l.errors = append(l.errors, errorMsg)
}

// HasErrors проверяет наличие ошибок
func (l *ErrorListener) HasErrors() bool {
	return len(l.errors) > 0
}

// ErrorMessages возвращает все сообщения об ошибках
func (l *ErrorListener) ErrorMessages() []string {
	return l.errors
}