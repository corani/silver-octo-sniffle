# test/test_009.md
## Source
```pascal
MODULE GlobalConsts;

CONST x = 1;
      y = 2;
      z = x + y;

BEGIN
    IF x = y THEN
        print("Equal")
    ELSE
        print("Not Equal")
    END;
    IF x + y = z THEN
        print("Equal")
    ELSE
        print("Not Equal")
    END
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
test/test_009.md:4:10:	integer	"2"	false	2	0.000000	(4, 10) -> (4, 11)
test/test_009.md:4:11:	semicolon	";"	false	0	0.000000	(4, 11) -> (4, 12)
test/test_009.md:5:6:	ident	"z"	false	0	0.000000	(5, 6) -> (5, 7)
test/test_009.md:5:8:	eq	"="	false	0	0.000000	(5, 8) -> (5, 9)
test/test_009.md:5:10:	ident	"x"	false	0	0.000000	(5, 10) -> (5, 11)
test/test_009.md:5:12:	plus	"+"	false	0	0.000000	(5, 12) -> (5, 13)
test/test_009.md:5:14:	ident	"y"	false	0	0.000000	(5, 14) -> (5, 15)
test/test_009.md:5:15:	semicolon	";"	false	0	0.000000	(5, 15) -> (5, 16)
test/test_009.md:7:0:	begin	"BEGIN"	false	0	0.000000	(7, 0) -> (7, 5)
test/test_009.md:8:4:	if	"IF"	false	0	0.000000	(8, 4) -> (8, 6)
test/test_009.md:8:7:	ident	"x"	false	0	0.000000	(8, 7) -> (8, 8)
test/test_009.md:8:9:	eq	"="	false	0	0.000000	(8, 9) -> (8, 10)
test/test_009.md:8:11:	ident	"y"	false	0	0.000000	(8, 11) -> (8, 12)
test/test_009.md:8:13:	then	"THEN"	false	0	0.000000	(8, 13) -> (8, 17)
test/test_009.md:9:8:	ident	"print"	false	0	0.000000	(9, 8) -> (9, 13)
test/test_009.md:9:13:	lparen	"("	false	0	0.000000	(9, 13) -> (9, 14)
test/test_009.md:9:14:	string	"Equal"	false	0	0.000000	(9, 14) -> (9, 21)
test/test_009.md:9:21:	rparen	")"	false	0	0.000000	(9, 21) -> (9, 22)
test/test_009.md:10:4:	else	"ELSE"	false	0	0.000000	(10, 4) -> (10, 8)
test/test_009.md:11:8:	ident	"print"	false	0	0.000000	(11, 8) -> (11, 13)
test/test_009.md:11:13:	lparen	"("	false	0	0.000000	(11, 13) -> (11, 14)
test/test_009.md:11:14:	string	"Not Equal"	false	0	0.000000	(11, 14) -> (11, 25)
test/test_009.md:11:25:	rparen	")"	false	0	0.000000	(11, 25) -> (11, 26)
test/test_009.md:12:4:	end	"END"	false	0	0.000000	(12, 4) -> (12, 7)
test/test_009.md:12:7:	semicolon	";"	false	0	0.000000	(12, 7) -> (12, 8)
test/test_009.md:13:4:	if	"IF"	false	0	0.000000	(13, 4) -> (13, 6)
test/test_009.md:13:7:	ident	"x"	false	0	0.000000	(13, 7) -> (13, 8)
test/test_009.md:13:9:	plus	"+"	false	0	0.000000	(13, 9) -> (13, 10)
test/test_009.md:13:11:	ident	"y"	false	0	0.000000	(13, 11) -> (13, 12)
test/test_009.md:13:13:	eq	"="	false	0	0.000000	(13, 13) -> (13, 14)
test/test_009.md:13:15:	ident	"z"	false	0	0.000000	(13, 15) -> (13, 16)
test/test_009.md:13:17:	then	"THEN"	false	0	0.000000	(13, 17) -> (13, 21)
test/test_009.md:14:8:	ident	"print"	false	0	0.000000	(14, 8) -> (14, 13)
test/test_009.md:14:13:	lparen	"("	false	0	0.000000	(14, 13) -> (14, 14)
test/test_009.md:14:14:	string	"Equal"	false	0	0.000000	(14, 14) -> (14, 21)
test/test_009.md:14:21:	rparen	")"	false	0	0.000000	(14, 21) -> (14, 22)
test/test_009.md:15:4:	else	"ELSE"	false	0	0.000000	(15, 4) -> (15, 8)
test/test_009.md:16:8:	ident	"print"	false	0	0.000000	(16, 8) -> (16, 13)
test/test_009.md:16:13:	lparen	"("	false	0	0.000000	(16, 13) -> (16, 14)
test/test_009.md:16:14:	string	"Not Equal"	false	0	0.000000	(16, 14) -> (16, 25)
test/test_009.md:16:25:	rparen	")"	false	0	0.000000	(16, 25) -> (16, 26)
test/test_009.md:17:4:	end	"END"	false	0	0.000000	(17, 4) -> (17, 7)
test/test_009.md:18:0:	end	"END"	false	0	0.000000	(18, 0) -> (18, 3)
test/test_009.md:18:4:	ident	"GlobalConsts"	false	0	0.000000	(18, 4) -> (18, 16)
test/test_009.md:18:16:	dot	"."	false	0	0.000000	(18, 16) -> (18, 17)
test/test_009.md:19:0:	eof	""	false	0	0.000000	(19, 0) -> (19, 0)
```
## AST
```scheme
(module "GlobalConsts"
  (consts
    (x [i64]
      (number [i64] 1)
    )
    (y [i64]
      (number [i64] 2)
    )
    (z [i64]
      (plus [i64]
        (constant [i64] "x")
        (constant [i64] "y")
      )
    )
  )
  (stmts
    (if
      (eq [boolean]
        (constant [i64] "x")
        (constant [i64] "y")
      )
      (stmts
        (expr2stmt
          (print [void]
            (string "Equal")
          )
        )
      )
      (stmts
        (expr2stmt
          (print [void]
            (string "Not Equal")
          )
        )
      )
    )
    (if
      (eq [boolean]
        (plus [i64]
          (constant [i64] "x")
          (constant [i64] "y")
        )
        (constant [i64] "z")
      )
      (stmts
        (expr2stmt
          (print [void]
            (string "Equal")
          )
        )
      )
      (stmts
        (expr2stmt
          (print [void]
            (string "Not Equal")
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

define i64 @main() {
entry:
	%0 = icmp eq i64 1, 2
	br i1 %0, label %1, label %4

1:
	%2 = getelementptr [6 x i8], [6 x i8]* @0, i64 0, i64 0
	%3 = call i64 @puts(i8* %2)
	br label %7

4:
	%5 = getelementptr [10 x i8], [10 x i8]* @1, i64 0, i64 0
	%6 = call i64 @puts(i8* %5)
	br label %7

7:
	%8 = add i64 1, 2
	%9 = icmp eq i64 %8, 3
	br i1 %9, label %10, label %13

10:
	%11 = getelementptr [6 x i8], [6 x i8]* @0, i64 0, i64 0
	%12 = call i64 @puts(i8* %11)
	br label %16

13:
	%14 = getelementptr [10 x i8], [10 x i8]* @1, i64 0, i64 0
	%15 = call i64 @puts(i8* %14)
	br label %16

16:
	ret i64 0
}

```
## Run
```bash
Not Equal
Equal
```
