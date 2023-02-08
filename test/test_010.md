# test/test_010.md
## Source
```pascal
MODULE IncDec;

VAR x: INTEGER;

BEGIN
  x := 1; print(x);
  INC(x); print(x);
  DEC(x); print(x);
END IncDec.
```
## Tokens
```tsv
test/test_010.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_010.md:1:8:	ident	"IncDec"	false	0	0.000000	(1, 8) -> (1, 14)
test/test_010.md:1:14:	semicolon	";"	false	0	0.000000	(1, 14) -> (1, 15)
test/test_010.md:3:0:	var	"VAR"	false	0	0.000000	(3, 0) -> (3, 3)
test/test_010.md:3:4:	ident	"x"	false	0	0.000000	(3, 4) -> (3, 5)
test/test_010.md:3:5:	colon	":"	false	0	0.000000	(3, 5) -> (3, 6)
test/test_010.md:3:7:	ident	"INTEGER"	false	0	0.000000	(3, 7) -> (3, 14)
test/test_010.md:3:14:	semicolon	";"	false	0	0.000000	(3, 14) -> (3, 15)
test/test_010.md:5:0:	begin	"BEGIN"	false	0	0.000000	(5, 0) -> (5, 5)
test/test_010.md:6:2:	ident	"x"	false	0	0.000000	(6, 2) -> (6, 3)
test/test_010.md:6:4:	assign	":="	false	0	0.000000	(6, 4) -> (6, 6)
test/test_010.md:6:7:	integer	"1"	false	1	0.000000	(6, 7) -> (6, 8)
test/test_010.md:6:8:	semicolon	";"	false	0	0.000000	(6, 8) -> (6, 9)
test/test_010.md:6:10:	ident	"print"	false	0	0.000000	(6, 10) -> (6, 15)
test/test_010.md:6:15:	lparen	"("	false	0	0.000000	(6, 15) -> (6, 16)
test/test_010.md:6:16:	ident	"x"	false	0	0.000000	(6, 16) -> (6, 17)
test/test_010.md:6:17:	rparen	")"	false	0	0.000000	(6, 17) -> (6, 18)
test/test_010.md:6:18:	semicolon	";"	false	0	0.000000	(6, 18) -> (6, 19)
test/test_010.md:7:2:	ident	"INC"	false	0	0.000000	(7, 2) -> (7, 5)
test/test_010.md:7:5:	lparen	"("	false	0	0.000000	(7, 5) -> (7, 6)
test/test_010.md:7:6:	ident	"x"	false	0	0.000000	(7, 6) -> (7, 7)
test/test_010.md:7:7:	rparen	")"	false	0	0.000000	(7, 7) -> (7, 8)
test/test_010.md:7:8:	semicolon	";"	false	0	0.000000	(7, 8) -> (7, 9)
test/test_010.md:7:10:	ident	"print"	false	0	0.000000	(7, 10) -> (7, 15)
test/test_010.md:7:15:	lparen	"("	false	0	0.000000	(7, 15) -> (7, 16)
test/test_010.md:7:16:	ident	"x"	false	0	0.000000	(7, 16) -> (7, 17)
test/test_010.md:7:17:	rparen	")"	false	0	0.000000	(7, 17) -> (7, 18)
test/test_010.md:7:18:	semicolon	";"	false	0	0.000000	(7, 18) -> (7, 19)
test/test_010.md:8:2:	ident	"DEC"	false	0	0.000000	(8, 2) -> (8, 5)
test/test_010.md:8:5:	lparen	"("	false	0	0.000000	(8, 5) -> (8, 6)
test/test_010.md:8:6:	ident	"x"	false	0	0.000000	(8, 6) -> (8, 7)
test/test_010.md:8:7:	rparen	")"	false	0	0.000000	(8, 7) -> (8, 8)
test/test_010.md:8:8:	semicolon	";"	false	0	0.000000	(8, 8) -> (8, 9)
test/test_010.md:8:10:	ident	"print"	false	0	0.000000	(8, 10) -> (8, 15)
test/test_010.md:8:15:	lparen	"("	false	0	0.000000	(8, 15) -> (8, 16)
test/test_010.md:8:16:	ident	"x"	false	0	0.000000	(8, 16) -> (8, 17)
test/test_010.md:8:17:	rparen	")"	false	0	0.000000	(8, 17) -> (8, 18)
test/test_010.md:8:18:	semicolon	";"	false	0	0.000000	(8, 18) -> (8, 19)
test/test_010.md:9:0:	end	"END"	false	0	0.000000	(9, 0) -> (9, 3)
test/test_010.md:9:4:	ident	"IncDec"	false	0	0.000000	(9, 4) -> (9, 10)
test/test_010.md:9:10:	dot	"."	false	0	0.000000	(9, 10) -> (9, 11)
test/test_010.md:10:0:	eof	""	false	0	0.000000	(10, 0) -> (10, 0)
```
## AST
```scheme
(module "IncDec"
  (vars
    (x [i64])
  )
  (stmts
    (assign x
      (number [i64] 1)
    )
    (expr2stmt
      (call "print" [void]
        (variable [i64] "x")
      )
    )
    (expr2stmt
      (call "INC" [void]
        (variable [i64] "x")
      )
    )
    (expr2stmt
      (call "print" [void]
        (variable [i64] "x")
      )
    )
    (expr2stmt
      (call "DEC" [void]
        (variable [i64] "x")
      )
    )
    (expr2stmt
      (call "print" [void]
        (variable [i64] "x")
      )
    )
  )
)
```
## IR
```llvm
@0 = global i64 0
@1 = global [4 x i8] c"%d\0A\00"

declare i64 @puts(i8* %str)

declare i64 @rand()

declare i64 @sprintf(i8* %buf, i8* %format, ...)

declare i64 @printf(i8* %format, ...)

define i64 @main() {
entry:
	store i64 1, i64* @0
	%0 = load i64, i64* @0
	%1 = getelementptr [4 x i8], [4 x i8]* @1, i64 0, i64 0
	%2 = call i64 (i8*, ...) @printf(i8* %1, i64 %0)
	%3 = load i64, i64* @0
	%4 = add i64 %3, 1
	store i64 %4, i64* @0
	%5 = load i64, i64* @0
	%6 = getelementptr [4 x i8], [4 x i8]* @1, i64 0, i64 0
	%7 = call i64 (i8*, ...) @printf(i8* %6, i64 %5)
	%8 = load i64, i64* @0
	%9 = sub i64 %8, 1
	store i64 %9, i64* @0
	%10 = load i64, i64* @0
	%11 = getelementptr [4 x i8], [4 x i8]* @1, i64 0, i64 0
	%12 = call i64 (i8*, ...) @printf(i8* %11, i64 %10)
	ret i64 0
}

```
## Run
```bash
1
2
1
```
