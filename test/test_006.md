# test/test_006.md
## Source
```
(* test trailing semi-colon *)
MODULE trailing;

BEGIN
  print(1);
END trailing.
```
## Tokens
```tsv
test/test_006.md:2:0:	module	"MODULE"	false	0	0.000000	(2, 0) -> (2, 6)
test/test_006.md:2:7:	ident	"trailing"	false	0	0.000000	(2, 7) -> (2, 15)
test/test_006.md:2:15:	semicolon	";"	false	0	0.000000	(2, 15) -> (2, 16)
test/test_006.md:4:0:	begin	"BEGIN"	false	0	0.000000	(4, 0) -> (4, 5)
test/test_006.md:5:2:	ident	"print"	false	0	0.000000	(5, 2) -> (5, 7)
test/test_006.md:5:7:	lparen	"("	false	0	0.000000	(5, 7) -> (5, 8)
test/test_006.md:5:8:	integer	"1"	false	1	0.000000	(5, 8) -> (5, 9)
test/test_006.md:5:9:	rparen	")"	false	0	0.000000	(5, 9) -> (5, 10)
test/test_006.md:5:10:	semicolon	";"	false	0	0.000000	(5, 10) -> (5, 11)
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
      (print [void]
        (number [i64] 1)
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
	%0 = getelementptr [4 x i8], [4 x i8]* @0, i32 0, i32 0
	%1 = call i32 (i8*, ...) @printf(i8* %0, i32 1)
	ret i32 0
}

```
## Run
```bash
1
```
