# test/test_016.in
```
print(11/3*3)
```
## Tokens
```tsv
test/test_016.in:1:1:	ident	"print"	0	(1, 1) -> (1, 6)
test/test_016.in:1:6:	lparen	"("	0	(1, 6) -> (1, 7)
test/test_016.in:1:7:	number	"11"	11	(1, 7) -> (1, 9)
test/test_016.in:1:9:	slash	"/"	0	(1, 9) -> (1, 10)
test/test_016.in:1:10:	number	"3"	3	(1, 10) -> (1, 11)
test/test_016.in:1:11:	asterisk	"*"	0	(1, 11) -> (1, 12)
test/test_016.in:1:12:	number	"3"	3	(1, 12) -> (1, 13)
test/test_016.in:1:13:	rparen	")"	0	(1, 13) -> (1, 14)
test/test_016.in:2:0:	eof	""	0	(2, 0) -> (2, 0)
```
## AST
```scheme
(module
  (expr2stmt
    (print [void]
      (asterisk [f64]
        (slash [f64]
          (number [i64] 11)
          (number [i64] 3)
        )
        (number [i64] 3)
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
	%3 = sitofp i32 3 to double
	%4 = fmul double %2, %3
	%5 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%6 = call i32 (i8*, ...) @printf(i8* %5, double %4)
	ret i32 0
}

```
## Run
```bash
11.000000
```
