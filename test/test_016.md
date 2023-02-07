# test/test_016.md
## Source
```pascal
MODULE Chars;

VAR x: CHAR;
    y: INTEGER;

BEGIN
    print("A");

    x := "B";
    print(x);

    y := ORD(x);
    print(y);

    INC(y);
    print(CHR(y));
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
test/test_016.md:7:4:	ident	"print"	false	0	0.000000	(7, 4) -> (7, 9)
test/test_016.md:7:9:	lparen	"("	false	0	0.000000	(7, 9) -> (7, 10)
test/test_016.md:7:10:	string	"A"	false	0	0.000000	(7, 10) -> (7, 13)
test/test_016.md:7:13:	rparen	")"	false	0	0.000000	(7, 13) -> (7, 14)
test/test_016.md:7:14:	semicolon	";"	false	0	0.000000	(7, 14) -> (7, 15)
test/test_016.md:9:4:	ident	"x"	false	0	0.000000	(9, 4) -> (9, 5)
test/test_016.md:9:6:	assign	":="	false	0	0.000000	(9, 6) -> (9, 8)
test/test_016.md:9:9:	string	"B"	false	0	0.000000	(9, 9) -> (9, 12)
test/test_016.md:9:12:	semicolon	";"	false	0	0.000000	(9, 12) -> (9, 13)
test/test_016.md:10:4:	ident	"print"	false	0	0.000000	(10, 4) -> (10, 9)
test/test_016.md:10:9:	lparen	"("	false	0	0.000000	(10, 9) -> (10, 10)
test/test_016.md:10:10:	ident	"x"	false	0	0.000000	(10, 10) -> (10, 11)
test/test_016.md:10:11:	rparen	")"	false	0	0.000000	(10, 11) -> (10, 12)
test/test_016.md:10:12:	semicolon	";"	false	0	0.000000	(10, 12) -> (10, 13)
test/test_016.md:12:4:	ident	"y"	false	0	0.000000	(12, 4) -> (12, 5)
test/test_016.md:12:6:	assign	":="	false	0	0.000000	(12, 6) -> (12, 8)
test/test_016.md:12:9:	ident	"ORD"	false	0	0.000000	(12, 9) -> (12, 12)
test/test_016.md:12:12:	lparen	"("	false	0	0.000000	(12, 12) -> (12, 13)
test/test_016.md:12:13:	ident	"x"	false	0	0.000000	(12, 13) -> (12, 14)
test/test_016.md:12:14:	rparen	")"	false	0	0.000000	(12, 14) -> (12, 15)
test/test_016.md:12:15:	semicolon	";"	false	0	0.000000	(12, 15) -> (12, 16)
test/test_016.md:13:4:	ident	"print"	false	0	0.000000	(13, 4) -> (13, 9)
test/test_016.md:13:9:	lparen	"("	false	0	0.000000	(13, 9) -> (13, 10)
test/test_016.md:13:10:	ident	"y"	false	0	0.000000	(13, 10) -> (13, 11)
test/test_016.md:13:11:	rparen	")"	false	0	0.000000	(13, 11) -> (13, 12)
test/test_016.md:13:12:	semicolon	";"	false	0	0.000000	(13, 12) -> (13, 13)
test/test_016.md:15:4:	ident	"INC"	false	0	0.000000	(15, 4) -> (15, 7)
test/test_016.md:15:7:	lparen	"("	false	0	0.000000	(15, 7) -> (15, 8)
test/test_016.md:15:8:	ident	"y"	false	0	0.000000	(15, 8) -> (15, 9)
test/test_016.md:15:9:	rparen	")"	false	0	0.000000	(15, 9) -> (15, 10)
test/test_016.md:15:10:	semicolon	";"	false	0	0.000000	(15, 10) -> (15, 11)
test/test_016.md:16:4:	ident	"print"	false	0	0.000000	(16, 4) -> (16, 9)
test/test_016.md:16:9:	lparen	"("	false	0	0.000000	(16, 9) -> (16, 10)
test/test_016.md:16:10:	ident	"CHR"	false	0	0.000000	(16, 10) -> (16, 13)
test/test_016.md:16:13:	lparen	"("	false	0	0.000000	(16, 13) -> (16, 14)
test/test_016.md:16:14:	ident	"y"	false	0	0.000000	(16, 14) -> (16, 15)
test/test_016.md:16:15:	rparen	")"	false	0	0.000000	(16, 15) -> (16, 16)
test/test_016.md:16:16:	rparen	")"	false	0	0.000000	(16, 16) -> (16, 17)
test/test_016.md:16:17:	semicolon	";"	false	0	0.000000	(16, 17) -> (16, 18)
test/test_016.md:17:0:	end	"END"	false	0	0.000000	(17, 0) -> (17, 3)
test/test_016.md:17:4:	ident	"Chars"	false	0	0.000000	(17, 4) -> (17, 9)
test/test_016.md:17:9:	dot	"."	false	0	0.000000	(17, 9) -> (17, 10)
test/test_016.md:18:0:	eof	""	false	0	0.000000	(18, 0) -> (18, 0)
```
## AST
```scheme
(module "Chars"
  (vars
    (x [char])
    (y [i64])
  )
  (stmts
    (expr2stmt
      (call "print" [void]
        (char "A")
      )
    )
    (assign x
      (char "B")
    )
    (expr2stmt
      (call "print" [void]
        (variable [char] "x")
      )
    )
    (assign y
      (call "ORD" [i64]
        (variable [char] "x")
      )
    )
    (expr2stmt
      (call "print" [void]
        (variable [i64] "y")
      )
    )
    (expr2stmt
      (call "INC" [void]
        (variable [i64] "y")
      )
    )
    (expr2stmt
      (call "print" [void]
        (call "CHR" [char]
          (variable [i64] "y")
        )
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

declare i64 @puts(i8* %str)

declare i64 @rand()

declare i64 @sprintf(i8* %buf, i8* %format, ...)

declare i64 @printf(i8* %format, ...)

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
	ret i64 0
}

```
## Run
```bash
A
B
66
C
```
