# test/test_024.in
```
print(1#2)
print(3/2#6/4)

```
## Tokens
```tsv
test/test_024.in:1:1:	ident	"print"	0	(1, 1) -> (1, 6)
test/test_024.in:1:6:	lparen	"("	0	(1, 6) -> (1, 7)
test/test_024.in:1:7:	number	"1"	1	(1, 7) -> (1, 8)
test/test_024.in:1:8:	ne	"#"	0	(1, 8) -> (1, 9)
test/test_024.in:1:9:	number	"2"	2	(1, 9) -> (1, 10)
test/test_024.in:1:10:	rparen	")"	0	(1, 10) -> (1, 11)
test/test_024.in:2:0:	ident	"print"	0	(2, 0) -> (2, 5)
test/test_024.in:2:5:	lparen	"("	0	(2, 5) -> (2, 6)
test/test_024.in:2:6:	number	"3"	3	(2, 6) -> (2, 7)
test/test_024.in:2:7:	slash	"/"	0	(2, 7) -> (2, 8)
test/test_024.in:2:8:	number	"2"	2	(2, 8) -> (2, 9)
test/test_024.in:2:9:	ne	"#"	0	(2, 9) -> (2, 10)
test/test_024.in:2:10:	number	"6"	6	(2, 10) -> (2, 11)
test/test_024.in:2:11:	slash	"/"	0	(2, 11) -> (2, 12)
test/test_024.in:2:12:	number	"4"	4	(2, 12) -> (2, 13)
test/test_024.in:2:13:	rparen	")"	0	(2, 13) -> (2, 14)
test/test_024.in:4:0:	eof	""	0	(4, 0) -> (4, 0)
```
## AST
```scheme
(module
  (expr2stmt
    (print [void]
      (ne [boolean]
        (number [i64] 1)
        (number [i64] 2)
      )
    )
  )
  (expr2stmt
    (print [void]
      (ne [boolean]
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
	%0 = icmp ne i32 1, 2
	%1 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, i1 %0)
	%3 = sitofp i32 3 to double
	%4 = sitofp i32 2 to double
	%5 = fdiv double %3, %4
	%6 = sitofp i32 6 to double
	%7 = sitofp i32 4 to double
	%8 = fdiv double %6, %7
	%9 = fcmp une double %5, %8
	%10 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%11 = call i32 (i8*, ...) @printf(i8* %10, i1 %9)
	ret i32 0
}

```
## Run
```bash
1
0
```
