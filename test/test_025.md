# test/test_025.md
## Source
```pascal
MODULE QualIdent;

BEGIN
    Texts.WriteInt(33); Texts.WriteLn;
END QualIdent.
```
## Tokens
```tsv
test/test_025.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_025.md:1:8:	ident	"QualIdent"	false	0	0.000000	(1, 8) -> (1, 17)
test/test_025.md:1:17:	semicolon	";"	false	0	0.000000	(1, 17) -> (1, 18)
test/test_025.md:3:0:	begin	"BEGIN"	false	0	0.000000	(3, 0) -> (3, 5)
test/test_025.md:4:4:	ident	"Texts"	false	0	0.000000	(4, 4) -> (4, 9)
test/test_025.md:4:9:	dot	"."	false	0	0.000000	(4, 9) -> (4, 10)
test/test_025.md:4:10:	ident	"WriteInt"	false	0	0.000000	(4, 10) -> (4, 18)
test/test_025.md:4:18:	lparen	"("	false	0	0.000000	(4, 18) -> (4, 19)
test/test_025.md:4:19:	integer	"33"	false	33	0.000000	(4, 19) -> (4, 21)
test/test_025.md:4:21:	rparen	")"	false	0	0.000000	(4, 21) -> (4, 22)
test/test_025.md:4:22:	semicolon	";"	false	0	0.000000	(4, 22) -> (4, 23)
test/test_025.md:4:24:	ident	"Texts"	false	0	0.000000	(4, 24) -> (4, 29)
test/test_025.md:4:29:	dot	"."	false	0	0.000000	(4, 29) -> (4, 30)
test/test_025.md:4:30:	ident	"WriteLn"	false	0	0.000000	(4, 30) -> (4, 37)
test/test_025.md:4:37:	semicolon	";"	false	0	0.000000	(4, 37) -> (4, 38)
test/test_025.md:5:0:	end	"END"	false	0	0.000000	(5, 0) -> (5, 3)
test/test_025.md:5:4:	ident	"QualIdent"	false	0	0.000000	(5, 4) -> (5, 13)
test/test_025.md:5:13:	dot	"."	false	0	0.000000	(5, 13) -> (5, 14)
test/test_025.md:6:0:	eof	""	false	0	0.000000	(6, 0) -> (6, 0)
```
## AST
```scheme
(module "QualIdent"
  (stmts
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteInt")
        (number [i64] 33)
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
	%0 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%1 = call i64 (i8*, ...) @printf(i8* %0, i64 33)
	%2 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%3 = call i64 @puts(i8* %2)
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
33
```
