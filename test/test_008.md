# test/test_008.md
## Source
```pascal
MODULE GlobalVars;

VAR x, y: INTEGER;
    res : BOOLEAN;

BEGIN
    x := 1;
    y := 2;
    res := x = y;

    IF res THEN
        print("Equal")
    ELSE
        print("Not Equal")
    END
END GlobalVars.
```
## Tokens
```tsv
test/test_008.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_008.md:1:8:	ident	"GlobalVars"	false	0	0.000000	(1, 8) -> (1, 18)
test/test_008.md:1:18:	semicolon	";"	false	0	0.000000	(1, 18) -> (1, 19)
test/test_008.md:3:0:	var	"VAR"	false	0	0.000000	(3, 0) -> (3, 3)
test/test_008.md:3:4:	ident	"x"	false	0	0.000000	(3, 4) -> (3, 5)
test/test_008.md:3:5:	comma	","	false	0	0.000000	(3, 5) -> (3, 6)
test/test_008.md:3:7:	ident	"y"	false	0	0.000000	(3, 7) -> (3, 8)
test/test_008.md:3:8:	colon	":"	false	0	0.000000	(3, 8) -> (3, 9)
test/test_008.md:3:10:	ident	"INTEGER"	false	0	0.000000	(3, 10) -> (3, 17)
test/test_008.md:3:17:	semicolon	";"	false	0	0.000000	(3, 17) -> (3, 18)
test/test_008.md:4:4:	ident	"res"	false	0	0.000000	(4, 4) -> (4, 7)
test/test_008.md:4:8:	colon	":"	false	0	0.000000	(4, 8) -> (4, 9)
test/test_008.md:4:10:	ident	"BOOLEAN"	false	0	0.000000	(4, 10) -> (4, 17)
test/test_008.md:4:17:	semicolon	";"	false	0	0.000000	(4, 17) -> (4, 18)
test/test_008.md:6:0:	begin	"BEGIN"	false	0	0.000000	(6, 0) -> (6, 5)
test/test_008.md:7:4:	ident	"x"	false	0	0.000000	(7, 4) -> (7, 5)
test/test_008.md:7:6:	assign	":="	false	0	0.000000	(7, 6) -> (7, 8)
test/test_008.md:7:9:	integer	"1"	false	1	0.000000	(7, 9) -> (7, 10)
test/test_008.md:7:10:	semicolon	";"	false	0	0.000000	(7, 10) -> (7, 11)
test/test_008.md:8:4:	ident	"y"	false	0	0.000000	(8, 4) -> (8, 5)
test/test_008.md:8:6:	assign	":="	false	0	0.000000	(8, 6) -> (8, 8)
test/test_008.md:8:9:	integer	"2"	false	2	0.000000	(8, 9) -> (8, 10)
test/test_008.md:8:10:	semicolon	";"	false	0	0.000000	(8, 10) -> (8, 11)
test/test_008.md:9:4:	ident	"res"	false	0	0.000000	(9, 4) -> (9, 7)
test/test_008.md:9:8:	assign	":="	false	0	0.000000	(9, 8) -> (9, 10)
test/test_008.md:9:11:	ident	"x"	false	0	0.000000	(9, 11) -> (9, 12)
test/test_008.md:9:13:	eq	"="	false	0	0.000000	(9, 13) -> (9, 14)
test/test_008.md:9:15:	ident	"y"	false	0	0.000000	(9, 15) -> (9, 16)
test/test_008.md:9:16:	semicolon	";"	false	0	0.000000	(9, 16) -> (9, 17)
test/test_008.md:11:4:	if	"IF"	false	0	0.000000	(11, 4) -> (11, 6)
test/test_008.md:11:7:	ident	"res"	false	0	0.000000	(11, 7) -> (11, 10)
test/test_008.md:11:11:	then	"THEN"	false	0	0.000000	(11, 11) -> (11, 15)
test/test_008.md:12:8:	ident	"print"	false	0	0.000000	(12, 8) -> (12, 13)
test/test_008.md:12:13:	lparen	"("	false	0	0.000000	(12, 13) -> (12, 14)
test/test_008.md:12:14:	string	"Equal"	false	0	0.000000	(12, 14) -> (12, 21)
test/test_008.md:12:21:	rparen	")"	false	0	0.000000	(12, 21) -> (12, 22)
test/test_008.md:13:4:	else	"ELSE"	false	0	0.000000	(13, 4) -> (13, 8)
test/test_008.md:14:8:	ident	"print"	false	0	0.000000	(14, 8) -> (14, 13)
test/test_008.md:14:13:	lparen	"("	false	0	0.000000	(14, 13) -> (14, 14)
test/test_008.md:14:14:	string	"Not Equal"	false	0	0.000000	(14, 14) -> (14, 25)
test/test_008.md:14:25:	rparen	")"	false	0	0.000000	(14, 25) -> (14, 26)
test/test_008.md:15:4:	end	"END"	false	0	0.000000	(15, 4) -> (15, 7)
test/test_008.md:16:0:	end	"END"	false	0	0.000000	(16, 0) -> (16, 3)
test/test_008.md:16:4:	ident	"GlobalVars"	false	0	0.000000	(16, 4) -> (16, 14)
test/test_008.md:16:14:	dot	"."	false	0	0.000000	(16, 14) -> (16, 15)
test/test_008.md:17:0:	eof	""	false	0	0.000000	(17, 0) -> (17, 0)
```
## AST
```scheme
(module "GlobalVars"
  (vars
    (x [i64])
    (y [i64])
    (res [boolean])
  )
  (stmts
    (assign x
      (number [i64] 1)
    )
    (assign y
      (number [i64] 2)
    )
    (assign res
      (eq [boolean]
        (variable [i64] "x")
        (variable [i64] "y")
      )
    )
    (if
      (variable [boolean] "res")
      (stmts
        (expr2stmt
          (call "print" [void]
            (string "Equal")
          )
        )
      )
      (stmts
        (expr2stmt
          (call "print" [void]
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
@0 = global i64 0
@1 = global i64 0
@2 = global i1 false
@3 = global [6 x i8] c"Equal\00"
@4 = global [10 x i8] c"Not Equal\00"

declare i64 @puts(i8* %str)

declare i64 @rand()

declare i64 @sprintf(i8* %buf, i8* %format, ...)

declare i64 @printf(i8* %format, ...)

define i64 @main() {
entry:
	store i64 1, i64* @0
	store i64 2, i64* @1
	%0 = load i64, i64* @0
	%1 = load i64, i64* @1
	%2 = icmp eq i64 %0, %1
	store i1 %2, i1* @2
	%3 = load i1, i1* @2
	br i1 %3, label %4, label %7

4:
	%5 = getelementptr [6 x i8], [6 x i8]* @3, i64 0, i64 0
	%6 = call i64 @puts(i8* %5)
	br label %10

7:
	%8 = getelementptr [10 x i8], [10 x i8]* @4, i64 0, i64 0
	%9 = call i64 @puts(i8* %8)
	br label %10

10:
	ret i64 0
}

```
## Run
```bash
Not Equal
```
