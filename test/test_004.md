# test/test_004.md
## Source
```pascal
MODULE BooleanOperators;

BEGIN
    Texts.WriteInt(ORD(~TRUE));         Texts.WriteLn;
    Texts.WriteInt(ORD(~FALSE));        Texts.WriteLn;
    Texts.WriteInt(ORD(TRUE & FALSE));  Texts.WriteLn;
    Texts.WriteInt(ORD(TRUE OR FALSE)); Texts.WriteLn;
    Texts.WriteInt(ORD(TRUE));          Texts.WriteLn;
    Texts.WriteInt(ORD(FALSE));         Texts.WriteLn;
END BooleanOperators.
```
## Tokens
```tsv
test/test_004.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_004.md:1:8:	ident	"BooleanOperators"	false	0	0.000000	(1, 8) -> (1, 24)
test/test_004.md:1:24:	semicolon	";"	false	0	0.000000	(1, 24) -> (1, 25)
test/test_004.md:3:0:	begin	"BEGIN"	false	0	0.000000	(3, 0) -> (3, 5)
test/test_004.md:4:4:	ident	"Texts"	false	0	0.000000	(4, 4) -> (4, 9)
test/test_004.md:4:9:	dot	"."	false	0	0.000000	(4, 9) -> (4, 10)
test/test_004.md:4:10:	ident	"WriteInt"	false	0	0.000000	(4, 10) -> (4, 18)
test/test_004.md:4:18:	lparen	"("	false	0	0.000000	(4, 18) -> (4, 19)
test/test_004.md:4:19:	ident	"ORD"	false	0	0.000000	(4, 19) -> (4, 22)
test/test_004.md:4:22:	lparen	"("	false	0	0.000000	(4, 22) -> (4, 23)
test/test_004.md:4:23:	tilde	"~"	false	0	0.000000	(4, 23) -> (4, 24)
test/test_004.md:4:24:	boolean	"TRUE"	true	0	0.000000	(4, 24) -> (4, 28)
test/test_004.md:4:28:	rparen	")"	false	0	0.000000	(4, 28) -> (4, 29)
test/test_004.md:4:29:	rparen	")"	false	0	0.000000	(4, 29) -> (4, 30)
test/test_004.md:4:30:	semicolon	";"	false	0	0.000000	(4, 30) -> (4, 31)
test/test_004.md:4:40:	ident	"Texts"	false	0	0.000000	(4, 40) -> (4, 45)
test/test_004.md:4:45:	dot	"."	false	0	0.000000	(4, 45) -> (4, 46)
test/test_004.md:4:46:	ident	"WriteLn"	false	0	0.000000	(4, 46) -> (4, 53)
test/test_004.md:4:53:	semicolon	";"	false	0	0.000000	(4, 53) -> (4, 54)
test/test_004.md:5:4:	ident	"Texts"	false	0	0.000000	(5, 4) -> (5, 9)
test/test_004.md:5:9:	dot	"."	false	0	0.000000	(5, 9) -> (5, 10)
test/test_004.md:5:10:	ident	"WriteInt"	false	0	0.000000	(5, 10) -> (5, 18)
test/test_004.md:5:18:	lparen	"("	false	0	0.000000	(5, 18) -> (5, 19)
test/test_004.md:5:19:	ident	"ORD"	false	0	0.000000	(5, 19) -> (5, 22)
test/test_004.md:5:22:	lparen	"("	false	0	0.000000	(5, 22) -> (5, 23)
test/test_004.md:5:23:	tilde	"~"	false	0	0.000000	(5, 23) -> (5, 24)
test/test_004.md:5:24:	boolean	"FALSE"	false	0	0.000000	(5, 24) -> (5, 29)
test/test_004.md:5:29:	rparen	")"	false	0	0.000000	(5, 29) -> (5, 30)
test/test_004.md:5:30:	rparen	")"	false	0	0.000000	(5, 30) -> (5, 31)
test/test_004.md:5:31:	semicolon	";"	false	0	0.000000	(5, 31) -> (5, 32)
test/test_004.md:5:40:	ident	"Texts"	false	0	0.000000	(5, 40) -> (5, 45)
test/test_004.md:5:45:	dot	"."	false	0	0.000000	(5, 45) -> (5, 46)
test/test_004.md:5:46:	ident	"WriteLn"	false	0	0.000000	(5, 46) -> (5, 53)
test/test_004.md:5:53:	semicolon	";"	false	0	0.000000	(5, 53) -> (5, 54)
test/test_004.md:6:4:	ident	"Texts"	false	0	0.000000	(6, 4) -> (6, 9)
test/test_004.md:6:9:	dot	"."	false	0	0.000000	(6, 9) -> (6, 10)
test/test_004.md:6:10:	ident	"WriteInt"	false	0	0.000000	(6, 10) -> (6, 18)
test/test_004.md:6:18:	lparen	"("	false	0	0.000000	(6, 18) -> (6, 19)
test/test_004.md:6:19:	ident	"ORD"	false	0	0.000000	(6, 19) -> (6, 22)
test/test_004.md:6:22:	lparen	"("	false	0	0.000000	(6, 22) -> (6, 23)
test/test_004.md:6:23:	boolean	"TRUE"	true	0	0.000000	(6, 23) -> (6, 27)
test/test_004.md:6:28:	ampersand	"&"	false	0	0.000000	(6, 28) -> (6, 29)
test/test_004.md:6:30:	boolean	"FALSE"	false	0	0.000000	(6, 30) -> (6, 35)
test/test_004.md:6:35:	rparen	")"	false	0	0.000000	(6, 35) -> (6, 36)
test/test_004.md:6:36:	rparen	")"	false	0	0.000000	(6, 36) -> (6, 37)
test/test_004.md:6:37:	semicolon	";"	false	0	0.000000	(6, 37) -> (6, 38)
test/test_004.md:6:40:	ident	"Texts"	false	0	0.000000	(6, 40) -> (6, 45)
test/test_004.md:6:45:	dot	"."	false	0	0.000000	(6, 45) -> (6, 46)
test/test_004.md:6:46:	ident	"WriteLn"	false	0	0.000000	(6, 46) -> (6, 53)
test/test_004.md:6:53:	semicolon	";"	false	0	0.000000	(6, 53) -> (6, 54)
test/test_004.md:7:4:	ident	"Texts"	false	0	0.000000	(7, 4) -> (7, 9)
test/test_004.md:7:9:	dot	"."	false	0	0.000000	(7, 9) -> (7, 10)
test/test_004.md:7:10:	ident	"WriteInt"	false	0	0.000000	(7, 10) -> (7, 18)
test/test_004.md:7:18:	lparen	"("	false	0	0.000000	(7, 18) -> (7, 19)
test/test_004.md:7:19:	ident	"ORD"	false	0	0.000000	(7, 19) -> (7, 22)
test/test_004.md:7:22:	lparen	"("	false	0	0.000000	(7, 22) -> (7, 23)
test/test_004.md:7:23:	boolean	"TRUE"	true	0	0.000000	(7, 23) -> (7, 27)
test/test_004.md:7:28:	or	"OR"	false	0	0.000000	(7, 28) -> (7, 30)
test/test_004.md:7:31:	boolean	"FALSE"	false	0	0.000000	(7, 31) -> (7, 36)
test/test_004.md:7:36:	rparen	")"	false	0	0.000000	(7, 36) -> (7, 37)
test/test_004.md:7:37:	rparen	")"	false	0	0.000000	(7, 37) -> (7, 38)
test/test_004.md:7:38:	semicolon	";"	false	0	0.000000	(7, 38) -> (7, 39)
test/test_004.md:7:40:	ident	"Texts"	false	0	0.000000	(7, 40) -> (7, 45)
test/test_004.md:7:45:	dot	"."	false	0	0.000000	(7, 45) -> (7, 46)
test/test_004.md:7:46:	ident	"WriteLn"	false	0	0.000000	(7, 46) -> (7, 53)
test/test_004.md:7:53:	semicolon	";"	false	0	0.000000	(7, 53) -> (7, 54)
test/test_004.md:8:4:	ident	"Texts"	false	0	0.000000	(8, 4) -> (8, 9)
test/test_004.md:8:9:	dot	"."	false	0	0.000000	(8, 9) -> (8, 10)
test/test_004.md:8:10:	ident	"WriteInt"	false	0	0.000000	(8, 10) -> (8, 18)
test/test_004.md:8:18:	lparen	"("	false	0	0.000000	(8, 18) -> (8, 19)
test/test_004.md:8:19:	ident	"ORD"	false	0	0.000000	(8, 19) -> (8, 22)
test/test_004.md:8:22:	lparen	"("	false	0	0.000000	(8, 22) -> (8, 23)
test/test_004.md:8:23:	boolean	"TRUE"	true	0	0.000000	(8, 23) -> (8, 27)
test/test_004.md:8:27:	rparen	")"	false	0	0.000000	(8, 27) -> (8, 28)
test/test_004.md:8:28:	rparen	")"	false	0	0.000000	(8, 28) -> (8, 29)
test/test_004.md:8:29:	semicolon	";"	false	0	0.000000	(8, 29) -> (8, 30)
test/test_004.md:8:40:	ident	"Texts"	false	0	0.000000	(8, 40) -> (8, 45)
test/test_004.md:8:45:	dot	"."	false	0	0.000000	(8, 45) -> (8, 46)
test/test_004.md:8:46:	ident	"WriteLn"	false	0	0.000000	(8, 46) -> (8, 53)
test/test_004.md:8:53:	semicolon	";"	false	0	0.000000	(8, 53) -> (8, 54)
test/test_004.md:9:4:	ident	"Texts"	false	0	0.000000	(9, 4) -> (9, 9)
test/test_004.md:9:9:	dot	"."	false	0	0.000000	(9, 9) -> (9, 10)
test/test_004.md:9:10:	ident	"WriteInt"	false	0	0.000000	(9, 10) -> (9, 18)
test/test_004.md:9:18:	lparen	"("	false	0	0.000000	(9, 18) -> (9, 19)
test/test_004.md:9:19:	ident	"ORD"	false	0	0.000000	(9, 19) -> (9, 22)
test/test_004.md:9:22:	lparen	"("	false	0	0.000000	(9, 22) -> (9, 23)
test/test_004.md:9:23:	boolean	"FALSE"	false	0	0.000000	(9, 23) -> (9, 28)
test/test_004.md:9:28:	rparen	")"	false	0	0.000000	(9, 28) -> (9, 29)
test/test_004.md:9:29:	rparen	")"	false	0	0.000000	(9, 29) -> (9, 30)
test/test_004.md:9:30:	semicolon	";"	false	0	0.000000	(9, 30) -> (9, 31)
test/test_004.md:9:40:	ident	"Texts"	false	0	0.000000	(9, 40) -> (9, 45)
test/test_004.md:9:45:	dot	"."	false	0	0.000000	(9, 45) -> (9, 46)
test/test_004.md:9:46:	ident	"WriteLn"	false	0	0.000000	(9, 46) -> (9, 53)
test/test_004.md:9:53:	semicolon	";"	false	0	0.000000	(9, 53) -> (9, 54)
test/test_004.md:10:0:	end	"END"	false	0	0.000000	(10, 0) -> (10, 3)
test/test_004.md:10:4:	ident	"BooleanOperators"	false	0	0.000000	(10, 4) -> (10, 20)
test/test_004.md:10:20:	dot	"."	false	0	0.000000	(10, 20) -> (10, 21)
test/test_004.md:11:0:	eof	""	false	0	0.000000	(11, 0) -> (11, 0)
```
## AST
```scheme
(module "BooleanOperators"
  (stmts
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteInt")
        (call
          (procedure [i64] "ORD")
          (not [boolean]
            #true
          )
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteInt")
        (call
          (procedure [i64] "ORD")
          (not [boolean]
            #false
          )
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteInt")
        (call
          (procedure [i64] "ORD")
          (ampersand [boolean]
            #true
            #false
          )
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteInt")
        (call
          (procedure [i64] "ORD")
          (or [boolean]
            #true
            #false
          )
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteInt")
        (call
          (procedure [i64] "ORD")
          #true
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteInt")
        (call
          (procedure [i64] "ORD")
          #false
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
@0 = global [3 x i8] c"%d\00"
@1 = global [1 x i8] c"\00"
@__argc = global i64 0
@__argv = global i8** inttoptr (i8 0 to i8**)
@__envp = global i8** inttoptr (i8 0 to i8**)

declare i64 @puts(i8* %str)

declare i64 @rand()

declare i64 @sprintf(i8* %buf, i8* %format, ...)

declare i64 @printf(i8* %format, ...)

declare i8* @malloc(i64 %size)

declare i8* @free(i8* %ptr)

define void @oberonMain() {
entry:
	%0 = icmp eq i1 true, false
	%1 = zext i1 %0 to i64
	%2 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%3 = call i64 (i8*, ...) @printf(i8* %2, i64 %1)
	%4 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%5 = call i64 @puts(i8* %4)
	%6 = icmp eq i1 false, false
	%7 = zext i1 %6 to i64
	%8 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%9 = call i64 (i8*, ...) @printf(i8* %8, i64 %7)
	%10 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%11 = call i64 @puts(i8* %10)
	%12 = and i1 true, false
	%13 = zext i1 %12 to i64
	%14 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%15 = call i64 (i8*, ...) @printf(i8* %14, i64 %13)
	%16 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%17 = call i64 @puts(i8* %16)
	%18 = or i1 true, false
	%19 = zext i1 %18 to i64
	%20 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%21 = call i64 (i8*, ...) @printf(i8* %20, i64 %19)
	%22 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%23 = call i64 @puts(i8* %22)
	%24 = zext i1 true to i64
	%25 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%26 = call i64 (i8*, ...) @printf(i8* %25, i64 %24)
	%27 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%28 = call i64 @puts(i8* %27)
	%29 = zext i1 false to i64
	%30 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%31 = call i64 (i8*, ...) @printf(i8* %30, i64 %29)
	%32 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%33 = call i64 @puts(i8* %32)
	ret void
}

define i64 @main(i64 %argc, i8** %argv, i8** %argp) {
entry:
	store i64 %argc, i64* @__argc
	store i8** %argv, i8*** @__argv
	store i8** %argp, i8*** @__envp
	call void @oberonMain()
	ret i64 0
}

```
## Run
```bash
0
1
0
1
1
0
```
