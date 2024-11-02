package types

import (
	"math/big"
)


type HLValue interface{}

type HLBigInt struct {
	Value *big.Int
}

type HLSysInt struct {
	Value int
}

type HLByte struct {
	Value uint8
}

type HLString struct {
	Value string
}

type HLSymbol struct {
	Value string
}

type HLKeyword struct {
	Value string
}


type HLLiteral string

const (
	HLNil   HLLiteral = "Nil"
	HLTrue  HLLiteral = "#t"
	HLFalse HLLiteral = "#f"
)

type HLList []HLValue

type HLVector []HLValue

type HLTable map[HLValue]HLValue

type HLFunction func(args ...HLValue) (HLValue, error)