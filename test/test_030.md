# test/test_030.in
```
print(33H, 0AEH)
```
## Tokens
```tsv
test/test_030.in:1:1:	ident	"print"	false	0	0.000000	(1, 1) -> (1, 6)
test/test_030.in:1:6:	lparen	"("	false	0	0.000000	(1, 6) -> (1, 7)
test/test_030.in:1:7:	integer	"33H"	false	51	0.000000	(1, 7) -> (1, 10)
test/test_030.in:1:10:	comma	","	false	0	0.000000	(1, 10) -> (1, 11)
test/test_030.in:1:12:	integer	"0AEH"	false	174	0.000000	(1, 12) -> (1, 16)
test/test_030.in:1:16:	rparen	")"	false	0	0.000000	(1, 16) -> (1, 17)
test/test_030.in:2:0:	eof	""	false	0	0.000000	(2, 0) -> (2, 0)
```
## AST
```scheme
(module
  (expr2stmt
    (print [void]
      (number [i64] 51)
      (number [i64] 174)
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
	%1 = call i32 (i8*, ...) @printf(i8* %0, i32 51)
	%2 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%3 = call i32 (i8*, ...) @printf(i8* %2, i32 174)
	ret i32 0
}

```
## Run
```bash
51
174
```
