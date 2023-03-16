# test/test_015.md
## Source
```pascal
MODULE ForBy;

VAR x: INTEGER;

BEGIN
  FOR x := 1 TO 10 BY 2 DO
    Texts.WriteInt(x);
  END;
  Texts.WriteLn;

  FOR x := 9 TO 1 BY -3 DO
    Texts.WriteInt(x);
  END;
  Texts.WriteLn;

  FOR x := 1 TO 3 DO 
    Texts.WriteInt(x);
  END;
  Texts.WriteLn
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
test/test_015.md:7:4:	ident	"Texts"	false	0	0.000000	(7, 4) -> (7, 9)
test/test_015.md:7:9:	dot	"."	false	0	0.000000	(7, 9) -> (7, 10)
test/test_015.md:7:10:	ident	"WriteInt"	false	0	0.000000	(7, 10) -> (7, 18)
test/test_015.md:7:18:	lparen	"("	false	0	0.000000	(7, 18) -> (7, 19)
test/test_015.md:7:19:	ident	"x"	false	0	0.000000	(7, 19) -> (7, 20)
test/test_015.md:7:20:	rparen	")"	false	0	0.000000	(7, 20) -> (7, 21)
test/test_015.md:7:21:	semicolon	";"	false	0	0.000000	(7, 21) -> (7, 22)
test/test_015.md:8:2:	end	"END"	false	0	0.000000	(8, 2) -> (8, 5)
test/test_015.md:8:5:	semicolon	";"	false	0	0.000000	(8, 5) -> (8, 6)
test/test_015.md:9:2:	ident	"Texts"	false	0	0.000000	(9, 2) -> (9, 7)
test/test_015.md:9:7:	dot	"."	false	0	0.000000	(9, 7) -> (9, 8)
test/test_015.md:9:8:	ident	"WriteLn"	false	0	0.000000	(9, 8) -> (9, 15)
test/test_015.md:9:15:	semicolon	";"	false	0	0.000000	(9, 15) -> (9, 16)
test/test_015.md:11:2:	for	"FOR"	false	0	0.000000	(11, 2) -> (11, 5)
test/test_015.md:11:6:	ident	"x"	false	0	0.000000	(11, 6) -> (11, 7)
test/test_015.md:11:8:	assign	":="	false	0	0.000000	(11, 8) -> (11, 10)
test/test_015.md:11:11:	integer	"9"	false	9	0.000000	(11, 11) -> (11, 12)
test/test_015.md:11:13:	to	"TO"	false	0	0.000000	(11, 13) -> (11, 15)
test/test_015.md:11:16:	integer	"1"	false	1	0.000000	(11, 16) -> (11, 17)
test/test_015.md:11:18:	by	"BY"	false	0	0.000000	(11, 18) -> (11, 20)
test/test_015.md:11:21:	minus	"-"	false	0	0.000000	(11, 21) -> (11, 22)
test/test_015.md:11:22:	integer	"3"	false	3	0.000000	(11, 22) -> (11, 23)
test/test_015.md:11:24:	do	"DO"	false	0	0.000000	(11, 24) -> (11, 26)
test/test_015.md:12:4:	ident	"Texts"	false	0	0.000000	(12, 4) -> (12, 9)
test/test_015.md:12:9:	dot	"."	false	0	0.000000	(12, 9) -> (12, 10)
test/test_015.md:12:10:	ident	"WriteInt"	false	0	0.000000	(12, 10) -> (12, 18)
test/test_015.md:12:18:	lparen	"("	false	0	0.000000	(12, 18) -> (12, 19)
test/test_015.md:12:19:	ident	"x"	false	0	0.000000	(12, 19) -> (12, 20)
test/test_015.md:12:20:	rparen	")"	false	0	0.000000	(12, 20) -> (12, 21)
test/test_015.md:12:21:	semicolon	";"	false	0	0.000000	(12, 21) -> (12, 22)
test/test_015.md:13:2:	end	"END"	false	0	0.000000	(13, 2) -> (13, 5)
test/test_015.md:13:5:	semicolon	";"	false	0	0.000000	(13, 5) -> (13, 6)
test/test_015.md:14:2:	ident	"Texts"	false	0	0.000000	(14, 2) -> (14, 7)
test/test_015.md:14:7:	dot	"."	false	0	0.000000	(14, 7) -> (14, 8)
test/test_015.md:14:8:	ident	"WriteLn"	false	0	0.000000	(14, 8) -> (14, 15)
test/test_015.md:14:15:	semicolon	";"	false	0	0.000000	(14, 15) -> (14, 16)
test/test_015.md:16:2:	for	"FOR"	false	0	0.000000	(16, 2) -> (16, 5)
test/test_015.md:16:6:	ident	"x"	false	0	0.000000	(16, 6) -> (16, 7)
test/test_015.md:16:8:	assign	":="	false	0	0.000000	(16, 8) -> (16, 10)
test/test_015.md:16:11:	integer	"1"	false	1	0.000000	(16, 11) -> (16, 12)
test/test_015.md:16:13:	to	"TO"	false	0	0.000000	(16, 13) -> (16, 15)
test/test_015.md:16:16:	integer	"3"	false	3	0.000000	(16, 16) -> (16, 17)
test/test_015.md:16:18:	do	"DO"	false	0	0.000000	(16, 18) -> (16, 20)
test/test_015.md:17:4:	ident	"Texts"	false	0	0.000000	(17, 4) -> (17, 9)
test/test_015.md:17:9:	dot	"."	false	0	0.000000	(17, 9) -> (17, 10)
test/test_015.md:17:10:	ident	"WriteInt"	false	0	0.000000	(17, 10) -> (17, 18)
test/test_015.md:17:18:	lparen	"("	false	0	0.000000	(17, 18) -> (17, 19)
test/test_015.md:17:19:	ident	"x"	false	0	0.000000	(17, 19) -> (17, 20)
test/test_015.md:17:20:	rparen	")"	false	0	0.000000	(17, 20) -> (17, 21)
test/test_015.md:17:21:	semicolon	";"	false	0	0.000000	(17, 21) -> (17, 22)
test/test_015.md:18:2:	end	"END"	false	0	0.000000	(18, 2) -> (18, 5)
test/test_015.md:18:5:	semicolon	";"	false	0	0.000000	(18, 5) -> (18, 6)
test/test_015.md:19:2:	ident	"Texts"	false	0	0.000000	(19, 2) -> (19, 7)
test/test_015.md:19:7:	dot	"."	false	0	0.000000	(19, 7) -> (19, 8)
test/test_015.md:19:8:	ident	"WriteLn"	false	0	0.000000	(19, 8) -> (19, 15)
test/test_015.md:20:0:	end	"END"	false	0	0.000000	(20, 0) -> (20, 3)
test/test_015.md:20:4:	ident	"ForBy"	false	0	0.000000	(20, 4) -> (20, 9)
test/test_015.md:20:9:	dot	"."	false	0	0.000000	(20, 9) -> (20, 10)
test/test_015.md:21:0:	eof	""	false	0	0.000000	(21, 0) -> (21, 0)
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
            (procedure [void] "Texts.WriteInt")
            (variable [i64] "x")
          )
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
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
            (procedure [void] "Texts.WriteInt")
            (variable [i64] "x")
          )
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
    (for x
      (number [i64] 1)
      (number [i64] 3)
      (number [i64] 1)
      (stmts
        (expr2stmt
          (call
            (procedure [void] "Texts.WriteInt")
            (variable [i64] "x")
          )
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
  )
)
```
## IR
```llvm
@0 = global i64 0
@1 = global [3 x i8] c"%d\00"
@2 = global [1 x i8] c"\00"
@__argc = global i64 0
@__argv = global i8** inttoptr (i8 0 to i8**)

declare i64 @puts(i8* %str)

declare i64 @rand()

declare i64 @sprintf(i8* %buf, i8* %format, ...)

declare i64 @printf(i8* %format, ...)

declare i8* @malloc(i64 %size)

declare i8* @free(i8* %ptr)

define void @oberonMain() {
entry:
	store i64 1, i64* @0
	br label %0

0:
	%1 = load i64, i64* @0
	%2 = getelementptr [3 x i8], [3 x i8]* @1, i64 0, i64 0
	%3 = call i64 (i8*, ...) @printf(i8* %2, i64 %1)
	%4 = load i64, i64* @0
	%5 = add i64 %4, 2
	store i64 %5, i64* @0
	%6 = icmp sle i64 %5, 10
	br i1 %6, label %0, label %7

7:
	%8 = getelementptr [1 x i8], [1 x i8]* @2, i64 0, i64 0
	%9 = call i64 @puts(i8* %8)
	%10 = sub i64 0, 3
	store i64 9, i64* @0
	br label %11

11:
	%12 = load i64, i64* @0
	%13 = getelementptr [3 x i8], [3 x i8]* @1, i64 0, i64 0
	%14 = call i64 (i8*, ...) @printf(i8* %13, i64 %12)
	%15 = load i64, i64* @0
	%16 = add i64 %15, %10
	store i64 %16, i64* @0
	%17 = icmp sge i64 %16, 1
	br i1 %17, label %11, label %18

18:
	%19 = getelementptr [1 x i8], [1 x i8]* @2, i64 0, i64 0
	%20 = call i64 @puts(i8* %19)
	store i64 1, i64* @0
	br label %21

21:
	%22 = load i64, i64* @0
	%23 = getelementptr [3 x i8], [3 x i8]* @1, i64 0, i64 0
	%24 = call i64 (i8*, ...) @printf(i8* %23, i64 %22)
	%25 = load i64, i64* @0
	%26 = add i64 %25, 1
	store i64 %26, i64* @0
	%27 = icmp sle i64 %26, 3
	br i1 %27, label %21, label %28

28:
	%29 = getelementptr [1 x i8], [1 x i8]* @2, i64 0, i64 0
	%30 = call i64 @puts(i8* %29)
	ret void
}

define i64 @main(i64 %argc, i8** %argv) {
entry:
	store i64 %argc, i64* @__argc
	store i8** %argv, i8*** @__argv
	call void @oberonMain()
	ret i64 0
}

```
## Run
```bash
13579
963
123
```
