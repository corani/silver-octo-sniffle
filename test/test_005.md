# test/test_005.in
```
print( -34 + 35 )
```
## Tokens
```tsv
test/test_005.in:1:1:	ident	"print"	0	(1, 1) -> (1, 6)
test/test_005.in:1:6:	lparen	"("	0	(1, 6) -> (1, 7)
test/test_005.in:1:8:	minus	"-"	0	(1, 8) -> (1, 9)
test/test_005.in:1:9:	number	"34"	34	(1, 9) -> (1, 11)
test/test_005.in:1:12:	plus	"+"	0	(1, 12) -> (1, 13)
test/test_005.in:1:14:	number	"35"	35	(1, 14) -> (1, 16)
test/test_005.in:1:17:	rparen	")"	0	(1, 17) -> (1, 18)
test/test_005.in:2:0:	eof	""	0	(2, 0) -> (2, 0)
```
## AST
```scheme
(module
  (expr2stmt
    (print [void]
      (plus [i64]
        (minus [i64]
          (number [i64] 0)
          (number [i64] 34)
        )
        (number [i64] 35)
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
	%0 = sub i32 0, 34
	%1 = add i32 %0, 35
	%2 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%3 = call i32 (i8*, ...) @printf(i8* %2, i32 %1)
	ret i32 0
}

```
## Run
```bash
1
```
