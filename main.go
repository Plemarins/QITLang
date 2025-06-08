func main() {
    code := `
    qubit q = |0>
    apply H q
    measure q
    `
    tokens := Lexer(code)
    nodes, err := Parse(tokens)
    if err != nil {
        fmt.Println("Parse error:", err)
        return
    }
    ctx := NewContext()
    if err := Evaluate(nodes, ctx); err != nil {
        fmt.Println("Evaluate error:", err)
        return
    }
}
