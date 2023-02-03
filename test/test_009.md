# test/test_009.md
## Source
```pascal
MODULE GlobalConsts;

CONST x = 1;
      y = 2;

BEGIN
    IF x = y THEN
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
test/test_009.md:6:0:	begin	"BEGIN"	false	0	0.000000	(6, 0) -> (6, 5)
test/test_009.md:7:4:	if	"IF"	false	0	0.000000	(7, 4) -> (7, 6)
test/test_009.md:7:7:	ident	"x"	false	0	0.000000	(7, 7) -> (7, 8)
test/test_009.md:7:9:	eq	"="	false	0	0.000000	(7, 9) -> (7, 10)
test/test_009.md:7:11:	ident	"y"	false	0	0.000000	(7, 11) -> (7, 12)
test/test_009.md:7:13:	then	"THEN"	false	0	0.000000	(7, 13) -> (7, 17)
test/test_009.md:8:8:	ident	"print"	false	0	0.000000	(8, 8) -> (8, 13)
test/test_009.md:8:13:	lparen	"("	false	0	0.000000	(8, 13) -> (8, 14)
test/test_009.md:8:14:	string	"Equal"	false	0	0.000000	(8, 14) -> (8, 21)
test/test_009.md:8:21:	rparen	")"	false	0	0.000000	(8, 21) -> (8, 22)
test/test_009.md:9:4:	else	"ELSE"	false	0	0.000000	(9, 4) -> (9, 8)
test/test_009.md:10:8:	ident	"print"	false	0	0.000000	(10, 8) -> (10, 13)
test/test_009.md:10:13:	lparen	"("	false	0	0.000000	(10, 13) -> (10, 14)
test/test_009.md:10:14:	string	"Not Equal"	false	0	0.000000	(10, 14) -> (10, 25)
test/test_009.md:10:25:	rparen	")"	false	0	0.000000	(10, 25) -> (10, 26)
test/test_009.md:11:4:	end	"END"	false	0	0.000000	(11, 4) -> (11, 7)
test/test_009.md:12:0:	end	"END"	false	0	0.000000	(12, 0) -> (12, 3)
test/test_009.md:12:4:	ident	"GlobalConsts"	false	0	0.000000	(12, 4) -> (12, 16)
test/test_009.md:12:16:	dot	"."	false	0	0.000000	(12, 16) -> (12, 17)
test/test_009.md:13:0:	eof	""	false	0	0.000000	(13, 0) -> (13, 0)
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
	ret i64 0
}

```
## Run
```bash
Not Equal
```
