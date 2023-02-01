# test/test_007.md
## Source
```pascal
MODULE ConstantIf;

BEGIN
  IF 1=3 THEN
    print(3)
  ELSIF 1=2 THEN
    print(2)
  ELSE
    print(1)
  END
END ConstantIf.
```
## Tokens
```tsv
test/test_007.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_007.md:1:8:	ident	"ConstantIf"	false	0	0.000000	(1, 8) -> (1, 18)
test/test_007.md:1:18:	semicolon	";"	false	0	0.000000	(1, 18) -> (1, 19)
test/test_007.md:3:0:	begin	"BEGIN"	false	0	0.000000	(3, 0) -> (3, 5)
test/test_007.md:4:2:	if	"IF"	false	0	0.000000	(4, 2) -> (4, 4)
test/test_007.md:4:5:	integer	"1"	false	1	0.000000	(4, 5) -> (4, 6)
test/test_007.md:4:6:	eq	"="	false	0	0.000000	(4, 6) -> (4, 7)
test/test_007.md:4:7:	integer	"3"	false	3	0.000000	(4, 7) -> (4, 8)
test/test_007.md:4:9:	then	"THEN"	false	0	0.000000	(4, 9) -> (4, 13)
test/test_007.md:5:4:	ident	"print"	false	0	0.000000	(5, 4) -> (5, 9)
test/test_007.md:5:9:	lparen	"("	false	0	0.000000	(5, 9) -> (5, 10)
test/test_007.md:5:10:	integer	"3"	false	3	0.000000	(5, 10) -> (5, 11)
test/test_007.md:5:11:	rparen	")"	false	0	0.000000	(5, 11) -> (5, 12)
test/test_007.md:6:2:	elsif	"ELSIF"	false	0	0.000000	(6, 2) -> (6, 7)
test/test_007.md:6:8:	integer	"1"	false	1	0.000000	(6, 8) -> (6, 9)
test/test_007.md:6:9:	eq	"="	false	0	0.000000	(6, 9) -> (6, 10)
test/test_007.md:6:10:	integer	"2"	false	2	0.000000	(6, 10) -> (6, 11)
test/test_007.md:6:12:	then	"THEN"	false	0	0.000000	(6, 12) -> (6, 16)
test/test_007.md:7:4:	ident	"print"	false	0	0.000000	(7, 4) -> (7, 9)
test/test_007.md:7:9:	lparen	"("	false	0	0.000000	(7, 9) -> (7, 10)
test/test_007.md:7:10:	integer	"2"	false	2	0.000000	(7, 10) -> (7, 11)
test/test_007.md:7:11:	rparen	")"	false	0	0.000000	(7, 11) -> (7, 12)
test/test_007.md:8:2:	else	"ELSE"	false	0	0.000000	(8, 2) -> (8, 6)
test/test_007.md:9:4:	ident	"print"	false	0	0.000000	(9, 4) -> (9, 9)
test/test_007.md:9:9:	lparen	"("	false	0	0.000000	(9, 9) -> (9, 10)
test/test_007.md:9:10:	integer	"1"	false	1	0.000000	(9, 10) -> (9, 11)
test/test_007.md:9:11:	rparen	")"	false	0	0.000000	(9, 11) -> (9, 12)
test/test_007.md:10:2:	end	"END"	false	0	0.000000	(10, 2) -> (10, 5)
test/test_007.md:11:0:	end	"END"	false	0	0.000000	(11, 0) -> (11, 3)
test/test_007.md:11:4:	ident	"ConstantIf"	false	0	0.000000	(11, 4) -> (11, 14)
test/test_007.md:11:14:	dot	"."	false	0	0.000000	(11, 14) -> (11, 15)
test/test_007.md:12:0:	eof	""	false	0	0.000000	(12, 0) -> (12, 0)
```
## AST
```scheme
(module "ConstantIf"
  (stmts
    (if
      (eq [boolean]
        (number [i64] 1)
        (number [i64] 3)
      )
      (stmts
        (expr2stmt
          (print [void]
            (number [i64] 3)
          )
        )
      )
      (if
        (eq [boolean]
          (number [i64] 1)
          (number [i64] 2)
        )
        (stmts
          (expr2stmt
            (print [void]
              (number [i64] 2)
            )
          )
        )
        (stmts
          (expr2stmt
            (print [void]
              (number [i64] 1)
            )
          )
        )
      )
    )
  )
)
```
## IR
```llvm
@0 = global [4 x i8] c"%d\0A\00"

declare i32 @puts(i8* %str)

declare i32 @rand()

declare i32 @sprintf(i8* %buf, i8* %format, ...)

declare i32 @printf(i8* %format, ...)

define i32 @main() {
entry:
	%0 = icmp eq i32 1, 3
	br i1 %0, label %1, label %4

1:
	%2 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%3 = call i32 (i8*, ...) @printf(i8* %2, i32 3)
	br label %6

4:
	%5 = icmp eq i32 1, 2
	br i1 %5, label %7, label %10

6:
	ret i32 0

7:
	%8 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%9 = call i32 (i8*, ...) @printf(i8* %8, i32 2)
	br label %13

10:
	%11 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%12 = call i32 (i8*, ...) @printf(i8* %11, i32 1)
	br label %13

13:
	br label %6
}

```
## Run
```bash
1
```
