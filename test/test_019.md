# test/test_019.md
## Source
```
print("Hello, world!")
```
## Tokens
```tsv
test/test_019.md:1:1:	ident	"print"	false	0	0.000000	(1, 1) -> (1, 6)
test/test_019.md:1:6:	lparen	"("	false	0	0.000000	(1, 6) -> (1, 7)
test/test_019.md:1:7:	string	"Hello, world!"	false	0	0.000000	(1, 7) -> (1, 22)
test/test_019.md:1:22:	rparen	")"	false	0	0.000000	(1, 22) -> (1, 23)
test/test_019.md:2:0:	eof	""	false	0	0.000000	(2, 0) -> (2, 0)
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
