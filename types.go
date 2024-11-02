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
	Value String
}

type HLSymbol struct {
	Value String
}

type HLKeyword struct {
	Value String
}


type HLLiteral struct {
	Value String
}

const (
	HLNil   HLLiteral = "nil"
	HLTrue  HLLiteral = "#t"
	HLFalse HLLiteral = "#f"
)

type HLList []HLValue

type HLVector []HLValue

type HLTable map[HLValue]HLValue

type HLFunction func(args ...HLType) (HLType, error)