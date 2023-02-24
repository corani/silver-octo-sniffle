# test/test_016.md
## Source
```pascal
MODULE Chars;

VAR x: CHAR;
    y: INTEGER;

BEGIN
    C.print("A");

    x := 42X;
    C.print(x);

    y := ORD(x);
    C.print(y);

    INC(y);
    C.print(CHR(y));

    C.print(x <= "B");

    x := "A";

    REPEAT
      C.print(x);
      x := CHR(ORD(x)+1)
    UNTIL x > "Z"
END Chars.
```
## Tokens
```tsv
test/test_016.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_016.md:1:8:	ident	"Chars"	false	0	0.000000	(1, 8) -> (1, 13)
test/test_016.md:1:13:	semicolon	";"	false	0	0.000000	(1, 13) -> (1, 14)
test/test_016.md:3:0:	var	"VAR"	false	0	0.000000	(3, 0) -> (3, 3)
test/test_016.md:3:4:	ident	"x"	false	0	0.000000	(3, 4) -> (3, 5)
test/test_016.md:3:5:	colon	":"	false	0	0.000000	(3, 5) -> (3, 6)
test/test_016.md:3:7:	ident	"CHAR"	false	0	0.000000	(3, 7) -> (3, 11)
test/test_016.md:3:11:	semicolon	";"	false	0	0.000000	(3, 11) -> (3, 12)
test/test_016.md:4:4:	ident	"y"	false	0	0.000000	(4, 4) -> (4, 5)
test/test_016.md:4:5:	colon	":"	false	0	0.000000	(4, 5) -> (4, 6)
test/test_016.md:4:7:	ident	"INTEGER"	false	0	0.000000	(4, 7) -> (4, 14)
test/test_016.md:4:14:	semicolon	";"	false	0	0.000000	(4, 14) -> (4, 15)
test/test_016.md:6:0:	begin	"BEGIN"	false	0	0.000000	(6, 0) -> (6, 5)
test/test_016.md:7:4:	ident	"C"	false	0	0.000000	(7, 4) -> (7, 5)
test/test_016.md:7:5:	dot	"."	false	0	0.000000	(7, 5) -> (7, 6)
test/test_016.md:7:6:	ident	"print"	false	0	0.000000	(7, 6) -> (7, 11)
test/test_016.md:7:11:	lparen	"("	false	0	0.000000	(7, 11) -> (7, 12)
test/test_016.md:7:12:	string	"A"	false	0	0.000000	(7, 12) -> (7, 15)
test/test_016.md:7:15:	rparen	")"	false	0	0.000000	(7, 15) -> (7, 16)
test/test_016.md:7:16:	semicolon	";"	false	0	0.000000	(7, 16) -> (7, 17)
test/test_016.md:9:4:	ident	"x"	false	0	0.000000	(9, 4) -> (9, 5)
test/test_016.md:9:6:	assign	":="	false	0	0.000000	(9, 6) -> (9, 8)
test/test_016.md:9:9:	string	"B"	false	0	0.000000	(9, 9) -> (9, 12)
test/test_016.md:9:12:	semicolon	";"	false	0	0.000000	(9, 12) -> (9, 13)
test/test_016.md:10:4:	ident	"C"	false	0	0.000000	(10, 4) -> (10, 5)
test/test_016.md:10:5:	dot	"."	false	0	0.000000	(10, 5) -> (10, 6)
test/test_016.md:10:6:	ident	"print"	false	0	0.000000	(10, 6) -> (10, 11)
test/test_016.md:10:11:	lparen	"("	false	0	0.000000	(10, 11) -> (10, 12)
test/test_016.md:10:12:	ident	"x"	false	0	0.000000	(10, 12) -> (10, 13)
test/test_016.md:10:13:	rparen	")"	false	0	0.000000	(10, 13) -> (10, 14)
test/test_016.md:10:14:	semicolon	";"	false	0	0.000000	(10, 14) -> (10, 15)
test/test_016.md:12:4:	ident	"y"	false	0	0.000000	(12, 4) -> (12, 5)
test/test_016.md:12:6:	assign	":="	false	0	0.000000	(12, 6) -> (12, 8)
test/test_016.md:12:9:	ident	"ORD"	false	0	0.000000	(12, 9) -> (12, 12)
test/test_016.md:12:12:	lparen	"("	false	0	0.000000	(12, 12) -> (12, 13)
test/test_016.md:12:13:	ident	"x"	false	0	0.000000	(12, 13) -> (12, 14)
test/test_016.md:12:14:	rparen	")"	false	0	0.000000	(12, 14) -> (12, 15)
test/test_016.md:12:15:	semicolon	";"	false	0	0.000000	(12, 15) -> (12, 16)
test/test_016.md:13:4:	ident	"C"	false	0	0.000000	(13, 4) -> (13, 5)
test/test_016.md:13:5:	dot	"."	false	0	0.000000	(13, 5) -> (13, 6)
test/test_016.md:13:6:	ident	"print"	false	0	0.000000	(13, 6) -> (13, 11)
test/test_016.md:13:11:	lparen	"("	false	0	0.000000	(13, 11) -> (13, 12)
test/test_016.md:13:12:	ident	"y"	false	0	0.000000	(13, 12) -> (13, 13)
test/test_016.md:13:13:	rparen	")"	false	0	0.000000	(13, 13) -> (13, 14)
test/test_016.md:13:14:	semicolon	";"	false	0	0.000000	(13, 14) -> (13, 15)
test/test_016.md:15:4:	ident	"INC"	false	0	0.000000	(15, 4) -> (15, 7)
test/test_016.md:15:7:	lparen	"("	false	0	0.000000	(15, 7) -> (15, 8)
test/test_016.md:15:8:	ident	"y"	false	0	0.000000	(15, 8) -> (15, 9)
test/test_016.md:15:9:	rparen	")"	false	0	0.000000	(15, 9) -> (15, 10)
test/test_016.md:15:10:	semicolon	";"	false	0	0.000000	(15, 10) -> (15, 11)
test/test_016.md:16:4:	ident	"C"	false	0	0.000000	(16, 4) -> (16, 5)
test/test_016.md:16:5:	dot	"."	false	0	0.000000	(16, 5) -> (16, 6)
test/test_016.md:16:6:	ident	"print"	false	0	0.000000	(16, 6) -> (16, 11)
test/test_016.md:16:11:	lparen	"("	false	0	0.000000	(16, 11) -> (16, 12)
test/test_016.md:16:12:	ident	"CHR"	false	0	0.000000	(16, 12) -> (16, 15)
test/test_016.md:16:15:	lparen	"("	false	0	0.000000	(16, 15) -> (16, 16)
test/test_016.md:16:16:	ident	"y"	false	0	0.000000	(16, 16) -> (16, 17)
test/test_016.md:16:17:	rparen	")"	false	0	0.000000	(16, 17) -> (16, 18)
test/test_016.md:16:18:	rparen	")"	false	0	0.000000	(16, 18) -> (16, 19)
test/test_016.md:16:19:	semicolon	";"	false	0	0.000000	(16, 19) -> (16, 20)
test/test_016.md:18:4:	ident	"C"	false	0	0.000000	(18, 4) -> (18, 5)
test/test_016.md:18:5:	dot	"."	false	0	0.000000	(18, 5) -> (18, 6)
test/test_016.md:18:6:	ident	"print"	false	0	0.000000	(18, 6) -> (18, 11)
test/test_016.md:18:11:	lparen	"("	false	0	0.000000	(18, 11) -> (18, 12)
test/test_016.md:18:12:	ident	"x"	false	0	0.000000	(18, 12) -> (18, 13)
test/test_016.md:18:14:	le	"<="	false	0	0.000000	(18, 14) -> (18, 16)
test/test_016.md:18:17:	string	"B"	false	0	0.000000	(18, 17) -> (18, 20)
test/test_016.md:18:20:	rparen	")"	false	0	0.000000	(18, 20) -> (18, 21)
test/test_016.md:18:21:	semicolon	";"	false	0	0.000000	(18, 21) -> (18, 22)
test/test_016.md:20:4:	ident	"x"	false	0	0.000000	(20, 4) -> (20, 5)
test/test_016.md:20:6:	assign	":="	false	0	0.000000	(20, 6) -> (20, 8)
test/test_016.md:20:9:	string	"A"	false	0	0.000000	(20, 9) -> (20, 12)
test/test_016.md:20:12:	semicolon	";"	false	0	0.000000	(20, 12) -> (20, 13)
test/test_016.md:22:4:	repeat	"REPEAT"	false	0	0.000000	(22, 4) -> (22, 10)
test/test_016.md:23:6:	ident	"C"	false	0	0.000000	(23, 6) -> (23, 7)
test/test_016.md:23:7:	dot	"."	false	0	0.000000	(23, 7) -> (23, 8)
test/test_016.md:23:8:	ident	"print"	false	0	0.000000	(23, 8) -> (23, 13)
test/test_016.md:23:13:	lparen	"("	false	0	0.000000	(23, 13) -> (23, 14)
test/test_016.md:23:14:	ident	"x"	false	0	0.000000	(23, 14) -> (23, 15)
test/test_016.md:23:15:	rparen	")"	false	0	0.000000	(23, 15) -> (23, 16)
test/test_016.md:23:16:	semicolon	";"	false	0	0.000000	(23, 16) -> (23, 17)
test/test_016.md:24:6:	ident	"x"	false	0	0.000000	(24, 6) -> (24, 7)
test/test_016.md:24:8:	assign	":="	false	0	0.000000	(24, 8) -> (24, 10)
test/test_016.md:24:11:	ident	"CHR"	false	0	0.000000	(24, 11) -> (24, 14)
test/test_016.md:24:14:	lparen	"("	false	0	0.000000	(24, 14) -> (24, 15)
test/test_016.md:24:15:	ident	"ORD"	false	0	0.000000	(24, 15) -> (24, 18)
test/test_016.md:24:18:	lparen	"("	false	0	0.000000	(24, 18) -> (24, 19)
test/test_016.md:24:19:	ident	"x"	false	0	0.000000	(24, 19) -> (24, 20)
test/test_016.md:24:20:	rparen	")"	false	0	0.000000	(24, 20) -> (24, 21)
test/test_016.md:24:21:	plus	"+"	false	0	0.000000	(24, 21) -> (24, 22)
test/test_016.md:24:22:	integer	"1"	false	1	0.000000	(24, 22) -> (24, 23)
test/test_016.md:24:23:	rparen	")"	false	0	0.000000	(24, 23) -> (24, 24)
test/test_016.md:25:4:	until	"UNTIL"	false	0	0.000000	(25, 4) -> (25, 9)
test/test_016.md:25:10:	ident	"x"	false	0	0.000000	(25, 10) -> (25, 11)
test/test_016.md:25:12:	gt	">"	false	0	0.000000	(25, 12) -> (25, 13)
test/test_016.md:25:14:	string	"Z"	false	0	0.000000	(25, 14) -> (25, 17)
test/test_016.md:26:0:	end	"END"	false	0	0.000000	(26, 0) -> (26, 3)
test/test_016.md:26:4:	ident	"Chars"	false	0	0.000000	(26, 4) -> (26, 9)
test/test_016.md:26:9:	dot	"."	false	0	0.000000	(26, 9) -> (26, 10)
test/test_016.md:27:0:	eof	""	false	0	0.000000	(27, 0) -> (27, 0)
```
## AST
```scheme
(module "Chars"
  (vars
    (x 
      (CHAR [char])
    )
    (y 
      (INTEGER [i64])
    )
  )
  (stmts
    (expr2stmt
      (call
        (procedure [void] "C.print")
        (char "A")
      )
    )
    (assign
      (variable [char] "x")
      (char "B")
    )
    (expr2stmt
      (call
        (procedure [void] "C.print")
        (variable [char] "x")
      )
    )
    (assign
      (variable [i64] "y")
      (call
        (procedure [i64] "ORD")
        (variable [char] "x")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "C.print")
        (variable [i64] "y")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "INC")
        (variable [i64] "y")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "C.print")
        (call
          (procedure [char] "CHR")
          (variable [i64] "y")
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "C.print")
        (le [boolean]
          (variable [char] "x")
          (char "B")
        )
      )
    )
    (assign
      (variable [char] "x")
      (char "A")
    )
    (repeat
      (stmts
        (expr2stmt
          (call
            (procedure [void] "C.print")
            (variable [char] "x")
          )
        )
        (assign
          (variable [char] "x")
          (call
            (procedure [char] "CHR")
            (plus [i64]
              (call
                (procedure [i64] "ORD")
                (variable [char] "x")
              )
              (number [i64] 1)
            )
          )
        )
      )
      (gt [boolean]
        (variable [char] "x")
        (char "Z")
      )
    )
  )
)
```
## IR
```llvm
@0 = global i8 0
@1 = global i64 0
@2 = global [4 x i8] c"%c\0A\00"
@3 = global [4 x i8] c"%d\0A\00"
@4 = global [5 x i8] c"TRUE\00"
@5 = global [6 x i8] c"FALSE\00"

declare i64 @puts(i8* %str)

declare i64 @rand()

declare i64 @sprintf(i8* %buf, i8* %format, ...)

declare i64 @printf(i8* %format, ...)

declare i8* @malloc(i64 %size)

declare i8* @free(i8* %ptr)

define i64 @main() {
entry:
	%0 = getelementptr [4 x i8], [4 x i8]* @2, i64 0, i64 0
	%1 = call i64 (i8*, ...) @printf(i8* %0, i8 65)
	store i8 66, i8* @0
	%2 = load i8, i8* @0
	%3 = getelementptr [4 x i8], [4 x i8]* @2, i64 0, i64 0
	%4 = call i64 (i8*, ...) @printf(i8* %3, i8 %2)
	%5 = load i8, i8* @0
	%6 = zext i8 %5 to i64
	store i64 %6, i64* @1
	%7 = load i64, i64* @1
	%8 = getelementptr [4 x i8], [4 x i8]* @3, i64 0, i64 0
	%9 = call i64 (i8*, ...) @printf(i8* %8, i64 %7)
	%10 = load i64, i64* @1
	%11 = add i64 %10, 1
	store i64 %11, i64* @1
	%12 = load i64, i64* @1
	%13 = trunc i64 %12 to i8
	%14 = getelementptr [4 x i8], [4 x i8]* @2, i64 0, i64 0
	%15 = call i64 (i8*, ...) @printf(i8* %14, i8 %13)
	%16 = load i8, i8* @0
	%17 = icmp sle i8 %16, 66
	br i1 %17, label %18, label %20

18:
	%19 = getelementptr [5 x i8], [5 x i8]* @4, i64 0, i64 0
	br label %22

20:
	%21 = getelementptr [6 x i8], [6 x i8]* @5, i64 0, i64 0
	br label %22

22:
	%23 = phi i8* [ %19, %18 ], [ %21, %20 ]
	%24 = call i64 @puts(i8* %23)
	store i8 65, i8* @0
	br label %25

25:
	%26 = load i8, i8* @0
	%27 = getelementptr [4 x i8], [4 x i8]* @2, i64 0, i64 0
	%28 = call i64 (i8*, ...) @printf(i8* %27, i8 %26)
	%29 = load i8, i8* @0
	%30 = zext i8 %29 to i64
	%31 = add i64 %30, 1
	%32 = trunc i64 %31 to i8
	store i8 %32, i8* @0
	%33 = load i8, i8* @0
	%34 = icmp sgt i8 %33, 90
	br i1 %34, label %35, label %25

35:
	ret i64 0
}

```
## Run
```bash
A
B
66
C
TRUE
A
B
C
D
E
F
G
H
I
J
K
L
M
N
O
P
Q
R
S
T
U
V
W
X
Y
Z
```
