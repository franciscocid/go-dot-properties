package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeSut(text string) *tokenizer {
	return &tokenizer{
		text:   text,
		cursor: 0,
	}
}

func TestTokenization(t *testing.T) {
	// 	t.Run("", func(t *testing.T) {
	// 		t.Skip()
	// 		testFile := `
	// # You are reading the ".properties" entry.
	// ! The exclamation mark can also mark text as comments.
	// 	# The key characters =, and : should be written with
	// # a preceding backslash to ensure that they are properly loaded.
	// # However, there is no need to precede the value characters =, and : by a backslash.
	// website = https://en.wikipedia.org/
	// language = English
	// # The backslash below tells the application to continue reading
	// # the value onto the next line.
	// message = Welcome to \
	// 			Wikipedia!
	// # But if the number of backslashes at the end of the line is even, the next line is not included in the value. In the following example, the value for "key" is "valueOverOneLine\"
	// key = valueOverOneLine\\
	// # This line is not included in the value for "key"
	// # Add spaces to the key
	// key\ with\ spaces = This is the value that could be looked up with the key "key with spaces".
	// # The characters = and : in the key must be escaped as well:
	// key\:with\=colonAndEqualsSign = This is the value for the key "key:with=colonAndEqualsSign"
	// # Unicode
	// tab : \u0009
	// # If you want your property to include a backslash, it should be escaped by another backslash
	// path=c:\\wiki\\templates
	// # However, some editors will handle this automatically
	// `
	// 		tokens := makeSut().Tokenize(testFile)
	// 		for _, token := range tokens {
	// 			fmt.Printf("%s ", token.Text)
	// 		}
	// 		fmt.Println()
	// 		fmt.Println()
	// 		fmt.Println(tokens)
	// 	})

	t.Run("Ignore commentaries (lines that begin w/ # or !)", func(t *testing.T) {
		testString := `
	# You are reading the ".properties" entry.
	! The exclamation mark can also mark text as comments.
		# The key characters =, and : should be written with
	    # a preceding backslash to ensure that they are properly loaded.
	# However, there is no need to precede the value characters =, and : by a backslash.
	`
		tokens := makeSut(testString).Tokenize()
		assert.Equal(t, 0, len(tokens), "No token should be created, but %d were made: %v", len(tokens), tokens)
	})

	t.Run("Tokenize identifier, separator and value w/ single line properties", func(t *testing.T) {
		testString := `
	website = https://en.wikipedia.org/
	language = English
	`
		tokens := makeSut(testString).Tokenize()
		assert.Equal(t, 6, len(tokens), "6 token should be created, but %d were made: %v", len(tokens), tokens)
	})

	// 	t.Run("Tokenize identifier, separator and value w/ multi line properties", func(t *testing.T) {
	// 		testString := `
	// # The backslash below tells the application to continue reading
	// # the value onto the next line.
	// message = Welcome to \
	// 			Wikipedia!
	// `
	// 		tokens := makeSut(testString).Tokenize()
	// 		fmt.Println(tokens)
	// 		assert.Equal(t, 3, len(tokens), "3 token should be created, but %d were made: %v", len(tokens), tokens)
	// 	})
}
