# test/test_027.in
```
print(1>2)
print(2/2>2/1)
print(2>=2)
print(2/1>=2/1)
```
## Tokens
```tsv
test/test_027.in:1:1:	ident	"print"	0	(1, 1) -> (1, 6)
test/test_027.in:1:6:	lparen	"("	0	(1, 6) -> (1, 7)
test/test_027.in:1:7:	integer	"1"	1	(1, 7) -> (1, 8)
test/test_027.in:1:8:	gt	">"	0	(1, 8) -> (1, 9)
test/test_027.in:1:9:	integer	"2"	2	(1, 9) -> (1, 10)
test/test_027.in:1:10:	rparen	")"	0	(1, 10) -> (1, 11)
test/test_027.in:2:0:	ident	"print"	0	(2, 0) -> (2, 5)
test/test_027.in:2:5:	lparen	"("	0	(2, 5) -> (2, 6)
test/test_027.in:2:6:	integer	"2"	2	(2, 6) -> (2, 7)
test/test_027.in:2:7:	slash	"/"	0	(2, 7) -> (2, 8)
test/test_027.in:2:8:	integer	"2"	2	(2, 8) -> (2, 9)
test/test_027.in:2:9:	gt	">"	0	(2, 9) -> (2, 10)
test/test_027.in:2:10:	integer	"2"	2	(2, 10) -> (2, 11)
test/test_027.in:2:11:	slash	"/"	0	(2, 11) -> (2, 12)
test/test_027.in:2:12:	integer	"1"	1	(2, 12) -> (2, 13)
test/test_027.in:2:13:	rparen	")"	0	(2, 13) -> (2, 14)
test/test_027.in:3:0:	ident	"print"	0	(3, 0) -> (3, 5)
test/test_027.in:3:5:	lparen	"("	0	(3, 5) -> (3, 6)
test/test_027.in:3:6:	integer	"2"	2	(3, 6) -> (3, 7)
test/test_027.in:3:7:	ge	">="	0	(3, 7) -> (3, 9)
test/test_027.in:3:9:	integer	"2"	2	(3, 9) -> (3, 10)
test/test_027.in:3:10:	rparen	")"	0	(3, 10) -> (3, 11)
test/test_027.in:4:0:	ident	"print"	0	(4, 0) -> (4, 5)
test/test_027.in:4:5:	lparen	"("	0	(4, 5) -> (4, 6)
test/test_027.in:4:6:	integer	"2"	2	(4, 6) -> (4, 7)
test/test_027.in:4:7:	slash	"/"	0	(4, 7) -> (4, 8)
test/test_027.in:4:8:	integer	"1"	1	(4, 8) -> (4, 9)
test/test_027.in:4:9:	ge	">="	0	(4, 9) -> (4, 11)
test/test_027.in:4:11:	integer	"2"	2	(4, 11) -> (4, 12)
test/test_027.in:4:12:	slash	"/"	0	(4, 12) -> (4, 13)
test/test_027.in:4:13:	integer	"1"	1	(4, 13) -> (4, 14)
test/test_027.in:4:14:	rparen	")"	0	(4, 14) -> (4, 15)
test/test_027.in:5:0:	eof	""	0	(5, 0) -> (5, 0)
```
## AST
```scheme
(module
  (expr2stmt
    (print [void]
      (gt [boolean]
        (number [i64] 1)
        (number [i64] 2)
      )
    )
  )
  (expr2stmt
    (print [void]
      (gt [boolean]
        (slash [f64]
          (number [i64] 2)
          (number [i64] 2)
        )
        (slash [f64]
          (number [i64] 2)
          (number [i64] 1)
        )
      )
    )
  )
  (expr2stmt
    (print [void]
      (ge [boolean]
        (number [i64] 2)
        (number [i64] 2)
      )
    )
  )
  (expr2stmt
    (print [void]
      (ge [boolean]
        (slash [f64]
          (number [i64] 2)
          (number [i64] 1)
        )
        (slash [f64]
          (number [i64] 2)
          (number [i64] 1)
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
	%0 = icmp sgt i32 1, 2
	%1 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, i1 %0)
	%3 = sitofp i32 2 to double
	%4 = sitofp i32 2 to double
	%5 = fdiv double %3, %4
	%6 = sitofp i32 2 to double
	%7 = sitofp i32 1 to double
	%8 = fdiv double %6, %7
	%9 = fcmp ugt double %5, %8
	%10 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%11 = call i32 (i8*, ...) @printf(i8* %10, i1 %9)
	%12 = icmp sge i32 2, 2
	%13 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%14 = call i32 (i8*, ...) @printf(i8* %13, i1 %12)
	%15 = sitofp i32 2 to double
	%16 = sitofp i32 1 to double
	%17 = fdiv double %15, %16
	%18 = sitofp i32 2 to double
	%19 = sitofp i32 1 to double
	%20 = fdiv double %18, %19
	%21 = fcmp uge double %17, %20
	%22 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%23 = call i32 (i8*, ...) @printf(i8* %22, i1 %21)
	ret i32 0
}

```
## Run
```bash
0
0
1
1
```
