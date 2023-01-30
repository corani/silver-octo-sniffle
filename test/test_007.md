# test/test_007.in
```
print( (34 + 35) * 33 )
```
## Tokens
```tsv
test/test_007.in:1:1:	ident	"print"	0	(1, 1) -> (1, 6)
test/test_007.in:1:6:	lparen	"("	0	(1, 6) -> (1, 7)
test/test_007.in:1:8:	lparen	"("	0	(1, 8) -> (1, 9)
test/test_007.in:1:9:	number	"34"	34	(1, 9) -> (1, 11)
test/test_007.in:1:12:	plus	"+"	0	(1, 12) -> (1, 13)
test/test_007.in:1:14:	number	"35"	35	(1, 14) -> (1, 16)
test/test_007.in:1:16:	rparen	")"	0	(1, 16) -> (1, 17)
test/test_007.in:1:18:	asterisk	"*"	0	(1, 18) -> (1, 19)
test/test_007.in:1:20:	number	"33"	33	(1, 20) -> (1, 22)
test/test_007.in:1:23:	rparen	")"	0	(1, 23) -> (1, 24)
test/test_007.in:2:0:	eof	""	0	(2, 0) -> (2, 0)
```
## AST
```scheme
(module
  (expr2stmt
    (print [void]
      (asterisk [i64]
        (plus [i64]
          (number [i64] 34)
          (number [i64] 35)
        )
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
	%0 = add i32 34, 35
	%1 = mul i32 %0, 33
	%2 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%3 = call i32 (i8*, ...) @printf(i8* %2, i32 %1)
	ret i32 0
}

```
## Run
```bash
2277
```
