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
        Texts.WriteString("Equal");     Texts.WriteLn
    ELSE
        Texts.WriteString("Not Equal"); Texts.WriteLn
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
test/test_008.md:12:8:	ident	"Texts"	false	0	0.000000	(12, 8) -> (12, 13)
test/test_008.md:12:13:	dot	"."	false	0	0.000000	(12, 13) -> (12, 14)
test/test_008.md:12:14:	ident	"WriteString"	false	0	0.000000	(12, 14) -> (12, 25)
test/test_008.md:12:25:	lparen	"("	false	0	0.000000	(12, 25) -> (12, 26)
test/test_008.md:12:26:	string	"Equal"	false	0	0.000000	(12, 26) -> (12, 33)
test/test_008.md:12:33:	rparen	")"	false	0	0.000000	(12, 33) -> (12, 34)
test/test_008.md:12:34:	semicolon	";"	false	0	0.000000	(12, 34) -> (12, 35)
test/test_008.md:12:40:	ident	"Texts"	false	0	0.000000	(12, 40) -> (12, 45)
test/test_008.md:12:45:	dot	"."	false	0	0.000000	(12, 45) -> (12, 46)
test/test_008.md:12:46:	ident	"WriteLn"	false	0	0.000000	(12, 46) -> (12, 53)
test/test_008.md:13:4:	else	"ELSE"	false	0	0.000000	(13, 4) -> (13, 8)
test/test_008.md:14:8:	ident	"Texts"	false	0	0.000000	(14, 8) -> (14, 13)
test/test_008.md:14:13:	dot	"."	false	0	0.000000	(14, 13) -> (14, 14)
test/test_008.md:14:14:	ident	"WriteString"	false	0	0.000000	(14, 14) -> (14, 25)
test/test_008.md:14:25:	lparen	"("	false	0	0.000000	(14, 25) -> (14, 26)
test/test_008.md:14:26:	string	"Not Equal"	false	0	0.000000	(14, 26) -> (14, 37)
test/test_008.md:14:37:	rparen	")"	false	0	0.000000	(14, 37) -> (14, 38)
test/test_008.md:14:38:	semicolon	";"	false	0	0.000000	(14, 38) -> (14, 39)
test/test_008.md:14:40:	ident	"Texts"	false	0	0.000000	(14, 40) -> (14, 45)
test/test_008.md:14:45:	dot	"."	false	0	0.000000	(14, 45) -> (14, 46)
test/test_008.md:14:46:	ident	"WriteLn"	false	0	0.000000	(14, 46) -> (14, 53)
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
    (x 
      (INTEGER [i64])
    )
    (y 
      (INTEGER [i64])
    )
    (res 
      (BOOLEAN [boolean])
    )
  )
  (stmts
    (assign
      (variable [i64] "x")
      (number [i64] 1)
    )
    (assign
      (variable [i64] "y")
      (number [i64] 2)
    )
    (assign
      (variable [boolean] "res")
      (eq [boolean]
        (variable [i64] "x")
        (variable [i64] "y")
      )
    )
    (if
      (variable [boolean] "res")
      (stmts
        (expr2stmt
          (call
            (procedure [void] "Texts.WriteString")
            (string "Equal")
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
            (procedure [void] "Texts.WriteString")
            (string "Not Equal")
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
)
```
## IR
```llvm
@0 = global i64 0
@1 = global i64 0
@2 = global i1 false
@3 = global [6 x i8] c"Equal\00"
@4 = global [1 x i8] c"\00"
@5 = global [10 x i8] c"Not Equal\00"
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
	store i64 2, i64* @1
	%0 = load i64, i64* @0
	%1 = load i64, i64* @1
	%2 = icmp eq i64 %0, %1
	store i1 %2, i1* @2
	%3 = load i1, i1* @2
	br i1 %3, label %4, label %9

4:
	%5 = getelementptr [6 x i8], [6 x i8]* @3, i64 0, i64 0
	%6 = call i64 (i8*, ...) @printf(i8* %5)
	%7 = getelementptr [1 x i8], [1 x i8]* @4, i64 0, i64 0
	%8 = call i64 @puts(i8* %7)
	br label %14

9:
	%10 = getelementptr [10 x i8], [10 x i8]* @5, i64 0, i64 0
	%11 = call i64 (i8*, ...) @printf(i8* %10)
	%12 = getelementptr [1 x i8], [1 x i8]* @4, i64 0, i64 0
	%13 = call i64 @puts(i8* %12)
	br label %14

14:
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
Not Equal
```
