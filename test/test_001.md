# test/test_001.in
```
print(42)
```
## Tokens
```
test/test_001.in:1:1:	ident	"print"	0	(1, 1) -> (1, 6)
test/test_001.in:1:6:	lparen	"("	0	(1, 6) -> (1, 7)
test/test_001.in:1:7:	number	"42"	42	(1, 7) -> (1, 9)
test/test_001.in:1:9:	rparen	")"	0	(1, 9) -> (1, 10)
test/test_001.in:2:0:	eof	""	0	(2, 0) -> (2, 0)
```
## AST
```scheme
(module
  (expr2stmt
    (print [void]
      (number [i64] 42)
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
	%0 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%1 = call i32 (i8*, ...) @printf(i8* %0, i32 42)
	ret i32 0
}

```
## Run
```bash
42
```