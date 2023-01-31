# test/test_028.in
```
print(33.42)
```
## Tokens
```tsv
test/test_028.in:1:1:	ident	"print"	false	0	0.000000	(1, 1) -> (1, 6)
test/test_028.in:1:6:	lparen	"("	false	0	0.000000	(1, 6) -> (1, 7)
test/test_028.in:1:7:	real	"33.42"	false	0	33.420000	(1, 7) -> (1, 12)
test/test_028.in:1:12:	rparen	")"	false	0	0.000000	(1, 12) -> (1, 13)
test/test_028.in:2:0:	eof	""	false	0	0.000000	(2, 0) -> (2, 0)
```
## AST
```scheme
(module
  (expr2stmt
    (print [void]
      (number [f64] 33.420000)
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
	%0 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%1 = call i32 (i8*, ...) @printf(i8* %0, double 0x4040B5C28F5C28F6)
	ret i32 0
}

```
## Run
```bash
33.420000
```
