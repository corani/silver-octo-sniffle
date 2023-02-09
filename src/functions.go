package main

import (
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
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

func (f function) ConstValue([]Expr) (*Value, error) {
	return nil, nil
}

type Function interface {
	Validate([]Expr) (Type, error)
	ConstValue([]Expr) (*Value, error)
	Generate(*Generator, []Expr) error
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

// ----- print ----------------------------------------------------------------

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

func (f *builtinPrint) Generate(g *Generator, args []Expr) error {
	arg := args[0]

	// NOTE(daniel): we could probably do something smarter here, but for now this works.
	switch arg.Type() {
	case TypeInt64:
		f.printInteger(g, arg)
	case TypeFloat64:
		f.printReal(g, arg)
	case TypeString:
		f.printString(g, arg)
	case TypeBoolean:
		f.printBoolean(g, arg)
	case TypeChar:
		f.printChar(g, arg)
	case TypeSet:
		f.printSet(g, arg)
	default:
		return fmt.Errorf("don't know how to print type: %s", arg.Type())
	}

	return nil
}

func (f *builtinPrint) printInteger(g *Generator, arg Expr) {
	number := g.visitAndReturnValue(arg)

	format := g.internString("%d\n\000")
	formatptr := g.currentBlock.NewGetElementPtr(format.ContentType, format, zero, zero)

	g.currentValue = g.currentBlock.NewCall(g.stdlib["printf"], formatptr, number)
}

func (f *builtinPrint) printReal(g *Generator, arg Expr) {
	number := g.visitAndReturnValue(arg)

	format := g.internString("%f\n\000")
	formatptr := g.currentBlock.NewGetElementPtr(format.ContentType, format, zero, zero)

	g.currentValue = g.currentBlock.NewCall(g.stdlib["printf"], formatptr, number)
}

func (f *builtinPrint) printString(g *Generator, arg Expr) {
	str := g.visitAndReturnValue(arg)

	g.currentValue = g.currentBlock.NewCall(g.stdlib["puts"], str)
}

func (f *builtinPrint) printBoolean(g *Generator, arg Expr) {
	TRUE := g.internString("TRUE\000")
	FALSE := g.internString("FALSE\000")

	boolean := g.visitAndReturnValue(arg)

	trueBlk := g.currentFunc.NewBlock("")
	falseBlk := g.currentFunc.NewBlock("")
	endBlk := g.currentFunc.NewBlock("")

	g.currentBlock.NewCondBr(boolean, trueBlk, falseBlk)

	g.currentBlock = trueBlk
	truePtr := g.currentBlock.NewGetElementPtr(TRUE.ContentType, TRUE, zero, zero)
	g.currentBlock.NewBr(endBlk)

	g.currentBlock = falseBlk
	falsePtr := g.currentBlock.NewGetElementPtr(FALSE.ContentType, FALSE, zero, zero)
	g.currentBlock.NewBr(endBlk)

	g.currentBlock = endBlk
	strptr := g.currentBlock.NewPhi(ir.NewIncoming(truePtr, trueBlk), ir.NewIncoming(falsePtr, falseBlk))
	g.currentValue = g.currentBlock.NewCall(g.stdlib["puts"], strptr)
}

func (f *builtinPrint) printChar(g *Generator, arg Expr) {
	v := g.visitAndReturnValue(arg)

	format := g.internString("%c\n\000")
	formatptr := g.currentBlock.NewGetElementPtr(format.ContentType, format, zero, zero)

	g.currentValue = g.currentBlock.NewCall(g.stdlib["printf"], formatptr, v)
}

func (f *builtinPrint) printSet(g *Generator, arg Expr) {
	// TODO(daniel): can we get a better print for sets?
	f.printInteger(g, arg)
}

// ----- ABS ------------------------------------------------------------------

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

func (f *builtinABS) ConstValue(args []Expr) (*Value, error) {
	c := args[0].ConstValue()
	if c == nil {
		return nil, nil
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
		}, nil
	case TypeFloat64:
		return &Value{
			typ:  args[0].Type(),
			real: c.Real(),
		}, nil
	default:
		return nil, nil
	}
}

func (f *builtinABS) Generate(g *Generator, args []Expr) error {
	return fmt.Errorf("builtinABS.Generate is not implemented")
}

// ----- ODD ------------------------------------------------------------------

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

func (f *builtinODD) ConstValue(args []Expr) (*Value, error) {
	c := args[0].ConstValue()
	if c == nil {
		return nil, nil
	}

	return &Value{
		typ:     TypeBoolean,
		boolean: c.Int()%2 == 1,
	}, nil
}

func (f *builtinODD) Generate(g *Generator, args []Expr) error {
	return fmt.Errorf("builtinODD.Generate is not implemented")
}

// ----- LSL ------------------------------------------------------------------

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

func (f *builtinLSL) ConstValue(args []Expr) (*Value, error) {
	return nil, fmt.Errorf("builtinLSL.ConstValue is not implemented")
}

func (f *builtinLSL) Generate(g *Generator, args []Expr) error {
	return fmt.Errorf("builtinLSL.Generate is not implemented")
}

// ----- ASR ------------------------------------------------------------------

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

func (f *builtinASR) ConstValue(args []Expr) (*Value, error) {
	return nil, fmt.Errorf("builtinASR.ConstValue is not implemented")
}

func (f *builtinASR) Generate(g *Generator, args []Expr) error {
	return fmt.Errorf("builtinASR.Generate is not implemented")
}

// ----- ROR ------------------------------------------------------------------

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

func (f *builtinROR) ConstValue(args []Expr) (*Value, error) {
	return nil, fmt.Errorf("builtinROR.ConstValue is not implemented")
}

func (f *builtinROR) Generate(g *Generator, args []Expr) error {
	return fmt.Errorf("builtinROR.Generate is not implemented")
}

// ----- LEN ------------------------------------------------------------------

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

func (f *builtinLEN) ConstValue(args []Expr) (*Value, error) {
	return nil, fmt.Errorf("builtinLEN.ConstValue is not implemented")
}

func (f *builtinLEN) Generate(g *Generator, args []Expr) error {
	return fmt.Errorf("builtinLEN.Generate is not implemented")
}

// ----- ORD ------------------------------------------------------------------

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

func (f *builtinORD) ConstValue(args []Expr) (*Value, error) {
	c := args[0].ConstValue()
	if c == nil {
		return nil, nil
	}

	switch args[0].Type() {
	case TypeBoolean, TypeChar:
		return &Value{
			typ:     TypeInt64,
			integer: c.Int(),
		}, nil
	default:
		return nil, nil
	}
}

func (f *builtinORD) Generate(g *Generator, args []Expr) error {
	switch args[0].Type() {
	case TypeBoolean:
		v := g.visitAndReturnValue(args[0])
		g.currentValue = g.currentBlock.NewZExt(v, types.I64)
	case TypeChar:
		v := g.visitAndReturnValue(args[0])
		g.currentValue = g.currentBlock.NewZExt(v, types.I64)
	}

	return nil
}

// ----- CHR ------------------------------------------------------------------

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

func (f *builtinCHR) ConstValue(args []Expr) (*Value, error) {
	if c := args[0].ConstValue(); c != nil {
		return &Value{
			typ:  TypeChar,
			char: c.Char(),
		}, nil
	}

	return nil, nil
}

func (f *builtinCHR) Generate(g *Generator, args []Expr) error {
	v := g.visitAndReturnValue(args[0])
	g.currentValue = g.currentBlock.NewTrunc(v, types.I8)

	return nil
}

// ----- FLOOR ----------------------------------------------------------------

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

func (f *builtinFLOOR) ConstValue(args []Expr) (*Value, error) {
	if c := args[0].ConstValue(); c != nil {
		return &Value{
			typ:     TypeInt64,
			integer: c.Int(),
		}, nil
	}

	return nil, nil
}

func (f *builtinFLOOR) Generate(g *Generator, args []Expr) error {
	v := g.visitAndReturnValue(args[0])
	g.currentValue = g.currentBlock.NewFPToSI(v, types.I64)

	return nil
}

// ----- FLT ------------------------------------------------------------------

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

func (f *builtinFLT) ConstValue(args []Expr) (*Value, error) {
	if c := args[0].ConstValue(); c != nil {
		return &Value{
			typ:  TypeFloat64,
			real: c.Real(),
		}, nil
	}

	return nil, nil
}

func (f *builtinFLT) Generate(g *Generator, args []Expr) error {
	v := g.visitAndReturnValue(args[0])
	g.currentValue = g.currentBlock.NewSIToFP(v, types.Double)

	return nil
}

// ----- INC ------------------------------------------------------------------

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

func (f *builtinINC) Generate(g *Generator, args []Expr) error {
	offset := 1

	if len(args) == 2 {
		offset = args[1].ConstValue().Int()
	}

	if offset > 0 {
		v := g.visitAndReturnValue(args[0])

		g.currentValue = g.currentBlock.NewAdd(v, constant.NewInt(types.I64, int64(offset)))
		g.currentBlock.NewStore(g.currentValue, g.vars[args[0].Token().Text])
	}

	return nil
}

// ----- DEC ------------------------------------------------------------------

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

func (f *builtinDEC) Generate(g *Generator, args []Expr) error {
	offset := 1

	if len(args) == 2 {
		offset = args[1].ConstValue().Int()
	}

	if offset > 0 {
		v := g.visitAndReturnValue(args[0])

		g.currentValue = g.currentBlock.NewSub(v, constant.NewInt(types.I64, int64(offset)))
		g.currentBlock.NewStore(g.currentValue, g.vars[args[0].Token().Text])
	}

	return nil
}

// ----- INCL -----------------------------------------------------------------

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

func (f *builtinINCL) Generate(g *Generator, args []Expr) error {
	v := g.visitAndReturnValue(args[0])
	b := g.visitAndReturnValue(args[1])
	g.currentValue = g.currentBlock.NewShl(constant.NewInt(types.I64, 1), b)
	g.currentValue = g.currentBlock.NewOr(v, g.currentValue)
	g.currentBlock.NewStore(g.currentValue, g.vars[args[0].Token().Text])

	return nil
}

// ----- EXCL -----------------------------------------------------------------

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

func (f *builtinEXCL) Generate(g *Generator, args []Expr) error {
	v := g.visitAndReturnValue(args[0])
	b := g.visitAndReturnValue(args[1])
	g.currentValue = g.currentBlock.NewShl(constant.NewInt(types.I64, 1), b)
	g.currentValue = g.currentBlock.NewXor(g.currentValue, constant.NewInt(types.I64, -1))
	g.currentValue = g.currentBlock.NewAnd(v, g.currentValue)
	g.currentBlock.NewStore(g.currentValue, g.vars[args[0].Token().Text])

	return nil
}

// ----- PACK -----------------------------------------------------------------

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

func (f *builtinPACK) Generate(g *Generator, args []Expr) error {
	return fmt.Errorf("builtinPACK.Generate is not implemented")
}

// ----- UNPK -----------------------------------------------------------------

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

func (f *builtinUNPK) Generate(g *Generator, args []Expr) error {
	return fmt.Errorf("builtinUNPK.Generate is not implemented")
}

// ----- NEW ------------------------------------------------------------------

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

func (f *builtinNEW) Generate(g *Generator, args []Expr) error {
	return fmt.Errorf("builtinNEW.Generate is not implemented")
}

// ----- ASSERT ---------------------------------------------------------------

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

func (f *builtinASSERT) Generate(g *Generator, args []Expr) error {
	return fmt.Errorf("builtinASSERT.Generate is not implemented")
}
