# test/test_003.md
## Source
```pascal
MODULE MultiLine;

BEGIN
    (*
      multi-line comments
      should work
    *)
    print("
      multi-line strings
      should work
    ")
END MultiLine.
```
## Tokens
```tsv
test/test_003.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_003.md:1:8:	ident	"MultiLine"	false	0	0.000000	(1, 8) -> (1, 17)
test/test_003.md:1:17:	semicolon	";"	false	0	0.000000	(1, 17) -> (1, 18)
test/test_003.md:3:0:	begin	"BEGIN"	false	0	0.000000	(3, 0) -> (3, 5)
test/test_003.md:8:4:	ident	"print"	false	0	0.000000	(8, 4) -> (8, 9)
test/test_003.md:8:9:	lparen	"("	false	0	0.000000	(8, 9) -> (8, 10)
test/test_003.md:8:10:	string	"\n      multi-line strings\n      should work\n    "	false	0	0.000000	(8, 10) -> (11, 5)
test/test_003.md:11:5:	rparen	")"	false	0	0.000000	(11, 5) -> (11, 6)
test/test_003.md:12:0:	end	"END"	false	0	0.000000	(12, 0) -> (12, 3)
test/test_003.md:12:4:	ident	"MultiLine"	false	0	0.000000	(12, 4) -> (12, 13)
test/test_003.md:12:13:	dot	"."	false	0	0.000000	(12, 13) -> (12, 14)
test/test_003.md:13:0:	eof	""	false	0	0.000000	(13, 0) -> (13, 0)
```
## AST
```scheme
(module "MultiLine"
  (stmts
    (expr2stmt
      (call
        (procedure [void] "print")
        (string "\n      multi-line strings\n      should work\n    ")
      )
    )
  )
)
```
## IR
```llvm
@0 = global [49 x i8] c"\0A      multi-line strings\0A      should work\0A    \00"

declare i64 @puts(i8* %str)

declare i64 @rand()

declare i64 @sprintf(i8* %buf, i8* %format, ...)

declare i64 @printf(i8* %format, ...)

declare i8* @malloc(i64 %size)

define i64 @main() {
entry:
	%0 = getelementptr [49 x i8], [49 x i8]* @0, i64 0, i64 0
	%1 = call i64 @puts(i8* %0)
	ret i64 0
}

```
## Run
```bash

      multi-line strings
      should work
    
```
