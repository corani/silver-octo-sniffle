# test/test_020.md
## Source
```pascal
MODULE CheckErrors;

BEGIN
  IF 1 = "one" THEN
    print("wut")
  END
END CheckErrors.
```
## Tokens
```tsv
test/test_020.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_020.md:1:8:	ident	"CheckErrors"	false	0	0.000000	(1, 8) -> (1, 19)
test/test_020.md:1:19:	semicolon	";"	false	0	0.000000	(1, 19) -> (1, 20)
test/test_020.md:3:0:	begin	"BEGIN"	false	0	0.000000	(3, 0) -> (3, 5)
test/test_020.md:4:2:	if	"IF"	false	0	0.000000	(4, 2) -> (4, 4)
test/test_020.md:4:5:	integer	"1"	false	1	0.000000	(4, 5) -> (4, 6)
test/test_020.md:4:7:	eq	"="	false	0	0.000000	(4, 7) -> (4, 8)
test/test_020.md:4:9:	string	"one"	false	0	0.000000	(4, 9) -> (4, 14)
test/test_020.md:4:15:	then	"THEN"	false	0	0.000000	(4, 15) -> (4, 19)
test/test_020.md:5:4:	ident	"print"	false	0	0.000000	(5, 4) -> (5, 9)
test/test_020.md:5:9:	lparen	"("	false	0	0.000000	(5, 9) -> (5, 10)
test/test_020.md:5:10:	string	"wut"	false	0	0.000000	(5, 10) -> (5, 15)
test/test_020.md:5:15:	rparen	")"	false	0	0.000000	(5, 15) -> (5, 16)
test/test_020.md:6:2:	end	"END"	false	0	0.000000	(6, 2) -> (6, 5)
test/test_020.md:7:0:	end	"END"	false	0	0.000000	(7, 0) -> (7, 3)
test/test_020.md:7:4:	ident	"CheckErrors"	false	0	0.000000	(7, 4) -> (7, 15)
test/test_020.md:7:15:	dot	"."	false	0	0.000000	(7, 15) -> (7, 16)
test/test_020.md:8:0:	eof	""	false	0	0.000000	(8, 0) -> (8, 0)
```
## AST
```scheme
(module "CheckErrors"
  (stmts
    (if
      (eq [void]
        (number [i64] 1)
        (string "one")
      )
      (stmts
        (expr2stmt
          (call
            (procedure [void] "print")
            (string "wut")
          )
        )
      )
      (stmts
      )
    )
  )
)
```
## Type Check errors
```
test/test_020.md:4:7: ERROR: different types for binary operator eq
test/test_020.md:4:2: ERROR: condition for IF must be boolean, got void
```
