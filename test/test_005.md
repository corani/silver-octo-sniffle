# test/test_005.md
## Source
```pascal
MODULE Comparisons;

BEGIN
    print(1=2);
    print(3.0/2.0=6.0/4.0);
    print(1#2);
    print(3 DIV 2#6 DIV 4);
    print(1<2);
    print(2.0/1.0<1.0/1.0);
    print(1<=1);
    print(1.0/1.0<=1.0/1.0);
    print(1>2);
    print(2.0/2.0>2.0/1.0);
    print(2>=2);
    print(2.0/1.0>=2.0/1.0)
END Comparisons.
```
## Tokens
```tsv
test/test_005.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_005.md:1:8:	ident	"Comparisons"	false	0	0.000000	(1, 8) -> (1, 19)
test/test_005.md:1:19:	semicolon	";"	false	0	0.000000	(1, 19) -> (1, 20)
test/test_005.md:3:0:	begin	"BEGIN"	false	0	0.000000	(3, 0) -> (3, 5)
test/test_005.md:4:4:	ident	"print"	false	0	0.000000	(4, 4) -> (4, 9)
test/test_005.md:4:9:	lparen	"("	false	0	0.000000	(4, 9) -> (4, 10)
test/test_005.md:4:10:	integer	"1"	false	1	0.000000	(4, 10) -> (4, 11)
test/test_005.md:4:11:	eq	"="	false	0	0.000000	(4, 11) -> (4, 12)
test/test_005.md:4:12:	integer	"2"	false	2	0.000000	(4, 12) -> (4, 13)
test/test_005.md:4:13:	rparen	")"	false	0	0.000000	(4, 13) -> (4, 14)
test/test_005.md:4:14:	semicolon	";"	false	0	0.000000	(4, 14) -> (4, 15)
test/test_005.md:5:4:	ident	"print"	false	0	0.000000	(5, 4) -> (5, 9)
test/test_005.md:5:9:	lparen	"("	false	0	0.000000	(5, 9) -> (5, 10)
test/test_005.md:5:10:	real	"3.0"	false	0	3.000000	(5, 10) -> (5, 13)
test/test_005.md:5:13:	slash	"/"	false	0	0.000000	(5, 13) -> (5, 14)
test/test_005.md:5:14:	real	"2.0"	false	0	2.000000	(5, 14) -> (5, 17)
test/test_005.md:5:17:	eq	"="	false	0	0.000000	(5, 17) -> (5, 18)
test/test_005.md:5:18:	real	"6.0"	false	0	6.000000	(5, 18) -> (5, 21)
test/test_005.md:5:21:	slash	"/"	false	0	0.000000	(5, 21) -> (5, 22)
test/test_005.md:5:22:	real	"4.0"	false	0	4.000000	(5, 22) -> (5, 25)
test/test_005.md:5:25:	rparen	")"	false	0	0.000000	(5, 25) -> (5, 26)
test/test_005.md:5:26:	semicolon	";"	false	0	0.000000	(5, 26) -> (5, 27)
test/test_005.md:6:4:	ident	"print"	false	0	0.000000	(6, 4) -> (6, 9)
test/test_005.md:6:9:	lparen	"("	false	0	0.000000	(6, 9) -> (6, 10)
test/test_005.md:6:10:	integer	"1"	false	1	0.000000	(6, 10) -> (6, 11)
test/test_005.md:6:11:	ne	"#"	false	0	0.000000	(6, 11) -> (6, 12)
test/test_005.md:6:12:	integer	"2"	false	2	0.000000	(6, 12) -> (6, 13)
test/test_005.md:6:13:	rparen	")"	false	0	0.000000	(6, 13) -> (6, 14)
test/test_005.md:6:14:	semicolon	";"	false	0	0.000000	(6, 14) -> (6, 15)
test/test_005.md:7:4:	ident	"print"	false	0	0.000000	(7, 4) -> (7, 9)
test/test_005.md:7:9:	lparen	"("	false	0	0.000000	(7, 9) -> (7, 10)
test/test_005.md:7:10:	integer	"3"	false	3	0.000000	(7, 10) -> (7, 11)
test/test_005.md:7:12:	div	"DIV"	false	0	0.000000	(7, 12) -> (7, 15)
test/test_005.md:7:16:	integer	"2"	false	2	0.000000	(7, 16) -> (7, 17)
test/test_005.md:7:17:	ne	"#"	false	0	0.000000	(7, 17) -> (7, 18)
test/test_005.md:7:18:	integer	"6"	false	6	0.000000	(7, 18) -> (7, 19)
test/test_005.md:7:20:	div	"DIV"	false	0	0.000000	(7, 20) -> (7, 23)
test/test_005.md:7:24:	integer	"4"	false	4	0.000000	(7, 24) -> (7, 25)
test/test_005.md:7:25:	rparen	")"	false	0	0.000000	(7, 25) -> (7, 26)
test/test_005.md:7:26:	semicolon	";"	false	0	0.000000	(7, 26) -> (7, 27)
test/test_005.md:8:4:	ident	"print"	false	0	0.000000	(8, 4) -> (8, 9)
test/test_005.md:8:9:	lparen	"("	false	0	0.000000	(8, 9) -> (8, 10)
test/test_005.md:8:10:	integer	"1"	false	1	0.000000	(8, 10) -> (8, 11)
test/test_005.md:8:11:	lt	"<"	false	0	0.000000	(8, 11) -> (8, 12)
test/test_005.md:8:12:	integer	"2"	false	2	0.000000	(8, 12) -> (8, 13)
test/test_005.md:8:13:	rparen	")"	false	0	0.000000	(8, 13) -> (8, 14)
test/test_005.md:8:14:	semicolon	";"	false	0	0.000000	(8, 14) -> (8, 15)
test/test_005.md:9:4:	ident	"print"	false	0	0.000000	(9, 4) -> (9, 9)
test/test_005.md:9:9:	lparen	"("	false	0	0.000000	(9, 9) -> (9, 10)
test/test_005.md:9:10:	real	"2.0"	false	0	2.000000	(9, 10) -> (9, 13)
test/test_005.md:9:13:	slash	"/"	false	0	0.000000	(9, 13) -> (9, 14)
test/test_005.md:9:14:	real	"1.0"	false	0	1.000000	(9, 14) -> (9, 17)
test/test_005.md:9:17:	lt	"<"	false	0	0.000000	(9, 17) -> (9, 18)
test/test_005.md:9:18:	real	"1.0"	false	0	1.000000	(9, 18) -> (9, 21)
test/test_005.md:9:21:	slash	"/"	false	0	0.000000	(9, 21) -> (9, 22)
test/test_005.md:9:22:	real	"1.0"	false	0	1.000000	(9, 22) -> (9, 25)
test/test_005.md:9:25:	rparen	")"	false	0	0.000000	(9, 25) -> (9, 26)
test/test_005.md:9:26:	semicolon	";"	false	0	0.000000	(9, 26) -> (9, 27)
test/test_005.md:10:4:	ident	"print"	false	0	0.000000	(10, 4) -> (10, 9)
test/test_005.md:10:9:	lparen	"("	false	0	0.000000	(10, 9) -> (10, 10)
test/test_005.md:10:10:	integer	"1"	false	1	0.000000	(10, 10) -> (10, 11)
test/test_005.md:10:11:	le	"<="	false	0	0.000000	(10, 11) -> (10, 13)
test/test_005.md:10:13:	integer	"1"	false	1	0.000000	(10, 13) -> (10, 14)
test/test_005.md:10:14:	rparen	")"	false	0	0.000000	(10, 14) -> (10, 15)
test/test_005.md:10:15:	semicolon	";"	false	0	0.000000	(10, 15) -> (10, 16)
test/test_005.md:11:4:	ident	"print"	false	0	0.000000	(11, 4) -> (11, 9)
test/test_005.md:11:9:	lparen	"("	false	0	0.000000	(11, 9) -> (11, 10)
test/test_005.md:11:10:	real	"1.0"	false	0	1.000000	(11, 10) -> (11, 13)
test/test_005.md:11:13:	slash	"/"	false	0	0.000000	(11, 13) -> (11, 14)
test/test_005.md:11:14:	real	"1.0"	false	0	1.000000	(11, 14) -> (11, 17)
test/test_005.md:11:17:	le	"<="	false	0	0.000000	(11, 17) -> (11, 19)
test/test_005.md:11:19:	real	"1.0"	false	0	1.000000	(11, 19) -> (11, 22)
test/test_005.md:11:22:	slash	"/"	false	0	0.000000	(11, 22) -> (11, 23)
test/test_005.md:11:23:	real	"1.0"	false	0	1.000000	(11, 23) -> (11, 26)
test/test_005.md:11:26:	rparen	")"	false	0	0.000000	(11, 26) -> (11, 27)
test/test_005.md:11:27:	semicolon	";"	false	0	0.000000	(11, 27) -> (11, 28)
test/test_005.md:12:4:	ident	"print"	false	0	0.000000	(12, 4) -> (12, 9)
test/test_005.md:12:9:	lparen	"("	false	0	0.000000	(12, 9) -> (12, 10)
test/test_005.md:12:10:	integer	"1"	false	1	0.000000	(12, 10) -> (12, 11)
test/test_005.md:12:11:	gt	">"	false	0	0.000000	(12, 11) -> (12, 12)
test/test_005.md:12:12:	integer	"2"	false	2	0.000000	(12, 12) -> (12, 13)
test/test_005.md:12:13:	rparen	")"	false	0	0.000000	(12, 13) -> (12, 14)
test/test_005.md:12:14:	semicolon	";"	false	0	0.000000	(12, 14) -> (12, 15)
test/test_005.md:13:4:	ident	"print"	false	0	0.000000	(13, 4) -> (13, 9)
test/test_005.md:13:9:	lparen	"("	false	0	0.000000	(13, 9) -> (13, 10)
test/test_005.md:13:10:	real	"2.0"	false	0	2.000000	(13, 10) -> (13, 13)
test/test_005.md:13:13:	slash	"/"	false	0	0.000000	(13, 13) -> (13, 14)
test/test_005.md:13:14:	real	"2.0"	false	0	2.000000	(13, 14) -> (13, 17)
test/test_005.md:13:17:	gt	">"	false	0	0.000000	(13, 17) -> (13, 18)
test/test_005.md:13:18:	real	"2.0"	false	0	2.000000	(13, 18) -> (13, 21)
test/test_005.md:13:21:	slash	"/"	false	0	0.000000	(13, 21) -> (13, 22)
test/test_005.md:13:22:	real	"1.0"	false	0	1.000000	(13, 22) -> (13, 25)
test/test_005.md:13:25:	rparen	")"	false	0	0.000000	(13, 25) -> (13, 26)
test/test_005.md:13:26:	semicolon	";"	false	0	0.000000	(13, 26) -> (13, 27)
test/test_005.md:14:4:	ident	"print"	false	0	0.000000	(14, 4) -> (14, 9)
test/test_005.md:14:9:	lparen	"("	false	0	0.000000	(14, 9) -> (14, 10)
test/test_005.md:14:10:	integer	"2"	false	2	0.000000	(14, 10) -> (14, 11)
test/test_005.md:14:11:	ge	">="	false	0	0.000000	(14, 11) -> (14, 13)
test/test_005.md:14:13:	integer	"2"	false	2	0.000000	(14, 13) -> (14, 14)
test/test_005.md:14:14:	rparen	")"	false	0	0.000000	(14, 14) -> (14, 15)
test/test_005.md:14:15:	semicolon	";"	false	0	0.000000	(14, 15) -> (14, 16)
test/test_005.md:15:4:	ident	"print"	false	0	0.000000	(15, 4) -> (15, 9)
test/test_005.md:15:9:	lparen	"("	false	0	0.000000	(15, 9) -> (15, 10)
test/test_005.md:15:10:	real	"2.0"	false	0	2.000000	(15, 10) -> (15, 13)
test/test_005.md:15:13:	slash	"/"	false	0	0.000000	(15, 13) -> (15, 14)
test/test_005.md:15:14:	real	"1.0"	false	0	1.000000	(15, 14) -> (15, 17)
test/test_005.md:15:17:	ge	">="	false	0	0.000000	(15, 17) -> (15, 19)
test/test_005.md:15:19:	real	"2.0"	false	0	2.000000	(15, 19) -> (15, 22)
test/test_005.md:15:22:	slash	"/"	false	0	0.000000	(15, 22) -> (15, 23)
test/test_005.md:15:23:	real	"1.0"	false	0	1.000000	(15, 23) -> (15, 26)
test/test_005.md:15:26:	rparen	")"	false	0	0.000000	(15, 26) -> (15, 27)
test/test_005.md:16:0:	end	"END"	false	0	0.000000	(16, 0) -> (16, 3)
test/test_005.md:16:4:	ident	"Comparisons"	false	0	0.000000	(16, 4) -> (16, 15)
test/test_005.md:16:15:	dot	"."	false	0	0.000000	(16, 15) -> (16, 16)
test/test_005.md:17:0:	eof	""	false	0	0.000000	(17, 0) -> (17, 0)
```
## AST
```scheme
(module "Comparisons"
  (stmts
    (expr2stmt
      (call
        (procedure [void] "print")
        (eq [boolean]
          (number [i64] 1)
          (number [i64] 2)
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "print")
        (eq [boolean]
          (slash [f64]
            (number [f64] 3.000000)
            (number [f64] 2.000000)
          )
          (slash [f64]
            (number [f64] 6.000000)
            (number [f64] 4.000000)
          )
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "print")
        (ne [boolean]
          (number [i64] 1)
          (number [i64] 2)
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "print")
        (ne [boolean]
          (div [i64]
            (number [i64] 3)
            (number [i64] 2)
          )
          (div [i64]
            (number [i64] 6)
            (number [i64] 4)
          )
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "print")
        (lt [boolean]
          (number [i64] 1)
          (number [i64] 2)
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "print")
        (lt [boolean]
          (slash [f64]
            (number [f64] 2.000000)
            (number [f64] 1.000000)
          )
          (slash [f64]
            (number [f64] 1.000000)
            (number [f64] 1.000000)
          )
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "print")
        (le [boolean]
          (number [i64] 1)
          (number [i64] 1)
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "print")
        (le [boolean]
          (slash [f64]
            (number [f64] 1.000000)
            (number [f64] 1.000000)
          )
          (slash [f64]
            (number [f64] 1.000000)
            (number [f64] 1.000000)
          )
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "print")
        (gt [boolean]
          (number [i64] 1)
          (number [i64] 2)
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "print")
        (gt [boolean]
          (slash [f64]
            (number [f64] 2.000000)
            (number [f64] 2.000000)
          )
          (slash [f64]
            (number [f64] 2.000000)
            (number [f64] 1.000000)
          )
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "print")
        (ge [boolean]
          (number [i64] 2)
          (number [i64] 2)
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "print")
        (ge [boolean]
          (slash [f64]
            (number [f64] 2.000000)
            (number [f64] 1.000000)
          )
          (slash [f64]
            (number [f64] 2.000000)
            (number [f64] 1.000000)
          )
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

declare i64 @puts(i8* %str)

declare i64 @rand()

declare i64 @sprintf(i8* %buf, i8* %format, ...)

declare i64 @printf(i8* %format, ...)

declare i8* @malloc(i64 %size)

declare i8* @free(i8* %ptr)

define i64 @main() {
entry:
	%0 = icmp eq i64 1, 2
	br i1 %0, label %1, label %3

1:
	%2 = getelementptr [5 x i8], [5 x i8]* @0, i64 0, i64 0
	br label %5

3:
	%4 = getelementptr [6 x i8], [6 x i8]* @1, i64 0, i64 0
	br label %5

5:
	%6 = phi i8* [ %2, %1 ], [ %4, %3 ]
	%7 = call i64 @puts(i8* %6)
	%8 = fdiv double 3.0, 2.0
	%9 = fdiv double 6.0, 4.0
	%10 = fcmp ueq double %8, %9
	br i1 %10, label %11, label %13

11:
	%12 = getelementptr [5 x i8], [5 x i8]* @0, i64 0, i64 0
	br label %15

13:
	%14 = getelementptr [6 x i8], [6 x i8]* @1, i64 0, i64 0
	br label %15

15:
	%16 = phi i8* [ %12, %11 ], [ %14, %13 ]
	%17 = call i64 @puts(i8* %16)
	%18 = icmp ne i64 1, 2
	br i1 %18, label %19, label %21

19:
	%20 = getelementptr [5 x i8], [5 x i8]* @0, i64 0, i64 0
	br label %23

21:
	%22 = getelementptr [6 x i8], [6 x i8]* @1, i64 0, i64 0
	br label %23

23:
	%24 = phi i8* [ %20, %19 ], [ %22, %21 ]
	%25 = call i64 @puts(i8* %24)
	%26 = sdiv i64 3, 2
	%27 = sdiv i64 6, 4
	%28 = icmp ne i64 %26, %27
	br i1 %28, label %29, label %31

29:
	%30 = getelementptr [5 x i8], [5 x i8]* @0, i64 0, i64 0
	br label %33

31:
	%32 = getelementptr [6 x i8], [6 x i8]* @1, i64 0, i64 0
	br label %33

33:
	%34 = phi i8* [ %30, %29 ], [ %32, %31 ]
	%35 = call i64 @puts(i8* %34)
	%36 = icmp slt i64 1, 2
	br i1 %36, label %37, label %39

37:
	%38 = getelementptr [5 x i8], [5 x i8]* @0, i64 0, i64 0
	br label %41

39:
	%40 = getelementptr [6 x i8], [6 x i8]* @1, i64 0, i64 0
	br label %41

41:
	%42 = phi i8* [ %38, %37 ], [ %40, %39 ]
	%43 = call i64 @puts(i8* %42)
	%44 = fdiv double 2.0, 1.0
	%45 = fdiv double 1.0, 1.0
	%46 = fcmp ult double %44, %45
	br i1 %46, label %47, label %49

47:
	%48 = getelementptr [5 x i8], [5 x i8]* @0, i64 0, i64 0
	br label %51

49:
	%50 = getelementptr [6 x i8], [6 x i8]* @1, i64 0, i64 0
	br label %51

51:
	%52 = phi i8* [ %48, %47 ], [ %50, %49 ]
	%53 = call i64 @puts(i8* %52)
	%54 = icmp sle i64 1, 1
	br i1 %54, label %55, label %57

55:
	%56 = getelementptr [5 x i8], [5 x i8]* @0, i64 0, i64 0
	br label %59

57:
	%58 = getelementptr [6 x i8], [6 x i8]* @1, i64 0, i64 0
	br label %59

59:
	%60 = phi i8* [ %56, %55 ], [ %58, %57 ]
	%61 = call i64 @puts(i8* %60)
	%62 = fdiv double 1.0, 1.0
	%63 = fdiv double 1.0, 1.0
	%64 = fcmp ule double %62, %63
	br i1 %64, label %65, label %67

65:
	%66 = getelementptr [5 x i8], [5 x i8]* @0, i64 0, i64 0
	br label %69

67:
	%68 = getelementptr [6 x i8], [6 x i8]* @1, i64 0, i64 0
	br label %69

69:
	%70 = phi i8* [ %66, %65 ], [ %68, %67 ]
	%71 = call i64 @puts(i8* %70)
	%72 = icmp sgt i64 1, 2
	br i1 %72, label %73, label %75

73:
	%74 = getelementptr [5 x i8], [5 x i8]* @0, i64 0, i64 0
	br label %77

75:
	%76 = getelementptr [6 x i8], [6 x i8]* @1, i64 0, i64 0
	br label %77

77:
	%78 = phi i8* [ %74, %73 ], [ %76, %75 ]
	%79 = call i64 @puts(i8* %78)
	%80 = fdiv double 2.0, 2.0
	%81 = fdiv double 2.0, 1.0
	%82 = fcmp ugt double %80, %81
	br i1 %82, label %83, label %85

83:
	%84 = getelementptr [5 x i8], [5 x i8]* @0, i64 0, i64 0
	br label %87

85:
	%86 = getelementptr [6 x i8], [6 x i8]* @1, i64 0, i64 0
	br label %87

87:
	%88 = phi i8* [ %84, %83 ], [ %86, %85 ]
	%89 = call i64 @puts(i8* %88)
	%90 = icmp sge i64 2, 2
	br i1 %90, label %91, label %93

91:
	%92 = getelementptr [5 x i8], [5 x i8]* @0, i64 0, i64 0
	br label %95

93:
	%94 = getelementptr [6 x i8], [6 x i8]* @1, i64 0, i64 0
	br label %95

95:
	%96 = phi i8* [ %92, %91 ], [ %94, %93 ]
	%97 = call i64 @puts(i8* %96)
	%98 = fdiv double 2.0, 1.0
	%99 = fdiv double 2.0, 1.0
	%100 = fcmp uge double %98, %99
	br i1 %100, label %101, label %103

101:
	%102 = getelementptr [5 x i8], [5 x i8]* @0, i64 0, i64 0
	br label %105

103:
	%104 = getelementptr [6 x i8], [6 x i8]* @1, i64 0, i64 0
	br label %105

105:
	%106 = phi i8* [ %102, %101 ], [ %104, %103 ]
	%107 = call i64 @puts(i8* %106)
	ret i64 0
}

```
## Run
```bash
FALSE
TRUE
TRUE
FALSE
TRUE
FALSE
TRUE
TRUE
FALSE
FALSE
TRUE
TRUE
```
