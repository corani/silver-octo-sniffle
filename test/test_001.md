# test/test_001.md
## Source
```pascal
MODULE SimpleLiterals;

BEGIN
    print(42);
    print(33.42);
    print(TRUE);
    print(FALSE);
    print("Hello, world!");
    print(33H);
    print(0AEH);
    print(48656C6C6F2C20776F726C6421X)
END SimpleLiterals.
```
## Tokens
```tsv
test/test_001.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_001.md:1:8:	ident	"SimpleLiterals"	false	0	0.000000	(1, 8) -> (1, 22)
test/test_001.md:1:22:	semicolon	";"	false	0	0.000000	(1, 22) -> (1, 23)
test/test_001.md:3:0:	begin	"BEGIN"	false	0	0.000000	(3, 0) -> (3, 5)
test/test_001.md:4:4:	ident	"print"	false	0	0.000000	(4, 4) -> (4, 9)
test/test_001.md:4:9:	lparen	"("	false	0	0.000000	(4, 9) -> (4, 10)
test/test_001.md:4:10:	integer	"42"	false	42	0.000000	(4, 10) -> (4, 12)
test/test_001.md:4:12:	rparen	")"	false	0	0.000000	(4, 12) -> (4, 13)
test/test_001.md:4:13:	semicolon	";"	false	0	0.000000	(4, 13) -> (4, 14)
test/test_001.md:5:4:	ident	"print"	false	0	0.000000	(5, 4) -> (5, 9)
test/test_001.md:5:9:	lparen	"("	false	0	0.000000	(5, 9) -> (5, 10)
test/test_001.md:5:10:	real	"33.42"	false	0	33.420000	(5, 10) -> (5, 15)
test/test_001.md:5:15:	rparen	")"	false	0	0.000000	(5, 15) -> (5, 16)
test/test_001.md:5:16:	semicolon	";"	false	0	0.000000	(5, 16) -> (5, 17)
test/test_001.md:6:4:	ident	"print"	false	0	0.000000	(6, 4) -> (6, 9)
test/test_001.md:6:9:	lparen	"("	false	0	0.000000	(6, 9) -> (6, 10)
test/test_001.md:6:10:	boolean	"TRUE"	true	0	0.000000	(6, 10) -> (6, 14)
test/test_001.md:6:14:	rparen	")"	false	0	0.000000	(6, 14) -> (6, 15)
test/test_001.md:6:15:	semicolon	";"	false	0	0.000000	(6, 15) -> (6, 16)
test/test_001.md:7:4:	ident	"print"	false	0	0.000000	(7, 4) -> (7, 9)
test/test_001.md:7:9:	lparen	"("	false	0	0.000000	(7, 9) -> (7, 10)
test/test_001.md:7:10:	boolean	"FALSE"	false	0	0.000000	(7, 10) -> (7, 15)
test/test_001.md:7:15:	rparen	")"	false	0	0.000000	(7, 15) -> (7, 16)
test/test_001.md:7:16:	semicolon	";"	false	0	0.000000	(7, 16) -> (7, 17)
test/test_001.md:8:4:	ident	"print"	false	0	0.000000	(8, 4) -> (8, 9)
test/test_001.md:8:9:	lparen	"("	false	0	0.000000	(8, 9) -> (8, 10)
test/test_001.md:8:10:	string	"Hello, world!"	false	0	0.000000	(8, 10) -> (8, 25)
test/test_001.md:8:25:	rparen	")"	false	0	0.000000	(8, 25) -> (8, 26)
test/test_001.md:8:26:	semicolon	";"	false	0	0.000000	(8, 26) -> (8, 27)
test/test_001.md:9:4:	ident	"print"	false	0	0.000000	(9, 4) -> (9, 9)
test/test_001.md:9:9:	lparen	"("	false	0	0.000000	(9, 9) -> (9, 10)
test/test_001.md:9:10:	integer	"33H"	false	51	0.000000	(9, 10) -> (9, 13)
test/test_001.md:9:13:	rparen	")"	false	0	0.000000	(9, 13) -> (9, 14)
test/test_001.md:9:14:	semicolon	";"	false	0	0.000000	(9, 14) -> (9, 15)
test/test_001.md:10:4:	ident	"print"	false	0	0.000000	(10, 4) -> (10, 9)
test/test_001.md:10:9:	lparen	"("	false	0	0.000000	(10, 9) -> (10, 10)
test/test_001.md:10:10:	integer	"0AEH"	false	174	0.000000	(10, 10) -> (10, 14)
test/test_001.md:10:14:	rparen	")"	false	0	0.000000	(10, 14) -> (10, 15)
test/test_001.md:10:15:	semicolon	";"	false	0	0.000000	(10, 15) -> (10, 16)
test/test_001.md:11:4:	ident	"print"	false	0	0.000000	(11, 4) -> (11, 9)
test/test_001.md:11:9:	lparen	"("	false	0	0.000000	(11, 9) -> (11, 10)
test/test_001.md:11:10:	string	"Hello, world!"	false	0	0.000000	(11, 10) -> (11, 37)
test/test_001.md:11:37:	rparen	")"	false	0	0.000000	(11, 37) -> (11, 38)
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
        (procedure [void] "print")
        (number [i64] 42)
      )
    )
    (expr2stmt
      (call
        (procedure [void] "print")
        (number [f64] 33.420000)
      )
    )
    (expr2stmt
      (call
        (procedure [void] "print")
        #true
      )
    )
    (expr2stmt
      (call
        (procedure [void] "print")
        #false
      )
    )
    (expr2stmt
      (call
        (procedure [void] "print")
        (string "Hello, world!")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "print")
        (number [i64] 51)
      )
    )
    (expr2stmt
      (call
        (procedure [void] "print")
        (number [i64] 174)
      )
    )
    (expr2stmt
      (call
        (procedure [void] "print")
        (string "Hello, world!")
      )
    )
  )
)
```
## IR
```llvm
@0 = global [4 x i8] c"%d\0A\00"
@1 = global [4 x i8] c"%f\0A\00"
@2 = global [5 x i8] c"TRUE\00"
@3 = global [6 x i8] c"FALSE\00"
@4 = global [14 x i8] c"Hello, world!\00"

declare i64 @puts(i8* %str)

declare i64 @rand()

declare i64 @sprintf(i8* %buf, i8* %format, ...)

declare i64 @printf(i8* %format, ...)

declare i8* @malloc(i64 %size)

define i64 @main() {
entry:
	%0 = getelementptr [4 x i8], [4 x i8]* @0, i64 0, i64 0
	%1 = call i64 (i8*, ...) @printf(i8* %0, i64 42)
	%2 = getelementptr [4 x i8], [4 x i8]* @1, i64 0, i64 0
	%3 = call i64 (i8*, ...) @printf(i8* %2, double 0x4040B5C28F5C28F6)
	br i1 true, label %4, label %6

4:
	%5 = getelementptr [5 x i8], [5 x i8]* @2, i64 0, i64 0
	br label %8

6:
	%7 = getelementptr [6 x i8], [6 x i8]* @3, i64 0, i64 0
	br label %8

8:
	%9 = phi i8* [ %5, %4 ], [ %7, %6 ]
	%10 = call i64 @puts(i8* %9)
	br i1 false, label %11, label %13

11:
	%12 = getelementptr [5 x i8], [5 x i8]* @2, i64 0, i64 0
	br label %15

13:
	%14 = getelementptr [6 x i8], [6 x i8]* @3, i64 0, i64 0
	br label %15

15:
	%16 = phi i8* [ %12, %11 ], [ %14, %13 ]
	%17 = call i64 @puts(i8* %16)
	%18 = getelementptr [14 x i8], [14 x i8]* @4, i64 0, i64 0
	%19 = call i64 @puts(i8* %18)
	%20 = getelementptr [4 x i8], [4 x i8]* @0, i64 0, i64 0
	%21 = call i64 (i8*, ...) @printf(i8* %20, i64 51)
	%22 = getelementptr [4 x i8], [4 x i8]* @0, i64 0, i64 0
	%23 = call i64 (i8*, ...) @printf(i8* %22, i64 174)
	%24 = getelementptr [14 x i8], [14 x i8]* @4, i64 0, i64 0
	%25 = call i64 @puts(i8* %24)
	ret i64 0
}

```
## Run
```bash
42
33.420000
TRUE
FALSE
Hello, world!
51
174
Hello, world!
```
