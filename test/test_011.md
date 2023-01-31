# test/test_011.md
## Source
```
print(11 DIV 3)
```
## Tokens
```tsv
test/test_011.md:1:1:	ident	"print"	false	0	0.000000	(1, 1) -> (1, 6)
test/test_011.md:1:6:	lparen	"("	false	0	0.000000	(1, 6) -> (1, 7)
test/test_011.md:1:7:	integer	"11"	false	11	0.000000	(1, 7) -> (1, 9)
test/test_011.md:1:10:	div	"DIV"	false	0	0.000000	(1, 10) -> (1, 13)
test/test_011.md:1:14:	integer	"3"	false	3	0.000000	(1, 14) -> (1, 15)
test/test_011.md:1:15:	rparen	")"	false	0	0.000000	(1, 15) -> (1, 16)
test/test_011.md:2:0:	eof	""	false	0	0.000000	(2, 0) -> (2, 0)
```
## AST
```scheme
(module
  (expr2stmt
    (print [void]
      (div [i64]
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
