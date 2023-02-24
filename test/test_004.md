# test/test_004.md
## Source
```pascal
MODULE BooleanOperators;

BEGIN
    C.print(~TRUE);
    C.print(~FALSE);
    C.print(TRUE & FALSE);
    C.print(TRUE OR FALSE);
    C.print(ORD(TRUE));
    C.print(ORD(FALSE))
END BooleanOperators.
```
## Tokens
```tsv
test/test_004.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_004.md:1:8:	ident	"BooleanOperators"	false	0	0.000000	(1, 8) -> (1, 24)
test/test_004.md:1:24:	semicolon	";"	false	0	0.000000	(1, 24) -> (1, 25)
test/test_004.md:3:0:	begin	"BEGIN"	false	0	0.000000	(3, 0) -> (3, 5)
test/test_004.md:4:4:	ident	"C"	false	0	0.000000	(4, 4) -> (4, 5)
test/test_004.md:4:5:	dot	"."	false	0	0.000000	(4, 5) -> (4, 6)
test/test_004.md:4:6:	ident	"print"	false	0	0.000000	(4, 6) -> (4, 11)
test/test_004.md:4:11:	lparen	"("	false	0	0.000000	(4, 11) -> (4, 12)
test/test_004.md:4:12:	tilde	"~"	false	0	0.000000	(4, 12) -> (4, 13)
test/test_004.md:4:13:	boolean	"TRUE"	true	0	0.000000	(4, 13) -> (4, 17)
test/test_004.md:4:17:	rparen	")"	false	0	0.000000	(4, 17) -> (4, 18)
test/test_004.md:4:18:	semicolon	";"	false	0	0.000000	(4, 18) -> (4, 19)
test/test_004.md:5:4:	ident	"C"	false	0	0.000000	(5, 4) -> (5, 5)
test/test_004.md:5:5:	dot	"."	false	0	0.000000	(5, 5) -> (5, 6)
test/test_004.md:5:6:	ident	"print"	false	0	0.000000	(5, 6) -> (5, 11)
test/test_004.md:5:11:	lparen	"("	false	0	0.000000	(5, 11) -> (5, 12)
test/test_004.md:5:12:	tilde	"~"	false	0	0.000000	(5, 12) -> (5, 13)
test/test_004.md:5:13:	boolean	"FALSE"	false	0	0.000000	(5, 13) -> (5, 18)
test/test_004.md:5:18:	rparen	")"	false	0	0.000000	(5, 18) -> (5, 19)
test/test_004.md:5:19:	semicolon	";"	false	0	0.000000	(5, 19) -> (5, 20)
test/test_004.md:6:4:	ident	"C"	false	0	0.000000	(6, 4) -> (6, 5)
test/test_004.md:6:5:	dot	"."	false	0	0.000000	(6, 5) -> (6, 6)
test/test_004.md:6:6:	ident	"print"	false	0	0.000000	(6, 6) -> (6, 11)
test/test_004.md:6:11:	lparen	"("	false	0	0.000000	(6, 11) -> (6, 12)
test/test_004.md:6:12:	boolean	"TRUE"	true	0	0.000000	(6, 12) -> (6, 16)
test/test_004.md:6:17:	ampersand	"&"	false	0	0.000000	(6, 17) -> (6, 18)
test/test_004.md:6:19:	boolean	"FALSE"	false	0	0.000000	(6, 19) -> (6, 24)
test/test_004.md:6:24:	rparen	")"	false	0	0.000000	(6, 24) -> (6, 25)
test/test_004.md:6:25:	semicolon	";"	false	0	0.000000	(6, 25) -> (6, 26)
test/test_004.md:7:4:	ident	"C"	false	0	0.000000	(7, 4) -> (7, 5)
test/test_004.md:7:5:	dot	"."	false	0	0.000000	(7, 5) -> (7, 6)
test/test_004.md:7:6:	ident	"print"	false	0	0.000000	(7, 6) -> (7, 11)
test/test_004.md:7:11:	lparen	"("	false	0	0.000000	(7, 11) -> (7, 12)
test/test_004.md:7:12:	boolean	"TRUE"	true	0	0.000000	(7, 12) -> (7, 16)
test/test_004.md:7:17:	or	"OR"	false	0	0.000000	(7, 17) -> (7, 19)
test/test_004.md:7:20:	boolean	"FALSE"	false	0	0.000000	(7, 20) -> (7, 25)
test/test_004.md:7:25:	rparen	")"	false	0	0.000000	(7, 25) -> (7, 26)
test/test_004.md:7:26:	semicolon	";"	false	0	0.000000	(7, 26) -> (7, 27)
test/test_004.md:8:4:	ident	"C"	false	0	0.000000	(8, 4) -> (8, 5)
test/test_004.md:8:5:	dot	"."	false	0	0.000000	(8, 5) -> (8, 6)
test/test_004.md:8:6:	ident	"print"	false	0	0.000000	(8, 6) -> (8, 11)
test/test_004.md:8:11:	lparen	"("	false	0	0.000000	(8, 11) -> (8, 12)
test/test_004.md:8:12:	ident	"ORD"	false	0	0.000000	(8, 12) -> (8, 15)
test/test_004.md:8:15:	lparen	"("	false	0	0.000000	(8, 15) -> (8, 16)
test/test_004.md:8:16:	boolean	"TRUE"	true	0	0.000000	(8, 16) -> (8, 20)
test/test_004.md:8:20:	rparen	")"	false	0	0.000000	(8, 20) -> (8, 21)
test/test_004.md:8:21:	rparen	")"	false	0	0.000000	(8, 21) -> (8, 22)
test/test_004.md:8:22:	semicolon	";"	false	0	0.000000	(8, 22) -> (8, 23)
test/test_004.md:9:4:	ident	"C"	false	0	0.000000	(9, 4) -> (9, 5)
test/test_004.md:9:5:	dot	"."	false	0	0.000000	(9, 5) -> (9, 6)
test/test_004.md:9:6:	ident	"print"	false	0	0.000000	(9, 6) -> (9, 11)
test/test_004.md:9:11:	lparen	"("	false	0	0.000000	(9, 11) -> (9, 12)
test/test_004.md:9:12:	ident	"ORD"	false	0	0.000000	(9, 12) -> (9, 15)
test/test_004.md:9:15:	lparen	"("	false	0	0.000000	(9, 15) -> (9, 16)
test/test_004.md:9:16:	boolean	"FALSE"	false	0	0.000000	(9, 16) -> (9, 21)
test/test_004.md:9:21:	rparen	")"	false	0	0.000000	(9, 21) -> (9, 22)
test/test_004.md:9:22:	rparen	")"	false	0	0.000000	(9, 22) -> (9, 23)
test/test_004.md:10:0:	end	"END"	false	0	0.000000	(10, 0) -> (10, 3)
test/test_004.md:10:4:	ident	"BooleanOperators"	false	0	0.000000	(10, 4) -> (10, 20)
test/test_004.md:10:20:	dot	"."	false	0	0.000000	(10, 20) -> (10, 21)
test/test_004.md:11:0:	eof	""	false	0	0.000000	(11, 0) -> (11, 0)
```
## AST
```scheme
(module "BooleanOperators"
  (stmts
    (expr2stmt
      (call
        (procedure [void] "C.print")
        (not [boolean]
          #true
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "C.print")
        (not [boolean]
          #false
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "C.print")
        (ampersand [boolean]
          #true
          #false
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "C.print")
        (or [boolean]
          #true
          #false
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "C.print")
        (call
          (procedure [i64] "ORD")
          #true
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "C.print")
        (call
          (procedure [i64] "ORD")
          #false
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
@2 = global [4 x i8] c"%d\0A\00"

declare i64 @puts(i8* %str)

declare i64 @rand()

declare i64 @sprintf(i8* %buf, i8* %format, ...)

declare i64 @printf(i8* %format, ...)

declare i8* @malloc(i64 %size)

declare i8* @free(i8* %ptr)

define i64 @main() {
entry:
	%0 = icmp eq i1 true, false
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
	%8 = icmp eq i1 false, false
	br i1 %8, label %9, label %11

9:
	%10 = getelementptr [5 x i8], [5 x i8]* @0, i64 0, i64 0
	br label %13

11:
	%12 = getelementptr [6 x i8], [6 x i8]* @1, i64 0, i64 0
	br label %13

13:
	%14 = phi i8* [ %10, %9 ], [ %12, %11 ]
	%15 = call i64 @puts(i8* %14)
	%16 = and i1 true, false
	br i1 %16, label %17, label %19

17:
	%18 = getelementptr [5 x i8], [5 x i8]* @0, i64 0, i64 0
	br label %21

19:
	%20 = getelementptr [6 x i8], [6 x i8]* @1, i64 0, i64 0
	br label %21

21:
	%22 = phi i8* [ %18, %17 ], [ %20, %19 ]
	%23 = call i64 @puts(i8* %22)
	%24 = or i1 true, false
	br i1 %24, label %25, label %27

25:
	%26 = getelementptr [5 x i8], [5 x i8]* @0, i64 0, i64 0
	br label %29

27:
	%28 = getelementptr [6 x i8], [6 x i8]* @1, i64 0, i64 0
	br label %29

29:
	%30 = phi i8* [ %26, %25 ], [ %28, %27 ]
	%31 = call i64 @puts(i8* %30)
	%32 = zext i1 true to i64
	%33 = getelementptr [4 x i8], [4 x i8]* @2, i64 0, i64 0
	%34 = call i64 (i8*, ...) @printf(i8* %33, i64 %32)
	%35 = zext i1 false to i64
	%36 = getelementptr [4 x i8], [4 x i8]* @2, i64 0, i64 0
	%37 = call i64 (i8*, ...) @printf(i8* %36, i64 %35)
	ret i64 0
}

```
## Run
```bash
FALSE
TRUE
FALSE
TRUE
1
0
```
