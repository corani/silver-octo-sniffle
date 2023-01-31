# test/test_010.md
## Source
```
print(12,34)
```
## Tokens
```tsv
test/test_010.md:1:1:	ident	"print"	false	0	0.000000	(1, 1) -> (1, 6)
test/test_010.md:1:6:	lparen	"("	false	0	0.000000	(1, 6) -> (1, 7)
test/test_010.md:1:7:	integer	"12"	false	12	0.000000	(1, 7) -> (1, 9)
test/test_010.md:1:9:	comma	","	false	0	0.000000	(1, 9) -> (1, 10)
test/test_010.md:1:10:	integer	"34"	false	34	0.000000	(1, 10) -> (1, 12)
test/test_010.md:1:12:	rparen	")"	false	0	0.000000	(1, 12) -> (1, 13)
test/test_010.md:2:0:	eof	""	false	0	0.000000	(2, 0) -> (2, 0)
```
## AST
```scheme
(module
  (expr2stmt
    (print [void]
      (number [i64] 12)
      (number [i64] 34)
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
	%1 = call i32 (i8*, ...) @printf(i8* %0, i32 12)
	%2 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%3 = call i32 (i8*, ...) @printf(i8* %2, i32 34)
	ret i32 0
}

```
## Run
```bash
12
34
```
