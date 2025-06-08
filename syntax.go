// SyntaxRule 構文規則を表す
type SyntaxRule struct {
    Language string
    RuleName string
    Sequence []string
}

var syntaxRules = []SyntaxRule{
    {Language: "qitlang", RuleName: "assign", Sequence: []string{"X", "=", "Y"}},
    {Language: "qitlang", RuleName: "apply", Sequence: []string{"apply", "G", "Q"}},
    {Language: "qitlang", RuleName: "measure", Sequence: []string{"measure", "Q"}},
}

// ValidateSyntax 構文の有効性をチェック
func ValidateSyntax(lang, rule string, tokens []string) bool {
    for _, r := range syntaxRules {
        if r.Language == lang && r.RuleName == rule {
            if len(r.Sequence) != len(tokens) {
                return false
            }
            for i, seq := range r.Sequence {
                if seq == "X" || seq == "Y" || seq == "G" || seq == "Q" {
                    continue // 変数は任意のトークンを許可
                }
                if seq != tokens[i] {
                    return false
                }
            }
            return true
        }
    }
    return false
}
