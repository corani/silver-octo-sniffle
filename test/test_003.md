# test/test_003.in
```
print(66-33)
```
## Tokens
```
test/test_003.in:1:1:	ident	"print"	0	(1, 1) -> (1, 6)
test/test_003.in:1:6:	lparen	"("	0	(1, 6) -> (1, 7)
test/test_003.in:1:7:	number	"66"	66	(1, 7) -> (1, 9)
test/test_003.in:1:9:	sub	"-"	0	(1, 9) -> (1, 10)
test/test_003.in:1:10:	number	"33"	33	(1, 10) -> (1, 12)
test/test_003.in:1:12:	rparen	")"	0	(1, 12) -> (1, 13)
test/test_003.in:2:0:	eof	""	0	(2, 0) -> (2, 0)
```
## AST
```scheme
(module
  (expr2stmt
    (print [void]
      (sub [i64]
        (number [i64] 66)
        (number [i64] 33)
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
	%0 = sub i32 66, 33
	%1 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, i32 %0)
	ret i32 0
}

```
## Run
```bash
33
```
