# test/test_014.md
## Source
```pascal
MODULE ExtendedWhile;

VAR x: INTEGER;
    y: INTEGER;

BEGIN
  x := 0; y := 1;
  WHILE x < 10 DO
    INC(x); print(x*y)
  ELSIF y < 10 DO 
    INC(y); x := 0;
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
test/test_014.md:9:12:	ident	"print"	false	0	0.000000	(9, 12) -> (9, 17)
test/test_014.md:9:17:	lparen	"("	false	0	0.000000	(9, 17) -> (9, 18)
test/test_014.md:9:18:	ident	"x"	false	0	0.000000	(9, 18) -> (9, 19)
test/test_014.md:9:19:	asterisk	"*"	false	0	0.000000	(9, 19) -> (9, 20)
test/test_014.md:9:20:	ident	"y"	false	0	0.000000	(9, 20) -> (9, 21)
test/test_014.md:9:21:	rparen	")"	false	0	0.000000	(9, 21) -> (9, 22)
test/test_014.md:10:2:	elsif	"ELSIF"	false	0	0.000000	(10, 2) -> (10, 7)
test/test_014.md:10:8:	ident	"y"	false	0	0.000000	(10, 8) -> (10, 9)
test/test_014.md:10:10:	lt	"<"	false	0	0.000000	(10, 10) -> (10, 11)
test/test_014.md:10:12:	integer	"10"	false	10	0.000000	(10, 12) -> (10, 14)
test/test_014.md:10:15:	do	"DO"	false	0	0.000000	(10, 15) -> (10, 17)
test/test_014.md:11:4:	ident	"INC"	false	0	0.000000	(11, 4) -> (11, 7)
test/test_014.md:11:7:	lparen	"("	false	0	0.000000	(11, 7) -> (11, 8)
test/test_014.md:11:8:	ident	"y"	false	0	0.000000	(11, 8) -> (11, 9)
test/test_014.md:11:9:	rparen	")"	false	0	0.000000	(11, 9) -> (11, 10)
test/test_014.md:11:10:	semicolon	";"	false	0	0.000000	(11, 10) -> (11, 11)
test/test_014.md:11:12:	ident	"x"	false	0	0.000000	(11, 12) -> (11, 13)
test/test_014.md:11:14:	assign	":="	false	0	0.000000	(11, 14) -> (11, 16)
test/test_014.md:11:17:	integer	"0"	false	0	0.000000	(11, 17) -> (11, 18)
test/test_014.md:11:18:	semicolon	";"	false	0	0.000000	(11, 18) -> (11, 19)
test/test_014.md:12:2:	end	"END"	false	0	0.000000	(12, 2) -> (12, 5)
test/test_014.md:13:0:	end	"END"	false	0	0.000000	(13, 0) -> (13, 3)
test/test_014.md:13:4:	ident	"ExtendedWhile"	false	0	0.000000	(13, 4) -> (13, 17)
test/test_014.md:13:17:	dot	"."	false	0	0.000000	(13, 17) -> (13, 18)
test/test_014.md:14:0:	eof	""	false	0	0.000000	(14, 0) -> (14, 0)
```
## AST
```scheme
(module "ExtendedWhile"
  (vars
    (x [i64])
    (y [i64])
  )
  (stmts
    (assign x
      (number [i64] 0)
    )
    (assign y
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
            (call "INC" [void]
              (variable [i64] "x")
            )
          )
          (expr2stmt
            (call "print" [void]
              (asterisk [i64]
                (variable [i64] "x")
                (variable [i64] "y")
              )
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
            (call "INC" [void]
              (variable [i64] "y")
            )
          )
          (assign x
            (number [i64] 0)
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
@2 = global [4 x i8] c"%d\0A\00"

declare i64 @puts(i8* %str)

declare i64 @rand()

declare i64 @sprintf(i8* %buf, i8* %format, ...)

declare i64 @printf(i8* %format, ...)

define i64 @main() {
entry:
	store i64 0, i64* @0
	store i64 1, i64* @1
	br label %0

0:
	%1 = load i64, i64* @0
	%2 = icmp slt i64 %1, 10
	br i1 %2, label %3, label %11

3:
	%4 = load i64, i64* @0
	%5 = add i64 %4, 1
	store i64 %5, i64* @0
	%6 = load i64, i64* @0
	%7 = load i64, i64* @1
	%8 = mul i64 %6, %7
	%9 = getelementptr [4 x i8], [4 x i8]* @2, i64 0, i64 0
	%10 = call i64 (i8*, ...) @printf(i8* %9, i64 %8)
	br label %0

11:
	%12 = load i64, i64* @1
	%13 = icmp slt i64 %12, 10
	br i1 %13, label %14, label %17

14:
	%15 = load i64, i64* @1
	%16 = add i64 %15, 1
	store i64 %16, i64* @1
	store i64 0, i64* @0
	br label %0

17:
	ret i64 0
}

```
## Run
```bash
1
2
3
4
5
6
7
8
9
10
2
4
6
8
10
12
14
16
18
20
3
6
9
12
15
18
21
24
27
30
4
8
12
16
20
24
28
32
36
40
5
10
15
20
25
30
35
40
45
50
6
12
18
24
30
36
42
48
54
60
7
14
21
28
35
42
49
56
63
70
8
16
24
32
40
48
56
64
72
80
9
18
27
36
45
54
63
72
81
90
10
20
30
40
50
60
70
80
90
100
```
