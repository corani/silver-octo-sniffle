# test/test_003.md
## Source
```pascal
(*
  multi-line comments
  should work
*)
print("
  multi-line strings
  should work
")
```
## Tokens
```tsv
test/test_003.md:5:0:	ident	"print"	false	0	0.000000	(5, 0) -> (5, 5)
test/test_003.md:5:5:	lparen	"("	false	0	0.000000	(5, 5) -> (5, 6)
test/test_003.md:5:6:	string	"\n  multi-line strings\n  should work\n"	false	0	0.000000	(5, 6) -> (8, 1)
test/test_003.md:8:1:	rparen	")"	false	0	0.000000	(8, 1) -> (8, 2)
test/test_003.md:9:0:	eof	""	false	0	0.000000	(9, 0) -> (9, 0)
```
## AST
```scheme
(module ""
  (stmts
    (expr2stmt
      (print [void]
        (string "\n  multi-line strings\n  should work\n")
      )
    )
  )
)
```
## IR
```llvm
@0 = global [36 x i8] c"\0A  multi-line strings\0A  should work\0A"

declare i32 @puts(i8* %str)

declare i32 @rand()

declare i32 @sprintf(i8* %buf, i8* %format, ...)

declare i32 @printf(i8* %format, ...)

define i32 @main() {
entry:
	%0 = getelementptr [36 x i8], [36 x i8]* @0, i32 0, i32 0
	%1 = call i32 @puts(i8* %0)
	ret i32 0
}

```
## Run
```bash

  multi-line strings
  should work

```
