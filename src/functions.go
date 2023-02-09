package main

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
	ReturnType Type
	Args       []functionArg
}

func (f functionOverload) matches(args []Expr) (Type, bool) {
	if len(f.Args) != len(args) {
		return TypeVoid, false
	}

	for i := 0; i < len(args); i++ {
		if !f.Args[i].matches(args[i]) {
			return TypeVoid, false
		}
	}

	return f.ReturnType, true
}

type function struct {
	Overloads []functionOverload
}

func newFunction(overloads []functionOverload) function {
	return function{overloads}
}

func (f function) Validate(args []Expr) (Type, error) {
	for _, overload := range f.Overloads {
		if t, ok := overload.matches(args); ok {
			return t, nil
		}
	}

	return TypeVoid, fmt.Errorf("no matching overload found")
}

type Function interface {
	Validate([]Expr) (Type, error)
	Visit(BuiltinVisitor, []Expr)
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

// ----- print ----------------------------------------------------------------

type BuiltinPrint struct {
	function
}

var _ Function = (*BuiltinPrint)(nil)

func newBuiltinPrint() *BuiltinPrint {
	return &BuiltinPrint{
		function: newFunction([]functionOverload{
			{
				ReturnType: TypeVoid,
				Args: []functionArg{
					{Type: TypeBoolean, Kind: KindUndefined},
				},
			},
			{
				ReturnType: TypeVoid,
				Args: []functionArg{
					{Type: TypeInt64, Kind: KindUndefined},
				},
			},
			{
				ReturnType: TypeVoid,
				Args: []functionArg{
					{Type: TypeFloat64, Kind: KindUndefined},
				},
			},
			{
				ReturnType: TypeVoid,
				Args: []functionArg{
					{Type: TypeChar, Kind: KindUndefined},
				},
			},
			{
				ReturnType: TypeVoid,
				Args: []functionArg{
					{Type: TypeString, Kind: KindUndefined},
				},
			},
			{
				ReturnType: TypeVoid,
				Args: []functionArg{
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
				ReturnType: TypeInt64,
				Args: []functionArg{
					{Type: TypeInt64, Kind: KindUndefined},
				},
			},
			{
				ReturnType: TypeFloat64,
				Args: []functionArg{
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
				ReturnType: TypeBoolean,
				Args: []functionArg{
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
				ReturnType: TypeInt64,
				Args: []functionArg{
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
				ReturnType: TypeInt64,
				Args: []functionArg{
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
				ReturnType: TypeInt64,
				Args: []functionArg{
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
				ReturnType: TypeInt64,
				Args: []functionArg{
					{Type: TypeArray, Kind: KindUndefined},
				},
			},
			{
				ReturnType: TypeInt64,
				Args: []functionArg{
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
				ReturnType: TypeInt64,
				Args: []functionArg{
					{Type: TypeChar, Kind: KindUndefined},
				},
			},
			{
				ReturnType: TypeInt64,
				Args: []functionArg{
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
				ReturnType: TypeChar,
				Args: []functionArg{
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
				ReturnType: TypeInt64,
				Args: []functionArg{
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
				ReturnType: TypeFloat64,
				Args: []functionArg{
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
				ReturnType: TypeVoid,
				Args: []functionArg{
					{Type: TypeInt64, Kind: KindVar},
				},
			},
			{
				ReturnType: TypeVoid,
				Args: []functionArg{
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
				ReturnType: TypeVoid,
				Args: []functionArg{
					{Type: TypeInt64, Kind: KindVar},
				},
			},
			{
				ReturnType: TypeVoid,
				Args: []functionArg{
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
				ReturnType: TypeVoid,
				Args: []functionArg{
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
				ReturnType: TypeVoid,
				Args: []functionArg{
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
				ReturnType: TypeVoid,
				Args: []functionArg{
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
				ReturnType: TypeVoid,
				Args: []functionArg{
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
				ReturnType: TypeVoid,
				Args: []functionArg{
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
				ReturnType: TypeVoid,
				Args: []functionArg{
					{Type: TypeBoolean, Kind: KindUndefined},
				},
			},
		}),
	}
}

func (f *BuiltinASSERT) Visit(v BuiltinVisitor, args []Expr) {
	v.VisitBuiltinASSERT(f, args)
}
