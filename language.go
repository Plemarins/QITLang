package main

// Language 言語を表す
type Language struct {
    Name   string
    Levels []string
}

// 言語とレベルの定義
var languages = []Language{
    {Name: "qitlang", Levels: []string{"syntax", "semantics", "pragmatics", "grammar", "stylistics", "meta_structure"}},
    {Name: "fortran", Levels: []string{"syntax", "semantics", "grammar"}},
}

// HasLevel 言語が特定のレベルを持つかチェック
func HasLevel(lang, level string) bool {
    for _, l := range languages {
        if l.Name == lang {
            for _, lvl := range l.Levels {
                if lvl == level {
                    return true
                }
            }
        }
    }
    return false
}
