# test/test_010.md
## Source
```pascal
MODULE IncDec;

VAR x: INTEGER;

BEGIN
  x := 1; Texts.WriteInt(x);    Texts.WriteLn;
  INC(x); Texts.WriteInt(x);    Texts.WriteLn;
  DEC(x); Texts.WriteInt(x);    Texts.WriteLn;
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
test/test_010.md:6:10:	ident	"Texts"	false	0	0.000000	(6, 10) -> (6, 15)
test/test_010.md:6:15:	dot	"."	false	0	0.000000	(6, 15) -> (6, 16)
test/test_010.md:6:16:	ident	"WriteInt"	false	0	0.000000	(6, 16) -> (6, 24)
test/test_010.md:6:24:	lparen	"("	false	0	0.000000	(6, 24) -> (6, 25)
test/test_010.md:6:25:	ident	"x"	false	0	0.000000	(6, 25) -> (6, 26)
test/test_010.md:6:26:	rparen	")"	false	0	0.000000	(6, 26) -> (6, 27)
test/test_010.md:6:27:	semicolon	";"	false	0	0.000000	(6, 27) -> (6, 28)
test/test_010.md:6:32:	ident	"Texts"	false	0	0.000000	(6, 32) -> (6, 37)
test/test_010.md:6:37:	dot	"."	false	0	0.000000	(6, 37) -> (6, 38)
test/test_010.md:6:38:	ident	"WriteLn"	false	0	0.000000	(6, 38) -> (6, 45)
test/test_010.md:6:45:	semicolon	";"	false	0	0.000000	(6, 45) -> (6, 46)
test/test_010.md:7:2:	ident	"INC"	false	0	0.000000	(7, 2) -> (7, 5)
test/test_010.md:7:5:	lparen	"("	false	0	0.000000	(7, 5) -> (7, 6)
test/test_010.md:7:6:	ident	"x"	false	0	0.000000	(7, 6) -> (7, 7)
test/test_010.md:7:7:	rparen	")"	false	0	0.000000	(7, 7) -> (7, 8)
test/test_010.md:7:8:	semicolon	";"	false	0	0.000000	(7, 8) -> (7, 9)
test/test_010.md:7:10:	ident	"Texts"	false	0	0.000000	(7, 10) -> (7, 15)
test/test_010.md:7:15:	dot	"."	false	0	0.000000	(7, 15) -> (7, 16)
test/test_010.md:7:16:	ident	"WriteInt"	false	0	0.000000	(7, 16) -> (7, 24)
test/test_010.md:7:24:	lparen	"("	false	0	0.000000	(7, 24) -> (7, 25)
test/test_010.md:7:25:	ident	"x"	false	0	0.000000	(7, 25) -> (7, 26)
test/test_010.md:7:26:	rparen	")"	false	0	0.000000	(7, 26) -> (7, 27)
test/test_010.md:7:27:	semicolon	";"	false	0	0.000000	(7, 27) -> (7, 28)
test/test_010.md:7:32:	ident	"Texts"	false	0	0.000000	(7, 32) -> (7, 37)
test/test_010.md:7:37:	dot	"."	false	0	0.000000	(7, 37) -> (7, 38)
test/test_010.md:7:38:	ident	"WriteLn"	false	0	0.000000	(7, 38) -> (7, 45)
test/test_010.md:7:45:	semicolon	";"	false	0	0.000000	(7, 45) -> (7, 46)
test/test_010.md:8:2:	ident	"DEC"	false	0	0.000000	(8, 2) -> (8, 5)
test/test_010.md:8:5:	lparen	"("	false	0	0.000000	(8, 5) -> (8, 6)
test/test_010.md:8:6:	ident	"x"	false	0	0.000000	(8, 6) -> (8, 7)
test/test_010.md:8:7:	rparen	")"	false	0	0.000000	(8, 7) -> (8, 8)
test/test_010.md:8:8:	semicolon	";"	false	0	0.000000	(8, 8) -> (8, 9)
test/test_010.md:8:10:	ident	"Texts"	false	0	0.000000	(8, 10) -> (8, 15)
test/test_010.md:8:15:	dot	"."	false	0	0.000000	(8, 15) -> (8, 16)
test/test_010.md:8:16:	ident	"WriteInt"	false	0	0.000000	(8, 16) -> (8, 24)
test/test_010.md:8:24:	lparen	"("	false	0	0.000000	(8, 24) -> (8, 25)
test/test_010.md:8:25:	ident	"x"	false	0	0.000000	(8, 25) -> (8, 26)
test/test_010.md:8:26:	rparen	")"	false	0	0.000000	(8, 26) -> (8, 27)
test/test_010.md:8:27:	semicolon	";"	false	0	0.000000	(8, 27) -> (8, 28)
test/test_010.md:8:32:	ident	"Texts"	false	0	0.000000	(8, 32) -> (8, 37)
test/test_010.md:8:37:	dot	"."	false	0	0.000000	(8, 37) -> (8, 38)
test/test_010.md:8:38:	ident	"WriteLn"	false	0	0.000000	(8, 38) -> (8, 45)
test/test_010.md:8:45:	semicolon	";"	false	0	0.000000	(8, 45) -> (8, 46)
test/test_010.md:9:0:	end	"END"	false	0	0.000000	(9, 0) -> (9, 3)
test/test_010.md:9:4:	ident	"IncDec"	false	0	0.000000	(9, 4) -> (9, 10)
test/test_010.md:9:10:	dot	"."	false	0	0.000000	(9, 10) -> (9, 11)
test/test_010.md:10:0:	eof	""	false	0	0.000000	(10, 0) -> (10, 0)
```
## AST
```scheme
(module "IncDec"
  (vars
    (x 
      (INTEGER [i64])
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
    (expr2stmt
      (call
        (procedure [void] "INC")
        (variable [i64] "x")
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
    (expr2stmt
      (call
        (procedure [void] "DEC")
        (variable [i64] "x")
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
@1 = global [3 x i8] c"%d\00"
@2 = global [1 x i8] c"\00"
@__argc = global i64 0
@__argv = global i8** inttoptr (i8 0 to i8**)
@__envp = global i8** inttoptr (i8 0 to i8**)

declare i64 @puts(i8* %str)

declare i64 @rand()

declare i64 @sprintf(i8* %buf, i8* %format, ...)

declare i64 @printf(i8* %format, ...)

declare i8* @malloc(i64 %size)

declare i8* @free(i8* %ptr)

define void @oberonMain() {
entry:
	store i64 1, i64* @0
	%0 = load i64, i64* @0
	%1 = getelementptr [3 x i8], [3 x i8]* @1, i64 0, i64 0
	%2 = call i64 (i8*, ...) @printf(i8* %1, i64 %0)
	%3 = getelementptr [1 x i8], [1 x i8]* @2, i64 0, i64 0
	%4 = call i64 @puts(i8* %3)
	%5 = load i64, i64* @0
	%6 = add i64 %5, 1
	store i64 %6, i64* @0
	%7 = load i64, i64* @0
	%8 = getelementptr [3 x i8], [3 x i8]* @1, i64 0, i64 0
	%9 = call i64 (i8*, ...) @printf(i8* %8, i64 %7)
	%10 = getelementptr [1 x i8], [1 x i8]* @2, i64 0, i64 0
	%11 = call i64 @puts(i8* %10)
	%12 = load i64, i64* @0
	%13 = sub i64 %12, 1
	store i64 %13, i64* @0
	%14 = load i64, i64* @0
	%15 = getelementptr [3 x i8], [3 x i8]* @1, i64 0, i64 0
	%16 = call i64 (i8*, ...) @printf(i8* %15, i64 %14)
	%17 = getelementptr [1 x i8], [1 x i8]* @2, i64 0, i64 0
	%18 = call i64 @puts(i8* %17)
	ret void
}

define i64 @main(i64 %argc, i8** %argv, i8** %argp) {
entry:
	store i64 %argc, i64* @__argc
	store i8** %argv, i8*** @__argv
	store i8** %argp, i8*** @__envp
	call void @oberonMain()
	ret i64 0
}

```
## Run
```bash
1
2
1
```
