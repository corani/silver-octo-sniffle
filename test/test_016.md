# test/test_016.md
## Source
```pascal
MODULE Chars;

VAR x: CHAR;
    y: INTEGER;

BEGIN
    Texts.Write("A");       Texts.WriteLn;

    x := 42X;
    Texts.Write(x);         Texts.WriteLn;

    y := ORD(x);
    Texts.WriteInt(y);      Texts.WriteLn;

    INC(y);
    Texts.Write(CHR(y));    Texts.WriteLn;

    Texts.WriteInt(ORD(x <= "B")); Texts.WriteLn;

    x := "A";

    REPEAT
      Texts.Write(x);
      x := CHR(ORD(x)+1)
    UNTIL x > "Z";

    Texts.WriteLn
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
test/test_016.md:7:4:	ident	"Texts"	false	0	0.000000	(7, 4) -> (7, 9)
test/test_016.md:7:9:	dot	"."	false	0	0.000000	(7, 9) -> (7, 10)
test/test_016.md:7:10:	ident	"Write"	false	0	0.000000	(7, 10) -> (7, 15)
test/test_016.md:7:15:	lparen	"("	false	0	0.000000	(7, 15) -> (7, 16)
test/test_016.md:7:16:	string	"A"	false	0	0.000000	(7, 16) -> (7, 19)
test/test_016.md:7:19:	rparen	")"	false	0	0.000000	(7, 19) -> (7, 20)
test/test_016.md:7:20:	semicolon	";"	false	0	0.000000	(7, 20) -> (7, 21)
test/test_016.md:7:28:	ident	"Texts"	false	0	0.000000	(7, 28) -> (7, 33)
test/test_016.md:7:33:	dot	"."	false	0	0.000000	(7, 33) -> (7, 34)
test/test_016.md:7:34:	ident	"WriteLn"	false	0	0.000000	(7, 34) -> (7, 41)
test/test_016.md:7:41:	semicolon	";"	false	0	0.000000	(7, 41) -> (7, 42)
test/test_016.md:9:4:	ident	"x"	false	0	0.000000	(9, 4) -> (9, 5)
test/test_016.md:9:6:	assign	":="	false	0	0.000000	(9, 6) -> (9, 8)
test/test_016.md:9:9:	string	"B"	false	0	0.000000	(9, 9) -> (9, 12)
test/test_016.md:9:12:	semicolon	";"	false	0	0.000000	(9, 12) -> (9, 13)
test/test_016.md:10:4:	ident	"Texts"	false	0	0.000000	(10, 4) -> (10, 9)
test/test_016.md:10:9:	dot	"."	false	0	0.000000	(10, 9) -> (10, 10)
test/test_016.md:10:10:	ident	"Write"	false	0	0.000000	(10, 10) -> (10, 15)
test/test_016.md:10:15:	lparen	"("	false	0	0.000000	(10, 15) -> (10, 16)
test/test_016.md:10:16:	ident	"x"	false	0	0.000000	(10, 16) -> (10, 17)
test/test_016.md:10:17:	rparen	")"	false	0	0.000000	(10, 17) -> (10, 18)
test/test_016.md:10:18:	semicolon	";"	false	0	0.000000	(10, 18) -> (10, 19)
test/test_016.md:10:28:	ident	"Texts"	false	0	0.000000	(10, 28) -> (10, 33)
test/test_016.md:10:33:	dot	"."	false	0	0.000000	(10, 33) -> (10, 34)
test/test_016.md:10:34:	ident	"WriteLn"	false	0	0.000000	(10, 34) -> (10, 41)
test/test_016.md:10:41:	semicolon	";"	false	0	0.000000	(10, 41) -> (10, 42)
test/test_016.md:12:4:	ident	"y"	false	0	0.000000	(12, 4) -> (12, 5)
test/test_016.md:12:6:	assign	":="	false	0	0.000000	(12, 6) -> (12, 8)
test/test_016.md:12:9:	ident	"ORD"	false	0	0.000000	(12, 9) -> (12, 12)
test/test_016.md:12:12:	lparen	"("	false	0	0.000000	(12, 12) -> (12, 13)
test/test_016.md:12:13:	ident	"x"	false	0	0.000000	(12, 13) -> (12, 14)
test/test_016.md:12:14:	rparen	")"	false	0	0.000000	(12, 14) -> (12, 15)
test/test_016.md:12:15:	semicolon	";"	false	0	0.000000	(12, 15) -> (12, 16)
test/test_016.md:13:4:	ident	"Texts"	false	0	0.000000	(13, 4) -> (13, 9)
test/test_016.md:13:9:	dot	"."	false	0	0.000000	(13, 9) -> (13, 10)
test/test_016.md:13:10:	ident	"WriteInt"	false	0	0.000000	(13, 10) -> (13, 18)
test/test_016.md:13:18:	lparen	"("	false	0	0.000000	(13, 18) -> (13, 19)
test/test_016.md:13:19:	ident	"y"	false	0	0.000000	(13, 19) -> (13, 20)
test/test_016.md:13:20:	rparen	")"	false	0	0.000000	(13, 20) -> (13, 21)
test/test_016.md:13:21:	semicolon	";"	false	0	0.000000	(13, 21) -> (13, 22)
test/test_016.md:13:28:	ident	"Texts"	false	0	0.000000	(13, 28) -> (13, 33)
test/test_016.md:13:33:	dot	"."	false	0	0.000000	(13, 33) -> (13, 34)
test/test_016.md:13:34:	ident	"WriteLn"	false	0	0.000000	(13, 34) -> (13, 41)
test/test_016.md:13:41:	semicolon	";"	false	0	0.000000	(13, 41) -> (13, 42)
test/test_016.md:15:4:	ident	"INC"	false	0	0.000000	(15, 4) -> (15, 7)
test/test_016.md:15:7:	lparen	"("	false	0	0.000000	(15, 7) -> (15, 8)
test/test_016.md:15:8:	ident	"y"	false	0	0.000000	(15, 8) -> (15, 9)
test/test_016.md:15:9:	rparen	")"	false	0	0.000000	(15, 9) -> (15, 10)
test/test_016.md:15:10:	semicolon	";"	false	0	0.000000	(15, 10) -> (15, 11)
test/test_016.md:16:4:	ident	"Texts"	false	0	0.000000	(16, 4) -> (16, 9)
test/test_016.md:16:9:	dot	"."	false	0	0.000000	(16, 9) -> (16, 10)
test/test_016.md:16:10:	ident	"Write"	false	0	0.000000	(16, 10) -> (16, 15)
test/test_016.md:16:15:	lparen	"("	false	0	0.000000	(16, 15) -> (16, 16)
test/test_016.md:16:16:	ident	"CHR"	false	0	0.000000	(16, 16) -> (16, 19)
test/test_016.md:16:19:	lparen	"("	false	0	0.000000	(16, 19) -> (16, 20)
test/test_016.md:16:20:	ident	"y"	false	0	0.000000	(16, 20) -> (16, 21)
test/test_016.md:16:21:	rparen	")"	false	0	0.000000	(16, 21) -> (16, 22)
test/test_016.md:16:22:	rparen	")"	false	0	0.000000	(16, 22) -> (16, 23)
test/test_016.md:16:23:	semicolon	";"	false	0	0.000000	(16, 23) -> (16, 24)
test/test_016.md:16:28:	ident	"Texts"	false	0	0.000000	(16, 28) -> (16, 33)
test/test_016.md:16:33:	dot	"."	false	0	0.000000	(16, 33) -> (16, 34)
test/test_016.md:16:34:	ident	"WriteLn"	false	0	0.000000	(16, 34) -> (16, 41)
test/test_016.md:16:41:	semicolon	";"	false	0	0.000000	(16, 41) -> (16, 42)
test/test_016.md:18:4:	ident	"Texts"	false	0	0.000000	(18, 4) -> (18, 9)
test/test_016.md:18:9:	dot	"."	false	0	0.000000	(18, 9) -> (18, 10)
test/test_016.md:18:10:	ident	"WriteInt"	false	0	0.000000	(18, 10) -> (18, 18)
test/test_016.md:18:18:	lparen	"("	false	0	0.000000	(18, 18) -> (18, 19)
test/test_016.md:18:19:	ident	"ORD"	false	0	0.000000	(18, 19) -> (18, 22)
test/test_016.md:18:22:	lparen	"("	false	0	0.000000	(18, 22) -> (18, 23)
test/test_016.md:18:23:	ident	"x"	false	0	0.000000	(18, 23) -> (18, 24)
test/test_016.md:18:25:	le	"<="	false	0	0.000000	(18, 25) -> (18, 27)
test/test_016.md:18:28:	string	"B"	false	0	0.000000	(18, 28) -> (18, 31)
test/test_016.md:18:31:	rparen	")"	false	0	0.000000	(18, 31) -> (18, 32)
test/test_016.md:18:32:	rparen	")"	false	0	0.000000	(18, 32) -> (18, 33)
test/test_016.md:18:33:	semicolon	";"	false	0	0.000000	(18, 33) -> (18, 34)
test/test_016.md:18:35:	ident	"Texts"	false	0	0.000000	(18, 35) -> (18, 40)
test/test_016.md:18:40:	dot	"."	false	0	0.000000	(18, 40) -> (18, 41)
test/test_016.md:18:41:	ident	"WriteLn"	false	0	0.000000	(18, 41) -> (18, 48)
test/test_016.md:18:48:	semicolon	";"	false	0	0.000000	(18, 48) -> (18, 49)
test/test_016.md:20:4:	ident	"x"	false	0	0.000000	(20, 4) -> (20, 5)
test/test_016.md:20:6:	assign	":="	false	0	0.000000	(20, 6) -> (20, 8)
test/test_016.md:20:9:	string	"A"	false	0	0.000000	(20, 9) -> (20, 12)
test/test_016.md:20:12:	semicolon	";"	false	0	0.000000	(20, 12) -> (20, 13)
test/test_016.md:22:4:	repeat	"REPEAT"	false	0	0.000000	(22, 4) -> (22, 10)
test/test_016.md:23:6:	ident	"Texts"	false	0	0.000000	(23, 6) -> (23, 11)
test/test_016.md:23:11:	dot	"."	false	0	0.000000	(23, 11) -> (23, 12)
test/test_016.md:23:12:	ident	"Write"	false	0	0.000000	(23, 12) -> (23, 17)
test/test_016.md:23:17:	lparen	"("	false	0	0.000000	(23, 17) -> (23, 18)
test/test_016.md:23:18:	ident	"x"	false	0	0.000000	(23, 18) -> (23, 19)
test/test_016.md:23:19:	rparen	")"	false	0	0.000000	(23, 19) -> (23, 20)
test/test_016.md:23:20:	semicolon	";"	false	0	0.000000	(23, 20) -> (23, 21)
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
test/test_016.md:25:17:	semicolon	";"	false	0	0.000000	(25, 17) -> (25, 18)
test/test_016.md:27:4:	ident	"Texts"	false	0	0.000000	(27, 4) -> (27, 9)
test/test_016.md:27:9:	dot	"."	false	0	0.000000	(27, 9) -> (27, 10)
test/test_016.md:27:10:	ident	"WriteLn"	false	0	0.000000	(27, 10) -> (27, 17)
test/test_016.md:28:0:	end	"END"	false	0	0.000000	(28, 0) -> (28, 3)
test/test_016.md:28:4:	ident	"Chars"	false	0	0.000000	(28, 4) -> (28, 9)
test/test_016.md:28:9:	dot	"."	false	0	0.000000	(28, 9) -> (28, 10)
test/test_016.md:29:0:	eof	""	false	0	0.000000	(29, 0) -> (29, 0)
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
        (procedure [void] "Texts.Write")
        (char "A")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
    (assign
      (variable [char] "x")
      (char "B")
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.Write")
        (variable [char] "x")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
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
        (procedure [void] "Texts.WriteInt")
        (variable [i64] "y")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
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
        (procedure [void] "Texts.Write")
        (call
          (procedure [char] "CHR")
          (variable [i64] "y")
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteInt")
        (call
          (procedure [i64] "ORD")
          (le [boolean]
            (variable [char] "x")
            (char "B")
          )
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
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
            (procedure [void] "Texts.Write")
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
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
  )
)
```
## IR
```llvm
@0 = global i8 0
@1 = global i64 0
@2 = global [3 x i8] c"%c\00"
@3 = global [1 x i8] c"\00"
@4 = global [3 x i8] c"%d\00"
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
	%0 = getelementptr [3 x i8], [3 x i8]* @2, i64 0, i64 0
	%1 = call i64 (i8*, ...) @printf(i8* %0, i8 65)
	%2 = getelementptr [1 x i8], [1 x i8]* @3, i64 0, i64 0
	%3 = call i64 @puts(i8* %2)
	store i8 66, i8* @0
	%4 = load i8, i8* @0
	%5 = getelementptr [3 x i8], [3 x i8]* @2, i64 0, i64 0
	%6 = call i64 (i8*, ...) @printf(i8* %5, i8 %4)
	%7 = getelementptr [1 x i8], [1 x i8]* @3, i64 0, i64 0
	%8 = call i64 @puts(i8* %7)
	%9 = load i8, i8* @0
	%10 = zext i8 %9 to i64
	store i64 %10, i64* @1
	%11 = load i64, i64* @1
	%12 = getelementptr [3 x i8], [3 x i8]* @4, i64 0, i64 0
	%13 = call i64 (i8*, ...) @printf(i8* %12, i64 %11)
	%14 = getelementptr [1 x i8], [1 x i8]* @3, i64 0, i64 0
	%15 = call i64 @puts(i8* %14)
	%16 = load i64, i64* @1
	%17 = add i64 %16, 1
	store i64 %17, i64* @1
	%18 = load i64, i64* @1
	%19 = trunc i64 %18 to i8
	%20 = getelementptr [3 x i8], [3 x i8]* @2, i64 0, i64 0
	%21 = call i64 (i8*, ...) @printf(i8* %20, i8 %19)
	%22 = getelementptr [1 x i8], [1 x i8]* @3, i64 0, i64 0
	%23 = call i64 @puts(i8* %22)
	%24 = load i8, i8* @0
	%25 = icmp sle i8 %24, 66
	%26 = zext i1 %25 to i64
	%27 = getelementptr [3 x i8], [3 x i8]* @4, i64 0, i64 0
	%28 = call i64 (i8*, ...) @printf(i8* %27, i64 %26)
	%29 = getelementptr [1 x i8], [1 x i8]* @3, i64 0, i64 0
	%30 = call i64 @puts(i8* %29)
	store i8 65, i8* @0
	br label %31

31:
	%32 = load i8, i8* @0
	%33 = getelementptr [3 x i8], [3 x i8]* @2, i64 0, i64 0
	%34 = call i64 (i8*, ...) @printf(i8* %33, i8 %32)
	%35 = load i8, i8* @0
	%36 = zext i8 %35 to i64
	%37 = add i64 %36, 1
	%38 = trunc i64 %37 to i8
	store i8 %38, i8* @0
	%39 = load i8, i8* @0
	%40 = icmp sgt i8 %39, 90
	br i1 %40, label %41, label %31

41:
	%42 = getelementptr [1 x i8], [1 x i8]* @3, i64 0, i64 0
	%43 = call i64 @puts(i8* %42)
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
A
B
66
C
1
ABCDEFGHIJKLMNOPQRSTUVWXYZ
```
