# test/test_029.in
```
print(69+0.420)
```
## Tokens
```tsv
test/test_029.in:1:1:	ident	"print"	0	(1, 1) -> (1, 6)
test/test_029.in:1:6:	lparen	"("	0	(1, 6) -> (1, 7)
test/test_029.in:1:7:	integer	"69"	69	(1, 7) -> (1, 9)
test/test_029.in:1:9:	plus	"+"	0	(1, 9) -> (1, 10)
test/test_029.in:1:10:	real	"0.420"	0	(1, 10) -> (1, 15)
test/test_029.in:1:15:	rparen	")"	0	(1, 15) -> (1, 16)
test/test_029.in:2:0:	eof	""	0	(2, 0) -> (2, 0)
```
## AST
```scheme
(module
  (expr2stmt
    (print [void]
      (plus [f64]
        (number [i64] 69)
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
	%0 = sitofp i32 69 to double
	%1 = fadd double %0, 0x3FDAE147AE147AE1
	%2 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%3 = call i32 (i8*, ...) @printf(i8* %2, double %1)
	ret i32 0
}

```
## Run
```bash
69.420000
```
