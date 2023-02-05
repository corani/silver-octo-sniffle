# test/test_005.md
## Source
```pascal
MODULE Comparisons;

BEGIN
    print(1=2);
    print(3/2=6/4);
    print(1#2);
    print(3/2#6/4);
    print(1<2);
    print(2/1<1/1);
    print(1<=1);
    print(1/1<=1/1);
    print(1>2);
    print(2/2>2/1);
    print(2>=2);
    print(2/1>=2/1)
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
test/test_005.md:5:10:	integer	"3"	false	3	0.000000	(5, 10) -> (5, 11)
test/test_005.md:5:11:	slash	"/"	false	0	0.000000	(5, 11) -> (5, 12)
test/test_005.md:5:12:	integer	"2"	false	2	0.000000	(5, 12) -> (5, 13)
test/test_005.md:5:13:	eq	"="	false	0	0.000000	(5, 13) -> (5, 14)
test/test_005.md:5:14:	integer	"6"	false	6	0.000000	(5, 14) -> (5, 15)
test/test_005.md:5:15:	slash	"/"	false	0	0.000000	(5, 15) -> (5, 16)
test/test_005.md:5:16:	integer	"4"	false	4	0.000000	(5, 16) -> (5, 17)
test/test_005.md:5:17:	rparen	")"	false	0	0.000000	(5, 17) -> (5, 18)
test/test_005.md:5:18:	semicolon	";"	false	0	0.000000	(5, 18) -> (5, 19)
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
test/test_005.md:7:11:	slash	"/"	false	0	0.000000	(7, 11) -> (7, 12)
test/test_005.md:7:12:	integer	"2"	false	2	0.000000	(7, 12) -> (7, 13)
test/test_005.md:7:13:	ne	"#"	false	0	0.000000	(7, 13) -> (7, 14)
test/test_005.md:7:14:	integer	"6"	false	6	0.000000	(7, 14) -> (7, 15)
test/test_005.md:7:15:	slash	"/"	false	0	0.000000	(7, 15) -> (7, 16)
test/test_005.md:7:16:	integer	"4"	false	4	0.000000	(7, 16) -> (7, 17)
test/test_005.md:7:17:	rparen	")"	false	0	0.000000	(7, 17) -> (7, 18)
test/test_005.md:7:18:	semicolon	";"	false	0	0.000000	(7, 18) -> (7, 19)
test/test_005.md:8:4:	ident	"print"	false	0	0.000000	(8, 4) -> (8, 9)
test/test_005.md:8:9:	lparen	"("	false	0	0.000000	(8, 9) -> (8, 10)
test/test_005.md:8:10:	integer	"1"	false	1	0.000000	(8, 10) -> (8, 11)
test/test_005.md:8:11:	lt	"<"	false	0	0.000000	(8, 11) -> (8, 12)
test/test_005.md:8:12:	integer	"2"	false	2	0.000000	(8, 12) -> (8, 13)
test/test_005.md:8:13:	rparen	")"	false	0	0.000000	(8, 13) -> (8, 14)
test/test_005.md:8:14:	semicolon	";"	false	0	0.000000	(8, 14) -> (8, 15)
test/test_005.md:9:4:	ident	"print"	false	0	0.000000	(9, 4) -> (9, 9)
test/test_005.md:9:9:	lparen	"("	false	0	0.000000	(9, 9) -> (9, 10)
test/test_005.md:9:10:	integer	"2"	false	2	0.000000	(9, 10) -> (9, 11)
test/test_005.md:9:11:	slash	"/"	false	0	0.000000	(9, 11) -> (9, 12)
test/test_005.md:9:12:	integer	"1"	false	1	0.000000	(9, 12) -> (9, 13)
test/test_005.md:9:13:	lt	"<"	false	0	0.000000	(9, 13) -> (9, 14)
test/test_005.md:9:14:	integer	"1"	false	1	0.000000	(9, 14) -> (9, 15)
test/test_005.md:9:15:	slash	"/"	false	0	0.000000	(9, 15) -> (9, 16)
test/test_005.md:9:16:	integer	"1"	false	1	0.000000	(9, 16) -> (9, 17)
test/test_005.md:9:17:	rparen	")"	false	0	0.000000	(9, 17) -> (9, 18)
test/test_005.md:9:18:	semicolon	";"	false	0	0.000000	(9, 18) -> (9, 19)
test/test_005.md:10:4:	ident	"print"	false	0	0.000000	(10, 4) -> (10, 9)
test/test_005.md:10:9:	lparen	"("	false	0	0.000000	(10, 9) -> (10, 10)
test/test_005.md:10:10:	integer	"1"	false	1	0.000000	(10, 10) -> (10, 11)
test/test_005.md:10:11:	le	"<="	false	0	0.000000	(10, 11) -> (10, 13)
test/test_005.md:10:13:	integer	"1"	false	1	0.000000	(10, 13) -> (10, 14)
test/test_005.md:10:14:	rparen	")"	false	0	0.000000	(10, 14) -> (10, 15)
test/test_005.md:10:15:	semicolon	";"	false	0	0.000000	(10, 15) -> (10, 16)
test/test_005.md:11:4:	ident	"print"	false	0	0.000000	(11, 4) -> (11, 9)
test/test_005.md:11:9:	lparen	"("	false	0	0.000000	(11, 9) -> (11, 10)
test/test_005.md:11:10:	integer	"1"	false	1	0.000000	(11, 10) -> (11, 11)
test/test_005.md:11:11:	slash	"/"	false	0	0.000000	(11, 11) -> (11, 12)
test/test_005.md:11:12:	integer	"1"	false	1	0.000000	(11, 12) -> (11, 13)
test/test_005.md:11:13:	le	"<="	false	0	0.000000	(11, 13) -> (11, 15)
test/test_005.md:11:15:	integer	"1"	false	1	0.000000	(11, 15) -> (11, 16)
test/test_005.md:11:16:	slash	"/"	false	0	0.000000	(11, 16) -> (11, 17)
test/test_005.md:11:17:	integer	"1"	false	1	0.000000	(11, 17) -> (11, 18)
test/test_005.md:11:18:	rparen	")"	false	0	0.000000	(11, 18) -> (11, 19)
test/test_005.md:11:19:	semicolon	";"	false	0	0.000000	(11, 19) -> (11, 20)
test/test_005.md:12:4:	ident	"print"	false	0	0.000000	(12, 4) -> (12, 9)
test/test_005.md:12:9:	lparen	"("	false	0	0.000000	(12, 9) -> (12, 10)
test/test_005.md:12:10:	integer	"1"	false	1	0.000000	(12, 10) -> (12, 11)
test/test_005.md:12:11:	gt	">"	false	0	0.000000	(12, 11) -> (12, 12)
test/test_005.md:12:12:	integer	"2"	false	2	0.000000	(12, 12) -> (12, 13)
test/test_005.md:12:13:	rparen	")"	false	0	0.000000	(12, 13) -> (12, 14)
test/test_005.md:12:14:	semicolon	";"	false	0	0.000000	(12, 14) -> (12, 15)
test/test_005.md:13:4:	ident	"print"	false	0	0.000000	(13, 4) -> (13, 9)
test/test_005.md:13:9:	lparen	"("	false	0	0.000000	(13, 9) -> (13, 10)
test/test_005.md:13:10:	integer	"2"	false	2	0.000000	(13, 10) -> (13, 11)
test/test_005.md:13:11:	slash	"/"	false	0	0.000000	(13, 11) -> (13, 12)
test/test_005.md:13:12:	integer	"2"	false	2	0.000000	(13, 12) -> (13, 13)
test/test_005.md:13:13:	gt	">"	false	0	0.000000	(13, 13) -> (13, 14)
test/test_005.md:13:14:	integer	"2"	false	2	0.000000	(13, 14) -> (13, 15)
test/test_005.md:13:15:	slash	"/"	false	0	0.000000	(13, 15) -> (13, 16)
test/test_005.md:13:16:	integer	"1"	false	1	0.000000	(13, 16) -> (13, 17)
test/test_005.md:13:17:	rparen	")"	false	0	0.000000	(13, 17) -> (13, 18)
test/test_005.md:13:18:	semicolon	";"	false	0	0.000000	(13, 18) -> (13, 19)
test/test_005.md:14:4:	ident	"print"	false	0	0.000000	(14, 4) -> (14, 9)
test/test_005.md:14:9:	lparen	"("	false	0	0.000000	(14, 9) -> (14, 10)
test/test_005.md:14:10:	integer	"2"	false	2	0.000000	(14, 10) -> (14, 11)
test/test_005.md:14:11:	ge	">="	false	0	0.000000	(14, 11) -> (14, 13)
test/test_005.md:14:13:	integer	"2"	false	2	0.000000	(14, 13) -> (14, 14)
test/test_005.md:14:14:	rparen	")"	false	0	0.000000	(14, 14) -> (14, 15)
test/test_005.md:14:15:	semicolon	";"	false	0	0.000000	(14, 15) -> (14, 16)
test/test_005.md:15:4:	ident	"print"	false	0	0.000000	(15, 4) -> (15, 9)
test/test_005.md:15:9:	lparen	"("	false	0	0.000000	(15, 9) -> (15, 10)
test/test_005.md:15:10:	integer	"2"	false	2	0.000000	(15, 10) -> (15, 11)
test/test_005.md:15:11:	slash	"/"	false	0	0.000000	(15, 11) -> (15, 12)
test/test_005.md:15:12:	integer	"1"	false	1	0.000000	(15, 12) -> (15, 13)
test/test_005.md:15:13:	ge	">="	false	0	0.000000	(15, 13) -> (15, 15)
test/test_005.md:15:15:	integer	"2"	false	2	0.000000	(15, 15) -> (15, 16)
test/test_005.md:15:16:	slash	"/"	false	0	0.000000	(15, 16) -> (15, 17)
test/test_005.md:15:17:	integer	"1"	false	1	0.000000	(15, 17) -> (15, 18)
test/test_005.md:15:18:	rparen	")"	false	0	0.000000	(15, 18) -> (15, 19)
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
      (call "print" [void]
        (eq [boolean]
          (number [i64] 1)
          (number [i64] 2)
        )
      )
    )
    (expr2stmt
      (call "print" [void]
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
    (expr2stmt
      (call "print" [void]
        (ne [boolean]
          (number [i64] 1)
          (number [i64] 2)
        )
      )
    )
    (expr2stmt
      (call "print" [void]
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
    (expr2stmt
      (call "print" [void]
        (lt [boolean]
          (number [i64] 1)
          (number [i64] 2)
        )
      )
    )
    (expr2stmt
      (call "print" [void]
        (lt [boolean]
          (slash [f64]
            (number [i64] 2)
            (number [i64] 1)
          )
          (slash [f64]
            (number [i64] 1)
            (number [i64] 1)
          )
        )
      )
    )
    (expr2stmt
      (call "print" [void]
        (le [boolean]
          (number [i64] 1)
          (number [i64] 1)
        )
      )
    )
    (expr2stmt
      (call "print" [void]
        (le [boolean]
          (slash [f64]
            (number [i64] 1)
            (number [i64] 1)
          )
          (slash [f64]
            (number [i64] 1)
            (number [i64] 1)
          )
        )
      )
    )
    (expr2stmt
      (call "print" [void]
        (gt [boolean]
          (number [i64] 1)
          (number [i64] 2)
        )
      )
    )
    (expr2stmt
      (call "print" [void]
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
      (call "print" [void]
        (ge [boolean]
          (number [i64] 2)
          (number [i64] 2)
        )
      )
    )
    (expr2stmt
      (call "print" [void]
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
	%8 = sitofp i64 3 to double
	%9 = sitofp i64 2 to double
	%10 = fdiv double %8, %9
	%11 = sitofp i64 6 to double
	%12 = sitofp i64 4 to double
	%13 = fdiv double %11, %12
	%14 = fcmp ueq double %10, %13
	br i1 %14, label %15, label %17

15:
	%16 = getelementptr [5 x i8], [5 x i8]* @0, i64 0, i64 0
	br label %19

17:
	%18 = getelementptr [6 x i8], [6 x i8]* @1, i64 0, i64 0
	br label %19

19:
	%20 = phi i8* [ %16, %15 ], [ %18, %17 ]
	%21 = call i64 @puts(i8* %20)
	%22 = icmp ne i64 1, 2
	br i1 %22, label %23, label %25

23:
	%24 = getelementptr [5 x i8], [5 x i8]* @0, i64 0, i64 0
	br label %27

25:
	%26 = getelementptr [6 x i8], [6 x i8]* @1, i64 0, i64 0
	br label %27

27:
	%28 = phi i8* [ %24, %23 ], [ %26, %25 ]
	%29 = call i64 @puts(i8* %28)
	%30 = sitofp i64 3 to double
	%31 = sitofp i64 2 to double
	%32 = fdiv double %30, %31
	%33 = sitofp i64 6 to double
	%34 = sitofp i64 4 to double
	%35 = fdiv double %33, %34
	%36 = fcmp une double %32, %35
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
	%44 = icmp slt i64 1, 2
	br i1 %44, label %45, label %47

45:
	%46 = getelementptr [5 x i8], [5 x i8]* @0, i64 0, i64 0
	br label %49

47:
	%48 = getelementptr [6 x i8], [6 x i8]* @1, i64 0, i64 0
	br label %49

49:
	%50 = phi i8* [ %46, %45 ], [ %48, %47 ]
	%51 = call i64 @puts(i8* %50)
	%52 = sitofp i64 2 to double
	%53 = sitofp i64 1 to double
	%54 = fdiv double %52, %53
	%55 = sitofp i64 1 to double
	%56 = sitofp i64 1 to double
	%57 = fdiv double %55, %56
	%58 = fcmp ult double %54, %57
	br i1 %58, label %59, label %61

59:
	%60 = getelementptr [5 x i8], [5 x i8]* @0, i64 0, i64 0
	br label %63

61:
	%62 = getelementptr [6 x i8], [6 x i8]* @1, i64 0, i64 0
	br label %63

63:
	%64 = phi i8* [ %60, %59 ], [ %62, %61 ]
	%65 = call i64 @puts(i8* %64)
	%66 = icmp sle i64 1, 1
	br i1 %66, label %67, label %69

67:
	%68 = getelementptr [5 x i8], [5 x i8]* @0, i64 0, i64 0
	br label %71

69:
	%70 = getelementptr [6 x i8], [6 x i8]* @1, i64 0, i64 0
	br label %71

71:
	%72 = phi i8* [ %68, %67 ], [ %70, %69 ]
	%73 = call i64 @puts(i8* %72)
	%74 = sitofp i64 1 to double
	%75 = sitofp i64 1 to double
	%76 = fdiv double %74, %75
	%77 = sitofp i64 1 to double
	%78 = sitofp i64 1 to double
	%79 = fdiv double %77, %78
	%80 = fcmp ule double %76, %79
	br i1 %80, label %81, label %83

81:
	%82 = getelementptr [5 x i8], [5 x i8]* @0, i64 0, i64 0
	br label %85

83:
	%84 = getelementptr [6 x i8], [6 x i8]* @1, i64 0, i64 0
	br label %85

85:
	%86 = phi i8* [ %82, %81 ], [ %84, %83 ]
	%87 = call i64 @puts(i8* %86)
	%88 = icmp sgt i64 1, 2
	br i1 %88, label %89, label %91

89:
	%90 = getelementptr [5 x i8], [5 x i8]* @0, i64 0, i64 0
	br label %93

91:
	%92 = getelementptr [6 x i8], [6 x i8]* @1, i64 0, i64 0
	br label %93

93:
	%94 = phi i8* [ %90, %89 ], [ %92, %91 ]
	%95 = call i64 @puts(i8* %94)
	%96 = sitofp i64 2 to double
	%97 = sitofp i64 2 to double
	%98 = fdiv double %96, %97
	%99 = sitofp i64 2 to double
	%100 = sitofp i64 1 to double
	%101 = fdiv double %99, %100
	%102 = fcmp ugt double %98, %101
	br i1 %102, label %103, label %105

103:
	%104 = getelementptr [5 x i8], [5 x i8]* @0, i64 0, i64 0
	br label %107

105:
	%106 = getelementptr [6 x i8], [6 x i8]* @1, i64 0, i64 0
	br label %107

107:
	%108 = phi i8* [ %104, %103 ], [ %106, %105 ]
	%109 = call i64 @puts(i8* %108)
	%110 = icmp sge i64 2, 2
	br i1 %110, label %111, label %113

111:
	%112 = getelementptr [5 x i8], [5 x i8]* @0, i64 0, i64 0
	br label %115

113:
	%114 = getelementptr [6 x i8], [6 x i8]* @1, i64 0, i64 0
	br label %115

115:
	%116 = phi i8* [ %112, %111 ], [ %114, %113 ]
	%117 = call i64 @puts(i8* %116)
	%118 = sitofp i64 2 to double
	%119 = sitofp i64 1 to double
	%120 = fdiv double %118, %119
	%121 = sitofp i64 2 to double
	%122 = sitofp i64 1 to double
	%123 = fdiv double %121, %122
	%124 = fcmp uge double %120, %123
	br i1 %124, label %125, label %127

125:
	%126 = getelementptr [5 x i8], [5 x i8]* @0, i64 0, i64 0
	br label %129

127:
	%128 = getelementptr [6 x i8], [6 x i8]* @1, i64 0, i64 0
	br label %129

129:
	%130 = phi i8* [ %126, %125 ], [ %128, %127 ]
	%131 = call i64 @puts(i8* %130)
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
