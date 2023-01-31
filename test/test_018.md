# test/test_018.md
## Source
```
(*
  multi-line comments
  should work
*)
print(1)
```
## Tokens
```tsv
test/test_018.md:5:0:	ident	"print"	false	0	0.000000	(5, 0) -> (5, 5)
test/test_018.md:5:5:	lparen	"("	false	0	0.000000	(5, 5) -> (5, 6)
test/test_018.md:5:6:	integer	"1"	false	1	0.000000	(5, 6) -> (5, 7)
test/test_018.md:5:7:	rparen	")"	false	0	0.000000	(5, 7) -> (5, 8)
test/test_018.md:6:0:	eof	""	false	0	0.000000	(6, 0) -> (6, 0)
```
## AST
```scheme
(module
  (expr2stmt
    (print [void]
      (number [i64] 1)
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
	%1 = call i32 (i8*, ...) @printf(i8* %0, i32 1)
	ret i32 0
}

```
## Run
```bash
1
```
