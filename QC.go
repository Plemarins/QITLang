package main

import (
    "fmt"
    "math"
    "math/rand"
    "strings"
)

// Qubit 量子ビットの状態（複素ベクトル）
type Qubit struct {
    State []complex128 // [α, β] for |ψ> = α|0> + β|1>
}

// Gate 量子ゲート（行列）
type Gate struct {
    Matrix [][]complex128
}

// Hadamardゲート
var H = Gate{
    Matrix: [][]complex128{
        {complex(1/math.Sqrt2, 0), complex(1/math.Sqrt2, 0)},
        {complex(1/math.Sqrt2, 0), complex(-1/math.Sqrt2, 0)},
    },
}

// ApplyGate ゲートを量子ビットに適用
func ApplyGate(g Gate, q *Qubit) {
    newState := make([]complex128, 2)
    for i := 0; i < 2; i++ {
        for j := 0; j < 2; j++ {
            newState[i] += g.Matrix[i][j] * q.State[j]
        }
    }
    q.State = newState
}

// Measure 量子ビットを測定
func Measure(q *Qubit) int {
    probs := []float64{
        real(q.State[0]*complex(real(q.State[0]), -imag(q.State[0]))),
        real(q.State[1]*complex(real(q.State[1]), -imag(q.State[1]))),
    }
    if rand.Float64() < probs[0] {
        q.State = []complex128{1, 0} // |0>
        return 0
    }
    q.State = []complex128{0, 1} // |1>
    return 1
}
