// Lexeme 語彙要素を表す
type Lexeme struct {
    Language string
    Word     string
    Category string
}

var lexemes = []Lexeme{
    {Language: "qitlang", Word: "qubit", Category: "type"},
    {Language: "qitlang", Word: "apply", Category: "verb"},
    {Language: "qitlang", Word: "H", Category: "gate"},
    {Language: "qitlang", Word: "measure", Category: "function"},
    {Language: "qitlang", Word: "|0>", Category: "state"},
}

// GetLexeme 語彙要素を取得
func GetLexeme(lang, word string) (Lexeme, bool) {
    for _, lex := range lexemes {
        if lex.Language == lang && lex.Word == word {
            return lex, true
        }
    }
    return Lexeme{}, false
}
