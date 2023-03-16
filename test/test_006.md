# test/test_006.md
## Source
```pascal
(* test trailing semi-colon *)
MODULE trailing;

BEGIN
  Texts.WriteInt(1);    Texts.WriteLn();
END trailing.
```
## Tokens
```tsv
test/test_006.md:2:0:	module	"MODULE"	false	0	0.000000	(2, 0) -> (2, 6)
test/test_006.md:2:7:	ident	"trailing"	false	0	0.000000	(2, 7) -> (2, 15)
test/test_006.md:2:15:	semicolon	";"	false	0	0.000000	(2, 15) -> (2, 16)
test/test_006.md:4:0:	begin	"BEGIN"	false	0	0.000000	(4, 0) -> (4, 5)
test/test_006.md:5:2:	ident	"Texts"	false	0	0.000000	(5, 2) -> (5, 7)
test/test_006.md:5:7:	dot	"."	false	0	0.000000	(5, 7) -> (5, 8)
test/test_006.md:5:8:	ident	"WriteInt"	false	0	0.000000	(5, 8) -> (5, 16)
test/test_006.md:5:16:	lparen	"("	false	0	0.000000	(5, 16) -> (5, 17)
test/test_006.md:5:17:	integer	"1"	false	1	0.000000	(5, 17) -> (5, 18)
test/test_006.md:5:18:	rparen	")"	false	0	0.000000	(5, 18) -> (5, 19)
test/test_006.md:5:19:	semicolon	";"	false	0	0.000000	(5, 19) -> (5, 20)
test/test_006.md:5:24:	ident	"Texts"	false	0	0.000000	(5, 24) -> (5, 29)
test/test_006.md:5:29:	dot	"."	false	0	0.000000	(5, 29) -> (5, 30)
test/test_006.md:5:30:	ident	"WriteLn"	false	0	0.000000	(5, 30) -> (5, 37)
test/test_006.md:5:37:	lparen	"("	false	0	0.000000	(5, 37) -> (5, 38)
test/test_006.md:5:38:	rparen	")"	false	0	0.000000	(5, 38) -> (5, 39)
test/test_006.md:5:39:	semicolon	";"	false	0	0.000000	(5, 39) -> (5, 40)
test/test_006.md:6:0:	end	"END"	false	0	0.000000	(6, 0) -> (6, 3)
test/test_006.md:6:4:	ident	"trailing"	false	0	0.000000	(6, 4) -> (6, 12)
test/test_006.md:6:12:	dot	"."	false	0	0.000000	(6, 12) -> (6, 13)
test/test_006.md:7:0:	eof	""	false	0	0.000000	(7, 0) -> (7, 0)
```
## AST
```scheme
(module "trailing"
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
```
## IR
```llvm
@0 = global [3 x i8] c"%d\00"
@1 = global [1 x i8] c"\00"
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
	%0 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%1 = call i64 (i8*, ...) @printf(i8* %0, i64 1)
	%2 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%3 = call i64 @puts(i8* %2)
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
```
