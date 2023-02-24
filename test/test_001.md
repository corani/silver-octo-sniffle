# test/test_001.md
## Source
```pascal
MODULE SimpleLiterals;

BEGIN
    C.print(42);
    C.print(33.42);
    C.print(TRUE);
    C.print(FALSE);
    C.print("Hello, world!");
    C.print(33H);
    C.print(0AEH);
    C.print(48656C6C6F2C20776F726C6421X)
END SimpleLiterals.
```
## Tokens
```tsv
test/test_001.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_001.md:1:8:	ident	"SimpleLiterals"	false	0	0.000000	(1, 8) -> (1, 22)
test/test_001.md:1:22:	semicolon	";"	false	0	0.000000	(1, 22) -> (1, 23)
test/test_001.md:3:0:	begin	"BEGIN"	false	0	0.000000	(3, 0) -> (3, 5)
test/test_001.md:4:4:	ident	"C"	false	0	0.000000	(4, 4) -> (4, 5)
test/test_001.md:4:5:	dot	"."	false	0	0.000000	(4, 5) -> (4, 6)
test/test_001.md:4:6:	ident	"print"	false	0	0.000000	(4, 6) -> (4, 11)
test/test_001.md:4:11:	lparen	"("	false	0	0.000000	(4, 11) -> (4, 12)
test/test_001.md:4:12:	integer	"42"	false	42	0.000000	(4, 12) -> (4, 14)
test/test_001.md:4:14:	rparen	")"	false	0	0.000000	(4, 14) -> (4, 15)
test/test_001.md:4:15:	semicolon	";"	false	0	0.000000	(4, 15) -> (4, 16)
test/test_001.md:5:4:	ident	"C"	false	0	0.000000	(5, 4) -> (5, 5)
test/test_001.md:5:5:	dot	"."	false	0	0.000000	(5, 5) -> (5, 6)
test/test_001.md:5:6:	ident	"print"	false	0	0.000000	(5, 6) -> (5, 11)
test/test_001.md:5:11:	lparen	"("	false	0	0.000000	(5, 11) -> (5, 12)
test/test_001.md:5:12:	real	"33.42"	false	0	33.420000	(5, 12) -> (5, 17)
test/test_001.md:5:17:	rparen	")"	false	0	0.000000	(5, 17) -> (5, 18)
test/test_001.md:5:18:	semicolon	";"	false	0	0.000000	(5, 18) -> (5, 19)
test/test_001.md:6:4:	ident	"C"	false	0	0.000000	(6, 4) -> (6, 5)
test/test_001.md:6:5:	dot	"."	false	0	0.000000	(6, 5) -> (6, 6)
test/test_001.md:6:6:	ident	"print"	false	0	0.000000	(6, 6) -> (6, 11)
test/test_001.md:6:11:	lparen	"("	false	0	0.000000	(6, 11) -> (6, 12)
test/test_001.md:6:12:	boolean	"TRUE"	true	0	0.000000	(6, 12) -> (6, 16)
test/test_001.md:6:16:	rparen	")"	false	0	0.000000	(6, 16) -> (6, 17)
test/test_001.md:6:17:	semicolon	";"	false	0	0.000000	(6, 17) -> (6, 18)
test/test_001.md:7:4:	ident	"C"	false	0	0.000000	(7, 4) -> (7, 5)
test/test_001.md:7:5:	dot	"."	false	0	0.000000	(7, 5) -> (7, 6)
test/test_001.md:7:6:	ident	"print"	false	0	0.000000	(7, 6) -> (7, 11)
test/test_001.md:7:11:	lparen	"("	false	0	0.000000	(7, 11) -> (7, 12)
test/test_001.md:7:12:	boolean	"FALSE"	false	0	0.000000	(7, 12) -> (7, 17)
test/test_001.md:7:17:	rparen	")"	false	0	0.000000	(7, 17) -> (7, 18)
test/test_001.md:7:18:	semicolon	";"	false	0	0.000000	(7, 18) -> (7, 19)
test/test_001.md:8:4:	ident	"C"	false	0	0.000000	(8, 4) -> (8, 5)
test/test_001.md:8:5:	dot	"."	false	0	0.000000	(8, 5) -> (8, 6)
test/test_001.md:8:6:	ident	"print"	false	0	0.000000	(8, 6) -> (8, 11)
test/test_001.md:8:11:	lparen	"("	false	0	0.000000	(8, 11) -> (8, 12)
test/test_001.md:8:12:	string	"Hello, world!"	false	0	0.000000	(8, 12) -> (8, 27)
test/test_001.md:8:27:	rparen	")"	false	0	0.000000	(8, 27) -> (8, 28)
test/test_001.md:8:28:	semicolon	";"	false	0	0.000000	(8, 28) -> (8, 29)
test/test_001.md:9:4:	ident	"C"	false	0	0.000000	(9, 4) -> (9, 5)
test/test_001.md:9:5:	dot	"."	false	0	0.000000	(9, 5) -> (9, 6)
test/test_001.md:9:6:	ident	"print"	false	0	0.000000	(9, 6) -> (9, 11)
test/test_001.md:9:11:	lparen	"("	false	0	0.000000	(9, 11) -> (9, 12)
test/test_001.md:9:12:	integer	"33H"	false	51	0.000000	(9, 12) -> (9, 15)
test/test_001.md:9:15:	rparen	")"	false	0	0.000000	(9, 15) -> (9, 16)
test/test_001.md:9:16:	semicolon	";"	false	0	0.000000	(9, 16) -> (9, 17)
test/test_001.md:10:4:	ident	"C"	false	0	0.000000	(10, 4) -> (10, 5)
test/test_001.md:10:5:	dot	"."	false	0	0.000000	(10, 5) -> (10, 6)
test/test_001.md:10:6:	ident	"print"	false	0	0.000000	(10, 6) -> (10, 11)
test/test_001.md:10:11:	lparen	"("	false	0	0.000000	(10, 11) -> (10, 12)
test/test_001.md:10:12:	integer	"0AEH"	false	174	0.000000	(10, 12) -> (10, 16)
test/test_001.md:10:16:	rparen	")"	false	0	0.000000	(10, 16) -> (10, 17)
test/test_001.md:10:17:	semicolon	";"	false	0	0.000000	(10, 17) -> (10, 18)
test/test_001.md:11:4:	ident	"C"	false	0	0.000000	(11, 4) -> (11, 5)
test/test_001.md:11:5:	dot	"."	false	0	0.000000	(11, 5) -> (11, 6)
test/test_001.md:11:6:	ident	"print"	false	0	0.000000	(11, 6) -> (11, 11)
test/test_001.md:11:11:	lparen	"("	false	0	0.000000	(11, 11) -> (11, 12)
test/test_001.md:11:12:	string	"Hello, world!"	false	0	0.000000	(11, 12) -> (11, 39)
test/test_001.md:11:39:	rparen	")"	false	0	0.000000	(11, 39) -> (11, 40)
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
        (procedure [void] "C.print")
        (number [i64] 42)
      )
    )
    (expr2stmt
      (call
        (procedure [void] "C.print")
        (number [f64] 33.420000)
      )
    )
    (expr2stmt
      (call
        (procedure [void] "C.print")
        #true
      )
    )
    (expr2stmt
      (call
        (procedure [void] "C.print")
        #false
      )
    )
    (expr2stmt
      (call
        (procedure [void] "C.print")
        (string "Hello, world!")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "C.print")
        (number [i64] 51)
      )
    )
    (expr2stmt
      (call
        (procedure [void] "C.print")
        (number [i64] 174)
      )
    )
    (expr2stmt
      (call
        (procedure [void] "C.print")
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

declare i8* @free(i8* %ptr)

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
