package main

import (
	"fmt"
)

type functionArg struct {
	Type Type
	Kind Kind
}

type functionOverload struct {
	ReturnType Type
	Args       []functionArg
}

type function struct {
	Overloads []functionOverload
}

func newFunction(overloads []functionOverload) function {
	return function{overloads}
}

func (f function) Validate(args []Expr) (Type, error) {
next:
	for _, overload := range f.Overloads {
		if len(overload.Args) == len(args) {
			for i := 0; i < len(args); i++ {
				if !argMatches(overload.Args[i], args[i]) {
					continue next
				}
			}

			return overload.ReturnType, nil
		}
	}

	return TypeVoid, fmt.Errorf("no matching overload found")
}

func argMatches(param functionArg, arg Expr) bool {
	if param.Type != arg.Type() {
		return false
	}

	if param.Kind == KindUndefined {
		return true
	}

	if designator, ok := arg.(*DesignatorExpr); ok {
		return param.Kind == designator.Kind()
	}

	return true
}

type Function interface {
	Validate([]Expr) (Type, error)
	ConstValue([]Expr) *Value
}

type Builtins map[string]Function

func (b Builtins) register(name string, function Function) {
	b[name] = function
}

//nolint:gochecknoglobals
var builtin = Builtins{}

func registerBuiltins() {
	builtin.register("print", newBuiltinPrint())
	builtin.register("ABS", newBuiltinABS())
	builtin.register("ODD", newBuiltinODD())
	builtin.register("LSL", newBuiltinLSL())
	builtin.register("ASR", newBuiltinASR())
	builtin.register("ROR", newBuiltinROR())
	builtin.register("LEN", newBuiltinLEN())
	builtin.register("ORD", newBuiltinORD())
	builtin.register("CHR", newBuiltinCHR())
	builtin.register("FLOOR", newBuiltinFLOOR())
	builtin.register("FLT", newBuiltinFLT())
	builtin.register("INC", newBuiltinINC())
	builtin.register("DEC", newBuiltinDEC())
	builtin.register("INCL", newBuiltinINCL())
	builtin.register("EXCL", newBuiltinEXCL())
	builtin.register("PACK", newBuiltinPACK())
	builtin.register("UNPK", newBuiltinUNPK())
	builtin.register("NEW", newBuiltinNEW())
	builtin.register("ASSERT", newBuiltinASSERT())
}

func init() {
	registerBuiltins()
}

type builtinPrint struct {
	function
}

var _ Function = (*builtinPrint)(nil)

func newBuiltinPrint() *builtinPrint {
	return &builtinPrint{
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

func (f *builtinPrint) ConstValue(args []Expr) *Value {
	return nil
}

type builtinABS struct {
	function
}

var _ Function = (*builtinABS)(nil)

func newBuiltinABS() *builtinABS {
	return &builtinABS{
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

func (f *builtinABS) ConstValue(args []Expr) *Value {
	c := args[0].ConstValue()
	if c == nil {
		return nil
	}

	if c.integer < 0 {
		c.integer = -c.integer
	}
	if c.real < 0 {
		c.real = -c.real
	}

	switch args[0].Type() {
	case TypeInt64:
		return &Value{
			typ:     args[0].Type(),
			integer: c.Int(),
		}
	case TypeFloat64:
		return &Value{
			typ:  args[0].Type(),
			real: c.Real(),
		}
	default:
		return nil
	}
}

type builtinODD struct {
	function
}

var _ Function = (*builtinODD)(nil)

func newBuiltinODD() *builtinODD {
	return &builtinODD{
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

func (f *builtinODD) ConstValue(args []Expr) *Value {
	c := args[0].ConstValue()
	if c == nil {
		return nil
	}

	return &Value{
		typ:     TypeBoolean,
		boolean: c.Int()%2 == 1,
	}
}

type builtinLSL struct {
	function
}

var _ Function = (*builtinLSL)(nil)

func newBuiltinLSL() *builtinLSL {
	return &builtinLSL{
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

func (f *builtinLSL) ConstValue(args []Expr) *Value {
	panic("builtinLSL.ConstValue is not implemented")
}

type builtinASR struct {
	function
}

var _ Function = (*builtinASR)(nil)

func newBuiltinASR() *builtinASR {
	return &builtinASR{
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

func (f *builtinASR) ConstValue(args []Expr) *Value {
	panic("builtinASR.ConstValue is not implemented")
}

type builtinROR struct {
	function
}

var _ Function = (*builtinROR)(nil)

func newBuiltinROR() *builtinROR {
	return &builtinROR{
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

func (f *builtinROR) ConstValue(args []Expr) *Value {
	panic("builtinROR.ConstValue is not implemented")
}

type builtinLEN struct {
	function
}

var _ Function = (*builtinLEN)(nil)

func newBuiltinLEN() *builtinLEN {
	return &builtinLEN{
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

func (f *builtinLEN) ConstValue(args []Expr) *Value {
	panic("builtinLEN.ConstValue is not implemented")
}

type builtinORD struct {
	function
}

var _ Function = (*builtinORD)(nil)

func newBuiltinORD() *builtinORD {
	return &builtinORD{
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

func (f *builtinORD) ConstValue(args []Expr) *Value {
	c := args[0].ConstValue()
	if c == nil {
		return nil
	}

	switch args[0].Type() {
	case TypeBoolean, TypeChar:
		return &Value{
			typ:     TypeInt64,
			integer: c.Int(),
		}
	default:
		return nil
	}
}

type builtinCHR struct {
	function
}

var _ Function = (*builtinCHR)(nil)

func newBuiltinCHR() *builtinCHR {
	return &builtinCHR{
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

func (f *builtinCHR) ConstValue(args []Expr) *Value {
	if c := args[0].ConstValue(); c != nil {
		return &Value{
			typ:  TypeChar,
			char: c.Char(),
		}
	}

	return nil
}

type builtinFLOOR struct {
	function
}

var _ Function = (*builtinFLOOR)(nil)

func newBuiltinFLOOR() *builtinFLOOR {
	return &builtinFLOOR{
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

func (f *builtinFLOOR) ConstValue(args []Expr) *Value {
	if c := args[0].ConstValue(); c != nil {
		return &Value{
			typ:     TypeInt64,
			integer: c.Int(),
		}
	}

	return nil
}

type builtinFLT struct {
	function
}

var _ Function = (*builtinFLT)(nil)

func newBuiltinFLT() *builtinFLT {
	return &builtinFLT{
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

func (f *builtinFLT) ConstValue(args []Expr) *Value {
	if c := args[0].ConstValue(); c != nil {
		return &Value{
			typ:  TypeFloat64,
			real: c.Real(),
		}
	}

	return nil
}

type builtinINC struct {
	function
}

var _ Function = (*builtinINC)(nil)

func newBuiltinINC() *builtinINC {
	return &builtinINC{
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

func (f *builtinINC) ConstValue(args []Expr) *Value {
	return nil
}

type builtinDEC struct {
	function
}

var _ Function = (*builtinDEC)(nil)

func newBuiltinDEC() *builtinDEC {
	return &builtinDEC{
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

func (f *builtinDEC) ConstValue(args []Expr) *Value {
	return nil
}

type builtinINCL struct {
	function
}

var _ Function = (*builtinINCL)(nil)

func newBuiltinINCL() *builtinINCL {
	return &builtinINCL{
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

func (f *builtinINCL) ConstValue(args []Expr) *Value {
	set := args[0].ConstValue()
	val := args[1].ConstValue()

	// TODO(daniel): I don't think this does anything?
	if set != nil && val != nil {
		set.integer |= 1 << val.Int()
	}

	return nil
}

type builtinEXCL struct {
	function
}

var _ Function = (*builtinEXCL)(nil)

func newBuiltinEXCL() *builtinEXCL {
	return &builtinEXCL{
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

func (f *builtinEXCL) ConstValue(args []Expr) *Value {
	set := args[0].ConstValue()
	val := args[1].ConstValue()

	// TODO(daniel): I don't think this does anything?
	if set != nil && val != nil {
		set.integer &= ^(1 << val.Int())
	}

	return nil
}

type builtinPACK struct {
	function
}

var _ Function = (*builtinPACK)(nil)

func newBuiltinPACK() *builtinPACK {
	return &builtinPACK{
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

func (f *builtinPACK) ConstValue(args []Expr) *Value {
	return nil
}

type builtinUNPK struct {
	function
}

var _ Function = (*builtinUNPK)(nil)

func newBuiltinUNPK() *builtinUNPK {
	return &builtinUNPK{
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

func (f *builtinUNPK) ConstValue(args []Expr) *Value {
	return nil
}

type builtinNEW struct {
	function
}

var _ Function = (*builtinNEW)(nil)

func newBuiltinNEW() *builtinNEW {
	return &builtinNEW{
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

func (f *builtinNEW) ConstValue(args []Expr) *Value {
	return nil
}

type builtinASSERT struct {
	function
}

var _ Function = (*builtinASSERT)(nil)

func newBuiltinASSERT() *builtinASSERT {
	return &builtinASSERT{
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

func (f *builtinASSERT) ConstValue(args []Expr) *Value {
	return nil
}
