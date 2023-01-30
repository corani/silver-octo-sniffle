# test/test_025.in
```
print(TRUE & FALSE)
print(TRUE OR FALSE)
```
## Tokens
```tsv
test/test_025.in:1:1:	ident	"print"	0	(1, 1) -> (1, 6)
test/test_025.in:1:6:	lparen	"("	0	(1, 6) -> (1, 7)
test/test_025.in:1:7:	boolean	"TRUE"	1	(1, 7) -> (1, 11)
test/test_025.in:1:12:	and	"&"	0	(1, 12) -> (1, 13)
test/test_025.in:1:14:	boolean	"FALSE"	0	(1, 14) -> (1, 19)
test/test_025.in:1:19:	rparen	")"	0	(1, 19) -> (1, 20)
test/test_025.in:2:0:	ident	"print"	0	(2, 0) -> (2, 5)
test/test_025.in:2:5:	lparen	"("	0	(2, 5) -> (2, 6)
test/test_025.in:2:6:	boolean	"TRUE"	1	(2, 6) -> (2, 10)
test/test_025.in:2:11:	or	"OR"	0	(2, 11) -> (2, 13)
test/test_025.in:2:14:	boolean	"FALSE"	0	(2, 14) -> (2, 19)
test/test_025.in:2:19:	rparen	")"	0	(2, 19) -> (2, 20)
test/test_025.in:3:0:	eof	""	0	(3, 0) -> (3, 0)
```
## AST
```scheme
(module
  (expr2stmt
    (print [void]
      (and [boolean]
        #true
        #false
      )
    )
  )
  (expr2stmt
    (print [void]
      (or [boolean]
        #true
        #false
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
	%0 = and i1 true, false
	%1 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, i1 %0)
	%3 = or i1 true, false
	%4 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%5 = call i32 (i8*, ...) @printf(i8* %4, i1 %3)
	ret i32 0
}

```
## Run
```bash
0
1
```
