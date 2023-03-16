# test/test_007.md
## Source
```pascal
MODULE ConstantIf;

BEGIN
  IF 1=3 THEN
    Texts.WriteInt(3);  Texts.WriteLn;
  ELSIF 1=2 THEN
    Texts.WriteInt(2);  Texts.WriteLn;
  ELSE
    Texts.WriteInt(1);  Texts.WriteLn;
  END;
  IF TRUE THEN
    Texts.WriteString("True");  Texts.WriteLn;
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
test/test_007.md:5:4:	ident	"Texts"	false	0	0.000000	(5, 4) -> (5, 9)
test/test_007.md:5:9:	dot	"."	false	0	0.000000	(5, 9) -> (5, 10)
test/test_007.md:5:10:	ident	"WriteInt"	false	0	0.000000	(5, 10) -> (5, 18)
test/test_007.md:5:18:	lparen	"("	false	0	0.000000	(5, 18) -> (5, 19)
test/test_007.md:5:19:	integer	"3"	false	3	0.000000	(5, 19) -> (5, 20)
test/test_007.md:5:20:	rparen	")"	false	0	0.000000	(5, 20) -> (5, 21)
test/test_007.md:5:21:	semicolon	";"	false	0	0.000000	(5, 21) -> (5, 22)
test/test_007.md:5:24:	ident	"Texts"	false	0	0.000000	(5, 24) -> (5, 29)
test/test_007.md:5:29:	dot	"."	false	0	0.000000	(5, 29) -> (5, 30)
test/test_007.md:5:30:	ident	"WriteLn"	false	0	0.000000	(5, 30) -> (5, 37)
test/test_007.md:5:37:	semicolon	";"	false	0	0.000000	(5, 37) -> (5, 38)
test/test_007.md:6:2:	elsif	"ELSIF"	false	0	0.000000	(6, 2) -> (6, 7)
test/test_007.md:6:8:	integer	"1"	false	1	0.000000	(6, 8) -> (6, 9)
test/test_007.md:6:9:	eq	"="	false	0	0.000000	(6, 9) -> (6, 10)
test/test_007.md:6:10:	integer	"2"	false	2	0.000000	(6, 10) -> (6, 11)
test/test_007.md:6:12:	then	"THEN"	false	0	0.000000	(6, 12) -> (6, 16)
test/test_007.md:7:4:	ident	"Texts"	false	0	0.000000	(7, 4) -> (7, 9)
test/test_007.md:7:9:	dot	"."	false	0	0.000000	(7, 9) -> (7, 10)
test/test_007.md:7:10:	ident	"WriteInt"	false	0	0.000000	(7, 10) -> (7, 18)
test/test_007.md:7:18:	lparen	"("	false	0	0.000000	(7, 18) -> (7, 19)
test/test_007.md:7:19:	integer	"2"	false	2	0.000000	(7, 19) -> (7, 20)
test/test_007.md:7:20:	rparen	")"	false	0	0.000000	(7, 20) -> (7, 21)
test/test_007.md:7:21:	semicolon	";"	false	0	0.000000	(7, 21) -> (7, 22)
test/test_007.md:7:24:	ident	"Texts"	false	0	0.000000	(7, 24) -> (7, 29)
test/test_007.md:7:29:	dot	"."	false	0	0.000000	(7, 29) -> (7, 30)
test/test_007.md:7:30:	ident	"WriteLn"	false	0	0.000000	(7, 30) -> (7, 37)
test/test_007.md:7:37:	semicolon	";"	false	0	0.000000	(7, 37) -> (7, 38)
test/test_007.md:8:2:	else	"ELSE"	false	0	0.000000	(8, 2) -> (8, 6)
test/test_007.md:9:4:	ident	"Texts"	false	0	0.000000	(9, 4) -> (9, 9)
test/test_007.md:9:9:	dot	"."	false	0	0.000000	(9, 9) -> (9, 10)
test/test_007.md:9:10:	ident	"WriteInt"	false	0	0.000000	(9, 10) -> (9, 18)
test/test_007.md:9:18:	lparen	"("	false	0	0.000000	(9, 18) -> (9, 19)
test/test_007.md:9:19:	integer	"1"	false	1	0.000000	(9, 19) -> (9, 20)
test/test_007.md:9:20:	rparen	")"	false	0	0.000000	(9, 20) -> (9, 21)
test/test_007.md:9:21:	semicolon	";"	false	0	0.000000	(9, 21) -> (9, 22)
test/test_007.md:9:24:	ident	"Texts"	false	0	0.000000	(9, 24) -> (9, 29)
test/test_007.md:9:29:	dot	"."	false	0	0.000000	(9, 29) -> (9, 30)
test/test_007.md:9:30:	ident	"WriteLn"	false	0	0.000000	(9, 30) -> (9, 37)
test/test_007.md:9:37:	semicolon	";"	false	0	0.000000	(9, 37) -> (9, 38)
test/test_007.md:10:2:	end	"END"	false	0	0.000000	(10, 2) -> (10, 5)
test/test_007.md:10:5:	semicolon	";"	false	0	0.000000	(10, 5) -> (10, 6)
test/test_007.md:11:2:	if	"IF"	false	0	0.000000	(11, 2) -> (11, 4)
test/test_007.md:11:5:	boolean	"TRUE"	true	0	0.000000	(11, 5) -> (11, 9)
test/test_007.md:11:10:	then	"THEN"	false	0	0.000000	(11, 10) -> (11, 14)
test/test_007.md:12:4:	ident	"Texts"	false	0	0.000000	(12, 4) -> (12, 9)
test/test_007.md:12:9:	dot	"."	false	0	0.000000	(12, 9) -> (12, 10)
test/test_007.md:12:10:	ident	"WriteString"	false	0	0.000000	(12, 10) -> (12, 21)
test/test_007.md:12:21:	lparen	"("	false	0	0.000000	(12, 21) -> (12, 22)
test/test_007.md:12:22:	string	"True"	false	0	0.000000	(12, 22) -> (12, 28)
test/test_007.md:12:28:	rparen	")"	false	0	0.000000	(12, 28) -> (12, 29)
test/test_007.md:12:29:	semicolon	";"	false	0	0.000000	(12, 29) -> (12, 30)
test/test_007.md:12:32:	ident	"Texts"	false	0	0.000000	(12, 32) -> (12, 37)
test/test_007.md:12:37:	dot	"."	false	0	0.000000	(12, 37) -> (12, 38)
test/test_007.md:12:38:	ident	"WriteLn"	false	0	0.000000	(12, 38) -> (12, 45)
test/test_007.md:12:45:	semicolon	";"	false	0	0.000000	(12, 45) -> (12, 46)
test/test_007.md:13:2:	end	"END"	false	0	0.000000	(13, 2) -> (13, 5)
test/test_007.md:14:0:	end	"END"	false	0	0.000000	(14, 0) -> (14, 3)
test/test_007.md:14:4:	ident	"ConstantIf"	false	0	0.000000	(14, 4) -> (14, 14)
test/test_007.md:14:14:	dot	"."	false	0	0.000000	(14, 14) -> (14, 15)
test/test_007.md:15:0:	eof	""	false	0	0.000000	(15, 0) -> (15, 0)
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
          (call
            (procedure [void] "Texts.WriteInt")
            (number [i64] 3)
          )
        )
        (expr2stmt
          (call
            (procedure [void] "Texts.WriteLn")
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
            (call
              (procedure [void] "Texts.WriteInt")
              (number [i64] 2)
            )
          )
          (expr2stmt
            (call
              (procedure [void] "Texts.WriteLn")
            )
          )
        )
        (stmts
          (expr2stmt
            (call
              (procedure [void] "Texts.WriteInt")
              (number [i64] 1)
            )
          )
          (expr2stmt
            (call
              (procedure [void] "Texts.WriteLn")
            )
          )
        )
      )
    )
    (if
      #true
      (stmts
        (expr2stmt
          (call
            (procedure [void] "Texts.WriteString")
            (string "True")
          )
        )
        (expr2stmt
          (call
            (procedure [void] "Texts.WriteLn")
          )
        )
      )
      (stmts
      )
    )
  )
)
```
## IR
```llvm
@0 = global [3 x i8] c"%d\00"
@1 = global [1 x i8] c"\00"
@2 = global [5 x i8] c"True\00"
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
	%0 = icmp eq i64 1, 3
	br i1 %0, label %1, label %6

1:
	%2 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%3 = call i64 (i8*, ...) @printf(i8* %2, i64 3)
	%4 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%5 = call i64 @puts(i8* %4)
	br label %8

6:
	%7 = icmp eq i64 1, 2
	br i1 %7, label %9, label %14

8:
	br i1 true, label %20, label %25

9:
	%10 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%11 = call i64 (i8*, ...) @printf(i8* %10, i64 2)
	%12 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%13 = call i64 @puts(i8* %12)
	br label %19

14:
	%15 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%16 = call i64 (i8*, ...) @printf(i8* %15, i64 1)
	%17 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%18 = call i64 @puts(i8* %17)
	br label %19

19:
	br label %8

20:
	%21 = getelementptr [5 x i8], [5 x i8]* @2, i64 0, i64 0
	%22 = call i64 (i8*, ...) @printf(i8* %21)
	%23 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%24 = call i64 @puts(i8* %23)
	br label %26

25:
	br label %26

26:
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
1
True
```
