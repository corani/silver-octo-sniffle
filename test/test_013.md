# test/test_013.md
## Source
```
print(11 / 3 / 2)
```
## Tokens
```tsv
test/test_013.md:1:1:	ident	"print"	false	0	0.000000	(1, 1) -> (1, 6)
test/test_013.md:1:6:	lparen	"("	false	0	0.000000	(1, 6) -> (1, 7)
test/test_013.md:1:7:	integer	"11"	false	11	0.000000	(1, 7) -> (1, 9)
test/test_013.md:1:10:	slash	"/"	false	0	0.000000	(1, 10) -> (1, 11)
test/test_013.md:1:12:	integer	"3"	false	3	0.000000	(1, 12) -> (1, 13)
test/test_013.md:1:14:	slash	"/"	false	0	0.000000	(1, 14) -> (1, 15)
test/test_013.md:1:16:	integer	"2"	false	2	0.000000	(1, 16) -> (1, 17)
test/test_013.md:1:17:	rparen	")"	false	0	0.000000	(1, 17) -> (1, 18)
test/test_013.md:2:0:	eof	""	false	0	0.000000	(2, 0) -> (2, 0)
```
## AST
```scheme
(module
  (expr2stmt
    (print [void]
      (slash [f64]
        (slash [f64]
          (number [i64] 11)
          (number [i64] 3)
        )
        (number [i64] 2)
      )
    )
  )
)
```
## IR
```llvm
@0 = global [4 x i8] c"%f\0A\00"

declare i32 @puts(i8* %str)

declare i32 @rand()

declare i32 @sprintf(i8* %buf, i8* %format, ...)

declare i32 @printf(i8* %format, ...)

define i32 @main() {
entry:
	%0 = sitofp i32 11 to double
	%1 = sitofp i32 3 to double
	%2 = fdiv double %0, %1
	%3 = sitofp i32 2 to double
	%4 = fdiv double %2, %3
	%5 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%6 = call i32 (i8*, ...) @printf(i8* %5, double %4)
	ret i32 0
}

```
## Run
```bash
1.833333
```
