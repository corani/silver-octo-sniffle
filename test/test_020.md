# test/test_020.in
```
print("
    multi-line
    strings?
")
```
## Tokens
```tsv
test/test_020.in:1:1:	ident	"print"	0	(1, 1) -> (1, 6)
test/test_020.in:1:6:	lparen	"("	0	(1, 6) -> (1, 7)
test/test_020.in:1:7:	string	"\n    multi-line\n    strings?\n"	0	(1, 7) -> (4, 1)
test/test_020.in:4:1:	rparen	")"	0	(4, 1) -> (4, 2)
test/test_020.in:5:0:	eof	""	0	(5, 0) -> (5, 0)
```
## AST
```scheme
(module
  (expr2stmt
    (print [void]
      (string "\n    multi-line\n    strings?\n")
    )
  )
)
```
## IR
```llvm
@0 = global [29 x i8] c"\0A    multi-line\0A    strings?\0A"

declare i32 @puts(i8* %str)

declare i32 @rand()

declare i32 @sprintf(i8* %buf, i8* %format, ...)

declare i32 @printf(i8* %format, ...)

define i32 @main() {
entry:
	%0 = getelementptr [29 x i8], [29 x i8]* @0, i32 0, i32 0
	%1 = call i32 @puts(i8* %0)
	ret i32 0
}

```
## Run
```bash

    multi-line
    strings?

```