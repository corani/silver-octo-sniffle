# test/test_021.md
## Source
```pascal
MODULE CheckErrors;

CONST x = 33;

VAR y : x;
    z : bla;

BEGIN
END CheckErrors.
```
## Tokens
```tsv
test/test_021.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_021.md:1:8:	ident	"CheckErrors"	false	0	0.000000	(1, 8) -> (1, 19)
test/test_021.md:1:19:	semicolon	";"	false	0	0.000000	(1, 19) -> (1, 20)
test/test_021.md:3:0:	const	"CONST"	false	0	0.000000	(3, 0) -> (3, 5)
test/test_021.md:3:6:	ident	"x"	false	0	0.000000	(3, 6) -> (3, 7)
test/test_021.md:3:8:	eq	"="	false	0	0.000000	(3, 8) -> (3, 9)
test/test_021.md:3:10:	integer	"33"	false	33	0.000000	(3, 10) -> (3, 12)
test/test_021.md:3:12:	semicolon	";"	false	0	0.000000	(3, 12) -> (3, 13)
test/test_021.md:5:0:	var	"VAR"	false	0	0.000000	(5, 0) -> (5, 3)
test/test_021.md:5:4:	ident	"y"	false	0	0.000000	(5, 4) -> (5, 5)
test/test_021.md:5:6:	colon	":"	false	0	0.000000	(5, 6) -> (5, 7)
test/test_021.md:5:8:	ident	"x"	false	0	0.000000	(5, 8) -> (5, 9)
test/test_021.md:5:9:	semicolon	";"	false	0	0.000000	(5, 9) -> (5, 10)
test/test_021.md:6:4:	ident	"z"	false	0	0.000000	(6, 4) -> (6, 5)
test/test_021.md:6:6:	colon	":"	false	0	0.000000	(6, 6) -> (6, 7)
test/test_021.md:6:8:	ident	"bla"	false	0	0.000000	(6, 8) -> (6, 11)
test/test_021.md:6:11:	semicolon	";"	false	0	0.000000	(6, 11) -> (6, 12)
test/test_021.md:8:0:	begin	"BEGIN"	false	0	0.000000	(8, 0) -> (8, 5)
test/test_021.md:9:0:	end	"END"	false	0	0.000000	(9, 0) -> (9, 3)
test/test_021.md:9:4:	ident	"CheckErrors"	false	0	0.000000	(9, 4) -> (9, 15)
test/test_021.md:9:15:	dot	"."	false	0	0.000000	(9, 15) -> (9, 16)
test/test_021.md:10:0:	eof	""	false	0	0.000000	(10, 0) -> (10, 0)
```
## AST
```scheme
(module "CheckErrors"
  (consts
    (x [i64]
      (number [i64] 33)
    )
  )
  (vars
    (y [void])
    (z [void])
  )
  (stmts
  )
)
```
## Type Check errors
```
test/test_021.md:5:8: ERROR: unknown type "x" for variable "y"
test/test_021.md:5:8: INFO : "x" is a constant, not a type
test/test_021.md:6:8: ERROR: unknown type "bla" for variable "z"
```
