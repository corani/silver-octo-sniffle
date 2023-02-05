# test/test_013.md
## Source
```pascal
MODULE While;

VAR x: INTEGER;
    y: INTEGER;

BEGIN
  x := 0; WHILE x < 10 DO
    INC(x);
    y := 0; WHILE y < 10 DO
      INC(y);
      print(x*y)
    END
  END
END While.
```
## Tokens
```tsv
test/test_013.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_013.md:1:8:	ident	"While"	false	0	0.000000	(1, 8) -> (1, 13)
test/test_013.md:1:13:	semicolon	";"	false	0	0.000000	(1, 13) -> (1, 14)
test/test_013.md:3:0:	var	"VAR"	false	0	0.000000	(3, 0) -> (3, 3)
test/test_013.md:3:4:	ident	"x"	false	0	0.000000	(3, 4) -> (3, 5)
test/test_013.md:3:5:	colon	":"	false	0	0.000000	(3, 5) -> (3, 6)
test/test_013.md:3:7:	ident	"INTEGER"	false	0	0.000000	(3, 7) -> (3, 14)
test/test_013.md:3:14:	semicolon	";"	false	0	0.000000	(3, 14) -> (3, 15)
test/test_013.md:4:4:	ident	"y"	false	0	0.000000	(4, 4) -> (4, 5)
test/test_013.md:4:5:	colon	":"	false	0	0.000000	(4, 5) -> (4, 6)
test/test_013.md:4:7:	ident	"INTEGER"	false	0	0.000000	(4, 7) -> (4, 14)
test/test_013.md:4:14:	semicolon	";"	false	0	0.000000	(4, 14) -> (4, 15)
test/test_013.md:6:0:	begin	"BEGIN"	false	0	0.000000	(6, 0) -> (6, 5)
test/test_013.md:7:2:	ident	"x"	false	0	0.000000	(7, 2) -> (7, 3)
test/test_013.md:7:4:	assign	":="	false	0	0.000000	(7, 4) -> (7, 6)
test/test_013.md:7:7:	integer	"0"	false	0	0.000000	(7, 7) -> (7, 8)
test/test_013.md:7:8:	semicolon	";"	false	0	0.000000	(7, 8) -> (7, 9)
test/test_013.md:7:10:	while	"WHILE"	false	0	0.000000	(7, 10) -> (7, 15)
test/test_013.md:7:16:	ident	"x"	false	0	0.000000	(7, 16) -> (7, 17)
test/test_013.md:7:18:	lt	"<"	false	0	0.000000	(7, 18) -> (7, 19)
test/test_013.md:7:20:	integer	"10"	false	10	0.000000	(7, 20) -> (7, 22)
test/test_013.md:7:23:	do	"DO"	false	0	0.000000	(7, 23) -> (7, 25)
test/test_013.md:8:4:	ident	"INC"	false	0	0.000000	(8, 4) -> (8, 7)
test/test_013.md:8:7:	lparen	"("	false	0	0.000000	(8, 7) -> (8, 8)
test/test_013.md:8:8:	ident	"x"	false	0	0.000000	(8, 8) -> (8, 9)
test/test_013.md:8:9:	rparen	")"	false	0	0.000000	(8, 9) -> (8, 10)
test/test_013.md:8:10:	semicolon	";"	false	0	0.000000	(8, 10) -> (8, 11)
test/test_013.md:9:4:	ident	"y"	false	0	0.000000	(9, 4) -> (9, 5)
test/test_013.md:9:6:	assign	":="	false	0	0.000000	(9, 6) -> (9, 8)
test/test_013.md:9:9:	integer	"0"	false	0	0.000000	(9, 9) -> (9, 10)
test/test_013.md:9:10:	semicolon	";"	false	0	0.000000	(9, 10) -> (9, 11)
test/test_013.md:9:12:	while	"WHILE"	false	0	0.000000	(9, 12) -> (9, 17)
test/test_013.md:9:18:	ident	"y"	false	0	0.000000	(9, 18) -> (9, 19)
test/test_013.md:9:20:	lt	"<"	false	0	0.000000	(9, 20) -> (9, 21)
test/test_013.md:9:22:	integer	"10"	false	10	0.000000	(9, 22) -> (9, 24)
test/test_013.md:9:25:	do	"DO"	false	0	0.000000	(9, 25) -> (9, 27)
test/test_013.md:10:6:	ident	"INC"	false	0	0.000000	(10, 6) -> (10, 9)
test/test_013.md:10:9:	lparen	"("	false	0	0.000000	(10, 9) -> (10, 10)
test/test_013.md:10:10:	ident	"y"	false	0	0.000000	(10, 10) -> (10, 11)
test/test_013.md:10:11:	rparen	")"	false	0	0.000000	(10, 11) -> (10, 12)
test/test_013.md:10:12:	semicolon	";"	false	0	0.000000	(10, 12) -> (10, 13)
test/test_013.md:11:6:	ident	"print"	false	0	0.000000	(11, 6) -> (11, 11)
test/test_013.md:11:11:	lparen	"("	false	0	0.000000	(11, 11) -> (11, 12)
test/test_013.md:11:12:	ident	"x"	false	0	0.000000	(11, 12) -> (11, 13)
test/test_013.md:11:13:	asterisk	"*"	false	0	0.000000	(11, 13) -> (11, 14)
test/test_013.md:11:14:	ident	"y"	false	0	0.000000	(11, 14) -> (11, 15)
test/test_013.md:11:15:	rparen	")"	false	0	0.000000	(11, 15) -> (11, 16)
test/test_013.md:12:4:	end	"END"	false	0	0.000000	(12, 4) -> (12, 7)
test/test_013.md:13:2:	end	"END"	false	0	0.000000	(13, 2) -> (13, 5)
test/test_013.md:14:0:	end	"END"	false	0	0.000000	(14, 0) -> (14, 3)
test/test_013.md:14:4:	ident	"While"	false	0	0.000000	(14, 4) -> (14, 9)
test/test_013.md:14:9:	dot	"."	false	0	0.000000	(14, 9) -> (14, 10)
test/test_013.md:15:0:	eof	""	false	0	0.000000	(15, 0) -> (15, 0)
```
## AST
```scheme
(module "While"
  (vars
    (x [i64])
    (y [i64])
  )
  (stmts
    (assign x
      (number [i64] 0)
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
          (assign y
            (number [i64] 0)
          )
          (while
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
	br label %0

0:
	%1 = load i64, i64* @0
	%2 = icmp slt i64 %1, 10
	br i1 %2, label %3, label %6

3:
	%4 = load i64, i64* @0
	%5 = add i64 %4, 1
	store i64 %5, i64* @0
	store i64 0, i64* @1
	br label %7

6:
	ret i64 0

7:
	%8 = load i64, i64* @1
	%9 = icmp slt i64 %8, 10
	br i1 %9, label %10, label %18

10:
	%11 = load i64, i64* @1
	%12 = add i64 %11, 1
	store i64 %12, i64* @1
	%13 = load i64, i64* @0
	%14 = load i64, i64* @1
	%15 = mul i64 %13, %14
	%16 = getelementptr [4 x i8], [4 x i8]* @2, i64 0, i64 0
	%17 = call i64 (i8*, ...) @printf(i8* %16, i64 %15)
	br label %7

18:
	br label %0
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
