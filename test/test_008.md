# test/test_008.in
```
print( (34 + 35) / 33 )
```
## Tokens
```
test/test_008.in:1:1:	ident	"print"	0	(1, 1) -> (1, 6)
test/test_008.in:1:6:	lparen	"("	0	(1, 6) -> (1, 7)
test/test_008.in:1:7:	lparen	"("	0	(1, 7) -> (1, 8)
test/test_008.in:1:8:	number	"34"	34	(1, 8) -> (1, 10)
test/test_008.in:1:10:	add	"+"	0	(1, 10) -> (1, 11)
test/test_008.in:1:11:	number	"35"	35	(1, 11) -> (1, 13)
test/test_008.in:1:13:	rparen	")"	0	(1, 13) -> (1, 14)
test/test_008.in:1:14:	div	"/"	0	(1, 14) -> (1, 15)
test/test_008.in:1:15:	number	"33"	33	(1, 15) -> (1, 17)
test/test_008.in:1:17:	rparen	")"	0	(1, 17) -> (1, 18)
test/test_008.in:2:0:	eof	""	0	(2, 0) -> (2, 0)
```
## AST
```scheme
(module
  (expr2stmt
    (print [void]
      (div [f64]
        (add [i64]
          (number [i64] 34)
          (number [i64] 35)
        )
        (number [i64] 33)
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
	%0 = add i32 34, 35
	%1 = sitofp i32 %0 to double
	%2 = sitofp i32 33 to double
	%3 = fdiv double %1, %2
	%4 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%5 = call i32 (i8*, ...) @printf(i8* %4, double %3)
	ret i32 0
}

```
## Run
```bash
2.090909
```