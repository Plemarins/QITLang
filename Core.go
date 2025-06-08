// Context 実行コンテキスト（変数と量子状態）
type Context struct {
    Variables map[string]interface{}
}

// NewContext コンテキスト初期化
func NewContext() *Context {
    return &Context{
        Variables: make(map[string]interface{}),
    }
}

// Lexer 字句解析
func Lexer(code string) []string {
    tokens := strings.Fields(code)
    for _, t := range tokens {
        if _, ok := GetLexeme("qitlang", t); !ok {
            if t != "=" && t != "|0>" {
                // 変数名は除外
                continue
            }
        }
    }
    return tokens
}

// Parser 構文解析（簡易AST生成）
type ASTNode struct {
    Type   string
    Value  string
    Args   []string
}

func Parse(tokens []string) ([]ASTNode, error) {
    var nodes []ASTNode
    i := 0
    for i < len(tokens) {
        if tokens[i] == "qubit" && i+3 < len(tokens) && tokens[i+2] == "=" {
            if !ValidateSyntax("qitlang", "assign", []string{tokens[i+1], "=", tokens[i+3]}) {
                return nil, fmt.Errorf("invalid syntax for assign")
            }
            nodes = append(nodes, ASTNode{Type: "assign", Value: tokens[i+1], Args: []string{tokens[i+3]}})
            i += 4
        } else if tokens[i] == "apply" && i+2 < len(tokens) {
            if !ValidateSyntax("qitlang", "apply", []string{"apply", tokens[i+1], tokens[i+2]}) {
                return nil, fmt.Errorf("invalid syntax for apply")
            }
            nodes = append(nodes, ASTNode{Type: "apply", Value: tokens[i+1], Args: []string{tokens[i+2]}})
            i += 3
        } else if tokens[i] == "measure" && i+1 < len(tokens) {
            if !ValidateSyntax("qitlang", "measure", []string{"measure", tokens[i+1]}) {
                return nil, fmt.Errorf("invalid syntax for measure")
            }
            nodes = append(nodes, ASTNode{Type: "measure", Value: tokens[i+1]})
            i += 2
        } else {
            return nil, fmt.Errorf("unknown token: %s", tokens[i])
        }
    }
    return nodes, nil
}

// Evaluate ASTを評価
func Evaluate(nodes []ASTNode, ctx *Context) error {
    for _, node := range nodes {
        if meaning, ok := GetSemanticMap("qitlang", node.Type); !ok || !HasLevel("qitlang", "semantics") {
            return fmt.Errorf("semantic error for %s", node.Type)
        } else {
            fmt.Printf("Semantic: %s -> %s\n", node.Type, meaning)
        }

        switch node.Type {
        case "assign":
            if node.Args[0] == "|0>" {
                ctx.Variables[node.Value] = &Qubit{State: []complex128{1, 0}}
            } else {
                return fmt.Errorf("unsupported state: %s", node.Args[0])
            }
        case "apply":
            if node.Value != "H" {
                return fmt.Errorf("unsupported gate: %s", node.Value)
            }
            q, ok := ctx.Variables[node.Args[0]].(*Qubit)
            if !ok {
                return fmt.Errorf("variable %s is not a qubit", node.Args[0])
            }
            if transformed, ok := ApplyTransform("qitlang", "inverse(apply(H))"); ok {
                fmt.Printf("Transformed: %s\n", transformed)
            }
            ApplyGate(H, q)
        case "measure":
            q, ok := ctx.Variables[node.Value].(*Qubit)
            if !ok {
                return fmt.Errorf("variable %s is not a qubit", node.Value)
            }
            result := Measure(q)
            fmt.Printf("Measurement: %d\n", result)
        }
    }
    return nil
}
