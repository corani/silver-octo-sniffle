# test/test_022.md
## Source
```pascal
MODULE Pointers;

TYPE ip = POINTER TO INTEGER;

VAR x : POINTER TO INTEGER;
    y : ip;

BEGIN
  NEW(x);
  y := x;
  x^ := 33;
  Texts.WriteInt(y^);   Texts.WriteLn;
  DELETE(x);
END Pointers.
```
## Tokens
```tsv
test/test_022.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_022.md:1:8:	ident	"Pointers"	false	0	0.000000	(1, 8) -> (1, 16)
test/test_022.md:1:16:	semicolon	";"	false	0	0.000000	(1, 16) -> (1, 17)
test/test_022.md:3:0:	type	"TYPE"	false	0	0.000000	(3, 0) -> (3, 4)
test/test_022.md:3:5:	ident	"ip"	false	0	0.000000	(3, 5) -> (3, 7)
test/test_022.md:3:8:	eq	"="	false	0	0.000000	(3, 8) -> (3, 9)
test/test_022.md:3:10:	pointer	"POINTER"	false	0	0.000000	(3, 10) -> (3, 17)
test/test_022.md:3:18:	to	"TO"	false	0	0.000000	(3, 18) -> (3, 20)
test/test_022.md:3:21:	ident	"INTEGER"	false	0	0.000000	(3, 21) -> (3, 28)
test/test_022.md:3:28:	semicolon	";"	false	0	0.000000	(3, 28) -> (3, 29)
test/test_022.md:5:0:	var	"VAR"	false	0	0.000000	(5, 0) -> (5, 3)
test/test_022.md:5:4:	ident	"x"	false	0	0.000000	(5, 4) -> (5, 5)
test/test_022.md:5:6:	colon	":"	false	0	0.000000	(5, 6) -> (5, 7)
test/test_022.md:5:8:	pointer	"POINTER"	false	0	0.000000	(5, 8) -> (5, 15)
test/test_022.md:5:16:	to	"TO"	false	0	0.000000	(5, 16) -> (5, 18)
test/test_022.md:5:19:	ident	"INTEGER"	false	0	0.000000	(5, 19) -> (5, 26)
test/test_022.md:5:26:	semicolon	";"	false	0	0.000000	(5, 26) -> (5, 27)
test/test_022.md:6:4:	ident	"y"	false	0	0.000000	(6, 4) -> (6, 5)
test/test_022.md:6:6:	colon	":"	false	0	0.000000	(6, 6) -> (6, 7)
test/test_022.md:6:8:	ident	"ip"	false	0	0.000000	(6, 8) -> (6, 10)
test/test_022.md:6:10:	semicolon	";"	false	0	0.000000	(6, 10) -> (6, 11)
test/test_022.md:8:0:	begin	"BEGIN"	false	0	0.000000	(8, 0) -> (8, 5)
test/test_022.md:9:2:	ident	"NEW"	false	0	0.000000	(9, 2) -> (9, 5)
test/test_022.md:9:5:	lparen	"("	false	0	0.000000	(9, 5) -> (9, 6)
test/test_022.md:9:6:	ident	"x"	false	0	0.000000	(9, 6) -> (9, 7)
test/test_022.md:9:7:	rparen	")"	false	0	0.000000	(9, 7) -> (9, 8)
test/test_022.md:9:8:	semicolon	";"	false	0	0.000000	(9, 8) -> (9, 9)
test/test_022.md:10:2:	ident	"y"	false	0	0.000000	(10, 2) -> (10, 3)
test/test_022.md:10:4:	assign	":="	false	0	0.000000	(10, 4) -> (10, 6)
test/test_022.md:10:7:	ident	"x"	false	0	0.000000	(10, 7) -> (10, 8)
test/test_022.md:10:8:	semicolon	";"	false	0	0.000000	(10, 8) -> (10, 9)
test/test_022.md:11:2:	ident	"x"	false	0	0.000000	(11, 2) -> (11, 3)
test/test_022.md:11:3:	caret	"^"	false	0	0.000000	(11, 3) -> (11, 4)
test/test_022.md:11:5:	assign	":="	false	0	0.000000	(11, 5) -> (11, 7)
test/test_022.md:11:8:	integer	"33"	false	33	0.000000	(11, 8) -> (11, 10)
test/test_022.md:11:10:	semicolon	";"	false	0	0.000000	(11, 10) -> (11, 11)
test/test_022.md:12:2:	ident	"Texts"	false	0	0.000000	(12, 2) -> (12, 7)
test/test_022.md:12:7:	dot	"."	false	0	0.000000	(12, 7) -> (12, 8)
test/test_022.md:12:8:	ident	"WriteInt"	false	0	0.000000	(12, 8) -> (12, 16)
test/test_022.md:12:16:	lparen	"("	false	0	0.000000	(12, 16) -> (12, 17)
test/test_022.md:12:17:	ident	"y"	false	0	0.000000	(12, 17) -> (12, 18)
test/test_022.md:12:18:	caret	"^"	false	0	0.000000	(12, 18) -> (12, 19)
test/test_022.md:12:19:	rparen	")"	false	0	0.000000	(12, 19) -> (12, 20)
test/test_022.md:12:20:	semicolon	";"	false	0	0.000000	(12, 20) -> (12, 21)
test/test_022.md:12:24:	ident	"Texts"	false	0	0.000000	(12, 24) -> (12, 29)
test/test_022.md:12:29:	dot	"."	false	0	0.000000	(12, 29) -> (12, 30)
test/test_022.md:12:30:	ident	"WriteLn"	false	0	0.000000	(12, 30) -> (12, 37)
test/test_022.md:12:37:	semicolon	";"	false	0	0.000000	(12, 37) -> (12, 38)
test/test_022.md:13:2:	ident	"DELETE"	false	0	0.000000	(13, 2) -> (13, 8)
test/test_022.md:13:8:	lparen	"("	false	0	0.000000	(13, 8) -> (13, 9)
test/test_022.md:13:9:	ident	"x"	false	0	0.000000	(13, 9) -> (13, 10)
test/test_022.md:13:10:	rparen	")"	false	0	0.000000	(13, 10) -> (13, 11)
test/test_022.md:13:11:	semicolon	";"	false	0	0.000000	(13, 11) -> (13, 12)
test/test_022.md:14:0:	end	"END"	false	0	0.000000	(14, 0) -> (14, 3)
test/test_022.md:14:4:	ident	"Pointers"	false	0	0.000000	(14, 4) -> (14, 12)
test/test_022.md:14:12:	dot	"."	false	0	0.000000	(14, 12) -> (14, 13)
test/test_022.md:15:0:	eof	""	false	0	0.000000	(15, 0) -> (15, 0)
```
## AST
```scheme
(module "Pointers"
  (types
    (ip
      (pointer
        (INTEGER [i64])
      )
    )
  )
  (vars
    (x 
      (pointer
        (INTEGER [i64])
      )
    )
    (y 
      (pointer
        (INTEGER [i64])
      )
    )
  )
  (stmts
    (expr2stmt
      (call
        (procedure [void] "NEW")
        (variable [pointer] "x")
      )
    )
    (assign
      (variable [pointer] "y")
      (variable [pointer] "x")
    )
    (assign
      (variable [i64] "x^")
      (number [i64] 33)
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteInt")
        (variable [i64] "y^")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "DELETE")
        (variable [pointer] "x")
      )
    )
  )
)
```
## IR
```llvm
@0 = global i64* inttoptr (i64 0 to i64*)
@1 = global i64* inttoptr (i64 0 to i64*)
@2 = global [3 x i8] c"%d\00"
@3 = global [1 x i8] c"\00"
@__argc = global i64 0
@__argv = global i8** inttoptr (i8 0 to i8**)

declare i64 @puts(i8* %str)

declare i64 @rand()

declare i64 @sprintf(i8* %buf, i8* %format, ...)

declare i64 @printf(i8* %format, ...)

declare i8* @malloc(i64 %size)

declare i8* @free(i8* %ptr)

define void @oberonMain() {
entry:
	%0 = call i8* @malloc(i64 8)
	%1 = bitcast i8* %0 to i64*
	store i64* %1, i64** @0
	%2 = load i64*, i64** @0
	store i64* %2, i64** @1
	%3 = load i64*, i64** @0
	store i64 33, i64* %3
	%4 = load i64*, i64** @1
	%5 = load i64, i64* %4
	%6 = getelementptr [3 x i8], [3 x i8]* @2, i64 0, i64 0
	%7 = call i64 (i8*, ...) @printf(i8* %6, i64 %5)
	%8 = getelementptr [1 x i8], [1 x i8]* @3, i64 0, i64 0
	%9 = call i64 @puts(i8* %8)
	%10 = load i64*, i64** @0
	%11 = bitcast i64* %10 to i8*
	%12 = call i8* @free(i8* %11)
	ret void
}

define i64 @main(i64 %argc, i8** %argv) {
entry:
	store i64 %argc, i64* @__argc
	store i8** %argv, i8*** @__argv
	call void @oberonMain()
	ret i64 0
}

```
## Run
```bash
33
```
