# test/test_017.md
## Source
```pascal
MODULE Sets;

VAR x: SET;
    y: SET;
    z: SET;

BEGIN
    x := {1, 2, 3, 10};
    y := {10..19};
  
    print(x);
    print(y);
    print(x+y);
    print(x-y);
    print(x*y);
    print(x/y);

    INCL(x, 4);
    EXCL(y, 30);

    z := x + y;

    print(z);

    INCL(z, 30);
    EXCL(z, 4);

    print(z);

    IF 30 IN z THEN
        print(30)
    END;

    IF ~(31 IN z) THEN
        print(31)
    END
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
test/test_017.md:8:17:	comma	","	false	0	0.000000	(8, 17) -> (8, 18)
test/test_017.md:8:19:	integer	"10"	false	10	0.000000	(8, 19) -> (8, 21)
test/test_017.md:8:21:	rbrace	"}"	false	0	0.000000	(8, 21) -> (8, 22)
test/test_017.md:8:22:	semicolon	";"	false	0	0.000000	(8, 22) -> (8, 23)
test/test_017.md:9:4:	ident	"y"	false	0	0.000000	(9, 4) -> (9, 5)
test/test_017.md:9:6:	assign	":="	false	0	0.000000	(9, 6) -> (9, 8)
test/test_017.md:9:9:	lbrace	"{"	false	0	0.000000	(9, 9) -> (9, 10)
test/test_017.md:9:10:	integer	"10"	false	10	0.000000	(9, 10) -> (9, 14)
test/test_017.md:9:14:	dotdot	".."	false	0	0.000000	(9, 14) -> (9, 16)
test/test_017.md:9:16:	integer	"19"	false	19	0.000000	(9, 16) -> (9, 18)
test/test_017.md:9:18:	rbrace	"}"	false	0	0.000000	(9, 18) -> (9, 19)
test/test_017.md:9:19:	semicolon	";"	false	0	0.000000	(9, 19) -> (9, 20)
test/test_017.md:11:4:	ident	"print"	false	0	0.000000	(11, 4) -> (11, 9)
test/test_017.md:11:9:	lparen	"("	false	0	0.000000	(11, 9) -> (11, 10)
test/test_017.md:11:10:	ident	"x"	false	0	0.000000	(11, 10) -> (11, 11)
test/test_017.md:11:11:	rparen	")"	false	0	0.000000	(11, 11) -> (11, 12)
test/test_017.md:11:12:	semicolon	";"	false	0	0.000000	(11, 12) -> (11, 13)
test/test_017.md:12:4:	ident	"print"	false	0	0.000000	(12, 4) -> (12, 9)
test/test_017.md:12:9:	lparen	"("	false	0	0.000000	(12, 9) -> (12, 10)
test/test_017.md:12:10:	ident	"y"	false	0	0.000000	(12, 10) -> (12, 11)
test/test_017.md:12:11:	rparen	")"	false	0	0.000000	(12, 11) -> (12, 12)
test/test_017.md:12:12:	semicolon	";"	false	0	0.000000	(12, 12) -> (12, 13)
test/test_017.md:13:4:	ident	"print"	false	0	0.000000	(13, 4) -> (13, 9)
test/test_017.md:13:9:	lparen	"("	false	0	0.000000	(13, 9) -> (13, 10)
test/test_017.md:13:10:	ident	"x"	false	0	0.000000	(13, 10) -> (13, 11)
test/test_017.md:13:11:	plus	"+"	false	0	0.000000	(13, 11) -> (13, 12)
test/test_017.md:13:12:	ident	"y"	false	0	0.000000	(13, 12) -> (13, 13)
test/test_017.md:13:13:	rparen	")"	false	0	0.000000	(13, 13) -> (13, 14)
test/test_017.md:13:14:	semicolon	";"	false	0	0.000000	(13, 14) -> (13, 15)
test/test_017.md:14:4:	ident	"print"	false	0	0.000000	(14, 4) -> (14, 9)
test/test_017.md:14:9:	lparen	"("	false	0	0.000000	(14, 9) -> (14, 10)
test/test_017.md:14:10:	ident	"x"	false	0	0.000000	(14, 10) -> (14, 11)
test/test_017.md:14:11:	minus	"-"	false	0	0.000000	(14, 11) -> (14, 12)
test/test_017.md:14:12:	ident	"y"	false	0	0.000000	(14, 12) -> (14, 13)
test/test_017.md:14:13:	rparen	")"	false	0	0.000000	(14, 13) -> (14, 14)
test/test_017.md:14:14:	semicolon	";"	false	0	0.000000	(14, 14) -> (14, 15)
test/test_017.md:15:4:	ident	"print"	false	0	0.000000	(15, 4) -> (15, 9)
test/test_017.md:15:9:	lparen	"("	false	0	0.000000	(15, 9) -> (15, 10)
test/test_017.md:15:10:	ident	"x"	false	0	0.000000	(15, 10) -> (15, 11)
test/test_017.md:15:11:	asterisk	"*"	false	0	0.000000	(15, 11) -> (15, 12)
test/test_017.md:15:12:	ident	"y"	false	0	0.000000	(15, 12) -> (15, 13)
test/test_017.md:15:13:	rparen	")"	false	0	0.000000	(15, 13) -> (15, 14)
test/test_017.md:15:14:	semicolon	";"	false	0	0.000000	(15, 14) -> (15, 15)
test/test_017.md:16:4:	ident	"print"	false	0	0.000000	(16, 4) -> (16, 9)
test/test_017.md:16:9:	lparen	"("	false	0	0.000000	(16, 9) -> (16, 10)
test/test_017.md:16:10:	ident	"x"	false	0	0.000000	(16, 10) -> (16, 11)
test/test_017.md:16:11:	slash	"/"	false	0	0.000000	(16, 11) -> (16, 12)
test/test_017.md:16:12:	ident	"y"	false	0	0.000000	(16, 12) -> (16, 13)
test/test_017.md:16:13:	rparen	")"	false	0	0.000000	(16, 13) -> (16, 14)
test/test_017.md:16:14:	semicolon	";"	false	0	0.000000	(16, 14) -> (16, 15)
test/test_017.md:18:4:	ident	"INCL"	false	0	0.000000	(18, 4) -> (18, 8)
test/test_017.md:18:8:	lparen	"("	false	0	0.000000	(18, 8) -> (18, 9)
test/test_017.md:18:9:	ident	"x"	false	0	0.000000	(18, 9) -> (18, 10)
test/test_017.md:18:10:	comma	","	false	0	0.000000	(18, 10) -> (18, 11)
test/test_017.md:18:12:	integer	"4"	false	4	0.000000	(18, 12) -> (18, 13)
test/test_017.md:18:13:	rparen	")"	false	0	0.000000	(18, 13) -> (18, 14)
test/test_017.md:18:14:	semicolon	";"	false	0	0.000000	(18, 14) -> (18, 15)
test/test_017.md:19:4:	ident	"EXCL"	false	0	0.000000	(19, 4) -> (19, 8)
test/test_017.md:19:8:	lparen	"("	false	0	0.000000	(19, 8) -> (19, 9)
test/test_017.md:19:9:	ident	"y"	false	0	0.000000	(19, 9) -> (19, 10)
test/test_017.md:19:10:	comma	","	false	0	0.000000	(19, 10) -> (19, 11)
test/test_017.md:19:12:	integer	"30"	false	30	0.000000	(19, 12) -> (19, 14)
test/test_017.md:19:14:	rparen	")"	false	0	0.000000	(19, 14) -> (19, 15)
test/test_017.md:19:15:	semicolon	";"	false	0	0.000000	(19, 15) -> (19, 16)
test/test_017.md:21:4:	ident	"z"	false	0	0.000000	(21, 4) -> (21, 5)
test/test_017.md:21:6:	assign	":="	false	0	0.000000	(21, 6) -> (21, 8)
test/test_017.md:21:9:	ident	"x"	false	0	0.000000	(21, 9) -> (21, 10)
test/test_017.md:21:11:	plus	"+"	false	0	0.000000	(21, 11) -> (21, 12)
test/test_017.md:21:13:	ident	"y"	false	0	0.000000	(21, 13) -> (21, 14)
test/test_017.md:21:14:	semicolon	";"	false	0	0.000000	(21, 14) -> (21, 15)
test/test_017.md:23:4:	ident	"print"	false	0	0.000000	(23, 4) -> (23, 9)
test/test_017.md:23:9:	lparen	"("	false	0	0.000000	(23, 9) -> (23, 10)
test/test_017.md:23:10:	ident	"z"	false	0	0.000000	(23, 10) -> (23, 11)
test/test_017.md:23:11:	rparen	")"	false	0	0.000000	(23, 11) -> (23, 12)
test/test_017.md:23:12:	semicolon	";"	false	0	0.000000	(23, 12) -> (23, 13)
test/test_017.md:25:4:	ident	"INCL"	false	0	0.000000	(25, 4) -> (25, 8)
test/test_017.md:25:8:	lparen	"("	false	0	0.000000	(25, 8) -> (25, 9)
test/test_017.md:25:9:	ident	"z"	false	0	0.000000	(25, 9) -> (25, 10)
test/test_017.md:25:10:	comma	","	false	0	0.000000	(25, 10) -> (25, 11)
test/test_017.md:25:12:	integer	"30"	false	30	0.000000	(25, 12) -> (25, 14)
test/test_017.md:25:14:	rparen	")"	false	0	0.000000	(25, 14) -> (25, 15)
test/test_017.md:25:15:	semicolon	";"	false	0	0.000000	(25, 15) -> (25, 16)
test/test_017.md:26:4:	ident	"EXCL"	false	0	0.000000	(26, 4) -> (26, 8)
test/test_017.md:26:8:	lparen	"("	false	0	0.000000	(26, 8) -> (26, 9)
test/test_017.md:26:9:	ident	"z"	false	0	0.000000	(26, 9) -> (26, 10)
test/test_017.md:26:10:	comma	","	false	0	0.000000	(26, 10) -> (26, 11)
test/test_017.md:26:12:	integer	"4"	false	4	0.000000	(26, 12) -> (26, 13)
test/test_017.md:26:13:	rparen	")"	false	0	0.000000	(26, 13) -> (26, 14)
test/test_017.md:26:14:	semicolon	";"	false	0	0.000000	(26, 14) -> (26, 15)
test/test_017.md:28:4:	ident	"print"	false	0	0.000000	(28, 4) -> (28, 9)
test/test_017.md:28:9:	lparen	"("	false	0	0.000000	(28, 9) -> (28, 10)
test/test_017.md:28:10:	ident	"z"	false	0	0.000000	(28, 10) -> (28, 11)
test/test_017.md:28:11:	rparen	")"	false	0	0.000000	(28, 11) -> (28, 12)
test/test_017.md:28:12:	semicolon	";"	false	0	0.000000	(28, 12) -> (28, 13)
test/test_017.md:30:4:	if	"IF"	false	0	0.000000	(30, 4) -> (30, 6)
test/test_017.md:30:7:	integer	"30"	false	30	0.000000	(30, 7) -> (30, 9)
test/test_017.md:30:10:	in	"IN"	false	0	0.000000	(30, 10) -> (30, 12)
test/test_017.md:30:13:	ident	"z"	false	0	0.000000	(30, 13) -> (30, 14)
test/test_017.md:30:15:	then	"THEN"	false	0	0.000000	(30, 15) -> (30, 19)
test/test_017.md:31:8:	ident	"print"	false	0	0.000000	(31, 8) -> (31, 13)
test/test_017.md:31:13:	lparen	"("	false	0	0.000000	(31, 13) -> (31, 14)
test/test_017.md:31:14:	integer	"30"	false	30	0.000000	(31, 14) -> (31, 16)
test/test_017.md:31:16:	rparen	")"	false	0	0.000000	(31, 16) -> (31, 17)
test/test_017.md:32:4:	end	"END"	false	0	0.000000	(32, 4) -> (32, 7)
test/test_017.md:32:7:	semicolon	";"	false	0	0.000000	(32, 7) -> (32, 8)
test/test_017.md:34:4:	if	"IF"	false	0	0.000000	(34, 4) -> (34, 6)
test/test_017.md:34:7:	tilde	"~"	false	0	0.000000	(34, 7) -> (34, 8)
test/test_017.md:34:8:	lparen	"("	false	0	0.000000	(34, 8) -> (34, 9)
test/test_017.md:34:9:	integer	"31"	false	31	0.000000	(34, 9) -> (34, 11)
test/test_017.md:34:12:	in	"IN"	false	0	0.000000	(34, 12) -> (34, 14)
test/test_017.md:34:15:	ident	"z"	false	0	0.000000	(34, 15) -> (34, 16)
test/test_017.md:34:16:	rparen	")"	false	0	0.000000	(34, 16) -> (34, 17)
test/test_017.md:34:18:	then	"THEN"	false	0	0.000000	(34, 18) -> (34, 22)
test/test_017.md:35:8:	ident	"print"	false	0	0.000000	(35, 8) -> (35, 13)
test/test_017.md:35:13:	lparen	"("	false	0	0.000000	(35, 13) -> (35, 14)
test/test_017.md:35:14:	integer	"31"	false	31	0.000000	(35, 14) -> (35, 16)
test/test_017.md:35:16:	rparen	")"	false	0	0.000000	(35, 16) -> (35, 17)
test/test_017.md:36:4:	end	"END"	false	0	0.000000	(36, 4) -> (36, 7)
test/test_017.md:37:0:	end	"END"	false	0	0.000000	(37, 0) -> (37, 3)
test/test_017.md:37:4:	ident	"Sets"	false	0	0.000000	(37, 4) -> (37, 8)
test/test_017.md:37:8:	dot	"."	false	0	0.000000	(37, 8) -> (37, 9)
test/test_017.md:38:0:	eof	""	false	0	0.000000	(38, 0) -> (38, 0)
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
      (set (1..3, 10))
    )
    (assign y
      (set (10..19))
    )
    (expr2stmt
      (call "print" [void]
        (variable [set] "x")
      )
    )
    (expr2stmt
      (call "print" [void]
        (variable [set] "y")
      )
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
      (call "print" [void]
        (minus [set]
          (variable [set] "x")
          (variable [set] "y")
        )
      )
    )
    (expr2stmt
      (call "print" [void]
        (asterisk [set]
          (variable [set] "x")
          (variable [set] "y")
        )
      )
    )
    (expr2stmt
      (call "print" [void]
        (slash [set]
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
    (if
      (in [boolean]
        (number [i64] 30)
        (variable [set] "z")
      )
      (stmts
        (expr2stmt
          (call "print" [void]
            (number [i64] 30)
          )
        )
      )
      (stmts
      )
    )
    (if
      (not [boolean]
        (in [boolean]
          (number [i64] 31)
          (variable [set] "z")
        )
      )
      (stmts
        (expr2stmt
          (call "print" [void]
            (number [i64] 31)
          )
        )
      )
      (stmts
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
	store i64 1038, i64* @0
	store i64 1047552, i64* @1
	%0 = load i64, i64* @0
	%1 = getelementptr [4 x i8], [4 x i8]* @3, i64 0, i64 0
	%2 = call i64 (i8*, ...) @printf(i8* %1, i64 %0)
	%3 = load i64, i64* @1
	%4 = getelementptr [4 x i8], [4 x i8]* @3, i64 0, i64 0
	%5 = call i64 (i8*, ...) @printf(i8* %4, i64 %3)
	%6 = load i64, i64* @0
	%7 = load i64, i64* @1
	%8 = or i64 %6, %7
	%9 = getelementptr [4 x i8], [4 x i8]* @3, i64 0, i64 0
	%10 = call i64 (i8*, ...) @printf(i8* %9, i64 %8)
	%11 = load i64, i64* @0
	%12 = load i64, i64* @1
	%13 = xor i64 %12, -1
	%14 = and i64 %11, %13
	%15 = getelementptr [4 x i8], [4 x i8]* @3, i64 0, i64 0
	%16 = call i64 (i8*, ...) @printf(i8* %15, i64 %14)
	%17 = load i64, i64* @0
	%18 = load i64, i64* @1
	%19 = and i64 %17, %18
	%20 = getelementptr [4 x i8], [4 x i8]* @3, i64 0, i64 0
	%21 = call i64 (i8*, ...) @printf(i8* %20, i64 %19)
	%22 = load i64, i64* @0
	%23 = load i64, i64* @1
	%24 = xor i64 %22, %23
	%25 = getelementptr [4 x i8], [4 x i8]* @3, i64 0, i64 0
	%26 = call i64 (i8*, ...) @printf(i8* %25, i64 %24)
	%27 = load i64, i64* @0
	%28 = shl i64 1, 4
	%29 = or i64 %27, %28
	store i64 %29, i64* @0
	%30 = load i64, i64* @1
	%31 = shl i64 1, 30
	%32 = xor i64 %31, -1
	%33 = and i64 %30, %32
	store i64 %33, i64* @1
	%34 = load i64, i64* @0
	%35 = load i64, i64* @1
	%36 = or i64 %34, %35
	store i64 %36, i64* @2
	%37 = load i64, i64* @2
	%38 = getelementptr [4 x i8], [4 x i8]* @3, i64 0, i64 0
	%39 = call i64 (i8*, ...) @printf(i8* %38, i64 %37)
	%40 = load i64, i64* @2
	%41 = shl i64 1, 30
	%42 = or i64 %40, %41
	store i64 %42, i64* @2
	%43 = load i64, i64* @2
	%44 = shl i64 1, 4
	%45 = xor i64 %44, -1
	%46 = and i64 %43, %45
	store i64 %46, i64* @2
	%47 = load i64, i64* @2
	%48 = getelementptr [4 x i8], [4 x i8]* @3, i64 0, i64 0
	%49 = call i64 (i8*, ...) @printf(i8* %48, i64 %47)
	%50 = load i64, i64* @2
	%51 = shl i64 1, 30
	%52 = and i64 %50, %51
	%53 = icmp ne i64 %52, 0
	br i1 %53, label %54, label %57

54:
	%55 = getelementptr [4 x i8], [4 x i8]* @3, i64 0, i64 0
	%56 = call i64 (i8*, ...) @printf(i8* %55, i64 30)
	br label %58

57:
	br label %58

58:
	%59 = load i64, i64* @2
	%60 = shl i64 1, 31
	%61 = and i64 %59, %60
	%62 = icmp ne i64 %61, 0
	%63 = icmp eq i1 %62, false
	br i1 %63, label %64, label %67

64:
	%65 = getelementptr [4 x i8], [4 x i8]* @3, i64 0, i64 0
	%66 = call i64 (i8*, ...) @printf(i8* %65, i64 31)
	br label %68

67:
	br label %68

68:
	ret i64 0
}

```
## Run
```bash
1038
1047552
1047566
14
1024
1046542
1047582
1074789390
30
31
```
