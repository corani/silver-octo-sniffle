# test/test_032.in
```
print(0.69E+2+420.E-3)
```
## Tokens
```tsv
test/test_032.in:1:1:	ident	"print"	0	(1, 1) -> (1, 6)
test/test_032.in:1:6:	lparen	"("	0	(1, 6) -> (1, 7)
test/test_032.in:1:7:	real	"0.69E+2"	0	(1, 7) -> (1, 14)
test/test_032.in:1:14:	plus	"+"	0	(1, 14) -> (1, 15)
test/test_032.in:1:15:	real	"420.E-3"	0	(1, 15) -> (1, 22)
test/test_032.in:1:22:	rparen	")"	0	(1, 22) -> (1, 23)
test/test_032.in:2:0:	eof	""	0	(2, 0) -> (2, 0)
```
## AST
```scheme
(module
  (expr2stmt
    (print [void]
      (plus [f64]
        (number [f64] 69.000000)
        (number [f64] 0.420000)
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
	%0 = fadd double 69.0, 0x3FDAE147AE147AE1
	%1 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, double %0)
	ret i32 0
}

```
## Run
```bash
69.420000
```
