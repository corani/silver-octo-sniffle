# test/test_022.in
```
print(~TRUE)
print(~FALSE)
```
## Tokens
```tsv
test/test_022.in:1:1:	ident	"print"	false	0	0.000000	(1, 1) -> (1, 6)
test/test_022.in:1:6:	lparen	"("	false	0	0.000000	(1, 6) -> (1, 7)
test/test_022.in:1:7:	not	"~"	false	0	0.000000	(1, 7) -> (1, 8)
test/test_022.in:1:8:	boolean	"TRUE"	true	0	0.000000	(1, 8) -> (1, 12)
test/test_022.in:1:12:	rparen	")"	false	0	0.000000	(1, 12) -> (1, 13)
test/test_022.in:2:0:	ident	"print"	false	0	0.000000	(2, 0) -> (2, 5)
test/test_022.in:2:5:	lparen	"("	false	0	0.000000	(2, 5) -> (2, 6)
test/test_022.in:2:6:	not	"~"	false	0	0.000000	(2, 6) -> (2, 7)
test/test_022.in:2:7:	boolean	"FALSE"	false	0	0.000000	(2, 7) -> (2, 12)
test/test_022.in:2:12:	rparen	")"	false	0	0.000000	(2, 12) -> (2, 13)
test/test_022.in:3:0:	eof	""	false	0	0.000000	(3, 0) -> (3, 0)
```
## AST
```scheme
(module
  (expr2stmt
    (print [void]
      (not [boolean]
        #true
      )
    )
  )
  (expr2stmt
    (print [void]
      (not [boolean]
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
	%0 = icmp eq i1 true, false
	%1 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, i1 %0)
	%3 = icmp eq i1 false, false
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
