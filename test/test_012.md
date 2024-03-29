# test/test_012.md
## Source
```pascal
MODULE Repeat;

VAR x: INTEGER;

BEGIN
  x := 10;
  REPEAT
    DEC(x);
    Texts.WriteInt(x); Texts.WriteLn;
  UNTIL x = 0;
END Repeat.
```
## Tokens
```tsv
test/test_012.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_012.md:1:8:	ident	"Repeat"	false	0	0.000000	(1, 8) -> (1, 14)
test/test_012.md:1:14:	semicolon	";"	false	0	0.000000	(1, 14) -> (1, 15)
test/test_012.md:3:0:	var	"VAR"	false	0	0.000000	(3, 0) -> (3, 3)
test/test_012.md:3:4:	ident	"x"	false	0	0.000000	(3, 4) -> (3, 5)
test/test_012.md:3:5:	colon	":"	false	0	0.000000	(3, 5) -> (3, 6)
test/test_012.md:3:7:	ident	"INTEGER"	false	0	0.000000	(3, 7) -> (3, 14)
test/test_012.md:3:14:	semicolon	";"	false	0	0.000000	(3, 14) -> (3, 15)
test/test_012.md:5:0:	begin	"BEGIN"	false	0	0.000000	(5, 0) -> (5, 5)
test/test_012.md:6:2:	ident	"x"	false	0	0.000000	(6, 2) -> (6, 3)
test/test_012.md:6:4:	assign	":="	false	0	0.000000	(6, 4) -> (6, 6)
test/test_012.md:6:7:	integer	"10"	false	10	0.000000	(6, 7) -> (6, 9)
test/test_012.md:6:9:	semicolon	";"	false	0	0.000000	(6, 9) -> (6, 10)
test/test_012.md:7:2:	repeat	"REPEAT"	false	0	0.000000	(7, 2) -> (7, 8)
test/test_012.md:8:4:	ident	"DEC"	false	0	0.000000	(8, 4) -> (8, 7)
test/test_012.md:8:7:	lparen	"("	false	0	0.000000	(8, 7) -> (8, 8)
test/test_012.md:8:8:	ident	"x"	false	0	0.000000	(8, 8) -> (8, 9)
test/test_012.md:8:9:	rparen	")"	false	0	0.000000	(8, 9) -> (8, 10)
test/test_012.md:8:10:	semicolon	";"	false	0	0.000000	(8, 10) -> (8, 11)
test/test_012.md:9:4:	ident	"Texts"	false	0	0.000000	(9, 4) -> (9, 9)
test/test_012.md:9:9:	dot	"."	false	0	0.000000	(9, 9) -> (9, 10)
test/test_012.md:9:10:	ident	"WriteInt"	false	0	0.000000	(9, 10) -> (9, 18)
test/test_012.md:9:18:	lparen	"("	false	0	0.000000	(9, 18) -> (9, 19)
test/test_012.md:9:19:	ident	"x"	false	0	0.000000	(9, 19) -> (9, 20)
test/test_012.md:9:20:	rparen	")"	false	0	0.000000	(9, 20) -> (9, 21)
test/test_012.md:9:21:	semicolon	";"	false	0	0.000000	(9, 21) -> (9, 22)
test/test_012.md:9:23:	ident	"Texts"	false	0	0.000000	(9, 23) -> (9, 28)
test/test_012.md:9:28:	dot	"."	false	0	0.000000	(9, 28) -> (9, 29)
test/test_012.md:9:29:	ident	"WriteLn"	false	0	0.000000	(9, 29) -> (9, 36)
test/test_012.md:9:36:	semicolon	";"	false	0	0.000000	(9, 36) -> (9, 37)
test/test_012.md:10:2:	until	"UNTIL"	false	0	0.000000	(10, 2) -> (10, 7)
test/test_012.md:10:8:	ident	"x"	false	0	0.000000	(10, 8) -> (10, 9)
test/test_012.md:10:10:	eq	"="	false	0	0.000000	(10, 10) -> (10, 11)
test/test_012.md:10:12:	integer	"0"	false	0	0.000000	(10, 12) -> (10, 13)
test/test_012.md:10:13:	semicolon	";"	false	0	0.000000	(10, 13) -> (10, 14)
test/test_012.md:11:0:	end	"END"	false	0	0.000000	(11, 0) -> (11, 3)
test/test_012.md:11:4:	ident	"Repeat"	false	0	0.000000	(11, 4) -> (11, 10)
test/test_012.md:11:10:	dot	"."	false	0	0.000000	(11, 10) -> (11, 11)
test/test_012.md:12:0:	eof	""	false	0	0.000000	(12, 0) -> (12, 0)
```
## AST
```scheme
(module "Repeat"
  (vars
    (x 
      (INTEGER [i64])
    )
  )
  (stmts
    (assign
      (variable [i64] "x")
      (number [i64] 10)
    )
    (repeat
      (stmts
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
      (eq [boolean]
        (variable [i64] "x")
        (number [i64] 0)
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
	store i64 10, i64* @0
	br label %0

0:
	%1 = load i64, i64* @0
	%2 = sub i64 %1, 1
	store i64 %2, i64* @0
	%3 = load i64, i64* @0
	%4 = getelementptr [3 x i8], [3 x i8]* @1, i64 0, i64 0
	%5 = call i64 (i8*, ...) @printf(i8* %4, i64 %3)
	%6 = getelementptr [1 x i8], [1 x i8]* @2, i64 0, i64 0
	%7 = call i64 @puts(i8* %6)
	%8 = load i64, i64* @0
	%9 = icmp eq i64 %8, 0
	br i1 %9, label %10, label %0

10:
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
9
8
7
6
5
4
3
2
1
0
```
