# test/test_005.md
## Source
```pascal
MODULE Comparisons;

BEGIN
    C.print(1=2);
    C.print(3.0/2.0=6.0/4.0);
    C.print(1#2);
    C.print(3 DIV 2#6 DIV 4);
    C.print(1<2);
    C.print(2.0/1.0<1.0/1.0);
    C.print(1<=1);
    C.print(1.0/1.0<=1.0/1.0);
    C.print(1>2);
    C.print(2.0/2.0>2.0/1.0);
    C.print(2>=2);
    C.print(2.0/1.0>=2.0/1.0)
END Comparisons.
```
## Tokens
```tsv
test/test_005.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_005.md:1:8:	ident	"Comparisons"	false	0	0.000000	(1, 8) -> (1, 19)
test/test_005.md:1:19:	semicolon	";"	false	0	0.000000	(1, 19) -> (1, 20)
test/test_005.md:3:0:	begin	"BEGIN"	false	0	0.000000	(3, 0) -> (3, 5)
test/test_005.md:4:4:	ident	"C"	false	0	0.000000	(4, 4) -> (4, 5)
test/test_005.md:4:5:	dot	"."	false	0	0.000000	(4, 5) -> (4, 6)
test/test_005.md:4:6:	ident	"print"	false	0	0.000000	(4, 6) -> (4, 11)
test/test_005.md:4:11:	lparen	"("	false	0	0.000000	(4, 11) -> (4, 12)
test/test_005.md:4:12:	integer	"1"	false	1	0.000000	(4, 12) -> (4, 13)
test/test_005.md:4:13:	eq	"="	false	0	0.000000	(4, 13) -> (4, 14)
test/test_005.md:4:14:	integer	"2"	false	2	0.000000	(4, 14) -> (4, 15)
test/test_005.md:4:15:	rparen	")"	false	0	0.000000	(4, 15) -> (4, 16)
test/test_005.md:4:16:	semicolon	";"	false	0	0.000000	(4, 16) -> (4, 17)
test/test_005.md:5:4:	ident	"C"	false	0	0.000000	(5, 4) -> (5, 5)
test/test_005.md:5:5:	dot	"."	false	0	0.000000	(5, 5) -> (5, 6)
test/test_005.md:5:6:	ident	"print"	false	0	0.000000	(5, 6) -> (5, 11)
test/test_005.md:5:11:	lparen	"("	false	0	0.000000	(5, 11) -> (5, 12)
test/test_005.md:5:12:	real	"3.0"	false	0	3.000000	(5, 12) -> (5, 15)
test/test_005.md:5:15:	slash	"/"	false	0	0.000000	(5, 15) -> (5, 16)
test/test_005.md:5:16:	real	"2.0"	false	0	2.000000	(5, 16) -> (5, 19)
test/test_005.md:5:19:	eq	"="	false	0	0.000000	(5, 19) -> (5, 20)
test/test_005.md:5:20:	real	"6.0"	false	0	6.000000	(5, 20) -> (5, 23)
test/test_005.md:5:23:	slash	"/"	false	0	0.000000	(5, 23) -> (5, 24)
test/test_005.md:5:24:	real	"4.0"	false	0	4.000000	(5, 24) -> (5, 27)
test/test_005.md:5:27:	rparen	")"	false	0	0.000000	(5, 27) -> (5, 28)
test/test_005.md:5:28:	semicolon	";"	false	0	0.000000	(5, 28) -> (5, 29)
test/test_005.md:6:4:	ident	"C"	false	0	0.000000	(6, 4) -> (6, 5)
test/test_005.md:6:5:	dot	"."	false	0	0.000000	(6, 5) -> (6, 6)
test/test_005.md:6:6:	ident	"print"	false	0	0.000000	(6, 6) -> (6, 11)
test/test_005.md:6:11:	lparen	"("	false	0	0.000000	(6, 11) -> (6, 12)
test/test_005.md:6:12:	integer	"1"	false	1	0.000000	(6, 12) -> (6, 13)
test/test_005.md:6:13:	ne	"#"	false	0	0.000000	(6, 13) -> (6, 14)
test/test_005.md:6:14:	integer	"2"	false	2	0.000000	(6, 14) -> (6, 15)
test/test_005.md:6:15:	rparen	")"	false	0	0.000000	(6, 15) -> (6, 16)
test/test_005.md:6:16:	semicolon	";"	false	0	0.000000	(6, 16) -> (6, 17)
test/test_005.md:7:4:	ident	"C"	false	0	0.000000	(7, 4) -> (7, 5)
test/test_005.md:7:5:	dot	"."	false	0	0.000000	(7, 5) -> (7, 6)
test/test_005.md:7:6:	ident	"print"	false	0	0.000000	(7, 6) -> (7, 11)
test/test_005.md:7:11:	lparen	"("	false	0	0.000000	(7, 11) -> (7, 12)
test/test_005.md:7:12:	integer	"3"	false	3	0.000000	(7, 12) -> (7, 13)
test/test_005.md:7:14:	div	"DIV"	false	0	0.000000	(7, 14) -> (7, 17)
test/test_005.md:7:18:	integer	"2"	false	2	0.000000	(7, 18) -> (7, 19)
test/test_005.md:7:19:	ne	"#"	false	0	0.000000	(7, 19) -> (7, 20)
test/test_005.md:7:20:	integer	"6"	false	6	0.000000	(7, 20) -> (7, 21)
test/test_005.md:7:22:	div	"DIV"	false	0	0.000000	(7, 22) -> (7, 25)
test/test_005.md:7:26:	integer	"4"	false	4	0.000000	(7, 26) -> (7, 27)
test/test_005.md:7:27:	rparen	")"	false	0	0.000000	(7, 27) -> (7, 28)
test/test_005.md:7:28:	semicolon	";"	false	0	0.000000	(7, 28) -> (7, 29)
test/test_005.md:8:4:	ident	"C"	false	0	0.000000	(8, 4) -> (8, 5)
test/test_005.md:8:5:	dot	"."	false	0	0.000000	(8, 5) -> (8, 6)
test/test_005.md:8:6:	ident	"print"	false	0	0.000000	(8, 6) -> (8, 11)
test/test_005.md:8:11:	lparen	"("	false	0	0.000000	(8, 11) -> (8, 12)
test/test_005.md:8:12:	integer	"1"	false	1	0.000000	(8, 12) -> (8, 13)
test/test_005.md:8:13:	lt	"<"	false	0	0.000000	(8, 13) -> (8, 14)
test/test_005.md:8:14:	integer	"2"	false	2	0.000000	(8, 14) -> (8, 15)
test/test_005.md:8:15:	rparen	")"	false	0	0.000000	(8, 15) -> (8, 16)
test/test_005.md:8:16:	semicolon	";"	false	0	0.000000	(8, 16) -> (8, 17)
test/test_005.md:9:4:	ident	"C"	false	0	0.000000	(9, 4) -> (9, 5)
test/test_005.md:9:5:	dot	"."	false	0	0.000000	(9, 5) -> (9, 6)
test/test_005.md:9:6:	ident	"print"	false	0	0.000000	(9, 6) -> (9, 11)
test/test_005.md:9:11:	lparen	"("	false	0	0.000000	(9, 11) -> (9, 12)
test/test_005.md:9:12:	real	"2.0"	false	0	2.000000	(9, 12) -> (9, 15)
test/test_005.md:9:15:	slash	"/"	false	0	0.000000	(9, 15) -> (9, 16)
test/test_005.md:9:16:	real	"1.0"	false	0	1.000000	(9, 16) -> (9, 19)
test/test_005.md:9:19:	lt	"<"	false	0	0.000000	(9, 19) -> (9, 20)
test/test_005.md:9:20:	real	"1.0"	false	0	1.000000	(9, 20) -> (9, 23)
test/test_005.md:9:23:	slash	"/"	false	0	0.000000	(9, 23) -> (9, 24)
test/test_005.md:9:24:	real	"1.0"	false	0	1.000000	(9, 24) -> (9, 27)
test/test_005.md:9:27:	rparen	")"	false	0	0.000000	(9, 27) -> (9, 28)
test/test_005.md:9:28:	semicolon	";"	false	0	0.000000	(9, 28) -> (9, 29)
test/test_005.md:10:4:	ident	"C"	false	0	0.000000	(10, 4) -> (10, 5)
test/test_005.md:10:5:	dot	"."	false	0	0.000000	(10, 5) -> (10, 6)
test/test_005.md:10:6:	ident	"print"	false	0	0.000000	(10, 6) -> (10, 11)
test/test_005.md:10:11:	lparen	"("	false	0	0.000000	(10, 11) -> (10, 12)
test/test_005.md:10:12:	integer	"1"	false	1	0.000000	(10, 12) -> (10, 13)
test/test_005.md:10:13:	le	"<="	false	0	0.000000	(10, 13) -> (10, 15)
test/test_005.md:10:15:	integer	"1"	false	1	0.000000	(10, 15) -> (10, 16)
test/test_005.md:10:16:	rparen	")"	false	0	0.000000	(10, 16) -> (10, 17)
test/test_005.md:10:17:	semicolon	";"	false	0	0.000000	(10, 17) -> (10, 18)
test/test_005.md:11:4:	ident	"C"	false	0	0.000000	(11, 4) -> (11, 5)
test/test_005.md:11:5:	dot	"."	false	0	0.000000	(11, 5) -> (11, 6)
test/test_005.md:11:6:	ident	"print"	false	0	0.000000	(11, 6) -> (11, 11)
test/test_005.md:11:11:	lparen	"("	false	0	0.000000	(11, 11) -> (11, 12)
test/test_005.md:11:12:	real	"1.0"	false	0	1.000000	(11, 12) -> (11, 15)
test/test_005.md:11:15:	slash	"/"	false	0	0.000000	(11, 15) -> (11, 16)
test/test_005.md:11:16:	real	"1.0"	false	0	1.000000	(11, 16) -> (11, 19)
test/test_005.md:11:19:	le	"<="	false	0	0.000000	(11, 19) -> (11, 21)
test/test_005.md:11:21:	real	"1.0"	false	0	1.000000	(11, 21) -> (11, 24)
test/test_005.md:11:24:	slash	"/"	false	0	0.000000	(11, 24) -> (11, 25)
test/test_005.md:11:25:	real	"1.0"	false	0	1.000000	(11, 25) -> (11, 28)
test/test_005.md:11:28:	rparen	")"	false	0	0.000000	(11, 28) -> (11, 29)
test/test_005.md:11:29:	semicolon	";"	false	0	0.000000	(11, 29) -> (11, 30)
test/test_005.md:12:4:	ident	"C"	false	0	0.000000	(12, 4) -> (12, 5)
test/test_005.md:12:5:	dot	"."	false	0	0.000000	(12, 5) -> (12, 6)
test/test_005.md:12:6:	ident	"print"	false	0	0.000000	(12, 6) -> (12, 11)
test/test_005.md:12:11:	lparen	"("	false	0	0.000000	(12, 11) -> (12, 12)
test/test_005.md:12:12:	integer	"1"	false	1	0.000000	(12, 12) -> (12, 13)
test/test_005.md:12:13:	gt	">"	false	0	0.000000	(12, 13) -> (12, 14)
test/test_005.md:12:14:	integer	"2"	false	2	0.000000	(12, 14) -> (12, 15)
test/test_005.md:12:15:	rparen	")"	false	0	0.000000	(12, 15) -> (12, 16)
test/test_005.md:12:16:	semicolon	";"	false	0	0.000000	(12, 16) -> (12, 17)
test/test_005.md:13:4:	ident	"C"	false	0	0.000000	(13, 4) -> (13, 5)
test/test_005.md:13:5:	dot	"."	false	0	0.000000	(13, 5) -> (13, 6)
test/test_005.md:13:6:	ident	"print"	false	0	0.000000	(13, 6) -> (13, 11)
test/test_005.md:13:11:	lparen	"("	false	0	0.000000	(13, 11) -> (13, 12)
test/test_005.md:13:12:	real	"2.0"	false	0	2.000000	(13, 12) -> (13, 15)
test/test_005.md:13:15:	slash	"/"	false	0	0.000000	(13, 15) -> (13, 16)
test/test_005.md:13:16:	real	"2.0"	false	0	2.000000	(13, 16) -> (13, 19)
test/test_005.md:13:19:	gt	">"	false	0	0.000000	(13, 19) -> (13, 20)
test/test_005.md:13:20:	real	"2.0"	false	0	2.000000	(13, 20) -> (13, 23)
test/test_005.md:13:23:	slash	"/"	false	0	0.000000	(13, 23) -> (13, 24)
test/test_005.md:13:24:	real	"1.0"	false	0	1.000000	(13, 24) -> (13, 27)
test/test_005.md:13:27:	rparen	")"	false	0	0.000000	(13, 27) -> (13, 28)
test/test_005.md:13:28:	semicolon	";"	false	0	0.000000	(13, 28) -> (13, 29)
test/test_005.md:14:4:	ident	"C"	false	0	0.000000	(14, 4) -> (14, 5)
test/test_005.md:14:5:	dot	"."	false	0	0.000000	(14, 5) -> (14, 6)
test/test_005.md:14:6:	ident	"print"	false	0	0.000000	(14, 6) -> (14, 11)
test/test_005.md:14:11:	lparen	"("	false	0	0.000000	(14, 11) -> (14, 12)
test/test_005.md:14:12:	integer	"2"	false	2	0.000000	(14, 12) -> (14, 13)
test/test_005.md:14:13:	ge	">="	false	0	0.000000	(14, 13) -> (14, 15)
test/test_005.md:14:15:	integer	"2"	false	2	0.000000	(14, 15) -> (14, 16)
test/test_005.md:14:16:	rparen	")"	false	0	0.000000	(14, 16) -> (14, 17)
test/test_005.md:14:17:	semicolon	";"	false	0	0.000000	(14, 17) -> (14, 18)
test/test_005.md:15:4:	ident	"C"	false	0	0.000000	(15, 4) -> (15, 5)
test/test_005.md:15:5:	dot	"."	false	0	0.000000	(15, 5) -> (15, 6)
test/test_005.md:15:6:	ident	"print"	false	0	0.000000	(15, 6) -> (15, 11)
test/test_005.md:15:11:	lparen	"("	false	0	0.000000	(15, 11) -> (15, 12)
test/test_005.md:15:12:	real	"2.0"	false	0	2.000000	(15, 12) -> (15, 15)
test/test_005.md:15:15:	slash	"/"	false	0	0.000000	(15, 15) -> (15, 16)
test/test_005.md:15:16:	real	"1.0"	false	0	1.000000	(15, 16) -> (15, 19)
test/test_005.md:15:19:	ge	">="	false	0	0.000000	(15, 19) -> (15, 21)
test/test_005.md:15:21:	real	"2.0"	false	0	2.000000	(15, 21) -> (15, 24)
test/test_005.md:15:24:	slash	"/"	false	0	0.000000	(15, 24) -> (15, 25)
test/test_005.md:15:25:	real	"1.0"	false	0	1.000000	(15, 25) -> (15, 28)
test/test_005.md:15:28:	rparen	")"	false	0	0.000000	(15, 28) -> (15, 29)
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
        (procedure [void] "C.print")
        (eq [boolean]
          (number [i64] 1)
          (number [i64] 2)
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "C.print")
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
        (procedure [void] "C.print")
        (ne [boolean]
          (number [i64] 1)
          (number [i64] 2)
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "C.print")
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
        (procedure [void] "C.print")
        (lt [boolean]
          (number [i64] 1)
          (number [i64] 2)
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "C.print")
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
        (procedure [void] "C.print")
        (le [boolean]
          (number [i64] 1)
          (number [i64] 1)
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "C.print")
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
        (procedure [void] "C.print")
        (gt [boolean]
          (number [i64] 1)
          (number [i64] 2)
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "C.print")
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
        (procedure [void] "C.print")
        (ge [boolean]
          (number [i64] 2)
          (number [i64] 2)
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "C.print")
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
