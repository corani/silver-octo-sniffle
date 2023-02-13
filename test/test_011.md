# test/test_011.md
## Source
```pascal
MODULE FltFloor;

VAR x: INTEGER;
    y: REAL;

BEGIN
  x := 1;        print(x);
  y := FLT(x);   print(y);
  y := y+1.1;    print(y);
  x := FLOOR(y); print(x);
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
test/test_011.md:7:17:	ident	"print"	false	0	0.000000	(7, 17) -> (7, 22)
test/test_011.md:7:22:	lparen	"("	false	0	0.000000	(7, 22) -> (7, 23)
test/test_011.md:7:23:	ident	"x"	false	0	0.000000	(7, 23) -> (7, 24)
test/test_011.md:7:24:	rparen	")"	false	0	0.000000	(7, 24) -> (7, 25)
test/test_011.md:7:25:	semicolon	";"	false	0	0.000000	(7, 25) -> (7, 26)
test/test_011.md:8:2:	ident	"y"	false	0	0.000000	(8, 2) -> (8, 3)
test/test_011.md:8:4:	assign	":="	false	0	0.000000	(8, 4) -> (8, 6)
test/test_011.md:8:7:	ident	"FLT"	false	0	0.000000	(8, 7) -> (8, 10)
test/test_011.md:8:10:	lparen	"("	false	0	0.000000	(8, 10) -> (8, 11)
test/test_011.md:8:11:	ident	"x"	false	0	0.000000	(8, 11) -> (8, 12)
test/test_011.md:8:12:	rparen	")"	false	0	0.000000	(8, 12) -> (8, 13)
test/test_011.md:8:13:	semicolon	";"	false	0	0.000000	(8, 13) -> (8, 14)
test/test_011.md:8:17:	ident	"print"	false	0	0.000000	(8, 17) -> (8, 22)
test/test_011.md:8:22:	lparen	"("	false	0	0.000000	(8, 22) -> (8, 23)
test/test_011.md:8:23:	ident	"y"	false	0	0.000000	(8, 23) -> (8, 24)
test/test_011.md:8:24:	rparen	")"	false	0	0.000000	(8, 24) -> (8, 25)
test/test_011.md:8:25:	semicolon	";"	false	0	0.000000	(8, 25) -> (8, 26)
test/test_011.md:9:2:	ident	"y"	false	0	0.000000	(9, 2) -> (9, 3)
test/test_011.md:9:4:	assign	":="	false	0	0.000000	(9, 4) -> (9, 6)
test/test_011.md:9:7:	ident	"y"	false	0	0.000000	(9, 7) -> (9, 8)
test/test_011.md:9:8:	plus	"+"	false	0	0.000000	(9, 8) -> (9, 9)
test/test_011.md:9:9:	real	"1.1"	false	0	1.100000	(9, 9) -> (9, 12)
test/test_011.md:9:12:	semicolon	";"	false	0	0.000000	(9, 12) -> (9, 13)
test/test_011.md:9:17:	ident	"print"	false	0	0.000000	(9, 17) -> (9, 22)
test/test_011.md:9:22:	lparen	"("	false	0	0.000000	(9, 22) -> (9, 23)
test/test_011.md:9:23:	ident	"y"	false	0	0.000000	(9, 23) -> (9, 24)
test/test_011.md:9:24:	rparen	")"	false	0	0.000000	(9, 24) -> (9, 25)
test/test_011.md:9:25:	semicolon	";"	false	0	0.000000	(9, 25) -> (9, 26)
test/test_011.md:10:2:	ident	"x"	false	0	0.000000	(10, 2) -> (10, 3)
test/test_011.md:10:4:	assign	":="	false	0	0.000000	(10, 4) -> (10, 6)
test/test_011.md:10:7:	ident	"FLOOR"	false	0	0.000000	(10, 7) -> (10, 12)
test/test_011.md:10:12:	lparen	"("	false	0	0.000000	(10, 12) -> (10, 13)
test/test_011.md:10:13:	ident	"y"	false	0	0.000000	(10, 13) -> (10, 14)
test/test_011.md:10:14:	rparen	")"	false	0	0.000000	(10, 14) -> (10, 15)
test/test_011.md:10:15:	semicolon	";"	false	0	0.000000	(10, 15) -> (10, 16)
test/test_011.md:10:17:	ident	"print"	false	0	0.000000	(10, 17) -> (10, 22)
test/test_011.md:10:22:	lparen	"("	false	0	0.000000	(10, 22) -> (10, 23)
test/test_011.md:10:23:	ident	"x"	false	0	0.000000	(10, 23) -> (10, 24)
test/test_011.md:10:24:	rparen	")"	false	0	0.000000	(10, 24) -> (10, 25)
test/test_011.md:10:25:	semicolon	";"	false	0	0.000000	(10, 25) -> (10, 26)
test/test_011.md:11:0:	end	"END"	false	0	0.000000	(11, 0) -> (11, 3)
test/test_011.md:11:4:	ident	"FltFloor"	false	0	0.000000	(11, 4) -> (11, 12)
test/test_011.md:11:12:	dot	"."	false	0	0.000000	(11, 12) -> (11, 13)
test/test_011.md:12:0:	eof	""	false	0	0.000000	(12, 0) -> (12, 0)
```
## AST
```scheme
(module "FltFloor"
  (vars
    (x [i64])
    (y [f64])
  )
  (stmts
    (assign
      (variable [i64] "x")
      (number [i64] 1)
    )
    (expr2stmt
      (call
        (procedure [void] "print")
        (variable [i64] "x")
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
        (procedure [void] "print")
        (variable [f64] "y")
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
        (procedure [void] "print")
        (variable [f64] "y")
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
        (procedure [void] "print")
        (variable [i64] "x")
      )
    )
  )
)
```
## IR
```llvm
@0 = global i64 0
@1 = global double 0.0
@2 = global [4 x i8] c"%d\0A\00"
@3 = global [4 x i8] c"%f\0A\00"

declare i64 @puts(i8* %str)

declare i64 @rand()

declare i64 @sprintf(i8* %buf, i8* %format, ...)

declare i64 @printf(i8* %format, ...)

define i64 @main() {
entry:
	store i64 1, i64* @0
	%0 = load i64, i64* @0
	%1 = getelementptr [4 x i8], [4 x i8]* @2, i64 0, i64 0
	%2 = call i64 (i8*, ...) @printf(i8* %1, i64 %0)
	%3 = load i64, i64* @0
	%4 = sitofp i64 %3 to double
	store double %4, double* @1
	%5 = load double, double* @1
	%6 = getelementptr [4 x i8], [4 x i8]* @3, i64 0, i64 0
	%7 = call i64 (i8*, ...) @printf(i8* %6, double %5)
	%8 = load double, double* @1
	%9 = fadd double %8, 0x3FF199999999999A
	store double %9, double* @1
	%10 = load double, double* @1
	%11 = getelementptr [4 x i8], [4 x i8]* @3, i64 0, i64 0
	%12 = call i64 (i8*, ...) @printf(i8* %11, double %10)
	%13 = load double, double* @1
	%14 = fptosi double %13 to i64
	store i64 %14, i64* @0
	%15 = load i64, i64* @0
	%16 = getelementptr [4 x i8], [4 x i8]* @2, i64 0, i64 0
	%17 = call i64 (i8*, ...) @printf(i8* %16, i64 %15)
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
