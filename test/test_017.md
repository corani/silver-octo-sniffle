# test/test_017.in
```
(* this is a comment *)
print(1)
```
## Tokens
```tsv
test/test_017.in:2:0:	ident	"print"	0	(2, 0) -> (2, 5)
test/test_017.in:2:5:	lparen	"("	0	(2, 5) -> (2, 6)
test/test_017.in:2:6:	integer	"1"	1	(2, 6) -> (2, 7)
test/test_017.in:2:7:	rparen	")"	0	(2, 7) -> (2, 8)
test/test_017.in:3:0:	eof	""	0	(3, 0) -> (3, 0)
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
