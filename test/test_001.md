# test/test_001.md
## Source
```pascal
MODULE SimpleLiterals;

BEGIN
    Texts.WriteInt(42);                             Texts.WriteLn;
    Texts.WriteReal(33.42);                         Texts.WriteLn;
    Texts.WriteInt(ORD(TRUE));                      Texts.WriteLn;
    Texts.WriteInt(ORD(FALSE));                     Texts.WriteLn;
    Texts.WriteString("Hello, world!");             Texts.WriteLn;
    Texts.WriteInt(33H);                            Texts.WriteLn;
    Texts.WriteInt(0AEH);                           Texts.WriteLn;
    Texts.WriteString(48656C6C6F2C20776F726C6421X); Texts.WriteLn;
END SimpleLiterals.
```
## Tokens
```tsv
test/test_001.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_001.md:1:8:	ident	"SimpleLiterals"	false	0	0.000000	(1, 8) -> (1, 22)
test/test_001.md:1:22:	semicolon	";"	false	0	0.000000	(1, 22) -> (1, 23)
test/test_001.md:3:0:	begin	"BEGIN"	false	0	0.000000	(3, 0) -> (3, 5)
test/test_001.md:4:4:	ident	"Texts"	false	0	0.000000	(4, 4) -> (4, 9)
test/test_001.md:4:9:	dot	"."	false	0	0.000000	(4, 9) -> (4, 10)
test/test_001.md:4:10:	ident	"WriteInt"	false	0	0.000000	(4, 10) -> (4, 18)
test/test_001.md:4:18:	lparen	"("	false	0	0.000000	(4, 18) -> (4, 19)
test/test_001.md:4:19:	integer	"42"	false	42	0.000000	(4, 19) -> (4, 21)
test/test_001.md:4:21:	rparen	")"	false	0	0.000000	(4, 21) -> (4, 22)
test/test_001.md:4:22:	semicolon	";"	false	0	0.000000	(4, 22) -> (4, 23)
test/test_001.md:4:52:	ident	"Texts"	false	0	0.000000	(4, 52) -> (4, 57)
test/test_001.md:4:57:	dot	"."	false	0	0.000000	(4, 57) -> (4, 58)
test/test_001.md:4:58:	ident	"WriteLn"	false	0	0.000000	(4, 58) -> (4, 65)
test/test_001.md:4:65:	semicolon	";"	false	0	0.000000	(4, 65) -> (4, 66)
test/test_001.md:5:4:	ident	"Texts"	false	0	0.000000	(5, 4) -> (5, 9)
test/test_001.md:5:9:	dot	"."	false	0	0.000000	(5, 9) -> (5, 10)
test/test_001.md:5:10:	ident	"WriteReal"	false	0	0.000000	(5, 10) -> (5, 19)
test/test_001.md:5:19:	lparen	"("	false	0	0.000000	(5, 19) -> (5, 20)
test/test_001.md:5:20:	real	"33.42"	false	0	33.420000	(5, 20) -> (5, 25)
test/test_001.md:5:25:	rparen	")"	false	0	0.000000	(5, 25) -> (5, 26)
test/test_001.md:5:26:	semicolon	";"	false	0	0.000000	(5, 26) -> (5, 27)
test/test_001.md:5:52:	ident	"Texts"	false	0	0.000000	(5, 52) -> (5, 57)
test/test_001.md:5:57:	dot	"."	false	0	0.000000	(5, 57) -> (5, 58)
test/test_001.md:5:58:	ident	"WriteLn"	false	0	0.000000	(5, 58) -> (5, 65)
test/test_001.md:5:65:	semicolon	";"	false	0	0.000000	(5, 65) -> (5, 66)
test/test_001.md:6:4:	ident	"Texts"	false	0	0.000000	(6, 4) -> (6, 9)
test/test_001.md:6:9:	dot	"."	false	0	0.000000	(6, 9) -> (6, 10)
test/test_001.md:6:10:	ident	"WriteInt"	false	0	0.000000	(6, 10) -> (6, 18)
test/test_001.md:6:18:	lparen	"("	false	0	0.000000	(6, 18) -> (6, 19)
test/test_001.md:6:19:	ident	"ORD"	false	0	0.000000	(6, 19) -> (6, 22)
test/test_001.md:6:22:	lparen	"("	false	0	0.000000	(6, 22) -> (6, 23)
test/test_001.md:6:23:	boolean	"TRUE"	true	0	0.000000	(6, 23) -> (6, 27)
test/test_001.md:6:27:	rparen	")"	false	0	0.000000	(6, 27) -> (6, 28)
test/test_001.md:6:28:	rparen	")"	false	0	0.000000	(6, 28) -> (6, 29)
test/test_001.md:6:29:	semicolon	";"	false	0	0.000000	(6, 29) -> (6, 30)
test/test_001.md:6:52:	ident	"Texts"	false	0	0.000000	(6, 52) -> (6, 57)
test/test_001.md:6:57:	dot	"."	false	0	0.000000	(6, 57) -> (6, 58)
test/test_001.md:6:58:	ident	"WriteLn"	false	0	0.000000	(6, 58) -> (6, 65)
test/test_001.md:6:65:	semicolon	";"	false	0	0.000000	(6, 65) -> (6, 66)
test/test_001.md:7:4:	ident	"Texts"	false	0	0.000000	(7, 4) -> (7, 9)
test/test_001.md:7:9:	dot	"."	false	0	0.000000	(7, 9) -> (7, 10)
test/test_001.md:7:10:	ident	"WriteInt"	false	0	0.000000	(7, 10) -> (7, 18)
test/test_001.md:7:18:	lparen	"("	false	0	0.000000	(7, 18) -> (7, 19)
test/test_001.md:7:19:	ident	"ORD"	false	0	0.000000	(7, 19) -> (7, 22)
test/test_001.md:7:22:	lparen	"("	false	0	0.000000	(7, 22) -> (7, 23)
test/test_001.md:7:23:	boolean	"FALSE"	false	0	0.000000	(7, 23) -> (7, 28)
test/test_001.md:7:28:	rparen	")"	false	0	0.000000	(7, 28) -> (7, 29)
test/test_001.md:7:29:	rparen	")"	false	0	0.000000	(7, 29) -> (7, 30)
test/test_001.md:7:30:	semicolon	";"	false	0	0.000000	(7, 30) -> (7, 31)
test/test_001.md:7:52:	ident	"Texts"	false	0	0.000000	(7, 52) -> (7, 57)
test/test_001.md:7:57:	dot	"."	false	0	0.000000	(7, 57) -> (7, 58)
test/test_001.md:7:58:	ident	"WriteLn"	false	0	0.000000	(7, 58) -> (7, 65)
test/test_001.md:7:65:	semicolon	";"	false	0	0.000000	(7, 65) -> (7, 66)
test/test_001.md:8:4:	ident	"Texts"	false	0	0.000000	(8, 4) -> (8, 9)
test/test_001.md:8:9:	dot	"."	false	0	0.000000	(8, 9) -> (8, 10)
test/test_001.md:8:10:	ident	"WriteString"	false	0	0.000000	(8, 10) -> (8, 21)
test/test_001.md:8:21:	lparen	"("	false	0	0.000000	(8, 21) -> (8, 22)
test/test_001.md:8:22:	string	"Hello, world!"	false	0	0.000000	(8, 22) -> (8, 37)
test/test_001.md:8:37:	rparen	")"	false	0	0.000000	(8, 37) -> (8, 38)
test/test_001.md:8:38:	semicolon	";"	false	0	0.000000	(8, 38) -> (8, 39)
test/test_001.md:8:52:	ident	"Texts"	false	0	0.000000	(8, 52) -> (8, 57)
test/test_001.md:8:57:	dot	"."	false	0	0.000000	(8, 57) -> (8, 58)
test/test_001.md:8:58:	ident	"WriteLn"	false	0	0.000000	(8, 58) -> (8, 65)
test/test_001.md:8:65:	semicolon	";"	false	0	0.000000	(8, 65) -> (8, 66)
test/test_001.md:9:4:	ident	"Texts"	false	0	0.000000	(9, 4) -> (9, 9)
test/test_001.md:9:9:	dot	"."	false	0	0.000000	(9, 9) -> (9, 10)
test/test_001.md:9:10:	ident	"WriteInt"	false	0	0.000000	(9, 10) -> (9, 18)
test/test_001.md:9:18:	lparen	"("	false	0	0.000000	(9, 18) -> (9, 19)
test/test_001.md:9:19:	integer	"33H"	false	51	0.000000	(9, 19) -> (9, 22)
test/test_001.md:9:22:	rparen	")"	false	0	0.000000	(9, 22) -> (9, 23)
test/test_001.md:9:23:	semicolon	";"	false	0	0.000000	(9, 23) -> (9, 24)
test/test_001.md:9:52:	ident	"Texts"	false	0	0.000000	(9, 52) -> (9, 57)
test/test_001.md:9:57:	dot	"."	false	0	0.000000	(9, 57) -> (9, 58)
test/test_001.md:9:58:	ident	"WriteLn"	false	0	0.000000	(9, 58) -> (9, 65)
test/test_001.md:9:65:	semicolon	";"	false	0	0.000000	(9, 65) -> (9, 66)
test/test_001.md:10:4:	ident	"Texts"	false	0	0.000000	(10, 4) -> (10, 9)
test/test_001.md:10:9:	dot	"."	false	0	0.000000	(10, 9) -> (10, 10)
test/test_001.md:10:10:	ident	"WriteInt"	false	0	0.000000	(10, 10) -> (10, 18)
test/test_001.md:10:18:	lparen	"("	false	0	0.000000	(10, 18) -> (10, 19)
test/test_001.md:10:19:	integer	"0AEH"	false	174	0.000000	(10, 19) -> (10, 23)
test/test_001.md:10:23:	rparen	")"	false	0	0.000000	(10, 23) -> (10, 24)
test/test_001.md:10:24:	semicolon	";"	false	0	0.000000	(10, 24) -> (10, 25)
test/test_001.md:10:52:	ident	"Texts"	false	0	0.000000	(10, 52) -> (10, 57)
test/test_001.md:10:57:	dot	"."	false	0	0.000000	(10, 57) -> (10, 58)
test/test_001.md:10:58:	ident	"WriteLn"	false	0	0.000000	(10, 58) -> (10, 65)
test/test_001.md:10:65:	semicolon	";"	false	0	0.000000	(10, 65) -> (10, 66)
test/test_001.md:11:4:	ident	"Texts"	false	0	0.000000	(11, 4) -> (11, 9)
test/test_001.md:11:9:	dot	"."	false	0	0.000000	(11, 9) -> (11, 10)
test/test_001.md:11:10:	ident	"WriteString"	false	0	0.000000	(11, 10) -> (11, 21)
test/test_001.md:11:21:	lparen	"("	false	0	0.000000	(11, 21) -> (11, 22)
test/test_001.md:11:22:	string	"Hello, world!"	false	0	0.000000	(11, 22) -> (11, 49)
test/test_001.md:11:49:	rparen	")"	false	0	0.000000	(11, 49) -> (11, 50)
test/test_001.md:11:50:	semicolon	";"	false	0	0.000000	(11, 50) -> (11, 51)
test/test_001.md:11:52:	ident	"Texts"	false	0	0.000000	(11, 52) -> (11, 57)
test/test_001.md:11:57:	dot	"."	false	0	0.000000	(11, 57) -> (11, 58)
test/test_001.md:11:58:	ident	"WriteLn"	false	0	0.000000	(11, 58) -> (11, 65)
test/test_001.md:11:65:	semicolon	";"	false	0	0.000000	(11, 65) -> (11, 66)
test/test_001.md:12:0:	end	"END"	false	0	0.000000	(12, 0) -> (12, 3)
test/test_001.md:12:4:	ident	"SimpleLiterals"	false	0	0.000000	(12, 4) -> (12, 18)
test/test_001.md:12:18:	dot	"."	false	0	0.000000	(12, 18) -> (12, 19)
test/test_001.md:13:0:	eof	""	false	0	0.000000	(13, 0) -> (13, 0)
```
## AST
```scheme
(module "SimpleLiterals"
  (stmts
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteInt")
        (number [i64] 42)
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteReal")
        (number [f64] 33.420000)
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
          #true
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
          #false
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
        (procedure [void] "Texts.WriteString")
        (string "Hello, world!")
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
        (number [i64] 51)
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
        (number [i64] 174)
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteString")
        (string "Hello, world!")
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
@0 = global [3 x i8] c"%d\00"
@1 = global [1 x i8] c"\00"
@2 = global [3 x i8] c"%f\00"
@3 = global [14 x i8] c"Hello, world!\00"
@__argc = global i64 0
@__argv = global i8** inttoptr (i8 0 to i8**)

declare i64 @puts(i8* %str)

declare i64 @rand()

declare i64 @sprintf(i8* %buf, i8* %format, ...)

declare i64 @printf(i8* %format, ...)

declare i8* @malloc(i64 %size)

declare i8* @free(i8* %ptr)

define void @oberonMain() {
entry:
	%0 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%1 = call i64 (i8*, ...) @printf(i8* %0, i64 42)
	%2 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%3 = call i64 @puts(i8* %2)
	%4 = getelementptr [3 x i8], [3 x i8]* @2, i64 0, i64 0
	%5 = call i64 (i8*, ...) @printf(i8* %4, double 0x4040B5C28F5C28F6)
	%6 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%7 = call i64 @puts(i8* %6)
	%8 = zext i1 true to i64
	%9 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%10 = call i64 (i8*, ...) @printf(i8* %9, i64 %8)
	%11 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%12 = call i64 @puts(i8* %11)
	%13 = zext i1 false to i64
	%14 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%15 = call i64 (i8*, ...) @printf(i8* %14, i64 %13)
	%16 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%17 = call i64 @puts(i8* %16)
	%18 = getelementptr [14 x i8], [14 x i8]* @3, i64 0, i64 0
	%19 = call i64 (i8*, ...) @printf(i8* %18)
	%20 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%21 = call i64 @puts(i8* %20)
	%22 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%23 = call i64 (i8*, ...) @printf(i8* %22, i64 51)
	%24 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%25 = call i64 @puts(i8* %24)
	%26 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%27 = call i64 (i8*, ...) @printf(i8* %26, i64 174)
	%28 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%29 = call i64 @puts(i8* %28)
	%30 = getelementptr [14 x i8], [14 x i8]* @3, i64 0, i64 0
	%31 = call i64 (i8*, ...) @printf(i8* %30)
	%32 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%33 = call i64 @puts(i8* %32)
	ret void
}

define i64 @main(i64 %argc, i8** %argv) {
entry:
	store i64 %argc, i64* @__argc
	store i8** %argv, i8*** @__argv
	call void @oberonMain()
	ret i64 0
}

```
## Run
```bash
42
33.420000
1
0
Hello, world!
51
174
Hello, world!
```
