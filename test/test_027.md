# test/test_027.in
```
print(1>2);
print(2/2>2/1);
print(2>=2);
print(2/1>=2/1)
```
## Tokens
```tsv
test/test_027.in:1:1:	ident	"print"	false	0	0.000000	(1, 1) -> (1, 6)
test/test_027.in:1:6:	lparen	"("	false	0	0.000000	(1, 6) -> (1, 7)
test/test_027.in:1:7:	integer	"1"	false	1	0.000000	(1, 7) -> (1, 8)
test/test_027.in:1:8:	gt	">"	false	0	0.000000	(1, 8) -> (1, 9)
test/test_027.in:1:9:	integer	"2"	false	2	0.000000	(1, 9) -> (1, 10)
test/test_027.in:1:10:	rparen	")"	false	0	0.000000	(1, 10) -> (1, 11)
test/test_027.in:1:11:	semicolon	";"	false	0	0.000000	(1, 11) -> (1, 12)
test/test_027.in:2:0:	ident	"print"	false	0	0.000000	(2, 0) -> (2, 5)
test/test_027.in:2:5:	lparen	"("	false	0	0.000000	(2, 5) -> (2, 6)
test/test_027.in:2:6:	integer	"2"	false	2	0.000000	(2, 6) -> (2, 7)
test/test_027.in:2:7:	slash	"/"	false	0	0.000000	(2, 7) -> (2, 8)
test/test_027.in:2:8:	integer	"2"	false	2	0.000000	(2, 8) -> (2, 9)
test/test_027.in:2:9:	gt	">"	false	0	0.000000	(2, 9) -> (2, 10)
test/test_027.in:2:10:	integer	"2"	false	2	0.000000	(2, 10) -> (2, 11)
test/test_027.in:2:11:	slash	"/"	false	0	0.000000	(2, 11) -> (2, 12)
test/test_027.in:2:12:	integer	"1"	false	1	0.000000	(2, 12) -> (2, 13)
test/test_027.in:2:13:	rparen	")"	false	0	0.000000	(2, 13) -> (2, 14)
test/test_027.in:2:14:	semicolon	";"	false	0	0.000000	(2, 14) -> (2, 15)
test/test_027.in:3:0:	ident	"print"	false	0	0.000000	(3, 0) -> (3, 5)
test/test_027.in:3:5:	lparen	"("	false	0	0.000000	(3, 5) -> (3, 6)
test/test_027.in:3:6:	integer	"2"	false	2	0.000000	(3, 6) -> (3, 7)
test/test_027.in:3:7:	ge	">="	false	0	0.000000	(3, 7) -> (3, 9)
test/test_027.in:3:9:	integer	"2"	false	2	0.000000	(3, 9) -> (3, 10)
test/test_027.in:3:10:	rparen	")"	false	0	0.000000	(3, 10) -> (3, 11)
test/test_027.in:3:11:	semicolon	";"	false	0	0.000000	(3, 11) -> (3, 12)
test/test_027.in:4:0:	ident	"print"	false	0	0.000000	(4, 0) -> (4, 5)
test/test_027.in:4:5:	lparen	"("	false	0	0.000000	(4, 5) -> (4, 6)
test/test_027.in:4:6:	integer	"2"	false	2	0.000000	(4, 6) -> (4, 7)
test/test_027.in:4:7:	slash	"/"	false	0	0.000000	(4, 7) -> (4, 8)
test/test_027.in:4:8:	integer	"1"	false	1	0.000000	(4, 8) -> (4, 9)
test/test_027.in:4:9:	ge	">="	false	0	0.000000	(4, 9) -> (4, 11)
test/test_027.in:4:11:	integer	"2"	false	2	0.000000	(4, 11) -> (4, 12)
test/test_027.in:4:12:	slash	"/"	false	0	0.000000	(4, 12) -> (4, 13)
test/test_027.in:4:13:	integer	"1"	false	1	0.000000	(4, 13) -> (4, 14)
test/test_027.in:4:14:	rparen	")"	false	0	0.000000	(4, 14) -> (4, 15)
test/test_027.in:5:0:	eof	""	false	0	0.000000	(5, 0) -> (5, 0)
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
@0 = global [5 x i8] c"TRUE\00"
@1 = global [6 x i8] c"FALSE\00"

declare i32 @puts(i8* %str)

declare i32 @rand()

declare i32 @sprintf(i8* %buf, i8* %format, ...)

declare i32 @printf(i8* %format, ...)

define i32 @main() {
entry:
	%0 = icmp sgt i32 1, 2
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
	%8 = sitofp i32 2 to double
	%9 = sitofp i32 2 to double
	%10 = fdiv double %8, %9
	%11 = sitofp i32 2 to double
	%12 = sitofp i32 1 to double
	%13 = fdiv double %11, %12
	%14 = fcmp ugt double %10, %13
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
	%22 = icmp sge i32 2, 2
	br i1 %22, label %23, label %25

23:
	%24 = getelementptr [5 x i8], [5 x i8]* @0, i32 0, i32 0
	br label %27

25:
	%26 = getelementptr [6 x i8], [6 x i8]* @1, i32 0, i32 0
	br label %27

27:
	%28 = phi i8* [ %24, %23 ], [ %26, %25 ]
	%29 = call i32 @puts(i8* %28)
	%30 = sitofp i32 2 to double
	%31 = sitofp i32 1 to double
	%32 = fdiv double %30, %31
	%33 = sitofp i32 2 to double
	%34 = sitofp i32 1 to double
	%35 = fdiv double %33, %34
	%36 = fcmp uge double %32, %35
	br i1 %36, label %37, label %39

37:
	%38 = getelementptr [5 x i8], [5 x i8]* @0, i32 0, i32 0
	br label %41

39:
	%40 = getelementptr [6 x i8], [6 x i8]* @1, i32 0, i32 0
	br label %41

41:
	%42 = phi i8* [ %38, %37 ], [ %40, %39 ]
	%43 = call i32 @puts(i8* %42)
	ret i32 0
}

```
## Run
```bash
FALSE
FALSE
TRUE
TRUE
```
