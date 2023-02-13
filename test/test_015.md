# test/test_015.md
## Source
```pascal
MODULE ForBy;

VAR x: INTEGER;

BEGIN
  FOR x := 1 TO 10 BY 2 DO
    print(x)
  END;
  FOR x := 9 TO 1 BY -3 DO
    print(x)
  END;
  FOR x := 1 TO 3 DO 
    print(x)
  END
END ForBy.
```
## Tokens
```tsv
test/test_015.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_015.md:1:8:	ident	"ForBy"	false	0	0.000000	(1, 8) -> (1, 13)
test/test_015.md:1:13:	semicolon	";"	false	0	0.000000	(1, 13) -> (1, 14)
test/test_015.md:3:0:	var	"VAR"	false	0	0.000000	(3, 0) -> (3, 3)
test/test_015.md:3:4:	ident	"x"	false	0	0.000000	(3, 4) -> (3, 5)
test/test_015.md:3:5:	colon	":"	false	0	0.000000	(3, 5) -> (3, 6)
test/test_015.md:3:7:	ident	"INTEGER"	false	0	0.000000	(3, 7) -> (3, 14)
test/test_015.md:3:14:	semicolon	";"	false	0	0.000000	(3, 14) -> (3, 15)
test/test_015.md:5:0:	begin	"BEGIN"	false	0	0.000000	(5, 0) -> (5, 5)
test/test_015.md:6:2:	for	"FOR"	false	0	0.000000	(6, 2) -> (6, 5)
test/test_015.md:6:6:	ident	"x"	false	0	0.000000	(6, 6) -> (6, 7)
test/test_015.md:6:8:	assign	":="	false	0	0.000000	(6, 8) -> (6, 10)
test/test_015.md:6:11:	integer	"1"	false	1	0.000000	(6, 11) -> (6, 12)
test/test_015.md:6:13:	to	"TO"	false	0	0.000000	(6, 13) -> (6, 15)
test/test_015.md:6:16:	integer	"10"	false	10	0.000000	(6, 16) -> (6, 18)
test/test_015.md:6:19:	by	"BY"	false	0	0.000000	(6, 19) -> (6, 21)
test/test_015.md:6:22:	integer	"2"	false	2	0.000000	(6, 22) -> (6, 23)
test/test_015.md:6:24:	do	"DO"	false	0	0.000000	(6, 24) -> (6, 26)
test/test_015.md:7:4:	ident	"print"	false	0	0.000000	(7, 4) -> (7, 9)
test/test_015.md:7:9:	lparen	"("	false	0	0.000000	(7, 9) -> (7, 10)
test/test_015.md:7:10:	ident	"x"	false	0	0.000000	(7, 10) -> (7, 11)
test/test_015.md:7:11:	rparen	")"	false	0	0.000000	(7, 11) -> (7, 12)
test/test_015.md:8:2:	end	"END"	false	0	0.000000	(8, 2) -> (8, 5)
test/test_015.md:8:5:	semicolon	";"	false	0	0.000000	(8, 5) -> (8, 6)
test/test_015.md:9:2:	for	"FOR"	false	0	0.000000	(9, 2) -> (9, 5)
test/test_015.md:9:6:	ident	"x"	false	0	0.000000	(9, 6) -> (9, 7)
test/test_015.md:9:8:	assign	":="	false	0	0.000000	(9, 8) -> (9, 10)
test/test_015.md:9:11:	integer	"9"	false	9	0.000000	(9, 11) -> (9, 12)
test/test_015.md:9:13:	to	"TO"	false	0	0.000000	(9, 13) -> (9, 15)
test/test_015.md:9:16:	integer	"1"	false	1	0.000000	(9, 16) -> (9, 17)
test/test_015.md:9:18:	by	"BY"	false	0	0.000000	(9, 18) -> (9, 20)
test/test_015.md:9:21:	minus	"-"	false	0	0.000000	(9, 21) -> (9, 22)
test/test_015.md:9:22:	integer	"3"	false	3	0.000000	(9, 22) -> (9, 23)
test/test_015.md:9:24:	do	"DO"	false	0	0.000000	(9, 24) -> (9, 26)
test/test_015.md:10:4:	ident	"print"	false	0	0.000000	(10, 4) -> (10, 9)
test/test_015.md:10:9:	lparen	"("	false	0	0.000000	(10, 9) -> (10, 10)
test/test_015.md:10:10:	ident	"x"	false	0	0.000000	(10, 10) -> (10, 11)
test/test_015.md:10:11:	rparen	")"	false	0	0.000000	(10, 11) -> (10, 12)
test/test_015.md:11:2:	end	"END"	false	0	0.000000	(11, 2) -> (11, 5)
test/test_015.md:11:5:	semicolon	";"	false	0	0.000000	(11, 5) -> (11, 6)
test/test_015.md:12:2:	for	"FOR"	false	0	0.000000	(12, 2) -> (12, 5)
test/test_015.md:12:6:	ident	"x"	false	0	0.000000	(12, 6) -> (12, 7)
test/test_015.md:12:8:	assign	":="	false	0	0.000000	(12, 8) -> (12, 10)
test/test_015.md:12:11:	integer	"1"	false	1	0.000000	(12, 11) -> (12, 12)
test/test_015.md:12:13:	to	"TO"	false	0	0.000000	(12, 13) -> (12, 15)
test/test_015.md:12:16:	integer	"3"	false	3	0.000000	(12, 16) -> (12, 17)
test/test_015.md:12:18:	do	"DO"	false	0	0.000000	(12, 18) -> (12, 20)
test/test_015.md:13:4:	ident	"print"	false	0	0.000000	(13, 4) -> (13, 9)
test/test_015.md:13:9:	lparen	"("	false	0	0.000000	(13, 9) -> (13, 10)
test/test_015.md:13:10:	ident	"x"	false	0	0.000000	(13, 10) -> (13, 11)
test/test_015.md:13:11:	rparen	")"	false	0	0.000000	(13, 11) -> (13, 12)
test/test_015.md:14:2:	end	"END"	false	0	0.000000	(14, 2) -> (14, 5)
test/test_015.md:15:0:	end	"END"	false	0	0.000000	(15, 0) -> (15, 3)
test/test_015.md:15:4:	ident	"ForBy"	false	0	0.000000	(15, 4) -> (15, 9)
test/test_015.md:15:9:	dot	"."	false	0	0.000000	(15, 9) -> (15, 10)
test/test_015.md:16:0:	eof	""	false	0	0.000000	(16, 0) -> (16, 0)
```
## AST
```scheme
(module "ForBy"
  (vars
    (x 
      (INTEGER [i64])
    )
  )
  (stmts
    (for x
      (number [i64] 1)
      (number [i64] 10)
      (number [i64] 2)
      (stmts
        (expr2stmt
          (call
            (procedure [void] "print")
            (variable [i64] "x")
          )
        )
      )
    )
    (for x
      (number [i64] 9)
      (number [i64] 1)
      (minus [i64]
        (number [i64] 0)
        (number [i64] 3)
      )
      (stmts
        (expr2stmt
          (call
            (procedure [void] "print")
            (variable [i64] "x")
          )
        )
      )
    )
    (for x
      (number [i64] 1)
      (number [i64] 3)
      (number [i64] 1)
      (stmts
        (expr2stmt
          (call
            (procedure [void] "print")
            (variable [i64] "x")
          )
        )
      )
    )
  )
)
```
## IR
```llvm
@0 = global i64 0
@1 = global [4 x i8] c"%d\0A\00"

declare i64 @puts(i8* %str)

declare i64 @rand()

declare i64 @sprintf(i8* %buf, i8* %format, ...)

declare i64 @printf(i8* %format, ...)

declare i8* @malloc(i64 %size)

define i64 @main() {
entry:
	store i64 1, i64* @0
	br label %0

0:
	%1 = load i64, i64* @0
	%2 = getelementptr [4 x i8], [4 x i8]* @1, i64 0, i64 0
	%3 = call i64 (i8*, ...) @printf(i8* %2, i64 %1)
	%4 = load i64, i64* @0
	%5 = add i64 %4, 2
	store i64 %5, i64* @0
	%6 = icmp sle i64 %5, 10
	br i1 %6, label %0, label %7

7:
	%8 = sub i64 0, 3
	store i64 9, i64* @0
	br label %9

9:
	%10 = load i64, i64* @0
	%11 = getelementptr [4 x i8], [4 x i8]* @1, i64 0, i64 0
	%12 = call i64 (i8*, ...) @printf(i8* %11, i64 %10)
	%13 = load i64, i64* @0
	%14 = add i64 %13, %8
	store i64 %14, i64* @0
	%15 = icmp sge i64 %14, 1
	br i1 %15, label %9, label %16

16:
	store i64 1, i64* @0
	br label %17

17:
	%18 = load i64, i64* @0
	%19 = getelementptr [4 x i8], [4 x i8]* @1, i64 0, i64 0
	%20 = call i64 (i8*, ...) @printf(i8* %19, i64 %18)
	%21 = load i64, i64* @0
	%22 = add i64 %21, 1
	store i64 %22, i64* @0
	%23 = icmp sle i64 %22, 3
	br i1 %23, label %17, label %24

24:
	ret i64 0
}

```
## Run
```bash
1
3
5
7
9
9
6
3
1
2
3
```
