# test/test_001.md
## Source
```
(* print simple literals *)
print(42);
print(33.42);
print(TRUE, FALSE);
print("Hello, world!");
print(33H, 0AEH);
print(48656C6C6F2C20776F726C6421X)
```
## Tokens
```tsv
test/test_001.md:2:0:	ident	"print"	false	0	0.000000	(2, 0) -> (2, 5)
test/test_001.md:2:5:	lparen	"("	false	0	0.000000	(2, 5) -> (2, 6)
test/test_001.md:2:6:	integer	"42"	false	42	0.000000	(2, 6) -> (2, 8)
test/test_001.md:2:8:	rparen	")"	false	0	0.000000	(2, 8) -> (2, 9)
test/test_001.md:2:9:	semicolon	";"	false	0	0.000000	(2, 9) -> (2, 10)
test/test_001.md:3:0:	ident	"print"	false	0	0.000000	(3, 0) -> (3, 5)
test/test_001.md:3:5:	lparen	"("	false	0	0.000000	(3, 5) -> (3, 6)
test/test_001.md:3:6:	real	"33.42"	false	0	33.420000	(3, 6) -> (3, 11)
test/test_001.md:3:11:	rparen	")"	false	0	0.000000	(3, 11) -> (3, 12)
test/test_001.md:3:12:	semicolon	";"	false	0	0.000000	(3, 12) -> (3, 13)
test/test_001.md:4:0:	ident	"print"	false	0	0.000000	(4, 0) -> (4, 5)
test/test_001.md:4:5:	lparen	"("	false	0	0.000000	(4, 5) -> (4, 6)
test/test_001.md:4:6:	boolean	"TRUE"	true	0	0.000000	(4, 6) -> (4, 10)
test/test_001.md:4:10:	comma	","	false	0	0.000000	(4, 10) -> (4, 11)
test/test_001.md:4:12:	boolean	"FALSE"	false	0	0.000000	(4, 12) -> (4, 17)
test/test_001.md:4:17:	rparen	")"	false	0	0.000000	(4, 17) -> (4, 18)
test/test_001.md:4:18:	semicolon	";"	false	0	0.000000	(4, 18) -> (4, 19)
test/test_001.md:5:0:	ident	"print"	false	0	0.000000	(5, 0) -> (5, 5)
test/test_001.md:5:5:	lparen	"("	false	0	0.000000	(5, 5) -> (5, 6)
test/test_001.md:5:6:	string	"Hello, world!"	false	0	0.000000	(5, 6) -> (5, 21)
test/test_001.md:5:21:	rparen	")"	false	0	0.000000	(5, 21) -> (5, 22)
test/test_001.md:5:22:	semicolon	";"	false	0	0.000000	(5, 22) -> (5, 23)
test/test_001.md:6:0:	ident	"print"	false	0	0.000000	(6, 0) -> (6, 5)
test/test_001.md:6:5:	lparen	"("	false	0	0.000000	(6, 5) -> (6, 6)
test/test_001.md:6:6:	integer	"33H"	false	51	0.000000	(6, 6) -> (6, 9)
test/test_001.md:6:9:	comma	","	false	0	0.000000	(6, 9) -> (6, 10)
test/test_001.md:6:11:	integer	"0AEH"	false	174	0.000000	(6, 11) -> (6, 15)
test/test_001.md:6:15:	rparen	")"	false	0	0.000000	(6, 15) -> (6, 16)
test/test_001.md:6:16:	semicolon	";"	false	0	0.000000	(6, 16) -> (6, 17)
test/test_001.md:7:0:	ident	"print"	false	0	0.000000	(7, 0) -> (7, 5)
test/test_001.md:7:5:	lparen	"("	false	0	0.000000	(7, 5) -> (7, 6)
test/test_001.md:7:6:	string	"Hello, world!"	false	0	0.000000	(7, 6) -> (7, 33)
test/test_001.md:7:33:	rparen	")"	false	0	0.000000	(7, 33) -> (7, 34)
test/test_001.md:8:0:	eof	""	false	0	0.000000	(8, 0) -> (8, 0)
```
## AST
```scheme
(module
  (stmts
    (expr2stmt
      (print [void]
        (number [i64] 42)
      )
    )
    (expr2stmt
      (print [void]
        (number [f64] 33.420000)
      )
    )
    (expr2stmt
      (print [void]
        #true
        #false
      )
    )
    (expr2stmt
      (print [void]
        (string "Hello, world!")
      )
    )
    (expr2stmt
      (print [void]
        (number [i64] 51)
        (number [i64] 174)
      )
    )
    (expr2stmt
      (print [void]
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
@4 = global [13 x i8] c"Hello, world!"

declare i32 @puts(i8* %str)

declare i32 @rand()

declare i32 @sprintf(i8* %buf, i8* %format, ...)

declare i32 @printf(i8* %format, ...)

define i32 @main() {
entry:
	%0 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%1 = call i32 (i8*, ...) @printf(i8* %0, i32 42)
	%2 = getelementptr [4 x i8], [4 x i8]* @1, i32 0, i32 0
	%3 = call i32 (i8*, ...) @printf(i8* %2, double 0x4040B5C28F5C28F6)
	br i1 true, label %4, label %6

4:
	%5 = getelementptr [5 x i8], [5 x i8]* @2, i32 0, i32 0
	br label %8

6:
	%7 = getelementptr [6 x i8], [6 x i8]* @3, i32 0, i32 0
	br label %8

8:
	%9 = phi i8* [ %5, %4 ], [ %7, %6 ]
	%10 = call i32 @puts(i8* %9)
	br i1 false, label %11, label %13

11:
	%12 = getelementptr [5 x i8], [5 x i8]* @2, i32 0, i32 0
	br label %15

13:
	%14 = getelementptr [6 x i8], [6 x i8]* @3, i32 0, i32 0
	br label %15

15:
	%16 = phi i8* [ %12, %11 ], [ %14, %13 ]
	%17 = call i32 @puts(i8* %16)
	%18 = getelementptr [13 x i8], [13 x i8]* @4, i32 0, i32 0
	%19 = call i32 @puts(i8* %18)
	%20 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%21 = call i32 (i8*, ...) @printf(i8* %20, i32 51)
	%22 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%23 = call i32 (i8*, ...) @printf(i8* %22, i32 174)
	%24 = getelementptr [13 x i8], [13 x i8]* @4, i32 0, i32 0
	%25 = call i32 @puts(i8* %24)
	ret i32 0
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
