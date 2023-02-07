# test/test_017.md
## Source
```pascal
MODULE Sets;

VAR x: SET;
    y: SET;
    z: SET;

BEGIN
    x := {1, 2, 3};
    y := {10..20, 30};
   
    print(x+y);

    INCL(x, 4);
    EXCL(y, 30);

    z := x + y;

    print(z);

    INCL(z, 30);
    EXCL(z, 4);

    print(z);
END Sets.
```
## Tokens
```tsv
test/test_017.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_017.md:1:8:	ident	"Sets"	false	0	0.000000	(1, 8) -> (1, 12)
test/test_017.md:1:12:	semicolon	";"	false	0	0.000000	(1, 12) -> (1, 13)
test/test_017.md:3:0:	var	"VAR"	false	0	0.000000	(3, 0) -> (3, 3)
test/test_017.md:3:4:	ident	"x"	false	0	0.000000	(3, 4) -> (3, 5)
test/test_017.md:3:5:	colon	":"	false	0	0.000000	(3, 5) -> (3, 6)
test/test_017.md:3:7:	ident	"SET"	false	0	0.000000	(3, 7) -> (3, 10)
test/test_017.md:3:10:	semicolon	";"	false	0	0.000000	(3, 10) -> (3, 11)
test/test_017.md:4:4:	ident	"y"	false	0	0.000000	(4, 4) -> (4, 5)
test/test_017.md:4:5:	colon	":"	false	0	0.000000	(4, 5) -> (4, 6)
test/test_017.md:4:7:	ident	"SET"	false	0	0.000000	(4, 7) -> (4, 10)
test/test_017.md:4:10:	semicolon	";"	false	0	0.000000	(4, 10) -> (4, 11)
test/test_017.md:5:4:	ident	"z"	false	0	0.000000	(5, 4) -> (5, 5)
test/test_017.md:5:5:	colon	":"	false	0	0.000000	(5, 5) -> (5, 6)
test/test_017.md:5:7:	ident	"SET"	false	0	0.000000	(5, 7) -> (5, 10)
test/test_017.md:5:10:	semicolon	";"	false	0	0.000000	(5, 10) -> (5, 11)
test/test_017.md:7:0:	begin	"BEGIN"	false	0	0.000000	(7, 0) -> (7, 5)
test/test_017.md:8:4:	ident	"x"	false	0	0.000000	(8, 4) -> (8, 5)
test/test_017.md:8:6:	assign	":="	false	0	0.000000	(8, 6) -> (8, 8)
test/test_017.md:8:9:	lbrace	"{"	false	0	0.000000	(8, 9) -> (8, 10)
test/test_017.md:8:10:	integer	"1"	false	1	0.000000	(8, 10) -> (8, 11)
test/test_017.md:8:11:	comma	","	false	0	0.000000	(8, 11) -> (8, 12)
test/test_017.md:8:13:	integer	"2"	false	2	0.000000	(8, 13) -> (8, 14)
test/test_017.md:8:14:	comma	","	false	0	0.000000	(8, 14) -> (8, 15)
test/test_017.md:8:16:	integer	"3"	false	3	0.000000	(8, 16) -> (8, 17)
test/test_017.md:8:17:	rbrace	"}"	false	0	0.000000	(8, 17) -> (8, 18)
test/test_017.md:8:18:	semicolon	";"	false	0	0.000000	(8, 18) -> (8, 19)
test/test_017.md:9:4:	ident	"y"	false	0	0.000000	(9, 4) -> (9, 5)
test/test_017.md:9:6:	assign	":="	false	0	0.000000	(9, 6) -> (9, 8)
test/test_017.md:9:9:	lbrace	"{"	false	0	0.000000	(9, 9) -> (9, 10)
test/test_017.md:9:10:	integer	"10"	false	10	0.000000	(9, 10) -> (9, 14)
test/test_017.md:9:14:	dotdot	".."	false	0	0.000000	(9, 14) -> (9, 16)
test/test_017.md:9:16:	integer	"20"	false	20	0.000000	(9, 16) -> (9, 18)
test/test_017.md:9:18:	comma	","	false	0	0.000000	(9, 18) -> (9, 19)
test/test_017.md:9:20:	integer	"30"	false	30	0.000000	(9, 20) -> (9, 22)
test/test_017.md:9:22:	rbrace	"}"	false	0	0.000000	(9, 22) -> (9, 23)
test/test_017.md:9:23:	semicolon	";"	false	0	0.000000	(9, 23) -> (9, 24)
test/test_017.md:11:4:	ident	"print"	false	0	0.000000	(11, 4) -> (11, 9)
test/test_017.md:11:9:	lparen	"("	false	0	0.000000	(11, 9) -> (11, 10)
test/test_017.md:11:10:	ident	"x"	false	0	0.000000	(11, 10) -> (11, 11)
test/test_017.md:11:11:	plus	"+"	false	0	0.000000	(11, 11) -> (11, 12)
test/test_017.md:11:12:	ident	"y"	false	0	0.000000	(11, 12) -> (11, 13)
test/test_017.md:11:13:	rparen	")"	false	0	0.000000	(11, 13) -> (11, 14)
test/test_017.md:11:14:	semicolon	";"	false	0	0.000000	(11, 14) -> (11, 15)
test/test_017.md:13:4:	ident	"INCL"	false	0	0.000000	(13, 4) -> (13, 8)
test/test_017.md:13:8:	lparen	"("	false	0	0.000000	(13, 8) -> (13, 9)
test/test_017.md:13:9:	ident	"x"	false	0	0.000000	(13, 9) -> (13, 10)
test/test_017.md:13:10:	comma	","	false	0	0.000000	(13, 10) -> (13, 11)
test/test_017.md:13:12:	integer	"4"	false	4	0.000000	(13, 12) -> (13, 13)
test/test_017.md:13:13:	rparen	")"	false	0	0.000000	(13, 13) -> (13, 14)
test/test_017.md:13:14:	semicolon	";"	false	0	0.000000	(13, 14) -> (13, 15)
test/test_017.md:14:4:	ident	"EXCL"	false	0	0.000000	(14, 4) -> (14, 8)
test/test_017.md:14:8:	lparen	"("	false	0	0.000000	(14, 8) -> (14, 9)
test/test_017.md:14:9:	ident	"y"	false	0	0.000000	(14, 9) -> (14, 10)
test/test_017.md:14:10:	comma	","	false	0	0.000000	(14, 10) -> (14, 11)
test/test_017.md:14:12:	integer	"30"	false	30	0.000000	(14, 12) -> (14, 14)
test/test_017.md:14:14:	rparen	")"	false	0	0.000000	(14, 14) -> (14, 15)
test/test_017.md:14:15:	semicolon	";"	false	0	0.000000	(14, 15) -> (14, 16)
test/test_017.md:16:4:	ident	"z"	false	0	0.000000	(16, 4) -> (16, 5)
test/test_017.md:16:6:	assign	":="	false	0	0.000000	(16, 6) -> (16, 8)
test/test_017.md:16:9:	ident	"x"	false	0	0.000000	(16, 9) -> (16, 10)
test/test_017.md:16:11:	plus	"+"	false	0	0.000000	(16, 11) -> (16, 12)
test/test_017.md:16:13:	ident	"y"	false	0	0.000000	(16, 13) -> (16, 14)
test/test_017.md:16:14:	semicolon	";"	false	0	0.000000	(16, 14) -> (16, 15)
test/test_017.md:18:4:	ident	"print"	false	0	0.000000	(18, 4) -> (18, 9)
test/test_017.md:18:9:	lparen	"("	false	0	0.000000	(18, 9) -> (18, 10)
test/test_017.md:18:10:	ident	"z"	false	0	0.000000	(18, 10) -> (18, 11)
test/test_017.md:18:11:	rparen	")"	false	0	0.000000	(18, 11) -> (18, 12)
test/test_017.md:18:12:	semicolon	";"	false	0	0.000000	(18, 12) -> (18, 13)
test/test_017.md:20:4:	ident	"INCL"	false	0	0.000000	(20, 4) -> (20, 8)
test/test_017.md:20:8:	lparen	"("	false	0	0.000000	(20, 8) -> (20, 9)
test/test_017.md:20:9:	ident	"z"	false	0	0.000000	(20, 9) -> (20, 10)
test/test_017.md:20:10:	comma	","	false	0	0.000000	(20, 10) -> (20, 11)
test/test_017.md:20:12:	integer	"30"	false	30	0.000000	(20, 12) -> (20, 14)
test/test_017.md:20:14:	rparen	")"	false	0	0.000000	(20, 14) -> (20, 15)
test/test_017.md:20:15:	semicolon	";"	false	0	0.000000	(20, 15) -> (20, 16)
test/test_017.md:21:4:	ident	"EXCL"	false	0	0.000000	(21, 4) -> (21, 8)
test/test_017.md:21:8:	lparen	"("	false	0	0.000000	(21, 8) -> (21, 9)
test/test_017.md:21:9:	ident	"z"	false	0	0.000000	(21, 9) -> (21, 10)
test/test_017.md:21:10:	comma	","	false	0	0.000000	(21, 10) -> (21, 11)
test/test_017.md:21:12:	integer	"4"	false	4	0.000000	(21, 12) -> (21, 13)
test/test_017.md:21:13:	rparen	")"	false	0	0.000000	(21, 13) -> (21, 14)
test/test_017.md:21:14:	semicolon	";"	false	0	0.000000	(21, 14) -> (21, 15)
test/test_017.md:23:4:	ident	"print"	false	0	0.000000	(23, 4) -> (23, 9)
test/test_017.md:23:9:	lparen	"("	false	0	0.000000	(23, 9) -> (23, 10)
test/test_017.md:23:10:	ident	"z"	false	0	0.000000	(23, 10) -> (23, 11)
test/test_017.md:23:11:	rparen	")"	false	0	0.000000	(23, 11) -> (23, 12)
test/test_017.md:23:12:	semicolon	";"	false	0	0.000000	(23, 12) -> (23, 13)
test/test_017.md:24:0:	end	"END"	false	0	0.000000	(24, 0) -> (24, 3)
test/test_017.md:24:4:	ident	"Sets"	false	0	0.000000	(24, 4) -> (24, 8)
test/test_017.md:24:8:	dot	"."	false	0	0.000000	(24, 8) -> (24, 9)
test/test_017.md:25:0:	eof	""	false	0	0.000000	(25, 0) -> (25, 0)
```
## AST
```scheme
(module "Sets"
  (vars
    (x [set])
    (y [set])
    (z [set])
  )
  (stmts
    (assign x
      (set (1..3))
    )
    (assign y
      (set (10..20, 30))
    )
    (expr2stmt
      (call "print" [void]
        (plus [set]
          (variable [set] "x")
          (variable [set] "y")
        )
      )
    )
    (expr2stmt
      (call "INCL" [void]
        (variable [set] "x")
        (number [i64] 4)
      )
    )
    (expr2stmt
      (call "EXCL" [void]
        (variable [set] "y")
        (number [i64] 30)
      )
    )
    (assign z
      (plus [set]
        (variable [set] "x")
        (variable [set] "y")
      )
    )
    (expr2stmt
      (call "print" [void]
        (variable [set] "z")
      )
    )
    (expr2stmt
      (call "INCL" [void]
        (variable [set] "z")
        (number [i64] 30)
      )
    )
    (expr2stmt
      (call "EXCL" [void]
        (variable [set] "z")
        (number [i64] 4)
      )
    )
    (expr2stmt
      (call "print" [void]
        (variable [set] "z")
      )
    )
  )
)
```
## IR
```llvm
@0 = global i64 0
@1 = global i64 0
@2 = global i64 0
@3 = global [4 x i8] c"%d\0A\00"

declare i64 @puts(i8* %str)

declare i64 @rand()

declare i64 @sprintf(i8* %buf, i8* %format, ...)

declare i64 @printf(i8* %format, ...)

define i64 @main() {
entry:
	store i64 14, i64* @0
	store i64 1075837952, i64* @1
	%0 = load i64, i64* @0
	%1 = load i64, i64* @1
	%2 = or i64 %0, %1
	%3 = getelementptr [4 x i8], [4 x i8]* @3, i64 0, i64 0
	%4 = call i64 (i8*, ...) @printf(i8* %3, i64 %2)
	%5 = load i64, i64* @0
	%6 = shl i64 1, 4
	%7 = or i64 %5, %6
	store i64 %7, i64* @0
	%8 = load i64, i64* @1
	%9 = shl i64 1, 30
	%10 = xor i64 %9, -1
	%11 = and i64 %8, %10
	store i64 %11, i64* @1
	%12 = load i64, i64* @0
	%13 = load i64, i64* @1
	%14 = or i64 %12, %13
	store i64 %14, i64* @2
	%15 = load i64, i64* @2
	%16 = getelementptr [4 x i8], [4 x i8]* @3, i64 0, i64 0
	%17 = call i64 (i8*, ...) @printf(i8* %16, i64 %15)
	%18 = load i64, i64* @2
	%19 = shl i64 1, 30
	%20 = or i64 %18, %19
	store i64 %20, i64* @2
	%21 = load i64, i64* @2
	%22 = shl i64 1, 4
	%23 = xor i64 %22, -1
	%24 = and i64 %21, %23
	store i64 %24, i64* @2
	%25 = load i64, i64* @2
	%26 = getelementptr [4 x i8], [4 x i8]* @3, i64 0, i64 0
	%27 = call i64 (i8*, ...) @printf(i8* %26, i64 %25)
	ret i64 0
}

```
## Run
```bash
1075837966
2096158
1075837966
```
