# test/test_031.in
```
print(48656C6C6F2C20776F726C6421X)
```
## Tokens
```tsv
test/test_031.in:1:1:	ident	"print"	0	(1, 1) -> (1, 6)
test/test_031.in:1:6:	lparen	"("	0	(1, 6) -> (1, 7)
test/test_031.in:1:7:	string	"Hello, world!"	0	(1, 7) -> (1, 34)
test/test_031.in:1:34:	rparen	")"	0	(1, 34) -> (1, 35)
test/test_031.in:2:0:	eof	""	0	(2, 0) -> (2, 0)
```
## AST
```scheme
(module
  (expr2stmt
    (print [void]
      (string "Hello, world!")
    )
  )
)
```
## IR
```llvm
@0 = global [13 x i8] c"Hello, world!"

declare i32 @puts(i8* %str)

declare i32 @rand()

declare i32 @sprintf(i8* %buf, i8* %format, ...)

declare i32 @printf(i8* %format, ...)

define i32 @main() {
entry:
	%0 = getelementptr [13 x i8], [13 x i8]* @0, i32 0, i32 0
	%1 = call i32 @puts(i8* %0)
	ret i32 0
}

```
## Run
```bash
Hello, world!
```
