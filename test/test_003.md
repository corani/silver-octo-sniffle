# test/test_003.in
```
print(66-33)
```
## Tokens
```tsv
test/test_003.in:1:1:	ident	"print"	false	0	0.000000	(1, 1) -> (1, 6)
test/test_003.in:1:6:	lparen	"("	false	0	0.000000	(1, 6) -> (1, 7)
test/test_003.in:1:7:	integer	"66"	false	66	0.000000	(1, 7) -> (1, 9)
test/test_003.in:1:9:	minus	"-"	false	0	0.000000	(1, 9) -> (1, 10)
test/test_003.in:1:10:	integer	"33"	false	33	0.000000	(1, 10) -> (1, 12)
test/test_003.in:1:12:	rparen	")"	false	0	0.000000	(1, 12) -> (1, 13)
test/test_003.in:2:0:	eof	""	false	0	0.000000	(2, 0) -> (2, 0)
```
## AST
```scheme
(module
  (expr2stmt
    (print [void]
      (minus [i64]
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
