// SemanticMap 意味写像を表す
type SemanticMap struct {
    Language string
    Word     string
    Meaning  string
}

var semanticMaps = []SemanticMap{
    {Language: "qitlang", Word: "apply", Meaning: "unitary_transform"},
    {Language: "qitlang", Word: "measure", Meaning: "probabilistic_collapse"},
    {Language: "qitlang", Word: "qubit", Meaning: "quantum_state"},
}

// GetSemanticMap 意味を取得
func GetSemanticMap(lang, word string) (string, bool) {
    for _, sm := range semanticMaps {
        if sm.Language == lang && sm.Word == word {
            return sm.Meaning, true
        }
    }
    return "", false
}
