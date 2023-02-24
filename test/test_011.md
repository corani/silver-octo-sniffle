# test/test_011.md
## Source
```pascal
MODULE FltFloor;

VAR x: INTEGER;
    y: REAL;

BEGIN
  x := 1;        Texts.WriteInt(x);     Texts.WriteLn;
  y := FLT(x);   Texts.WriteReal(y);    Texts.WriteLn;
  y := y+1.1;    Texts.WriteReal(y);    Texts.WriteLn;
  x := FLOOR(y); Texts.WriteInt(x);     Texts.WriteLn;
END FltFloor.
```
## Tokens
```tsv
test/test_011.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_011.md:1:8:	ident	"FltFloor"	false	0	0.000000	(1, 8) -> (1, 16)
test/test_011.md:1:16:	semicolon	";"	false	0	0.000000	(1, 16) -> (1, 17)
test/test_011.md:3:0:	var	"VAR"	false	0	0.000000	(3, 0) -> (3, 3)
test/test_011.md:3:4:	ident	"x"	false	0	0.000000	(3, 4) -> (3, 5)
test/test_011.md:3:5:	colon	":"	false	0	0.000000	(3, 5) -> (3, 6)
test/test_011.md:3:7:	ident	"INTEGER"	false	0	0.000000	(3, 7) -> (3, 14)
test/test_011.md:3:14:	semicolon	";"	false	0	0.000000	(3, 14) -> (3, 15)
test/test_011.md:4:4:	ident	"y"	false	0	0.000000	(4, 4) -> (4, 5)
test/test_011.md:4:5:	colon	":"	false	0	0.000000	(4, 5) -> (4, 6)
test/test_011.md:4:7:	ident	"REAL"	false	0	0.000000	(4, 7) -> (4, 11)
test/test_011.md:4:11:	semicolon	";"	false	0	0.000000	(4, 11) -> (4, 12)
test/test_011.md:6:0:	begin	"BEGIN"	false	0	0.000000	(6, 0) -> (6, 5)
test/test_011.md:7:2:	ident	"x"	false	0	0.000000	(7, 2) -> (7, 3)
test/test_011.md:7:4:	assign	":="	false	0	0.000000	(7, 4) -> (7, 6)
test/test_011.md:7:7:	integer	"1"	false	1	0.000000	(7, 7) -> (7, 8)
test/test_011.md:7:8:	semicolon	";"	false	0	0.000000	(7, 8) -> (7, 9)
test/test_011.md:7:17:	ident	"Texts"	false	0	0.000000	(7, 17) -> (7, 22)
test/test_011.md:7:22:	dot	"."	false	0	0.000000	(7, 22) -> (7, 23)
test/test_011.md:7:23:	ident	"WriteInt"	false	0	0.000000	(7, 23) -> (7, 31)
test/test_011.md:7:31:	lparen	"("	false	0	0.000000	(7, 31) -> (7, 32)
test/test_011.md:7:32:	ident	"x"	false	0	0.000000	(7, 32) -> (7, 33)
test/test_011.md:7:33:	rparen	")"	false	0	0.000000	(7, 33) -> (7, 34)
test/test_011.md:7:34:	semicolon	";"	false	0	0.000000	(7, 34) -> (7, 35)
test/test_011.md:7:40:	ident	"Texts"	false	0	0.000000	(7, 40) -> (7, 45)
test/test_011.md:7:45:	dot	"."	false	0	0.000000	(7, 45) -> (7, 46)
test/test_011.md:7:46:	ident	"WriteLn"	false	0	0.000000	(7, 46) -> (7, 53)
test/test_011.md:7:53:	semicolon	";"	false	0	0.000000	(7, 53) -> (7, 54)
test/test_011.md:8:2:	ident	"y"	false	0	0.000000	(8, 2) -> (8, 3)
test/test_011.md:8:4:	assign	":="	false	0	0.000000	(8, 4) -> (8, 6)
test/test_011.md:8:7:	ident	"FLT"	false	0	0.000000	(8, 7) -> (8, 10)
test/test_011.md:8:10:	lparen	"("	false	0	0.000000	(8, 10) -> (8, 11)
test/test_011.md:8:11:	ident	"x"	false	0	0.000000	(8, 11) -> (8, 12)
test/test_011.md:8:12:	rparen	")"	false	0	0.000000	(8, 12) -> (8, 13)
test/test_011.md:8:13:	semicolon	";"	false	0	0.000000	(8, 13) -> (8, 14)
test/test_011.md:8:17:	ident	"Texts"	false	0	0.000000	(8, 17) -> (8, 22)
test/test_011.md:8:22:	dot	"."	false	0	0.000000	(8, 22) -> (8, 23)
test/test_011.md:8:23:	ident	"WriteReal"	false	0	0.000000	(8, 23) -> (8, 32)
test/test_011.md:8:32:	lparen	"("	false	0	0.000000	(8, 32) -> (8, 33)
test/test_011.md:8:33:	ident	"y"	false	0	0.000000	(8, 33) -> (8, 34)
test/test_011.md:8:34:	rparen	")"	false	0	0.000000	(8, 34) -> (8, 35)
test/test_011.md:8:35:	semicolon	";"	false	0	0.000000	(8, 35) -> (8, 36)
test/test_011.md:8:40:	ident	"Texts"	false	0	0.000000	(8, 40) -> (8, 45)
test/test_011.md:8:45:	dot	"."	false	0	0.000000	(8, 45) -> (8, 46)
test/test_011.md:8:46:	ident	"WriteLn"	false	0	0.000000	(8, 46) -> (8, 53)
test/test_011.md:8:53:	semicolon	";"	false	0	0.000000	(8, 53) -> (8, 54)
test/test_011.md:9:2:	ident	"y"	false	0	0.000000	(9, 2) -> (9, 3)
test/test_011.md:9:4:	assign	":="	false	0	0.000000	(9, 4) -> (9, 6)
test/test_011.md:9:7:	ident	"y"	false	0	0.000000	(9, 7) -> (9, 8)
test/test_011.md:9:8:	plus	"+"	false	0	0.000000	(9, 8) -> (9, 9)
test/test_011.md:9:9:	real	"1.1"	false	0	1.100000	(9, 9) -> (9, 12)
test/test_011.md:9:12:	semicolon	";"	false	0	0.000000	(9, 12) -> (9, 13)
test/test_011.md:9:17:	ident	"Texts"	false	0	0.000000	(9, 17) -> (9, 22)
test/test_011.md:9:22:	dot	"."	false	0	0.000000	(9, 22) -> (9, 23)
test/test_011.md:9:23:	ident	"WriteReal"	false	0	0.000000	(9, 23) -> (9, 32)
test/test_011.md:9:32:	lparen	"("	false	0	0.000000	(9, 32) -> (9, 33)
test/test_011.md:9:33:	ident	"y"	false	0	0.000000	(9, 33) -> (9, 34)
test/test_011.md:9:34:	rparen	")"	false	0	0.000000	(9, 34) -> (9, 35)
test/test_011.md:9:35:	semicolon	";"	false	0	0.000000	(9, 35) -> (9, 36)
test/test_011.md:9:40:	ident	"Texts"	false	0	0.000000	(9, 40) -> (9, 45)
test/test_011.md:9:45:	dot	"."	false	0	0.000000	(9, 45) -> (9, 46)
test/test_011.md:9:46:	ident	"WriteLn"	false	0	0.000000	(9, 46) -> (9, 53)
test/test_011.md:9:53:	semicolon	";"	false	0	0.000000	(9, 53) -> (9, 54)
test/test_011.md:10:2:	ident	"x"	false	0	0.000000	(10, 2) -> (10, 3)
test/test_011.md:10:4:	assign	":="	false	0	0.000000	(10, 4) -> (10, 6)
test/test_011.md:10:7:	ident	"FLOOR"	false	0	0.000000	(10, 7) -> (10, 12)
test/test_011.md:10:12:	lparen	"("	false	0	0.000000	(10, 12) -> (10, 13)
test/test_011.md:10:13:	ident	"y"	false	0	0.000000	(10, 13) -> (10, 14)
test/test_011.md:10:14:	rparen	")"	false	0	0.000000	(10, 14) -> (10, 15)
test/test_011.md:10:15:	semicolon	";"	false	0	0.000000	(10, 15) -> (10, 16)
test/test_011.md:10:17:	ident	"Texts"	false	0	0.000000	(10, 17) -> (10, 22)
test/test_011.md:10:22:	dot	"."	false	0	0.000000	(10, 22) -> (10, 23)
test/test_011.md:10:23:	ident	"WriteInt"	false	0	0.000000	(10, 23) -> (10, 31)
test/test_011.md:10:31:	lparen	"("	false	0	0.000000	(10, 31) -> (10, 32)
test/test_011.md:10:32:	ident	"x"	false	0	0.000000	(10, 32) -> (10, 33)
test/test_011.md:10:33:	rparen	")"	false	0	0.000000	(10, 33) -> (10, 34)
test/test_011.md:10:34:	semicolon	";"	false	0	0.000000	(10, 34) -> (10, 35)
test/test_011.md:10:40:	ident	"Texts"	false	0	0.000000	(10, 40) -> (10, 45)
test/test_011.md:10:45:	dot	"."	false	0	0.000000	(10, 45) -> (10, 46)
test/test_011.md:10:46:	ident	"WriteLn"	false	0	0.000000	(10, 46) -> (10, 53)
test/test_011.md:10:53:	semicolon	";"	false	0	0.000000	(10, 53) -> (10, 54)
test/test_011.md:11:0:	end	"END"	false	0	0.000000	(11, 0) -> (11, 3)
test/test_011.md:11:4:	ident	"FltFloor"	false	0	0.000000	(11, 4) -> (11, 12)
test/test_011.md:11:12:	dot	"."	false	0	0.000000	(11, 12) -> (11, 13)
test/test_011.md:12:0:	eof	""	false	0	0.000000	(12, 0) -> (12, 0)
```
## AST
```scheme
(module "FltFloor"
  (vars
    (x 
      (INTEGER [i64])
    )
    (y 
      (REAL [f64])
    )
  )
  (stmts
    (assign
      (variable [i64] "x")
      (number [i64] 1)
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteInt")
        (variable [i64] "x")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
    (assign
      (variable [f64] "y")
      (call
        (procedure [f64] "FLT")
        (variable [i64] "x")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteReal")
        (variable [f64] "y")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
    (assign
      (variable [f64] "y")
      (plus [f64]
        (variable [f64] "y")
        (number [f64] 1.100000)
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteReal")
        (variable [f64] "y")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
    (assign
      (variable [i64] "x")
      (call
        (procedure [i64] "FLOOR")
        (variable [f64] "y")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteInt")
        (variable [i64] "x")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
  )
)
```
## IR
```llvm
@0 = global i64 0
@1 = global double 0.0
@2 = global [3 x i8] c"%d\00"
@3 = global [1 x i8] c"\00"
@4 = global [3 x i8] c"%f\00"

declare i64 @puts(i8* %str)

declare i64 @rand()

declare i64 @sprintf(i8* %buf, i8* %format, ...)

declare i64 @printf(i8* %format, ...)

declare i8* @malloc(i64 %size)

declare i8* @free(i8* %ptr)

define i64 @main() {
entry:
	store i64 1, i64* @0
	%0 = load i64, i64* @0
	%1 = getelementptr [3 x i8], [3 x i8]* @2, i64 0, i64 0
	%2 = call i64 (i8*, ...) @printf(i8* %1, i64 %0)
	%3 = getelementptr [1 x i8], [1 x i8]* @3, i64 0, i64 0
	%4 = call i64 @puts(i8* %3)
	%5 = load i64, i64* @0
	%6 = sitofp i64 %5 to double
	store double %6, double* @1
	%7 = load double, double* @1
	%8 = getelementptr [3 x i8], [3 x i8]* @4, i64 0, i64 0
	%9 = call i64 (i8*, ...) @printf(i8* %8, double %7)
	%10 = getelementptr [1 x i8], [1 x i8]* @3, i64 0, i64 0
	%11 = call i64 @puts(i8* %10)
	%12 = load double, double* @1
	%13 = fadd double %12, 0x3FF199999999999A
	store double %13, double* @1
	%14 = load double, double* @1
	%15 = getelementptr [3 x i8], [3 x i8]* @4, i64 0, i64 0
	%16 = call i64 (i8*, ...) @printf(i8* %15, double %14)
	%17 = getelementptr [1 x i8], [1 x i8]* @3, i64 0, i64 0
	%18 = call i64 @puts(i8* %17)
	%19 = load double, double* @1
	%20 = fptosi double %19 to i64
	store i64 %20, i64* @0
	%21 = load i64, i64* @0
	%22 = getelementptr [3 x i8], [3 x i8]* @2, i64 0, i64 0
	%23 = call i64 (i8*, ...) @printf(i8* %22, i64 %21)
	%24 = getelementptr [1 x i8], [1 x i8]* @3, i64 0, i64 0
	%25 = call i64 @puts(i8* %24)
	ret i64 0
}

```
## Run
```bash
1
1.000000
2.100000
2
```
