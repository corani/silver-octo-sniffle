# test/test_009.md
## Source
```pascal
MODULE GlobalConsts;

CONST x = 1;
      y = 2.2;
      z = FLT(x) + y;
      ne = (FLT(x) # y);

BEGIN
    IF FLT(x) = y THEN
        print("Equal")
    ELSE
        print("Not Equal")
    END;

    IF FLT(x) + y = z THEN
        print("Equal")
    ELSE
        print("Not Equal")
    END;

    IF ne THEN
        print("Not Equal")
    ELSE
        print("Equal")
    END;
END GlobalConsts.
```
## Tokens
```tsv
test/test_009.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_009.md:1:8:	ident	"GlobalConsts"	false	0	0.000000	(1, 8) -> (1, 20)
test/test_009.md:1:20:	semicolon	";"	false	0	0.000000	(1, 20) -> (1, 21)
test/test_009.md:3:0:	const	"CONST"	false	0	0.000000	(3, 0) -> (3, 5)
test/test_009.md:3:6:	ident	"x"	false	0	0.000000	(3, 6) -> (3, 7)
test/test_009.md:3:8:	eq	"="	false	0	0.000000	(3, 8) -> (3, 9)
test/test_009.md:3:10:	integer	"1"	false	1	0.000000	(3, 10) -> (3, 11)
test/test_009.md:3:11:	semicolon	";"	false	0	0.000000	(3, 11) -> (3, 12)
test/test_009.md:4:6:	ident	"y"	false	0	0.000000	(4, 6) -> (4, 7)
test/test_009.md:4:8:	eq	"="	false	0	0.000000	(4, 8) -> (4, 9)
test/test_009.md:4:10:	real	"2.2"	false	0	2.200000	(4, 10) -> (4, 13)
test/test_009.md:4:13:	semicolon	";"	false	0	0.000000	(4, 13) -> (4, 14)
test/test_009.md:5:6:	ident	"z"	false	0	0.000000	(5, 6) -> (5, 7)
test/test_009.md:5:8:	eq	"="	false	0	0.000000	(5, 8) -> (5, 9)
test/test_009.md:5:10:	ident	"FLT"	false	0	0.000000	(5, 10) -> (5, 13)
test/test_009.md:5:13:	lparen	"("	false	0	0.000000	(5, 13) -> (5, 14)
test/test_009.md:5:14:	ident	"x"	false	0	0.000000	(5, 14) -> (5, 15)
test/test_009.md:5:15:	rparen	")"	false	0	0.000000	(5, 15) -> (5, 16)
test/test_009.md:5:17:	plus	"+"	false	0	0.000000	(5, 17) -> (5, 18)
test/test_009.md:5:19:	ident	"y"	false	0	0.000000	(5, 19) -> (5, 20)
test/test_009.md:5:20:	semicolon	";"	false	0	0.000000	(5, 20) -> (5, 21)
test/test_009.md:6:6:	ident	"ne"	false	0	0.000000	(6, 6) -> (6, 8)
test/test_009.md:6:9:	eq	"="	false	0	0.000000	(6, 9) -> (6, 10)
test/test_009.md:6:11:	lparen	"("	false	0	0.000000	(6, 11) -> (6, 12)
test/test_009.md:6:12:	ident	"FLT"	false	0	0.000000	(6, 12) -> (6, 15)
test/test_009.md:6:15:	lparen	"("	false	0	0.000000	(6, 15) -> (6, 16)
test/test_009.md:6:16:	ident	"x"	false	0	0.000000	(6, 16) -> (6, 17)
test/test_009.md:6:17:	rparen	")"	false	0	0.000000	(6, 17) -> (6, 18)
test/test_009.md:6:19:	ne	"#"	false	0	0.000000	(6, 19) -> (6, 20)
test/test_009.md:6:21:	ident	"y"	false	0	0.000000	(6, 21) -> (6, 22)
test/test_009.md:6:22:	rparen	")"	false	0	0.000000	(6, 22) -> (6, 23)
test/test_009.md:6:23:	semicolon	";"	false	0	0.000000	(6, 23) -> (6, 24)
test/test_009.md:8:0:	begin	"BEGIN"	false	0	0.000000	(8, 0) -> (8, 5)
test/test_009.md:9:4:	if	"IF"	false	0	0.000000	(9, 4) -> (9, 6)
test/test_009.md:9:7:	ident	"FLT"	false	0	0.000000	(9, 7) -> (9, 10)
test/test_009.md:9:10:	lparen	"("	false	0	0.000000	(9, 10) -> (9, 11)
test/test_009.md:9:11:	ident	"x"	false	0	0.000000	(9, 11) -> (9, 12)
test/test_009.md:9:12:	rparen	")"	false	0	0.000000	(9, 12) -> (9, 13)
test/test_009.md:9:14:	eq	"="	false	0	0.000000	(9, 14) -> (9, 15)
test/test_009.md:9:16:	ident	"y"	false	0	0.000000	(9, 16) -> (9, 17)
test/test_009.md:9:18:	then	"THEN"	false	0	0.000000	(9, 18) -> (9, 22)
test/test_009.md:10:8:	ident	"print"	false	0	0.000000	(10, 8) -> (10, 13)
test/test_009.md:10:13:	lparen	"("	false	0	0.000000	(10, 13) -> (10, 14)
test/test_009.md:10:14:	string	"Equal"	false	0	0.000000	(10, 14) -> (10, 21)
test/test_009.md:10:21:	rparen	")"	false	0	0.000000	(10, 21) -> (10, 22)
test/test_009.md:11:4:	else	"ELSE"	false	0	0.000000	(11, 4) -> (11, 8)
test/test_009.md:12:8:	ident	"print"	false	0	0.000000	(12, 8) -> (12, 13)
test/test_009.md:12:13:	lparen	"("	false	0	0.000000	(12, 13) -> (12, 14)
test/test_009.md:12:14:	string	"Not Equal"	false	0	0.000000	(12, 14) -> (12, 25)
test/test_009.md:12:25:	rparen	")"	false	0	0.000000	(12, 25) -> (12, 26)
test/test_009.md:13:4:	end	"END"	false	0	0.000000	(13, 4) -> (13, 7)
test/test_009.md:13:7:	semicolon	";"	false	0	0.000000	(13, 7) -> (13, 8)
test/test_009.md:15:4:	if	"IF"	false	0	0.000000	(15, 4) -> (15, 6)
test/test_009.md:15:7:	ident	"FLT"	false	0	0.000000	(15, 7) -> (15, 10)
test/test_009.md:15:10:	lparen	"("	false	0	0.000000	(15, 10) -> (15, 11)
test/test_009.md:15:11:	ident	"x"	false	0	0.000000	(15, 11) -> (15, 12)
test/test_009.md:15:12:	rparen	")"	false	0	0.000000	(15, 12) -> (15, 13)
test/test_009.md:15:14:	plus	"+"	false	0	0.000000	(15, 14) -> (15, 15)
test/test_009.md:15:16:	ident	"y"	false	0	0.000000	(15, 16) -> (15, 17)
test/test_009.md:15:18:	eq	"="	false	0	0.000000	(15, 18) -> (15, 19)
test/test_009.md:15:20:	ident	"z"	false	0	0.000000	(15, 20) -> (15, 21)
test/test_009.md:15:22:	then	"THEN"	false	0	0.000000	(15, 22) -> (15, 26)
test/test_009.md:16:8:	ident	"print"	false	0	0.000000	(16, 8) -> (16, 13)
test/test_009.md:16:13:	lparen	"("	false	0	0.000000	(16, 13) -> (16, 14)
test/test_009.md:16:14:	string	"Equal"	false	0	0.000000	(16, 14) -> (16, 21)
test/test_009.md:16:21:	rparen	")"	false	0	0.000000	(16, 21) -> (16, 22)
test/test_009.md:17:4:	else	"ELSE"	false	0	0.000000	(17, 4) -> (17, 8)
test/test_009.md:18:8:	ident	"print"	false	0	0.000000	(18, 8) -> (18, 13)
test/test_009.md:18:13:	lparen	"("	false	0	0.000000	(18, 13) -> (18, 14)
test/test_009.md:18:14:	string	"Not Equal"	false	0	0.000000	(18, 14) -> (18, 25)
test/test_009.md:18:25:	rparen	")"	false	0	0.000000	(18, 25) -> (18, 26)
test/test_009.md:19:4:	end	"END"	false	0	0.000000	(19, 4) -> (19, 7)
test/test_009.md:19:7:	semicolon	";"	false	0	0.000000	(19, 7) -> (19, 8)
test/test_009.md:21:4:	if	"IF"	false	0	0.000000	(21, 4) -> (21, 6)
test/test_009.md:21:7:	ident	"ne"	false	0	0.000000	(21, 7) -> (21, 9)
test/test_009.md:21:10:	then	"THEN"	false	0	0.000000	(21, 10) -> (21, 14)
test/test_009.md:22:8:	ident	"print"	false	0	0.000000	(22, 8) -> (22, 13)
test/test_009.md:22:13:	lparen	"("	false	0	0.000000	(22, 13) -> (22, 14)
test/test_009.md:22:14:	string	"Not Equal"	false	0	0.000000	(22, 14) -> (22, 25)
test/test_009.md:22:25:	rparen	")"	false	0	0.000000	(22, 25) -> (22, 26)
test/test_009.md:23:4:	else	"ELSE"	false	0	0.000000	(23, 4) -> (23, 8)
test/test_009.md:24:8:	ident	"print"	false	0	0.000000	(24, 8) -> (24, 13)
test/test_009.md:24:13:	lparen	"("	false	0	0.000000	(24, 13) -> (24, 14)
test/test_009.md:24:14:	string	"Equal"	false	0	0.000000	(24, 14) -> (24, 21)
test/test_009.md:24:21:	rparen	")"	false	0	0.000000	(24, 21) -> (24, 22)
test/test_009.md:25:4:	end	"END"	false	0	0.000000	(25, 4) -> (25, 7)
test/test_009.md:25:7:	semicolon	";"	false	0	0.000000	(25, 7) -> (25, 8)
test/test_009.md:26:0:	end	"END"	false	0	0.000000	(26, 0) -> (26, 3)
test/test_009.md:26:4:	ident	"GlobalConsts"	false	0	0.000000	(26, 4) -> (26, 16)
test/test_009.md:26:16:	dot	"."	false	0	0.000000	(26, 16) -> (26, 17)
test/test_009.md:27:0:	eof	""	false	0	0.000000	(27, 0) -> (27, 0)
```
## AST
```scheme
(module "GlobalConsts"
  (consts
    (x [i64]
      (number [i64] 1)
    )
    (y [f64]
      (number [f64] 2.200000)
    )
    (z [f64]
      (plus [f64]
        (call
          (procedure [f64] "FLT")
          (constant [i64] "x")
        )
        (constant [f64] "y")
      )
    )
    (ne [boolean]
      (ne [boolean]
        (call
          (procedure [f64] "FLT")
          (constant [i64] "x")
        )
        (constant [f64] "y")
      )
    )
  )
  (stmts
    (if
      (eq [boolean]
        (call
          (procedure [f64] "FLT")
          (constant [i64] "x")
        )
        (constant [f64] "y")
      )
      (stmts
        (expr2stmt
          (call
            (procedure [void] "print")
            (string "Equal")
          )
        )
      )
      (stmts
        (expr2stmt
          (call
            (procedure [void] "print")
            (string "Not Equal")
          )
        )
      )
    )
    (if
      (eq [boolean]
        (plus [f64]
          (call
            (procedure [f64] "FLT")
            (constant [i64] "x")
          )
          (constant [f64] "y")
        )
        (constant [f64] "z")
      )
      (stmts
        (expr2stmt
          (call
            (procedure [void] "print")
            (string "Equal")
          )
        )
      )
      (stmts
        (expr2stmt
          (call
            (procedure [void] "print")
            (string "Not Equal")
          )
        )
      )
    )
    (if
      (constant [boolean] "ne")
      (stmts
        (expr2stmt
          (call
            (procedure [void] "print")
            (string "Not Equal")
          )
        )
      )
      (stmts
        (expr2stmt
          (call
            (procedure [void] "print")
            (string "Equal")
          )
        )
      )
    )
  )
)
```
## IR
```llvm
@0 = global [6 x i8] c"Equal\00"
@1 = global [10 x i8] c"Not Equal\00"

declare i64 @puts(i8* %str)

declare i64 @rand()

declare i64 @sprintf(i8* %buf, i8* %format, ...)

declare i64 @printf(i8* %format, ...)

declare i8* @malloc(i64 %size)

define i64 @main() {
entry:
	%0 = sitofp i64 1 to double
	%1 = fcmp ueq double %0, 0x400199999999999A
	br i1 %1, label %2, label %5

2:
	%3 = getelementptr [6 x i8], [6 x i8]* @0, i64 0, i64 0
	%4 = call i64 @puts(i8* %3)
	br label %8

5:
	%6 = getelementptr [10 x i8], [10 x i8]* @1, i64 0, i64 0
	%7 = call i64 @puts(i8* %6)
	br label %8

8:
	%9 = sitofp i64 1 to double
	%10 = fadd double %9, 0x400199999999999A
	%11 = fcmp ueq double %10, 0x400999999999999A
	br i1 %11, label %12, label %15

12:
	%13 = getelementptr [6 x i8], [6 x i8]* @0, i64 0, i64 0
	%14 = call i64 @puts(i8* %13)
	br label %18

15:
	%16 = getelementptr [10 x i8], [10 x i8]* @1, i64 0, i64 0
	%17 = call i64 @puts(i8* %16)
	br label %18

18:
	br i1 true, label %19, label %22

19:
	%20 = getelementptr [10 x i8], [10 x i8]* @1, i64 0, i64 0
	%21 = call i64 @puts(i8* %20)
	br label %25

22:
	%23 = getelementptr [6 x i8], [6 x i8]* @0, i64 0, i64 0
	%24 = call i64 @puts(i8* %23)
	br label %25

25:
	ret i64 0
}

```
## Run
```bash
Not Equal
Equal
Not Equal
```
