package ast

import (
	"fmt"
)

type functionArg struct {
	Type Type
	Kind Kind
}

func (f functionArg) matches(arg Expr) bool {
	if f.Type != arg.Type() {
		return false
	}

	if f.Kind == KindUndefined {
		return true
	}

	if designator, ok := arg.(*DesignatorExpr); ok {
		return f.Kind == designator.Kind()
	}

	return true
}

type functionOverload struct {
	returnType Type
	args       []functionArg
}

func (f functionOverload) matches(args []Expr) (Type, bool) {
	if len(f.args) != len(args) {
		return TypeVoid, false
	}

	for i := 0; i < len(args); i++ {
		if !f.args[i].matches(args[i]) {
			return TypeVoid, false
		}
	}

	return f.returnType, true
}

type function struct {
	overloads []functionOverload
}

func newFunction(overloads []functionOverload) function {
	return function{overloads}
}

func (f function) Validate(args []Expr) (Type, error) {
	for _, overload := range f.overloads {
		if t, ok := overload.matches(args); ok {
			return t, nil
		}
	}

	return TypeVoid, fmt.Errorf("no matching overload found")
}

func (f function) Type() Type {
	// TODO(daniel): since the builtin functions may have a different return type based on the
	// overload, this is incorrect!
	if len(f.overloads) == 0 {
		return TypeVoid
	}

	return f.overloads[0].returnType
}

type Function interface {
	Validate([]Expr) (Type, error)
	Visit(BuiltinVisitor, []Expr)
	Type() Type
}

type BuiltinVisitor interface {
	VisitBuiltinPrint(*BuiltinPrint, []Expr)
	VisitBuiltinABS(*BuiltinABS, []Expr)
	VisitBuiltinODD(*BuiltinODD, []Expr)
	VisitBuiltinLSL(*BuiltinLSL, []Expr)
	VisitBuiltinASR(*BuiltinASR, []Expr)
	VisitBuiltinROR(*BuiltinROR, []Expr)
	VisitBuiltinLEN(*BuiltinLEN, []Expr)
	VisitBuiltinORD(*BuiltinORD, []Expr)
	VisitBuiltinCHR(*BuiltinCHR, []Expr)
	VisitBuiltinFLOOR(*BuiltinFLOOR, []Expr)
	VisitBuiltinFLT(*BuiltinFLT, []Expr)
	VisitBuiltinINC(*BuiltinINC, []Expr)
	VisitBuiltinDEC(*BuiltinDEC, []Expr)
	VisitBuiltinINCL(*BuiltinINCL, []Expr)
	VisitBuiltinEXCL(*BuiltinEXCL, []Expr)
	VisitBuiltinPACK(*BuiltinPACK, []Expr)
	VisitBuiltinUNPK(*BuiltinUNPK, []Expr)
	VisitBuiltinNEW(*BuiltinNEW, []Expr)
	VisitBuiltinASSERT(*BuiltinASSERT, []Expr)
}

func GetBuiltinFunctions() map[string]Function {
	return map[string]Function{
		"print":  newBuiltinPrint(),
		"ABS":    newBuiltinABS(),
		"ODD":    newBuiltinODD(),
		"LSL":    newBuiltinLSL(),
		"ASR":    newBuiltinASR(),
		"ROR":    newBuiltinROR(),
		"LEN":    newBuiltinLEN(),
		"ORD":    newBuiltinORD(),
		"CHR":    newBuiltinCHR(),
		"FLOOR":  newBuiltinFLOOR(),
		"FLT":    newBuiltinFLT(),
		"INC":    newBuiltinINC(),
		"DEC":    newBuiltinDEC(),
		"INCL":   newBuiltinINCL(),
		"EXCL":   newBuiltinEXCL(),
		"PACK":   newBuiltinPACK(),
		"UNPK":   newBuiltinUNPK(),
		"NEW":    newBuiltinNEW(),
		"ASSERT": newBuiltinASSERT(),
	}
}

// ----- print ----------------------------------------------------------------

type BuiltinPrint struct {
	function
}

var _ Function = (*BuiltinPrint)(nil)

func newBuiltinPrint() *BuiltinPrint {
	return &BuiltinPrint{
		function: newFunction([]functionOverload{
			{
				returnType: TypeVoid,
				args: []functionArg{
					{Type: TypeBoolean, Kind: KindUndefined},
				},
			},
			{
				returnType: TypeVoid,
				args: []functionArg{
					{Type: TypeInt64, Kind: KindUndefined},
				},
			},
			{
				returnType: TypeVoid,
				args: []functionArg{
					{Type: TypeFloat64, Kind: KindUndefined},
				},
			},
			{
				returnType: TypeVoid,
				args: []functionArg{
					{Type: TypeChar, Kind: KindUndefined},
				},
			},
			{
				returnType: TypeVoid,
				args: []functionArg{
					{Type: TypeString, Kind: KindUndefined},
				},
			},
			{
				returnType: TypeVoid,
				args: []functionArg{
					{Type: TypeSet, Kind: KindUndefined},
				},
			},
		}),
	}
}

func (f *BuiltinPrint) Visit(v BuiltinVisitor, args []Expr) {
	v.VisitBuiltinPrint(f, args)
}

// ----- ABS ------------------------------------------------------------------

type BuiltinABS struct {
	function
}

var _ Function = (*BuiltinABS)(nil)

func newBuiltinABS() *BuiltinABS {
	return &BuiltinABS{
		function: newFunction([]functionOverload{
			{
				returnType: TypeInt64,
				args: []functionArg{
					{Type: TypeInt64, Kind: KindUndefined},
				},
			},
			{
				returnType: TypeFloat64,
				args: []functionArg{
					{Type: TypeFloat64, Kind: KindUndefined},
				},
			},
		}),
	}
}

func (f *BuiltinABS) Visit(v BuiltinVisitor, args []Expr) {
	v.VisitBuiltinABS(f, args)
}

// ----- ODD ------------------------------------------------------------------

type BuiltinODD struct {
	function
}

var _ Function = (*BuiltinODD)(nil)

func newBuiltinODD() *BuiltinODD {
	return &BuiltinODD{
		function: newFunction([]functionOverload{
			{
				returnType: TypeBoolean,
				args: []functionArg{
					{Type: TypeInt64, Kind: KindUndefined},
				},
			},
		}),
	}
}

func (f *BuiltinODD) Visit(v BuiltinVisitor, args []Expr) {
	v.VisitBuiltinODD(f, args)
}

// ----- LSL ------------------------------------------------------------------

type BuiltinLSL struct {
	function
}

var _ Function = (*BuiltinLSL)(nil)

func newBuiltinLSL() *BuiltinLSL {
	return &BuiltinLSL{
		function: newFunction([]functionOverload{
			{
				returnType: TypeInt64,
				args: []functionArg{
					{Type: TypeInt64, Kind: KindUndefined},
					{Type: TypeInt64, Kind: KindUndefined},
				},
			},
		}),
	}
}

func (f *BuiltinLSL) Visit(v BuiltinVisitor, args []Expr) {
	v.VisitBuiltinLSL(f, args)
}

// ----- ASR ------------------------------------------------------------------

type BuiltinASR struct {
	function
}

var _ Function = (*BuiltinASR)(nil)

func newBuiltinASR() *BuiltinASR {
	return &BuiltinASR{
		function: newFunction([]functionOverload{
			{
				returnType: TypeInt64,
				args: []functionArg{
					{Type: TypeInt64, Kind: KindUndefined},
					{Type: TypeInt64, Kind: KindUndefined},
				},
			},
		}),
	}
}

func (f *BuiltinASR) Visit(v BuiltinVisitor, args []Expr) {
	v.VisitBuiltinASR(f, args)
}

// ----- ROR ------------------------------------------------------------------

type BuiltinROR struct {
	function
}

var _ Function = (*BuiltinROR)(nil)

func newBuiltinROR() *BuiltinROR {
	return &BuiltinROR{
		function: newFunction([]functionOverload{
			{
				returnType: TypeInt64,
				args: []functionArg{
					{Type: TypeInt64, Kind: KindUndefined},
					{Type: TypeInt64, Kind: KindUndefined},
				},
			},
		}),
	}
}

func (f *BuiltinROR) Visit(v BuiltinVisitor, args []Expr) {
	v.VisitBuiltinROR(f, args)
}

// ----- LEN ------------------------------------------------------------------

type BuiltinLEN struct {
	function
}

var _ Function = (*BuiltinLEN)(nil)

func newBuiltinLEN() *BuiltinLEN {
	return &BuiltinLEN{
		function: newFunction([]functionOverload{
			{
				returnType: TypeInt64,
				args: []functionArg{
					{Type: TypeArray, Kind: KindUndefined},
				},
			},
			{
				returnType: TypeInt64,
				args: []functionArg{
					{Type: TypeArray, Kind: KindUndefined},
					{Type: TypeInt64, Kind: KindUndefined},
				},
			},
		}),
	}
}

func (f *BuiltinLEN) Visit(v BuiltinVisitor, args []Expr) {
	v.VisitBuiltinLEN(f, args)
}

// ----- ORD ------------------------------------------------------------------

type BuiltinORD struct {
	function
}

var _ Function = (*BuiltinORD)(nil)

func newBuiltinORD() *BuiltinORD {
	return &BuiltinORD{
		function: newFunction([]functionOverload{
			{
				returnType: TypeInt64,
				args: []functionArg{
					{Type: TypeChar, Kind: KindUndefined},
				},
			},
			{
				returnType: TypeInt64,
				args: []functionArg{
					{Type: TypeBoolean, Kind: KindUndefined},
				},
			},
		}),
	}
}

func (f *BuiltinORD) Visit(v BuiltinVisitor, args []Expr) {
	v.VisitBuiltinORD(f, args)
}

// ----- CHR ------------------------------------------------------------------

type BuiltinCHR struct {
	function
}

var _ Function = (*BuiltinCHR)(nil)

func newBuiltinCHR() *BuiltinCHR {
	return &BuiltinCHR{
		function: newFunction([]functionOverload{
			{
				returnType: TypeChar,
				args: []functionArg{
					{Type: TypeInt64, Kind: KindUndefined},
				},
			},
		}),
	}
}

func (f *BuiltinCHR) Visit(v BuiltinVisitor, args []Expr) {
	v.VisitBuiltinCHR(f, args)
}

// ----- FLOOR ----------------------------------------------------------------

type BuiltinFLOOR struct {
	function
}

var _ Function = (*BuiltinFLOOR)(nil)

func newBuiltinFLOOR() *BuiltinFLOOR {
	return &BuiltinFLOOR{
		function: newFunction([]functionOverload{
			{
				returnType: TypeInt64,
				args: []functionArg{
					{Type: TypeFloat64, Kind: KindUndefined},
				},
			},
		}),
	}
}

func (f *BuiltinFLOOR) Visit(v BuiltinVisitor, args []Expr) {
	v.VisitBuiltinFLOOR(f, args)
}

// ----- FLT ------------------------------------------------------------------

type BuiltinFLT struct {
	function
}

var _ Function = (*BuiltinFLT)(nil)

func newBuiltinFLT() *BuiltinFLT {
	return &BuiltinFLT{
		function: newFunction([]functionOverload{
			{
				returnType: TypeFloat64,
				args: []functionArg{
					{Type: TypeInt64, Kind: KindUndefined},
				},
			},
		}),
	}
}

func (f *BuiltinFLT) Visit(v BuiltinVisitor, args []Expr) {
	v.VisitBuiltinFLT(f, args)
}

// ----- INC ------------------------------------------------------------------

type BuiltinINC struct {
	function
}

var _ Function = (*BuiltinINC)(nil)

func newBuiltinINC() *BuiltinINC {
	return &BuiltinINC{
		function: newFunction([]functionOverload{
			{
				returnType: TypeVoid,
				args: []functionArg{
					{Type: TypeInt64, Kind: KindVar},
				},
			},
			{
				returnType: TypeVoid,
				args: []functionArg{
					{Type: TypeInt64, Kind: KindVar},
					{Type: TypeInt64, Kind: KindUndefined},
				},
			},
		}),
	}
}

func (f *BuiltinINC) Visit(v BuiltinVisitor, args []Expr) {
	v.VisitBuiltinINC(f, args)
}

// ----- DEC ------------------------------------------------------------------

type BuiltinDEC struct {
	function
}

var _ Function = (*BuiltinDEC)(nil)

func newBuiltinDEC() *BuiltinDEC {
	return &BuiltinDEC{
		function: newFunction([]functionOverload{
			{
				returnType: TypeVoid,
				args: []functionArg{
					{Type: TypeInt64, Kind: KindVar},
				},
			},
			{
				returnType: TypeVoid,
				args: []functionArg{
					{Type: TypeInt64, Kind: KindVar},
					{Type: TypeInt64, Kind: KindUndefined},
				},
			},
		}),
	}
}

func (f *BuiltinDEC) Visit(v BuiltinVisitor, args []Expr) {
	v.VisitBuiltinDEC(f, args)
}

// ----- INCL -----------------------------------------------------------------

type BuiltinINCL struct {
	function
}

var _ Function = (*BuiltinINCL)(nil)

func newBuiltinINCL() *BuiltinINCL {
	return &BuiltinINCL{
		function: newFunction([]functionOverload{
			{
				returnType: TypeVoid,
				args: []functionArg{
					{Type: TypeSet, Kind: KindVar},
					{Type: TypeInt64, Kind: KindUndefined},
				},
			},
		}),
	}
}

func (f *BuiltinINCL) Visit(v BuiltinVisitor, args []Expr) {
	v.VisitBuiltinINCL(f, args)
}

// ----- EXCL -----------------------------------------------------------------

type BuiltinEXCL struct {
	function
}

var _ Function = (*BuiltinEXCL)(nil)

func newBuiltinEXCL() *BuiltinEXCL {
	return &BuiltinEXCL{
		function: newFunction([]functionOverload{
			{
				returnType: TypeVoid,
				args: []functionArg{
					{Type: TypeSet, Kind: KindVar},
					{Type: TypeInt64, Kind: KindUndefined},
				},
			},
		}),
	}
}

func (f *BuiltinEXCL) Visit(v BuiltinVisitor, args []Expr) {
	v.VisitBuiltinEXCL(f, args)
}

// ----- PACK -----------------------------------------------------------------

type BuiltinPACK struct {
	function
}

var _ Function = (*BuiltinPACK)(nil)

func newBuiltinPACK() *BuiltinPACK {
	return &BuiltinPACK{
		function: newFunction([]functionOverload{
			{
				returnType: TypeVoid,
				args: []functionArg{
					{Type: TypeFloat64, Kind: KindVar},
					{Type: TypeInt64, Kind: KindUndefined},
				},
			},
		}),
	}
}

func (f *BuiltinPACK) Visit(v BuiltinVisitor, args []Expr) {
	v.VisitBuiltinPACK(f, args)
}

// ----- UNPK -----------------------------------------------------------------

type BuiltinUNPK struct {
	function
}

var _ Function = (*BuiltinUNPK)(nil)

func newBuiltinUNPK() *BuiltinUNPK {
	return &BuiltinUNPK{
		function: newFunction([]functionOverload{
			{
				returnType: TypeVoid,
				args: []functionArg{
					{Type: TypeFloat64, Kind: KindVar},
					{Type: TypeInt64, Kind: KindUndefined},
				},
			},
		}),
	}
}

func (f *BuiltinUNPK) Visit(v BuiltinVisitor, args []Expr) {
	v.VisitBuiltinUNPK(f, args)
}

// ----- NEW ------------------------------------------------------------------

type BuiltinNEW struct {
	function
}

var _ Function = (*BuiltinNEW)(nil)

func newBuiltinNEW() *BuiltinNEW {
	return &BuiltinNEW{
		function: newFunction([]functionOverload{
			{
				returnType: TypeVoid,
				args: []functionArg{
					{Type: TypePointer, Kind: KindVar},
				},
			},
		}),
	}
}

func (f *BuiltinNEW) Visit(v BuiltinVisitor, args []Expr) {
	v.VisitBuiltinNEW(f, args)
}

// ----- ASSERT ---------------------------------------------------------------

type BuiltinASSERT struct {
	function
}

var _ Function = (*BuiltinASSERT)(nil)

func newBuiltinASSERT() *BuiltinASSERT {
	return &BuiltinASSERT{
		function: newFunction([]functionOverload{
			{
				returnType: TypeVoid,
				args: []functionArg{
					{Type: TypeBoolean, Kind: KindUndefined},
				},
			},
		}),
	}
}

func (f *BuiltinASSERT) Visit(v BuiltinVisitor, args []Expr) {
	v.VisitBuiltinASSERT(f, args)
}
