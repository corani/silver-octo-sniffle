# test/test_021.in
```
print(TRUE, FALSE)
```
## Tokens
```tsv
test/test_021.in:1:1:	ident	"print"	0	(1, 1) -> (1, 6)
test/test_021.in:1:6:	lparen	"("	0	(1, 6) -> (1, 7)
test/test_021.in:1:7:	boolean	"TRUE"	1	(1, 7) -> (1, 11)
test/test_021.in:1:11:	comma	","	0	(1, 11) -> (1, 12)
test/test_021.in:1:13:	boolean	"FALSE"	0	(1, 13) -> (1, 18)
test/test_021.in:1:18:	rparen	")"	0	(1, 18) -> (1, 19)
test/test_021.in:2:0:	eof	""	0	(2, 0) -> (2, 0)
```
## AST
```scheme
(module
  (expr2stmt
    (print [void]
      #true
      #false
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
	%1 = call i32 (i8*, ...) @printf(i8* %0, i1 true)
	%2 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%3 = call i32 (i8*, ...) @printf(i8* %2, i1 false)
	ret i32 0
}

```
## Run
```bash
1
0
```
