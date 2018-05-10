// Code generated by gocc; DO NOT EDIT.

package parser

import "github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/direction"

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Func	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Func : Func1	<<  >>`,
		Id:         "Func",
		NTType:     1,
		Index:      1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Func : Expr	<<  >>`,
		Id:         "Func",
		NTType:     1,
		Index:      2,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Func1 : function_name "(" Args ")"	<< direction.NewFunction(X[0], X[2]) >>`,
		Id:         "Func1",
		NTType:     2,
		Index:      3,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewFunction(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Func1 : function_name "()"	<< direction.NewFunction(X[0], "") >>`,
		Id:         "Func1",
		NTType:     2,
		Index:      4,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewFunction(X[0], "")
		},
	},
	ProdTabEntry{
		String: `Args : DoubleQString	<< direction.NewArgument(X[0]) >>`,
		Id:         "Args",
		NTType:     3,
		Index:      5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewArgument(X[0])
		},
	},
	ProdTabEntry{
		String: `Args : SingleQString	<< direction.NewArgument(X[0]) >>`,
		Id:         "Args",
		NTType:     3,
		Index:      6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewArgument(X[0])
		},
	},
	ProdTabEntry{
		String: `Args : Int	<< direction.NewArgument(X[0]) >>`,
		Id:         "Args",
		NTType:     3,
		Index:      7,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewArgument(X[0])
		},
	},
	ProdTabEntry{
		String: `Args : Float	<< direction.NewArgument(X[0]) >>`,
		Id:         "Args",
		NTType:     3,
		Index:      8,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewArgument(X[0])
		},
	},
	ProdTabEntry{
		String: `Args : Bool	<< direction.NewArgument(X[0]) >>`,
		Id:         "Args",
		NTType:     3,
		Index:      9,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewArgument(X[0])
		},
	},
	ProdTabEntry{
		String: `Args : MappingRef	<< direction.NewArgument(X[0]) >>`,
		Id:         "Args",
		NTType:     3,
		Index:      10,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewArgument(X[0])
		},
	},
	ProdTabEntry{
		String: `Args : Func1	<< direction.NewArgument(X[0]) >>`,
		Id:         "Args",
		NTType:     3,
		Index:      11,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewArgument(X[0])
		},
	},
	ProdTabEntry{
		String: `Args : Args delimitor_param Args	<< direction.NewArguments(X[0], X[2]) >>`,
		Id:         "Args",
		NTType:     3,
		Index:      12,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewArguments(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `DoubleQString : doublequotes_string	<< direction.NewDoubleQuoteStringLit(X[0]) >>`,
		Id:         "DoubleQString",
		NTType:     4,
		Index:      13,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewDoubleQuoteStringLit(X[0])
		},
	},
	ProdTabEntry{
		String: `SingleQString : singlequote_string	<< direction.NewSingleQuoteStringLit(X[0]) >>`,
		Id:         "SingleQString",
		NTType:     5,
		Index:      14,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewSingleQuoteStringLit(X[0])
		},
	},
	ProdTabEntry{
		String: `Int : number	<< direction.NewIntLit(X[0]) >>`,
		Id:         "Int",
		NTType:     6,
		Index:      15,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewIntLit(X[0])
		},
	},
	ProdTabEntry{
		String: `MappingRef : argument	<< direction.NewMappingRef(X[0]) >>`,
		Id:         "MappingRef",
		NTType:     7,
		Index:      16,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewMappingRef(X[0])
		},
	},
	ProdTabEntry{
		String: `Bool : "true"	<< direction.NewBool(X[0]) >>`,
		Id:         "Bool",
		NTType:     8,
		Index:      17,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewBool(X[0])
		},
	},
	ProdTabEntry{
		String: `Bool : "false"	<< direction.NewBool(X[0]) >>`,
		Id:         "Bool",
		NTType:     8,
		Index:      18,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewBool(X[0])
		},
	},
	ProdTabEntry{
		String: `Float : float	<< direction.NewFloatLit(X[0]) >>`,
		Id:         "Float",
		NTType:     9,
		Index:      19,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewFloatLit(X[0])
		},
	},
	ProdTabEntry{
		String: `Expr : Int	<< direction.NewExpressionField(X[0]) >>`,
		Id:         "Expr",
		NTType:     10,
		Index:      20,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewExpressionField(X[0])
		},
	},
	ProdTabEntry{
		String: `Expr : Float	<< direction.NewExpressionField(X[0]) >>`,
		Id:         "Expr",
		NTType:     10,
		Index:      21,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewExpressionField(X[0])
		},
	},
	ProdTabEntry{
		String: `Expr : DoubleQString	<< direction.NewExpressionField(X[0]) >>`,
		Id:         "Expr",
		NTType:     10,
		Index:      22,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewExpressionField(X[0])
		},
	},
	ProdTabEntry{
		String: `Expr : SingleQString	<< direction.NewExpressionField(X[0]) >>`,
		Id:         "Expr",
		NTType:     10,
		Index:      23,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewExpressionField(X[0])
		},
	},
	ProdTabEntry{
		String: `Expr : MappingRef	<< direction.NewExpressionField(X[0]) >>`,
		Id:         "Expr",
		NTType:     10,
		Index:      24,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewExpressionField(X[0])
		},
	},
	ProdTabEntry{
		String: `Expr : Func1	<< direction.NewExpressionField(X[0]) >>`,
		Id:         "Expr",
		NTType:     10,
		Index:      25,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewExpressionField(X[0])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr Operator Expr	<< direction.NewExpression(X[0], X[1], X[2]) >>`,
		Id:         "Expr",
		NTType:     10,
		Index:      26,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewExpression(X[0], X[1], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : TernaryExp	<<  >>`,
		Id:         "Expr",
		NTType:     10,
		Index:      27,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expr : "(" Expr ")"	<< direction.NewExpressionField(X[1]) >>`,
		Id:         "Expr",
		NTType:     10,
		Index:      28,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewExpressionField(X[1])
		},
	},
	ProdTabEntry{
		String: `Operator : operator_charactor	<<  >>`,
		Id:         "Operator",
		NTType:     11,
		Index:      29,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `TernaryExp : Expr	<< direction.NewExpressionField(X[0]) >>`,
		Id:         "TernaryExp",
		NTType:     12,
		Index:      30,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewExpressionField(X[0])
		},
	},
	ProdTabEntry{
		String: `TernaryExp : TernaryExp "?" TernaryParam ":" TernaryParam	<< direction.NewTernaryExpression(X[0], X[2], X[4]) >>`,
		Id:         "TernaryExp",
		NTType:     12,
		Index:      31,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewTernaryExpression(X[0], X[2], X[4])
		},
	},
	ProdTabEntry{
		String: `TernaryParam : Int	<< direction.NewExpressionField(X[0]) >>`,
		Id:         "TernaryParam",
		NTType:     13,
		Index:      32,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewExpressionField(X[0])
		},
	},
	ProdTabEntry{
		String: `TernaryParam : Float	<< direction.NewExpressionField(X[0]) >>`,
		Id:         "TernaryParam",
		NTType:     13,
		Index:      33,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewExpressionField(X[0])
		},
	},
	ProdTabEntry{
		String: `TernaryParam : DoubleQString	<< direction.NewExpressionField(X[0]) >>`,
		Id:         "TernaryParam",
		NTType:     13,
		Index:      34,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewExpressionField(X[0])
		},
	},
	ProdTabEntry{
		String: `TernaryParam : SingleQString	<< direction.NewExpressionField(X[0]) >>`,
		Id:         "TernaryParam",
		NTType:     13,
		Index:      35,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewExpressionField(X[0])
		},
	},
	ProdTabEntry{
		String: `TernaryParam : MappingRef	<< direction.NewExpressionField(X[0]) >>`,
		Id:         "TernaryParam",
		NTType:     13,
		Index:      36,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return direction.NewExpressionField(X[0])
		},
	},
	ProdTabEntry{
		String: `TernaryParam : Func	<<  >>`,
		Id:         "TernaryParam",
		NTType:     13,
		Index:      37,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
}
