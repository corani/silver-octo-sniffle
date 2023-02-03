# test/test_004.md
## Source
```pascal
MODULE BooleanOperators;

BEGIN
    print(~TRUE);
    print(~FALSE);
    print(TRUE & FALSE);
    print(TRUE OR FALSE)
END BooleanOperators.
```
## Tokens
```tsv
test/test_004.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_004.md:1:8:	ident	"BooleanOperators"	false	0	0.000000	(1, 8) -> (1, 24)
test/test_004.md:1:24:	semicolon	";"	false	0	0.000000	(1, 24) -> (1, 25)
test/test_004.md:3:0:	begin	"BEGIN"	false	0	0.000000	(3, 0) -> (3, 5)
test/test_004.md:4:4:	ident	"print"	false	0	0.000000	(4, 4) -> (4, 9)
test/test_004.md:4:9:	lparen	"("	false	0	0.000000	(4, 9) -> (4, 10)
test/test_004.md:4:10:	tilde	"~"	false	0	0.000000	(4, 10) -> (4, 11)
test/test_004.md:4:11:	boolean	"TRUE"	true	0	0.000000	(4, 11) -> (4, 15)
test/test_004.md:4:15:	rparen	")"	false	0	0.000000	(4, 15) -> (4, 16)
test/test_004.md:4:16:	semicolon	";"	false	0	0.000000	(4, 16) -> (4, 17)
test/test_004.md:5:4:	ident	"print"	false	0	0.000000	(5, 4) -> (5, 9)
test/test_004.md:5:9:	lparen	"("	false	0	0.000000	(5, 9) -> (5, 10)
test/test_004.md:5:10:	tilde	"~"	false	0	0.000000	(5, 10) -> (5, 11)
test/test_004.md:5:11:	boolean	"FALSE"	false	0	0.000000	(5, 11) -> (5, 16)
test/test_004.md:5:16:	rparen	")"	false	0	0.000000	(5, 16) -> (5, 17)
test/test_004.md:5:17:	semicolon	";"	false	0	0.000000	(5, 17) -> (5, 18)
test/test_004.md:6:4:	ident	"print"	false	0	0.000000	(6, 4) -> (6, 9)
test/test_004.md:6:9:	lparen	"("	false	0	0.000000	(6, 9) -> (6, 10)
test/test_004.md:6:10:	boolean	"TRUE"	true	0	0.000000	(6, 10) -> (6, 14)
test/test_004.md:6:15:	ampersand	"&"	false	0	0.000000	(6, 15) -> (6, 16)
test/test_004.md:6:17:	boolean	"FALSE"	false	0	0.000000	(6, 17) -> (6, 22)
test/test_004.md:6:22:	rparen	")"	false	0	0.000000	(6, 22) -> (6, 23)
test/test_004.md:6:23:	semicolon	";"	false	0	0.000000	(6, 23) -> (6, 24)
test/test_004.md:7:4:	ident	"print"	false	0	0.000000	(7, 4) -> (7, 9)
test/test_004.md:7:9:	lparen	"("	false	0	0.000000	(7, 9) -> (7, 10)
test/test_004.md:7:10:	boolean	"TRUE"	true	0	0.000000	(7, 10) -> (7, 14)
test/test_004.md:7:15:	or	"OR"	false	0	0.000000	(7, 15) -> (7, 17)
test/test_004.md:7:18:	boolean	"FALSE"	false	0	0.000000	(7, 18) -> (7, 23)
test/test_004.md:7:23:	rparen	")"	false	0	0.000000	(7, 23) -> (7, 24)
test/test_004.md:8:0:	end	"END"	false	0	0.000000	(8, 0) -> (8, 3)
test/test_004.md:8:4:	ident	"BooleanOperators"	false	0	0.000000	(8, 4) -> (8, 20)
test/test_004.md:8:20:	dot	"."	false	0	0.000000	(8, 20) -> (8, 21)
test/test_004.md:9:0:	eof	""	false	0	0.000000	(9, 0) -> (9, 0)
```
## AST
```scheme
(module "BooleanOperators"
  (stmts
    (expr2stmt
      (print [void]
        (not [boolean]
          #true
        )
      )
    )
    (expr2stmt
      (print [void]
        (not [boolean]
          #false
        )
      )
    )
    (expr2stmt
      (print [void]
        (ampersand [boolean]
          #true
          #false
        )
      )
    )
    (expr2stmt
      (print [void]
        (or [boolean]
          #true
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

declare i64 @puts(i8* %str)

declare i64 @rand()

declare i64 @sprintf(i8* %buf, i8* %format, ...)

declare i64 @printf(i8* %format, ...)

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
	ret i64 0
}

```
## Run
```bash
FALSE
TRUE
FALSE
TRUE
```
