# test/test_006.md
## Source
```
print( (34 + 35) + 42000 )
```
## Tokens
```tsv
test/test_006.md:1:1:	ident	"print"	false	0	0.000000	(1, 1) -> (1, 6)
test/test_006.md:1:6:	lparen	"("	false	0	0.000000	(1, 6) -> (1, 7)
test/test_006.md:1:8:	lparen	"("	false	0	0.000000	(1, 8) -> (1, 9)
test/test_006.md:1:9:	integer	"34"	false	34	0.000000	(1, 9) -> (1, 11)
test/test_006.md:1:12:	plus	"+"	false	0	0.000000	(1, 12) -> (1, 13)
test/test_006.md:1:14:	integer	"35"	false	35	0.000000	(1, 14) -> (1, 16)
test/test_006.md:1:16:	rparen	")"	false	0	0.000000	(1, 16) -> (1, 17)
test/test_006.md:1:18:	plus	"+"	false	0	0.000000	(1, 18) -> (1, 19)
test/test_006.md:1:20:	integer	"42000"	false	42000	0.000000	(1, 20) -> (1, 25)
test/test_006.md:1:26:	rparen	")"	false	0	0.000000	(1, 26) -> (1, 27)
test/test_006.md:2:0:	eof	""	false	0	0.000000	(2, 0) -> (2, 0)
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
