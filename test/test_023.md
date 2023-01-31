# test/test_023.in
```
print(1=2)
print(3/2=6/4)
```
## Tokens
```tsv
test/test_023.in:1:1:	ident	"print"	false	0	0.000000	(1, 1) -> (1, 6)
test/test_023.in:1:6:	lparen	"("	false	0	0.000000	(1, 6) -> (1, 7)
test/test_023.in:1:7:	integer	"1"	false	1	0.000000	(1, 7) -> (1, 8)
test/test_023.in:1:8:	eq	"="	false	0	0.000000	(1, 8) -> (1, 9)
test/test_023.in:1:9:	integer	"2"	false	2	0.000000	(1, 9) -> (1, 10)
test/test_023.in:1:10:	rparen	")"	false	0	0.000000	(1, 10) -> (1, 11)
test/test_023.in:2:0:	ident	"print"	false	0	0.000000	(2, 0) -> (2, 5)
test/test_023.in:2:5:	lparen	"("	false	0	0.000000	(2, 5) -> (2, 6)
test/test_023.in:2:6:	integer	"3"	false	3	0.000000	(2, 6) -> (2, 7)
test/test_023.in:2:7:	slash	"/"	false	0	0.000000	(2, 7) -> (2, 8)
test/test_023.in:2:8:	integer	"2"	false	2	0.000000	(2, 8) -> (2, 9)
test/test_023.in:2:9:	eq	"="	false	0	0.000000	(2, 9) -> (2, 10)
test/test_023.in:2:10:	integer	"6"	false	6	0.000000	(2, 10) -> (2, 11)
test/test_023.in:2:11:	slash	"/"	false	0	0.000000	(2, 11) -> (2, 12)
test/test_023.in:2:12:	integer	"4"	false	4	0.000000	(2, 12) -> (2, 13)
test/test_023.in:2:13:	rparen	")"	false	0	0.000000	(2, 13) -> (2, 14)
test/test_023.in:3:0:	eof	""	false	0	0.000000	(3, 0) -> (3, 0)
```
## AST
```scheme
(module
  (expr2stmt
    (print [void]
      (eq [boolean]
        (number [i64] 1)
        (number [i64] 2)
      )
    )
  )
  (expr2stmt
    (print [void]
      (eq [boolean]
        (slash [f64]
          (number [i64] 3)
          (number [i64] 2)
        )
        (slash [f64]
          (number [i64] 6)
          (number [i64] 4)
        )
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
	%0 = icmp eq i32 1, 2
	%1 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, i1 %0)
	%3 = sitofp i32 3 to double
	%4 = sitofp i32 2 to double
	%5 = fdiv double %3, %4
	%6 = sitofp i32 6 to double
	%7 = sitofp i32 4 to double
	%8 = fdiv double %6, %7
	%9 = fcmp ueq double %5, %8
	%10 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%11 = call i32 (i8*, ...) @printf(i8* %10, i1 %9)
	ret i32 0
}

```
## Run
```bash
0
1
```
