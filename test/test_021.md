# test/test_021.in
```
print(TRUE, FALSE)
```
## Tokens
```tsv
test/test_021.in:1:1:	ident	"print"	false	0	0.000000	(1, 1) -> (1, 6)
test/test_021.in:1:6:	lparen	"("	false	0	0.000000	(1, 6) -> (1, 7)
test/test_021.in:1:7:	boolean	"TRUE"	true	0	0.000000	(1, 7) -> (1, 11)
test/test_021.in:1:11:	comma	","	false	0	0.000000	(1, 11) -> (1, 12)
test/test_021.in:1:13:	boolean	"FALSE"	false	0	0.000000	(1, 13) -> (1, 18)
test/test_021.in:1:18:	rparen	")"	false	0	0.000000	(1, 18) -> (1, 19)
test/test_021.in:2:0:	eof	""	false	0	0.000000	(2, 0) -> (2, 0)
```
## AST
```scheme
(module
  (expr2stmt
    (print [void]
      #true
      #false
    )
  )
)
```
## IR
```llvm
@0 = global [5 x i8] c"TRUE\00"
@1 = global [6 x i8] c"FALSE\00"

declare i32 @puts(i8* %str)

declare i32 @rand()

declare i32 @sprintf(i8* %buf, i8* %format, ...)

declare i32 @printf(i8* %format, ...)

define i32 @main() {
entry:
	br i1 true, label %0, label %2

0:
	%1 = getelementptr [5 x i8], [5 x i8]* @0, i32 0, i32 0
	br label %4

2:
	%3 = getelementptr [6 x i8], [6 x i8]* @1, i32 0, i32 0
	br label %4

4:
	%5 = phi i8* [ %1, %0 ], [ %3, %2 ]
	%6 = call i32 @puts(i8* %5)
	br i1 false, label %7, label %9

7:
	%8 = getelementptr [5 x i8], [5 x i8]* @0, i32 0, i32 0
	br label %11

9:
	%10 = getelementptr [6 x i8], [6 x i8]* @1, i32 0, i32 0
	br label %11

11:
	%12 = phi i8* [ %8, %7 ], [ %10, %9 ]
	%13 = call i32 @puts(i8* %12)
	ret i32 0
}

```
## Run
```bash
TRUE
FALSE
```
