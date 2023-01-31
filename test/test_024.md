# test/test_024.md
## Source
```
print(1#2);
print(3/2#6/4)
```
## Tokens
```tsv
test/test_024.md:1:1:	ident	"print"	false	0	0.000000	(1, 1) -> (1, 6)
test/test_024.md:1:6:	lparen	"("	false	0	0.000000	(1, 6) -> (1, 7)
test/test_024.md:1:7:	integer	"1"	false	1	0.000000	(1, 7) -> (1, 8)
test/test_024.md:1:8:	ne	"#"	false	0	0.000000	(1, 8) -> (1, 9)
test/test_024.md:1:9:	integer	"2"	false	2	0.000000	(1, 9) -> (1, 10)
test/test_024.md:1:10:	rparen	")"	false	0	0.000000	(1, 10) -> (1, 11)
test/test_024.md:1:11:	semicolon	";"	false	0	0.000000	(1, 11) -> (1, 12)
test/test_024.md:2:0:	ident	"print"	false	0	0.000000	(2, 0) -> (2, 5)
test/test_024.md:2:5:	lparen	"("	false	0	0.000000	(2, 5) -> (2, 6)
test/test_024.md:2:6:	integer	"3"	false	3	0.000000	(2, 6) -> (2, 7)
test/test_024.md:2:7:	slash	"/"	false	0	0.000000	(2, 7) -> (2, 8)
test/test_024.md:2:8:	integer	"2"	false	2	0.000000	(2, 8) -> (2, 9)
test/test_024.md:2:9:	ne	"#"	false	0	0.000000	(2, 9) -> (2, 10)
test/test_024.md:2:10:	integer	"6"	false	6	0.000000	(2, 10) -> (2, 11)
test/test_024.md:2:11:	slash	"/"	false	0	0.000000	(2, 11) -> (2, 12)
test/test_024.md:2:12:	integer	"4"	false	4	0.000000	(2, 12) -> (2, 13)
test/test_024.md:2:13:	rparen	")"	false	0	0.000000	(2, 13) -> (2, 14)
test/test_024.md:3:0:	eof	""	false	0	0.000000	(3, 0) -> (3, 0)
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
@0 = global [5 x i8] c"TRUE\00"
@1 = global [6 x i8] c"FALSE\00"

declare i32 @puts(i8* %str)

declare i32 @rand()

declare i32 @sprintf(i8* %buf, i8* %format, ...)

declare i32 @printf(i8* %format, ...)

define i32 @main() {
entry:
	%0 = icmp ne i32 1, 2
	br i1 %0, label %1, label %3

1:
	%2 = getelementptr [5 x i8], [5 x i8]* @0, i32 0, i32 0
	br label %5

3:
	%4 = getelementptr [6 x i8], [6 x i8]* @1, i32 0, i32 0
	br label %5

5:
	%6 = phi i8* [ %2, %1 ], [ %4, %3 ]
	%7 = call i32 @puts(i8* %6)
	%8 = sitofp i32 3 to double
	%9 = sitofp i32 2 to double
	%10 = fdiv double %8, %9
	%11 = sitofp i32 6 to double
	%12 = sitofp i32 4 to double
	%13 = fdiv double %11, %12
	%14 = fcmp une double %10, %13
	br i1 %14, label %15, label %17

15:
	%16 = getelementptr [5 x i8], [5 x i8]* @0, i32 0, i32 0
	br label %19

17:
	%18 = getelementptr [6 x i8], [6 x i8]* @1, i32 0, i32 0
	br label %19

19:
	%20 = phi i8* [ %16, %15 ], [ %18, %17 ]
	%21 = call i32 @puts(i8* %20)
	ret i32 0
}

```
## Run
```bash
TRUE
FALSE
```
