# test/test_032.md
## Source
```
print(0.69E+2+420.E-3)
```
## Tokens
```tsv
test/test_032.md:1:1:	ident	"print"	false	0	0.000000	(1, 1) -> (1, 6)
test/test_032.md:1:6:	lparen	"("	false	0	0.000000	(1, 6) -> (1, 7)
test/test_032.md:1:7:	real	"0.69E+2"	false	0	69.000000	(1, 7) -> (1, 14)
test/test_032.md:1:14:	plus	"+"	false	0	0.000000	(1, 14) -> (1, 15)
test/test_032.md:1:15:	real	"420.E-3"	false	0	0.420000	(1, 15) -> (1, 22)
test/test_032.md:1:22:	rparen	")"	false	0	0.000000	(1, 22) -> (1, 23)
test/test_032.md:2:0:	eof	""	false	0	0.000000	(2, 0) -> (2, 0)
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
