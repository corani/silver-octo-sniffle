# test/test_017.md
## Source
```pascal
MODULE Sets;

CONST xc = {1, 2, 3, 10};
      yc = {10..19};

VAR x: SET;
    y: SET;
    z: SET;

BEGIN
    print(xc);
    print(yc);
    print(xc+yc);
    print(xc-yc);
    print(xc*yc);
    print(xc/yc);

    x := xc; y := yc;

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
test/test_017.md:3:0:	const	"CONST"	false	0	0.000000	(3, 0) -> (3, 5)
test/test_017.md:3:6:	ident	"xc"	false	0	0.000000	(3, 6) -> (3, 8)
test/test_017.md:3:9:	eq	"="	false	0	0.000000	(3, 9) -> (3, 10)
test/test_017.md:3:11:	lbrace	"{"	false	0	0.000000	(3, 11) -> (3, 12)
test/test_017.md:3:12:	integer	"1"	false	1	0.000000	(3, 12) -> (3, 13)
test/test_017.md:3:13:	comma	","	false	0	0.000000	(3, 13) -> (3, 14)
test/test_017.md:3:15:	integer	"2"	false	2	0.000000	(3, 15) -> (3, 16)
test/test_017.md:3:16:	comma	","	false	0	0.000000	(3, 16) -> (3, 17)
test/test_017.md:3:18:	integer	"3"	false	3	0.000000	(3, 18) -> (3, 19)
test/test_017.md:3:19:	comma	","	false	0	0.000000	(3, 19) -> (3, 20)
test/test_017.md:3:21:	integer	"10"	false	10	0.000000	(3, 21) -> (3, 23)
test/test_017.md:3:23:	rbrace	"}"	false	0	0.000000	(3, 23) -> (3, 24)
test/test_017.md:3:24:	semicolon	";"	false	0	0.000000	(3, 24) -> (3, 25)
test/test_017.md:4:6:	ident	"yc"	false	0	0.000000	(4, 6) -> (4, 8)
test/test_017.md:4:9:	eq	"="	false	0	0.000000	(4, 9) -> (4, 10)
test/test_017.md:4:11:	lbrace	"{"	false	0	0.000000	(4, 11) -> (4, 12)
test/test_017.md:4:12:	integer	"10"	false	10	0.000000	(4, 12) -> (4, 16)
test/test_017.md:4:16:	dotdot	".."	false	0	0.000000	(4, 16) -> (4, 18)
test/test_017.md:4:18:	integer	"19"	false	19	0.000000	(4, 18) -> (4, 20)
test/test_017.md:4:20:	rbrace	"}"	false	0	0.000000	(4, 20) -> (4, 21)
test/test_017.md:4:21:	semicolon	";"	false	0	0.000000	(4, 21) -> (4, 22)
test/test_017.md:6:0:	var	"VAR"	false	0	0.000000	(6, 0) -> (6, 3)
test/test_017.md:6:4:	ident	"x"	false	0	0.000000	(6, 4) -> (6, 5)
test/test_017.md:6:5:	colon	":"	false	0	0.000000	(6, 5) -> (6, 6)
test/test_017.md:6:7:	ident	"SET"	false	0	0.000000	(6, 7) -> (6, 10)
test/test_017.md:6:10:	semicolon	";"	false	0	0.000000	(6, 10) -> (6, 11)
test/test_017.md:7:4:	ident	"y"	false	0	0.000000	(7, 4) -> (7, 5)
test/test_017.md:7:5:	colon	":"	false	0	0.000000	(7, 5) -> (7, 6)
test/test_017.md:7:7:	ident	"SET"	false	0	0.000000	(7, 7) -> (7, 10)
test/test_017.md:7:10:	semicolon	";"	false	0	0.000000	(7, 10) -> (7, 11)
test/test_017.md:8:4:	ident	"z"	false	0	0.000000	(8, 4) -> (8, 5)
test/test_017.md:8:5:	colon	":"	false	0	0.000000	(8, 5) -> (8, 6)
test/test_017.md:8:7:	ident	"SET"	false	0	0.000000	(8, 7) -> (8, 10)
test/test_017.md:8:10:	semicolon	";"	false	0	0.000000	(8, 10) -> (8, 11)
test/test_017.md:10:0:	begin	"BEGIN"	false	0	0.000000	(10, 0) -> (10, 5)
test/test_017.md:11:4:	ident	"print"	false	0	0.000000	(11, 4) -> (11, 9)
test/test_017.md:11:9:	lparen	"("	false	0	0.000000	(11, 9) -> (11, 10)
test/test_017.md:11:10:	ident	"xc"	false	0	0.000000	(11, 10) -> (11, 12)
test/test_017.md:11:12:	rparen	")"	false	0	0.000000	(11, 12) -> (11, 13)
test/test_017.md:11:13:	semicolon	";"	false	0	0.000000	(11, 13) -> (11, 14)
test/test_017.md:12:4:	ident	"print"	false	0	0.000000	(12, 4) -> (12, 9)
test/test_017.md:12:9:	lparen	"("	false	0	0.000000	(12, 9) -> (12, 10)
test/test_017.md:12:10:	ident	"yc"	false	0	0.000000	(12, 10) -> (12, 12)
test/test_017.md:12:12:	rparen	")"	false	0	0.000000	(12, 12) -> (12, 13)
test/test_017.md:12:13:	semicolon	";"	false	0	0.000000	(12, 13) -> (12, 14)
test/test_017.md:13:4:	ident	"print"	false	0	0.000000	(13, 4) -> (13, 9)
test/test_017.md:13:9:	lparen	"("	false	0	0.000000	(13, 9) -> (13, 10)
test/test_017.md:13:10:	ident	"xc"	false	0	0.000000	(13, 10) -> (13, 12)
test/test_017.md:13:12:	plus	"+"	false	0	0.000000	(13, 12) -> (13, 13)
test/test_017.md:13:13:	ident	"yc"	false	0	0.000000	(13, 13) -> (13, 15)
test/test_017.md:13:15:	rparen	")"	false	0	0.000000	(13, 15) -> (13, 16)
test/test_017.md:13:16:	semicolon	";"	false	0	0.000000	(13, 16) -> (13, 17)
test/test_017.md:14:4:	ident	"print"	false	0	0.000000	(14, 4) -> (14, 9)
test/test_017.md:14:9:	lparen	"("	false	0	0.000000	(14, 9) -> (14, 10)
test/test_017.md:14:10:	ident	"xc"	false	0	0.000000	(14, 10) -> (14, 12)
test/test_017.md:14:12:	minus	"-"	false	0	0.000000	(14, 12) -> (14, 13)
test/test_017.md:14:13:	ident	"yc"	false	0	0.000000	(14, 13) -> (14, 15)
test/test_017.md:14:15:	rparen	")"	false	0	0.000000	(14, 15) -> (14, 16)
test/test_017.md:14:16:	semicolon	";"	false	0	0.000000	(14, 16) -> (14, 17)
test/test_017.md:15:4:	ident	"print"	false	0	0.000000	(15, 4) -> (15, 9)
test/test_017.md:15:9:	lparen	"("	false	0	0.000000	(15, 9) -> (15, 10)
test/test_017.md:15:10:	ident	"xc"	false	0	0.000000	(15, 10) -> (15, 12)
test/test_017.md:15:12:	asterisk	"*"	false	0	0.000000	(15, 12) -> (15, 13)
test/test_017.md:15:13:	ident	"yc"	false	0	0.000000	(15, 13) -> (15, 15)
test/test_017.md:15:15:	rparen	")"	false	0	0.000000	(15, 15) -> (15, 16)
test/test_017.md:15:16:	semicolon	";"	false	0	0.000000	(15, 16) -> (15, 17)
test/test_017.md:16:4:	ident	"print"	false	0	0.000000	(16, 4) -> (16, 9)
test/test_017.md:16:9:	lparen	"("	false	0	0.000000	(16, 9) -> (16, 10)
test/test_017.md:16:10:	ident	"xc"	false	0	0.000000	(16, 10) -> (16, 12)
test/test_017.md:16:12:	slash	"/"	false	0	0.000000	(16, 12) -> (16, 13)
test/test_017.md:16:13:	ident	"yc"	false	0	0.000000	(16, 13) -> (16, 15)
test/test_017.md:16:15:	rparen	")"	false	0	0.000000	(16, 15) -> (16, 16)
test/test_017.md:16:16:	semicolon	";"	false	0	0.000000	(16, 16) -> (16, 17)
test/test_017.md:18:4:	ident	"x"	false	0	0.000000	(18, 4) -> (18, 5)
test/test_017.md:18:6:	assign	":="	false	0	0.000000	(18, 6) -> (18, 8)
test/test_017.md:18:9:	ident	"xc"	false	0	0.000000	(18, 9) -> (18, 11)
test/test_017.md:18:11:	semicolon	";"	false	0	0.000000	(18, 11) -> (18, 12)
test/test_017.md:18:13:	ident	"y"	false	0	0.000000	(18, 13) -> (18, 14)
test/test_017.md:18:15:	assign	":="	false	0	0.000000	(18, 15) -> (18, 17)
test/test_017.md:18:18:	ident	"yc"	false	0	0.000000	(18, 18) -> (18, 20)
test/test_017.md:18:20:	semicolon	";"	false	0	0.000000	(18, 20) -> (18, 21)
test/test_017.md:20:4:	ident	"INCL"	false	0	0.000000	(20, 4) -> (20, 8)
test/test_017.md:20:8:	lparen	"("	false	0	0.000000	(20, 8) -> (20, 9)
test/test_017.md:20:9:	ident	"x"	false	0	0.000000	(20, 9) -> (20, 10)
test/test_017.md:20:10:	comma	","	false	0	0.000000	(20, 10) -> (20, 11)
test/test_017.md:20:12:	integer	"4"	false	4	0.000000	(20, 12) -> (20, 13)
test/test_017.md:20:13:	rparen	")"	false	0	0.000000	(20, 13) -> (20, 14)
test/test_017.md:20:14:	semicolon	";"	false	0	0.000000	(20, 14) -> (20, 15)
test/test_017.md:21:4:	ident	"EXCL"	false	0	0.000000	(21, 4) -> (21, 8)
test/test_017.md:21:8:	lparen	"("	false	0	0.000000	(21, 8) -> (21, 9)
test/test_017.md:21:9:	ident	"y"	false	0	0.000000	(21, 9) -> (21, 10)
test/test_017.md:21:10:	comma	","	false	0	0.000000	(21, 10) -> (21, 11)
test/test_017.md:21:12:	integer	"30"	false	30	0.000000	(21, 12) -> (21, 14)
test/test_017.md:21:14:	rparen	")"	false	0	0.000000	(21, 14) -> (21, 15)
test/test_017.md:21:15:	semicolon	";"	false	0	0.000000	(21, 15) -> (21, 16)
test/test_017.md:23:4:	ident	"z"	false	0	0.000000	(23, 4) -> (23, 5)
test/test_017.md:23:6:	assign	":="	false	0	0.000000	(23, 6) -> (23, 8)
test/test_017.md:23:9:	ident	"x"	false	0	0.000000	(23, 9) -> (23, 10)
test/test_017.md:23:11:	plus	"+"	false	0	0.000000	(23, 11) -> (23, 12)
test/test_017.md:23:13:	ident	"y"	false	0	0.000000	(23, 13) -> (23, 14)
test/test_017.md:23:14:	semicolon	";"	false	0	0.000000	(23, 14) -> (23, 15)
test/test_017.md:25:4:	ident	"print"	false	0	0.000000	(25, 4) -> (25, 9)
test/test_017.md:25:9:	lparen	"("	false	0	0.000000	(25, 9) -> (25, 10)
test/test_017.md:25:10:	ident	"z"	false	0	0.000000	(25, 10) -> (25, 11)
test/test_017.md:25:11:	rparen	")"	false	0	0.000000	(25, 11) -> (25, 12)
test/test_017.md:25:12:	semicolon	";"	false	0	0.000000	(25, 12) -> (25, 13)
test/test_017.md:27:4:	ident	"INCL"	false	0	0.000000	(27, 4) -> (27, 8)
test/test_017.md:27:8:	lparen	"("	false	0	0.000000	(27, 8) -> (27, 9)
test/test_017.md:27:9:	ident	"z"	false	0	0.000000	(27, 9) -> (27, 10)
test/test_017.md:27:10:	comma	","	false	0	0.000000	(27, 10) -> (27, 11)
test/test_017.md:27:12:	integer	"30"	false	30	0.000000	(27, 12) -> (27, 14)
test/test_017.md:27:14:	rparen	")"	false	0	0.000000	(27, 14) -> (27, 15)
test/test_017.md:27:15:	semicolon	";"	false	0	0.000000	(27, 15) -> (27, 16)
test/test_017.md:28:4:	ident	"EXCL"	false	0	0.000000	(28, 4) -> (28, 8)
test/test_017.md:28:8:	lparen	"("	false	0	0.000000	(28, 8) -> (28, 9)
test/test_017.md:28:9:	ident	"z"	false	0	0.000000	(28, 9) -> (28, 10)
test/test_017.md:28:10:	comma	","	false	0	0.000000	(28, 10) -> (28, 11)
test/test_017.md:28:12:	integer	"4"	false	4	0.000000	(28, 12) -> (28, 13)
test/test_017.md:28:13:	rparen	")"	false	0	0.000000	(28, 13) -> (28, 14)
test/test_017.md:28:14:	semicolon	";"	false	0	0.000000	(28, 14) -> (28, 15)
test/test_017.md:30:4:	ident	"print"	false	0	0.000000	(30, 4) -> (30, 9)
test/test_017.md:30:9:	lparen	"("	false	0	0.000000	(30, 9) -> (30, 10)
test/test_017.md:30:10:	ident	"z"	false	0	0.000000	(30, 10) -> (30, 11)
test/test_017.md:30:11:	rparen	")"	false	0	0.000000	(30, 11) -> (30, 12)
test/test_017.md:30:12:	semicolon	";"	false	0	0.000000	(30, 12) -> (30, 13)
test/test_017.md:32:4:	if	"IF"	false	0	0.000000	(32, 4) -> (32, 6)
test/test_017.md:32:7:	integer	"30"	false	30	0.000000	(32, 7) -> (32, 9)
test/test_017.md:32:10:	in	"IN"	false	0	0.000000	(32, 10) -> (32, 12)
test/test_017.md:32:13:	ident	"z"	false	0	0.000000	(32, 13) -> (32, 14)
test/test_017.md:32:15:	then	"THEN"	false	0	0.000000	(32, 15) -> (32, 19)
test/test_017.md:33:8:	ident	"print"	false	0	0.000000	(33, 8) -> (33, 13)
test/test_017.md:33:13:	lparen	"("	false	0	0.000000	(33, 13) -> (33, 14)
test/test_017.md:33:14:	integer	"30"	false	30	0.000000	(33, 14) -> (33, 16)
test/test_017.md:33:16:	rparen	")"	false	0	0.000000	(33, 16) -> (33, 17)
test/test_017.md:34:4:	end	"END"	false	0	0.000000	(34, 4) -> (34, 7)
test/test_017.md:34:7:	semicolon	";"	false	0	0.000000	(34, 7) -> (34, 8)
test/test_017.md:36:4:	if	"IF"	false	0	0.000000	(36, 4) -> (36, 6)
test/test_017.md:36:7:	tilde	"~"	false	0	0.000000	(36, 7) -> (36, 8)
test/test_017.md:36:8:	lparen	"("	false	0	0.000000	(36, 8) -> (36, 9)
test/test_017.md:36:9:	integer	"31"	false	31	0.000000	(36, 9) -> (36, 11)
test/test_017.md:36:12:	in	"IN"	false	0	0.000000	(36, 12) -> (36, 14)
test/test_017.md:36:15:	ident	"z"	false	0	0.000000	(36, 15) -> (36, 16)
test/test_017.md:36:16:	rparen	")"	false	0	0.000000	(36, 16) -> (36, 17)
test/test_017.md:36:18:	then	"THEN"	false	0	0.000000	(36, 18) -> (36, 22)
test/test_017.md:37:8:	ident	"print"	false	0	0.000000	(37, 8) -> (37, 13)
test/test_017.md:37:13:	lparen	"("	false	0	0.000000	(37, 13) -> (37, 14)
test/test_017.md:37:14:	integer	"31"	false	31	0.000000	(37, 14) -> (37, 16)
test/test_017.md:37:16:	rparen	")"	false	0	0.000000	(37, 16) -> (37, 17)
test/test_017.md:38:4:	end	"END"	false	0	0.000000	(38, 4) -> (38, 7)
test/test_017.md:39:0:	end	"END"	false	0	0.000000	(39, 0) -> (39, 3)
test/test_017.md:39:4:	ident	"Sets"	false	0	0.000000	(39, 4) -> (39, 8)
test/test_017.md:39:8:	dot	"."	false	0	0.000000	(39, 8) -> (39, 9)
test/test_017.md:40:0:	eof	""	false	0	0.000000	(40, 0) -> (40, 0)
```
## AST
```scheme
(module "Sets"
  (consts
    (xc [set]
      (set (1..3, 10))
    )
    (yc [set]
      (set (10..19))
    )
  )
  (vars
    (x [set])
    (y [set])
    (z [set])
  )
  (stmts
    (expr2stmt
      (call "print" [void]
        (constant [set] "xc")
      )
    )
    (expr2stmt
      (call "print" [void]
        (constant [set] "yc")
      )
    )
    (expr2stmt
      (call "print" [void]
        (plus [set]
          (constant [set] "xc")
          (constant [set] "yc")
        )
      )
    )
    (expr2stmt
      (call "print" [void]
        (minus [set]
          (constant [set] "xc")
          (constant [set] "yc")
        )
      )
    )
    (expr2stmt
      (call "print" [void]
        (asterisk [set]
          (constant [set] "xc")
          (constant [set] "yc")
        )
      )
    )
    (expr2stmt
      (call "print" [void]
        (slash [set]
          (constant [set] "xc")
          (constant [set] "yc")
        )
      )
    )
    (assign x
      (constant [set] "xc")
    )
    (assign y
      (constant [set] "yc")
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
	%0 = getelementptr [4 x i8], [4 x i8]* @3, i64 0, i64 0
	%1 = call i64 (i8*, ...) @printf(i8* %0, i64 1038)
	%2 = getelementptr [4 x i8], [4 x i8]* @3, i64 0, i64 0
	%3 = call i64 (i8*, ...) @printf(i8* %2, i64 1047552)
	%4 = or i64 1038, 1047552
	%5 = getelementptr [4 x i8], [4 x i8]* @3, i64 0, i64 0
	%6 = call i64 (i8*, ...) @printf(i8* %5, i64 %4)
	%7 = xor i64 1047552, -1
	%8 = and i64 1038, %7
	%9 = getelementptr [4 x i8], [4 x i8]* @3, i64 0, i64 0
	%10 = call i64 (i8*, ...) @printf(i8* %9, i64 %8)
	%11 = and i64 1038, 1047552
	%12 = getelementptr [4 x i8], [4 x i8]* @3, i64 0, i64 0
	%13 = call i64 (i8*, ...) @printf(i8* %12, i64 %11)
	%14 = xor i64 1038, 1047552
	%15 = getelementptr [4 x i8], [4 x i8]* @3, i64 0, i64 0
	%16 = call i64 (i8*, ...) @printf(i8* %15, i64 %14)
	store i64 1038, i64* @0
	store i64 1047552, i64* @1
	%17 = load i64, i64* @0
	%18 = shl i64 1, 4
	%19 = or i64 %17, %18
	store i64 %19, i64* @0
	%20 = load i64, i64* @1
	%21 = shl i64 1, 30
	%22 = xor i64 %21, -1
	%23 = and i64 %20, %22
	store i64 %23, i64* @1
	%24 = load i64, i64* @0
	%25 = load i64, i64* @1
	%26 = or i64 %24, %25
	store i64 %26, i64* @2
	%27 = load i64, i64* @2
	%28 = getelementptr [4 x i8], [4 x i8]* @3, i64 0, i64 0
	%29 = call i64 (i8*, ...) @printf(i8* %28, i64 %27)
	%30 = load i64, i64* @2
	%31 = shl i64 1, 30
	%32 = or i64 %30, %31
	store i64 %32, i64* @2
	%33 = load i64, i64* @2
	%34 = shl i64 1, 4
	%35 = xor i64 %34, -1
	%36 = and i64 %33, %35
	store i64 %36, i64* @2
	%37 = load i64, i64* @2
	%38 = getelementptr [4 x i8], [4 x i8]* @3, i64 0, i64 0
	%39 = call i64 (i8*, ...) @printf(i8* %38, i64 %37)
	%40 = load i64, i64* @2
	%41 = shl i64 1, 30
	%42 = and i64 %40, %41
	%43 = icmp ne i64 %42, 0
	br i1 %43, label %44, label %47

44:
	%45 = getelementptr [4 x i8], [4 x i8]* @3, i64 0, i64 0
	%46 = call i64 (i8*, ...) @printf(i8* %45, i64 30)
	br label %48

47:
	br label %48

48:
	%49 = load i64, i64* @2
	%50 = shl i64 1, 31
	%51 = and i64 %49, %50
	%52 = icmp ne i64 %51, 0
	%53 = icmp eq i1 %52, false
	br i1 %53, label %54, label %57

54:
	%55 = getelementptr [4 x i8], [4 x i8]* @3, i64 0, i64 0
	%56 = call i64 (i8*, ...) @printf(i8* %55, i64 31)
	br label %58

57:
	br label %58

58:
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
