# test/test_005.md
## Source
```pascal
(* comparisons *)
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
```
## Tokens
```tsv
test/test_005.md:2:0:	ident	"print"	false	0	0.000000	(2, 0) -> (2, 5)
test/test_005.md:2:5:	lparen	"("	false	0	0.000000	(2, 5) -> (2, 6)
test/test_005.md:2:6:	integer	"1"	false	1	0.000000	(2, 6) -> (2, 7)
test/test_005.md:2:7:	eq	"="	false	0	0.000000	(2, 7) -> (2, 8)
test/test_005.md:2:8:	integer	"2"	false	2	0.000000	(2, 8) -> (2, 9)
test/test_005.md:2:9:	rparen	")"	false	0	0.000000	(2, 9) -> (2, 10)
test/test_005.md:2:10:	semicolon	";"	false	0	0.000000	(2, 10) -> (2, 11)
test/test_005.md:3:0:	ident	"print"	false	0	0.000000	(3, 0) -> (3, 5)
test/test_005.md:3:5:	lparen	"("	false	0	0.000000	(3, 5) -> (3, 6)
test/test_005.md:3:6:	integer	"3"	false	3	0.000000	(3, 6) -> (3, 7)
test/test_005.md:3:7:	slash	"/"	false	0	0.000000	(3, 7) -> (3, 8)
test/test_005.md:3:8:	integer	"2"	false	2	0.000000	(3, 8) -> (3, 9)
test/test_005.md:3:9:	eq	"="	false	0	0.000000	(3, 9) -> (3, 10)
test/test_005.md:3:10:	integer	"6"	false	6	0.000000	(3, 10) -> (3, 11)
test/test_005.md:3:11:	slash	"/"	false	0	0.000000	(3, 11) -> (3, 12)
test/test_005.md:3:12:	integer	"4"	false	4	0.000000	(3, 12) -> (3, 13)
test/test_005.md:3:13:	rparen	")"	false	0	0.000000	(3, 13) -> (3, 14)
test/test_005.md:3:14:	semicolon	";"	false	0	0.000000	(3, 14) -> (3, 15)
test/test_005.md:4:0:	ident	"print"	false	0	0.000000	(4, 0) -> (4, 5)
test/test_005.md:4:5:	lparen	"("	false	0	0.000000	(4, 5) -> (4, 6)
test/test_005.md:4:6:	integer	"1"	false	1	0.000000	(4, 6) -> (4, 7)
test/test_005.md:4:7:	ne	"#"	false	0	0.000000	(4, 7) -> (4, 8)
test/test_005.md:4:8:	integer	"2"	false	2	0.000000	(4, 8) -> (4, 9)
test/test_005.md:4:9:	rparen	")"	false	0	0.000000	(4, 9) -> (4, 10)
test/test_005.md:4:10:	semicolon	";"	false	0	0.000000	(4, 10) -> (4, 11)
test/test_005.md:5:0:	ident	"print"	false	0	0.000000	(5, 0) -> (5, 5)
test/test_005.md:5:5:	lparen	"("	false	0	0.000000	(5, 5) -> (5, 6)
test/test_005.md:5:6:	integer	"3"	false	3	0.000000	(5, 6) -> (5, 7)
test/test_005.md:5:7:	slash	"/"	false	0	0.000000	(5, 7) -> (5, 8)
test/test_005.md:5:8:	integer	"2"	false	2	0.000000	(5, 8) -> (5, 9)
test/test_005.md:5:9:	ne	"#"	false	0	0.000000	(5, 9) -> (5, 10)
test/test_005.md:5:10:	integer	"6"	false	6	0.000000	(5, 10) -> (5, 11)
test/test_005.md:5:11:	slash	"/"	false	0	0.000000	(5, 11) -> (5, 12)
test/test_005.md:5:12:	integer	"4"	false	4	0.000000	(5, 12) -> (5, 13)
test/test_005.md:5:13:	rparen	")"	false	0	0.000000	(5, 13) -> (5, 14)
test/test_005.md:5:14:	semicolon	";"	false	0	0.000000	(5, 14) -> (5, 15)
test/test_005.md:6:0:	ident	"print"	false	0	0.000000	(6, 0) -> (6, 5)
test/test_005.md:6:5:	lparen	"("	false	0	0.000000	(6, 5) -> (6, 6)
test/test_005.md:6:6:	integer	"1"	false	1	0.000000	(6, 6) -> (6, 7)
test/test_005.md:6:7:	lt	"<"	false	0	0.000000	(6, 7) -> (6, 8)
test/test_005.md:6:8:	integer	"2"	false	2	0.000000	(6, 8) -> (6, 9)
test/test_005.md:6:9:	rparen	")"	false	0	0.000000	(6, 9) -> (6, 10)
test/test_005.md:6:10:	semicolon	";"	false	0	0.000000	(6, 10) -> (6, 11)
test/test_005.md:7:0:	ident	"print"	false	0	0.000000	(7, 0) -> (7, 5)
test/test_005.md:7:5:	lparen	"("	false	0	0.000000	(7, 5) -> (7, 6)
test/test_005.md:7:6:	integer	"2"	false	2	0.000000	(7, 6) -> (7, 7)
test/test_005.md:7:7:	slash	"/"	false	0	0.000000	(7, 7) -> (7, 8)
test/test_005.md:7:8:	integer	"1"	false	1	0.000000	(7, 8) -> (7, 9)
test/test_005.md:7:9:	lt	"<"	false	0	0.000000	(7, 9) -> (7, 10)
test/test_005.md:7:10:	integer	"1"	false	1	0.000000	(7, 10) -> (7, 11)
test/test_005.md:7:11:	slash	"/"	false	0	0.000000	(7, 11) -> (7, 12)
test/test_005.md:7:12:	integer	"1"	false	1	0.000000	(7, 12) -> (7, 13)
test/test_005.md:7:13:	rparen	")"	false	0	0.000000	(7, 13) -> (7, 14)
test/test_005.md:7:14:	semicolon	";"	false	0	0.000000	(7, 14) -> (7, 15)
test/test_005.md:8:0:	ident	"print"	false	0	0.000000	(8, 0) -> (8, 5)
test/test_005.md:8:5:	lparen	"("	false	0	0.000000	(8, 5) -> (8, 6)
test/test_005.md:8:6:	integer	"1"	false	1	0.000000	(8, 6) -> (8, 7)
test/test_005.md:8:7:	le	"<="	false	0	0.000000	(8, 7) -> (8, 9)
test/test_005.md:8:9:	integer	"1"	false	1	0.000000	(8, 9) -> (8, 10)
test/test_005.md:8:10:	rparen	")"	false	0	0.000000	(8, 10) -> (8, 11)
test/test_005.md:8:11:	semicolon	";"	false	0	0.000000	(8, 11) -> (8, 12)
test/test_005.md:9:0:	ident	"print"	false	0	0.000000	(9, 0) -> (9, 5)
test/test_005.md:9:5:	lparen	"("	false	0	0.000000	(9, 5) -> (9, 6)
test/test_005.md:9:6:	integer	"1"	false	1	0.000000	(9, 6) -> (9, 7)
test/test_005.md:9:7:	slash	"/"	false	0	0.000000	(9, 7) -> (9, 8)
test/test_005.md:9:8:	integer	"1"	false	1	0.000000	(9, 8) -> (9, 9)
test/test_005.md:9:9:	le	"<="	false	0	0.000000	(9, 9) -> (9, 11)
test/test_005.md:9:11:	integer	"1"	false	1	0.000000	(9, 11) -> (9, 12)
test/test_005.md:9:12:	slash	"/"	false	0	0.000000	(9, 12) -> (9, 13)
test/test_005.md:9:13:	integer	"1"	false	1	0.000000	(9, 13) -> (9, 14)
test/test_005.md:9:14:	rparen	")"	false	0	0.000000	(9, 14) -> (9, 15)
test/test_005.md:9:15:	semicolon	";"	false	0	0.000000	(9, 15) -> (9, 16)
test/test_005.md:10:0:	ident	"print"	false	0	0.000000	(10, 0) -> (10, 5)
test/test_005.md:10:5:	lparen	"("	false	0	0.000000	(10, 5) -> (10, 6)
test/test_005.md:10:6:	integer	"1"	false	1	0.000000	(10, 6) -> (10, 7)
test/test_005.md:10:7:	gt	">"	false	0	0.000000	(10, 7) -> (10, 8)
test/test_005.md:10:8:	integer	"2"	false	2	0.000000	(10, 8) -> (10, 9)
test/test_005.md:10:9:	rparen	")"	false	0	0.000000	(10, 9) -> (10, 10)
test/test_005.md:10:10:	semicolon	";"	false	0	0.000000	(10, 10) -> (10, 11)
test/test_005.md:11:0:	ident	"print"	false	0	0.000000	(11, 0) -> (11, 5)
test/test_005.md:11:5:	lparen	"("	false	0	0.000000	(11, 5) -> (11, 6)
test/test_005.md:11:6:	integer	"2"	false	2	0.000000	(11, 6) -> (11, 7)
test/test_005.md:11:7:	slash	"/"	false	0	0.000000	(11, 7) -> (11, 8)
test/test_005.md:11:8:	integer	"2"	false	2	0.000000	(11, 8) -> (11, 9)
test/test_005.md:11:9:	gt	">"	false	0	0.000000	(11, 9) -> (11, 10)
test/test_005.md:11:10:	integer	"2"	false	2	0.000000	(11, 10) -> (11, 11)
test/test_005.md:11:11:	slash	"/"	false	0	0.000000	(11, 11) -> (11, 12)
test/test_005.md:11:12:	integer	"1"	false	1	0.000000	(11, 12) -> (11, 13)
test/test_005.md:11:13:	rparen	")"	false	0	0.000000	(11, 13) -> (11, 14)
test/test_005.md:11:14:	semicolon	";"	false	0	0.000000	(11, 14) -> (11, 15)
test/test_005.md:12:0:	ident	"print"	false	0	0.000000	(12, 0) -> (12, 5)
test/test_005.md:12:5:	lparen	"("	false	0	0.000000	(12, 5) -> (12, 6)
test/test_005.md:12:6:	integer	"2"	false	2	0.000000	(12, 6) -> (12, 7)
test/test_005.md:12:7:	ge	">="	false	0	0.000000	(12, 7) -> (12, 9)
test/test_005.md:12:9:	integer	"2"	false	2	0.000000	(12, 9) -> (12, 10)
test/test_005.md:12:10:	rparen	")"	false	0	0.000000	(12, 10) -> (12, 11)
test/test_005.md:12:11:	semicolon	";"	false	0	0.000000	(12, 11) -> (12, 12)
test/test_005.md:13:0:	ident	"print"	false	0	0.000000	(13, 0) -> (13, 5)
test/test_005.md:13:5:	lparen	"("	false	0	0.000000	(13, 5) -> (13, 6)
test/test_005.md:13:6:	integer	"2"	false	2	0.000000	(13, 6) -> (13, 7)
test/test_005.md:13:7:	slash	"/"	false	0	0.000000	(13, 7) -> (13, 8)
test/test_005.md:13:8:	integer	"1"	false	1	0.000000	(13, 8) -> (13, 9)
test/test_005.md:13:9:	ge	">="	false	0	0.000000	(13, 9) -> (13, 11)
test/test_005.md:13:11:	integer	"2"	false	2	0.000000	(13, 11) -> (13, 12)
test/test_005.md:13:12:	slash	"/"	false	0	0.000000	(13, 12) -> (13, 13)
test/test_005.md:13:13:	integer	"1"	false	1	0.000000	(13, 13) -> (13, 14)
test/test_005.md:13:14:	rparen	")"	false	0	0.000000	(13, 14) -> (13, 15)
test/test_005.md:14:0:	eof	""	false	0	0.000000	(14, 0) -> (14, 0)
```
## AST
```scheme
(module ""
  (stmts
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
    (expr2stmt
      (print [void]
        (lt [boolean]
          (number [i64] 1)
          (number [i64] 2)
        )
      )
    )
    (expr2stmt
      (print [void]
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
      (print [void]
        (le [boolean]
          (number [i64] 1)
          (number [i64] 1)
        )
      )
    )
    (expr2stmt
      (print [void]
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
	%0 = icmp eq i32 1, 2
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
	%14 = fcmp ueq double %10, %13
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
	%22 = icmp ne i32 1, 2
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
	%30 = sitofp i32 3 to double
	%31 = sitofp i32 2 to double
	%32 = fdiv double %30, %31
	%33 = sitofp i32 6 to double
	%34 = sitofp i32 4 to double
	%35 = fdiv double %33, %34
	%36 = fcmp une double %32, %35
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
	%44 = icmp slt i32 1, 2
	br i1 %44, label %45, label %47

45:
	%46 = getelementptr [5 x i8], [5 x i8]* @0, i32 0, i32 0
	br label %49

47:
	%48 = getelementptr [6 x i8], [6 x i8]* @1, i32 0, i32 0
	br label %49

49:
	%50 = phi i8* [ %46, %45 ], [ %48, %47 ]
	%51 = call i32 @puts(i8* %50)
	%52 = sitofp i32 2 to double
	%53 = sitofp i32 1 to double
	%54 = fdiv double %52, %53
	%55 = sitofp i32 1 to double
	%56 = sitofp i32 1 to double
	%57 = fdiv double %55, %56
	%58 = fcmp ult double %54, %57
	br i1 %58, label %59, label %61

59:
	%60 = getelementptr [5 x i8], [5 x i8]* @0, i32 0, i32 0
	br label %63

61:
	%62 = getelementptr [6 x i8], [6 x i8]* @1, i32 0, i32 0
	br label %63

63:
	%64 = phi i8* [ %60, %59 ], [ %62, %61 ]
	%65 = call i32 @puts(i8* %64)
	%66 = icmp sle i32 1, 1
	br i1 %66, label %67, label %69

67:
	%68 = getelementptr [5 x i8], [5 x i8]* @0, i32 0, i32 0
	br label %71

69:
	%70 = getelementptr [6 x i8], [6 x i8]* @1, i32 0, i32 0
	br label %71

71:
	%72 = phi i8* [ %68, %67 ], [ %70, %69 ]
	%73 = call i32 @puts(i8* %72)
	%74 = sitofp i32 1 to double
	%75 = sitofp i32 1 to double
	%76 = fdiv double %74, %75
	%77 = sitofp i32 1 to double
	%78 = sitofp i32 1 to double
	%79 = fdiv double %77, %78
	%80 = fcmp ule double %76, %79
	br i1 %80, label %81, label %83

81:
	%82 = getelementptr [5 x i8], [5 x i8]* @0, i32 0, i32 0
	br label %85

83:
	%84 = getelementptr [6 x i8], [6 x i8]* @1, i32 0, i32 0
	br label %85

85:
	%86 = phi i8* [ %82, %81 ], [ %84, %83 ]
	%87 = call i32 @puts(i8* %86)
	%88 = icmp sgt i32 1, 2
	br i1 %88, label %89, label %91

89:
	%90 = getelementptr [5 x i8], [5 x i8]* @0, i32 0, i32 0
	br label %93

91:
	%92 = getelementptr [6 x i8], [6 x i8]* @1, i32 0, i32 0
	br label %93

93:
	%94 = phi i8* [ %90, %89 ], [ %92, %91 ]
	%95 = call i32 @puts(i8* %94)
	%96 = sitofp i32 2 to double
	%97 = sitofp i32 2 to double
	%98 = fdiv double %96, %97
	%99 = sitofp i32 2 to double
	%100 = sitofp i32 1 to double
	%101 = fdiv double %99, %100
	%102 = fcmp ugt double %98, %101
	br i1 %102, label %103, label %105

103:
	%104 = getelementptr [5 x i8], [5 x i8]* @0, i32 0, i32 0
	br label %107

105:
	%106 = getelementptr [6 x i8], [6 x i8]* @1, i32 0, i32 0
	br label %107

107:
	%108 = phi i8* [ %104, %103 ], [ %106, %105 ]
	%109 = call i32 @puts(i8* %108)
	%110 = icmp sge i32 2, 2
	br i1 %110, label %111, label %113

111:
	%112 = getelementptr [5 x i8], [5 x i8]* @0, i32 0, i32 0
	br label %115

113:
	%114 = getelementptr [6 x i8], [6 x i8]* @1, i32 0, i32 0
	br label %115

115:
	%116 = phi i8* [ %112, %111 ], [ %114, %113 ]
	%117 = call i32 @puts(i8* %116)
	%118 = sitofp i32 2 to double
	%119 = sitofp i32 1 to double
	%120 = fdiv double %118, %119
	%121 = sitofp i32 2 to double
	%122 = sitofp i32 1 to double
	%123 = fdiv double %121, %122
	%124 = fcmp uge double %120, %123
	br i1 %124, label %125, label %127

125:
	%126 = getelementptr [5 x i8], [5 x i8]* @0, i32 0, i32 0
	br label %129

127:
	%128 = getelementptr [6 x i8], [6 x i8]* @1, i32 0, i32 0
	br label %129

129:
	%130 = phi i8* [ %126, %125 ], [ %128, %127 ]
	%131 = call i32 @puts(i8* %130)
	ret i32 0
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
