# test/test_002.in
```
print( 34 + 35 + 42000 )
```
## Tokens
```tsv
test/test_002.in:1:1:	ident	"print"	0	(1, 1) -> (1, 6)
test/test_002.in:1:6:	lparen	"("	0	(1, 6) -> (1, 7)
test/test_002.in:1:8:	integer	"34"	34	(1, 8) -> (1, 10)
test/test_002.in:1:11:	plus	"+"	0	(1, 11) -> (1, 12)
test/test_002.in:1:13:	integer	"35"	35	(1, 13) -> (1, 15)
test/test_002.in:1:16:	plus	"+"	0	(1, 16) -> (1, 17)
test/test_002.in:1:18:	integer	"42000"	42000	(1, 18) -> (1, 23)
test/test_002.in:1:24:	rparen	")"	0	(1, 24) -> (1, 25)
test/test_002.in:2:0:	eof	""	0	(2, 0) -> (2, 0)
```
## AST
```scheme
(module
  (expr2stmt
    (print [void]
      (plus [i64]
        (plus [i64]
          (number [i64] 34)
          (number [i64] 35)
        )
        (number [i64] 42000)
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
	%0 = add i32 34, 35
	%1 = add i32 %0, 42000
	%2 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%3 = call i32 (i8*, ...) @printf(i8* %2, i32 %1)
	ret i32 0
}

```
## Run
```bash
42069
```
