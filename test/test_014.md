# test/test_014.md
## Source
```pascal
MODULE ExtendedWhile;

VAR x: INTEGER;
    y: INTEGER;

BEGIN
  x := 0; y := 1;
  WHILE x < 10 DO
    INC(x); 
    Texts.WriteInt(x*y);  Texts.Write(",")
  ELSIF y < 10 DO 
    INC(y); x := 0;
    Texts.WriteLn
  END
END ExtendedWhile.
```
## Tokens
```tsv
test/test_014.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_014.md:1:8:	ident	"ExtendedWhile"	false	0	0.000000	(1, 8) -> (1, 21)
test/test_014.md:1:21:	semicolon	";"	false	0	0.000000	(1, 21) -> (1, 22)
test/test_014.md:3:0:	var	"VAR"	false	0	0.000000	(3, 0) -> (3, 3)
test/test_014.md:3:4:	ident	"x"	false	0	0.000000	(3, 4) -> (3, 5)
test/test_014.md:3:5:	colon	":"	false	0	0.000000	(3, 5) -> (3, 6)
test/test_014.md:3:7:	ident	"INTEGER"	false	0	0.000000	(3, 7) -> (3, 14)
test/test_014.md:3:14:	semicolon	";"	false	0	0.000000	(3, 14) -> (3, 15)
test/test_014.md:4:4:	ident	"y"	false	0	0.000000	(4, 4) -> (4, 5)
test/test_014.md:4:5:	colon	":"	false	0	0.000000	(4, 5) -> (4, 6)
test/test_014.md:4:7:	ident	"INTEGER"	false	0	0.000000	(4, 7) -> (4, 14)
test/test_014.md:4:14:	semicolon	";"	false	0	0.000000	(4, 14) -> (4, 15)
test/test_014.md:6:0:	begin	"BEGIN"	false	0	0.000000	(6, 0) -> (6, 5)
test/test_014.md:7:2:	ident	"x"	false	0	0.000000	(7, 2) -> (7, 3)
test/test_014.md:7:4:	assign	":="	false	0	0.000000	(7, 4) -> (7, 6)
test/test_014.md:7:7:	integer	"0"	false	0	0.000000	(7, 7) -> (7, 8)
test/test_014.md:7:8:	semicolon	";"	false	0	0.000000	(7, 8) -> (7, 9)
test/test_014.md:7:10:	ident	"y"	false	0	0.000000	(7, 10) -> (7, 11)
test/test_014.md:7:12:	assign	":="	false	0	0.000000	(7, 12) -> (7, 14)
test/test_014.md:7:15:	integer	"1"	false	1	0.000000	(7, 15) -> (7, 16)
test/test_014.md:7:16:	semicolon	";"	false	0	0.000000	(7, 16) -> (7, 17)
test/test_014.md:8:2:	while	"WHILE"	false	0	0.000000	(8, 2) -> (8, 7)
test/test_014.md:8:8:	ident	"x"	false	0	0.000000	(8, 8) -> (8, 9)
test/test_014.md:8:10:	lt	"<"	false	0	0.000000	(8, 10) -> (8, 11)
test/test_014.md:8:12:	integer	"10"	false	10	0.000000	(8, 12) -> (8, 14)
test/test_014.md:8:15:	do	"DO"	false	0	0.000000	(8, 15) -> (8, 17)
test/test_014.md:9:4:	ident	"INC"	false	0	0.000000	(9, 4) -> (9, 7)
test/test_014.md:9:7:	lparen	"("	false	0	0.000000	(9, 7) -> (9, 8)
test/test_014.md:9:8:	ident	"x"	false	0	0.000000	(9, 8) -> (9, 9)
test/test_014.md:9:9:	rparen	")"	false	0	0.000000	(9, 9) -> (9, 10)
test/test_014.md:9:10:	semicolon	";"	false	0	0.000000	(9, 10) -> (9, 11)
test/test_014.md:10:4:	ident	"Texts"	false	0	0.000000	(10, 4) -> (10, 9)
test/test_014.md:10:9:	dot	"."	false	0	0.000000	(10, 9) -> (10, 10)
test/test_014.md:10:10:	ident	"WriteInt"	false	0	0.000000	(10, 10) -> (10, 18)
test/test_014.md:10:18:	lparen	"("	false	0	0.000000	(10, 18) -> (10, 19)
test/test_014.md:10:19:	ident	"x"	false	0	0.000000	(10, 19) -> (10, 20)
test/test_014.md:10:20:	asterisk	"*"	false	0	0.000000	(10, 20) -> (10, 21)
test/test_014.md:10:21:	ident	"y"	false	0	0.000000	(10, 21) -> (10, 22)
test/test_014.md:10:22:	rparen	")"	false	0	0.000000	(10, 22) -> (10, 23)
test/test_014.md:10:23:	semicolon	";"	false	0	0.000000	(10, 23) -> (10, 24)
test/test_014.md:10:26:	ident	"Texts"	false	0	0.000000	(10, 26) -> (10, 31)
test/test_014.md:10:31:	dot	"."	false	0	0.000000	(10, 31) -> (10, 32)
test/test_014.md:10:32:	ident	"Write"	false	0	0.000000	(10, 32) -> (10, 37)
test/test_014.md:10:37:	lparen	"("	false	0	0.000000	(10, 37) -> (10, 38)
test/test_014.md:10:38:	string	","	false	0	0.000000	(10, 38) -> (10, 41)
test/test_014.md:10:41:	rparen	")"	false	0	0.000000	(10, 41) -> (10, 42)
test/test_014.md:11:2:	elsif	"ELSIF"	false	0	0.000000	(11, 2) -> (11, 7)
test/test_014.md:11:8:	ident	"y"	false	0	0.000000	(11, 8) -> (11, 9)
test/test_014.md:11:10:	lt	"<"	false	0	0.000000	(11, 10) -> (11, 11)
test/test_014.md:11:12:	integer	"10"	false	10	0.000000	(11, 12) -> (11, 14)
test/test_014.md:11:15:	do	"DO"	false	0	0.000000	(11, 15) -> (11, 17)
test/test_014.md:12:4:	ident	"INC"	false	0	0.000000	(12, 4) -> (12, 7)
test/test_014.md:12:7:	lparen	"("	false	0	0.000000	(12, 7) -> (12, 8)
test/test_014.md:12:8:	ident	"y"	false	0	0.000000	(12, 8) -> (12, 9)
test/test_014.md:12:9:	rparen	")"	false	0	0.000000	(12, 9) -> (12, 10)
test/test_014.md:12:10:	semicolon	";"	false	0	0.000000	(12, 10) -> (12, 11)
test/test_014.md:12:12:	ident	"x"	false	0	0.000000	(12, 12) -> (12, 13)
test/test_014.md:12:14:	assign	":="	false	0	0.000000	(12, 14) -> (12, 16)
test/test_014.md:12:17:	integer	"0"	false	0	0.000000	(12, 17) -> (12, 18)
test/test_014.md:12:18:	semicolon	";"	false	0	0.000000	(12, 18) -> (12, 19)
test/test_014.md:13:4:	ident	"Texts"	false	0	0.000000	(13, 4) -> (13, 9)
test/test_014.md:13:9:	dot	"."	false	0	0.000000	(13, 9) -> (13, 10)
test/test_014.md:13:10:	ident	"WriteLn"	false	0	0.000000	(13, 10) -> (13, 17)
test/test_014.md:14:2:	end	"END"	false	0	0.000000	(14, 2) -> (14, 5)
test/test_014.md:15:0:	end	"END"	false	0	0.000000	(15, 0) -> (15, 3)
test/test_014.md:15:4:	ident	"ExtendedWhile"	false	0	0.000000	(15, 4) -> (15, 17)
test/test_014.md:15:17:	dot	"."	false	0	0.000000	(15, 17) -> (15, 18)
test/test_014.md:16:0:	eof	""	false	0	0.000000	(16, 0) -> (16, 0)
```
## AST
```scheme
(module "ExtendedWhile"
  (vars
    (x 
      (INTEGER [i64])
    )
    (y 
      (INTEGER [i64])
    )
  )
  (stmts
    (assign
      (variable [i64] "x")
      (number [i64] 0)
    )
    (assign
      (variable [i64] "y")
      (number [i64] 1)
    )
    (while
      (cond
        (lt [boolean]
          (variable [i64] "x")
          (number [i64] 10)
        )
        (stmts
          (expr2stmt
            (call
              (procedure [void] "INC")
              (variable [i64] "x")
            )
          )
          (expr2stmt
            (call
              (procedure [void] "Texts.WriteInt")
              (asterisk [i64]
                (variable [i64] "x")
                (variable [i64] "y")
              )
            )
          )
          (expr2stmt
            (call
              (procedure [void] "Texts.Write")
              (char ",")
            )
          )
        )
      )
      (cond
        (lt [boolean]
          (variable [i64] "y")
          (number [i64] 10)
        )
        (stmts
          (expr2stmt
            (call
              (procedure [void] "INC")
              (variable [i64] "y")
            )
          )
          (assign
            (variable [i64] "x")
            (number [i64] 0)
          )
          (expr2stmt
            (call
              (procedure [void] "Texts.WriteLn")
            )
          )
        )
      )
    )
  )
)
```
## IR
```llvm
@0 = global i64 0
@1 = global i64 0
@2 = global [3 x i8] c"%d\00"
@3 = global [3 x i8] c"%c\00"
@4 = global [1 x i8] c"\00"
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
	store i64 0, i64* @0
	store i64 1, i64* @1
	br label %0

0:
	%1 = load i64, i64* @0
	%2 = icmp slt i64 %1, 10
	br i1 %2, label %3, label %13

3:
	%4 = load i64, i64* @0
	%5 = add i64 %4, 1
	store i64 %5, i64* @0
	%6 = load i64, i64* @0
	%7 = load i64, i64* @1
	%8 = mul i64 %6, %7
	%9 = getelementptr [3 x i8], [3 x i8]* @2, i64 0, i64 0
	%10 = call i64 (i8*, ...) @printf(i8* %9, i64 %8)
	%11 = getelementptr [3 x i8], [3 x i8]* @3, i64 0, i64 0
	%12 = call i64 (i8*, ...) @printf(i8* %11, i8 44)
	br label %0

13:
	%14 = load i64, i64* @1
	%15 = icmp slt i64 %14, 10
	br i1 %15, label %16, label %21

16:
	%17 = load i64, i64* @1
	%18 = add i64 %17, 1
	store i64 %18, i64* @1
	store i64 0, i64* @0
	%19 = getelementptr [1 x i8], [1 x i8]* @4, i64 0, i64 0
	%20 = call i64 @puts(i8* %19)
	br label %0

21:
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
1,2,3,4,5,6,7,8,9,10,
2,4,6,8,10,12,14,16,18,20,
3,6,9,12,15,18,21,24,27,30,
4,8,12,16,20,24,28,32,36,40,
5,10,15,20,25,30,35,40,45,50,
6,12,18,24,30,36,42,48,54,60,
7,14,21,28,35,42,49,56,63,70,
8,16,24,32,40,48,56,64,72,80,
9,18,27,36,45,54,63,72,81,90,
10,20,30,40,50,60,70,80,90,100,```
