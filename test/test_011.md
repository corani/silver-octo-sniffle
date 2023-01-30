# test/test_011.in
```
print(11 DIV 3)
```
## Tokens
```tsv
test/test_011.in:1:1:	ident	"print"	0	(1, 1) -> (1, 6)
test/test_011.in:1:6:	lparen	"("	0	(1, 6) -> (1, 7)
test/test_011.in:1:7:	number	"11"	11	(1, 7) -> (1, 9)
test/test_011.in:1:10:	idiv	"DIV"	0	(1, 10) -> (1, 13)
test/test_011.in:1:14:	number	"3"	3	(1, 14) -> (1, 15)
test/test_011.in:1:15:	rparen	")"	0	(1, 15) -> (1, 16)
test/test_011.in:2:0:	eof	""	0	(2, 0) -> (2, 0)
```
## AST
```scheme
(module
  (expr2stmt
    (print [void]
      (idiv [i64]
        (number [i64] 11)
        (number [i64] 3)
      )
    )
  )
)
```
## IR
```llvm
@0 = global [4 x i8] c"%d\0A\00"

declare i32 @puts(i8* %str)

declare i32 @rand()

declare i32 @sprintf(i8* %buf, i8* %format, ...)

declare i32 @printf(i8* %format, ...)

define i32 @main() {
entry:
	%0 = sdiv i32 11, 3
	%1 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, i32 %0)
	ret i32 0
}

```
## Run
```bash
3
```
