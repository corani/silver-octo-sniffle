# test/test_003.md
## Source
```pascal
MODULE MultiLine;

BEGIN
    (*
      multi-line comments
      should work
    *)
    Texts.WriteString("
      multi-line strings
      should work
    ");
    Texts.WriteLn()
END MultiLine.
```
## Tokens
```tsv
test/test_003.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_003.md:1:8:	ident	"MultiLine"	false	0	0.000000	(1, 8) -> (1, 17)
test/test_003.md:1:17:	semicolon	";"	false	0	0.000000	(1, 17) -> (1, 18)
test/test_003.md:3:0:	begin	"BEGIN"	false	0	0.000000	(3, 0) -> (3, 5)
test/test_003.md:8:4:	ident	"Texts"	false	0	0.000000	(8, 4) -> (8, 9)
test/test_003.md:8:9:	dot	"."	false	0	0.000000	(8, 9) -> (8, 10)
test/test_003.md:8:10:	ident	"WriteString"	false	0	0.000000	(8, 10) -> (8, 21)
test/test_003.md:8:21:	lparen	"("	false	0	0.000000	(8, 21) -> (8, 22)
test/test_003.md:8:22:	string	"\n      multi-line strings\n      should work\n    "	false	0	0.000000	(8, 22) -> (11, 5)
test/test_003.md:11:5:	rparen	")"	false	0	0.000000	(11, 5) -> (11, 6)
test/test_003.md:11:6:	semicolon	";"	false	0	0.000000	(11, 6) -> (11, 7)
test/test_003.md:12:4:	ident	"Texts"	false	0	0.000000	(12, 4) -> (12, 9)
test/test_003.md:12:9:	dot	"."	false	0	0.000000	(12, 9) -> (12, 10)
test/test_003.md:12:10:	ident	"WriteLn"	false	0	0.000000	(12, 10) -> (12, 17)
test/test_003.md:12:17:	lparen	"("	false	0	0.000000	(12, 17) -> (12, 18)
test/test_003.md:12:18:	rparen	")"	false	0	0.000000	(12, 18) -> (12, 19)
test/test_003.md:13:0:	end	"END"	false	0	0.000000	(13, 0) -> (13, 3)
test/test_003.md:13:4:	ident	"MultiLine"	false	0	0.000000	(13, 4) -> (13, 13)
test/test_003.md:13:13:	dot	"."	false	0	0.000000	(13, 13) -> (13, 14)
test/test_003.md:14:0:	eof	""	false	0	0.000000	(14, 0) -> (14, 0)
```
## AST
```scheme
(module "MultiLine"
  (stmts
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteString")
        (string "\n      multi-line strings\n      should work\n    ")
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
@0 = global [49 x i8] c"\0A      multi-line strings\0A      should work\0A    \00"
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
	%0 = getelementptr [49 x i8], [49 x i8]* @0, i64 0, i64 0
	%1 = call i64 (i8*, ...) @printf(i8* %0)
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

      multi-line strings
      should work
    
```
