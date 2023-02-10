# test/test_019.md
## Source
```pascal
MODULE LexErrors;

PROCEDURE Dummy
END

BEGIN
  print(3+
END Lex.
```
## Tokens
```tsv
test/test_019.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_019.md:1:8:	ident	"LexErrors"	false	0	0.000000	(1, 8) -> (1, 17)
test/test_019.md:1:17:	semicolon	";"	false	0	0.000000	(1, 17) -> (1, 18)
test/test_019.md:3:0:	procedure	"PROCEDURE"	false	0	0.000000	(3, 0) -> (3, 9)
test/test_019.md:3:10:	ident	"Dummy"	false	0	0.000000	(3, 10) -> (3, 15)
test/test_019.md:4:0:	end	"END"	false	0	0.000000	(4, 0) -> (4, 3)
test/test_019.md:6:0:	begin	"BEGIN"	false	0	0.000000	(6, 0) -> (6, 5)
test/test_019.md:7:2:	ident	"print"	false	0	0.000000	(7, 2) -> (7, 7)
test/test_019.md:7:7:	lparen	"("	false	0	0.000000	(7, 7) -> (7, 8)
test/test_019.md:7:8:	integer	"3"	false	3	0.000000	(7, 8) -> (7, 9)
test/test_019.md:7:9:	plus	"+"	false	0	0.000000	(7, 9) -> (7, 10)
test/test_019.md:8:0:	end	"END"	false	0	0.000000	(8, 0) -> (8, 3)
test/test_019.md:8:4:	ident	"Lex"	false	0	0.000000	(8, 4) -> (8, 7)
test/test_019.md:8:7:	dot	"."	false	0	0.000000	(8, 7) -> (8, 8)
test/test_019.md:9:0:	eof	""	false	0	0.000000	(9, 0) -> (9, 0)
```
## Parser errors
```
test/test_019.md:3:10: ERROR: PROCEDURE is unimplemented
test/test_019.md:8:0: ERROR: unexpected token in factor: END
test/test_019.md:8:0: INFO : expected an expression
test/test_019.md:8:0: ERROR: unexpected token: END
test/test_019.md:8:0: INFO : expected: rparen
test/test_019.md:8:7: ERROR: unexpected identifier "Lex" at end of module
test/test_019.md:8:7: INFO : module name is "LexErrors"
```
