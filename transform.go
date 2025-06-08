// TransformRule 変換ルールを表す
type TransformRule struct {
    Language string
    Input    string
    Output   string
}

var transformRules = []TransformRule{
    {Language: "qitlang", Input: "inverse(apply(H))", Output: "apply(H)"},
}

// ApplyTransform 変換ルールを適用
func ApplyTransform(lang, input string) (string, bool) {
    for _, tr := range transformRules {
        if tr.Language == lang && tr.Input == input {
            return tr.Output, true
        }
    }
    return "", false
}
