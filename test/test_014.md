# test/test_014.in
```
print((11 / 3) + (7 / 2))
```
## Tokens
```tsv
test/test_014.in:1:1:	ident	"print"	0	(1, 1) -> (1, 6)
test/test_014.in:1:6:	lparen	"("	0	(1, 6) -> (1, 7)
test/test_014.in:1:7:	lparen	"("	0	(1, 7) -> (1, 8)
test/test_014.in:1:8:	number	"11"	11	(1, 8) -> (1, 10)
test/test_014.in:1:11:	slash	"/"	0	(1, 11) -> (1, 12)
test/test_014.in:1:13:	number	"3"	3	(1, 13) -> (1, 14)
test/test_014.in:1:14:	rparen	")"	0	(1, 14) -> (1, 15)
test/test_014.in:1:16:	plus	"+"	0	(1, 16) -> (1, 17)
test/test_014.in:1:18:	lparen	"("	0	(1, 18) -> (1, 19)
test/test_014.in:1:19:	number	"7"	7	(1, 19) -> (1, 20)
test/test_014.in:1:21:	slash	"/"	0	(1, 21) -> (1, 22)
test/test_014.in:1:23:	number	"2"	2	(1, 23) -> (1, 24)
test/test_014.in:1:24:	rparen	")"	0	(1, 24) -> (1, 25)
test/test_014.in:1:25:	rparen	")"	0	(1, 25) -> (1, 26)
test/test_014.in:2:0:	eof	""	0	(2, 0) -> (2, 0)
```
## AST
```scheme
(module
  (expr2stmt
    (print [void]
      (plus [f64]
        (slash [f64]
          (number [i64] 11)
          (number [i64] 3)
        )
        (slash [f64]
          (number [i64] 7)
          (number [i64] 2)
        )
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
	%3 = sitofp i32 7 to double
	%4 = sitofp i32 2 to double
	%5 = fdiv double %3, %4
	%6 = fadd double %2, %5
	%7 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%8 = call i32 (i8*, ...) @printf(i8* %7, double %6)
	ret i32 0
}

```
## Run
```bash
7.166667
```
