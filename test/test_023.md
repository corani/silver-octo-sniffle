# test/test_023.md
## Source
```pascal
MODULE PointerErrors;

VAR x : POINTER;

BEGIN
  NEW(x);
  x.y[33] := 1;
END PointerErrors.
```
## Tokens
```tsv
test/test_023.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_023.md:1:8:	ident	"PointerErrors"	false	0	0.000000	(1, 8) -> (1, 21)
test/test_023.md:1:21:	semicolon	";"	false	0	0.000000	(1, 21) -> (1, 22)
test/test_023.md:3:0:	var	"VAR"	false	0	0.000000	(3, 0) -> (3, 3)
test/test_023.md:3:4:	ident	"x"	false	0	0.000000	(3, 4) -> (3, 5)
test/test_023.md:3:6:	colon	":"	false	0	0.000000	(3, 6) -> (3, 7)
test/test_023.md:3:8:	pointer	"POINTER"	false	0	0.000000	(3, 8) -> (3, 15)
test/test_023.md:3:15:	semicolon	";"	false	0	0.000000	(3, 15) -> (3, 16)
test/test_023.md:5:0:	begin	"BEGIN"	false	0	0.000000	(5, 0) -> (5, 5)
test/test_023.md:6:2:	ident	"NEW"	false	0	0.000000	(6, 2) -> (6, 5)
test/test_023.md:6:5:	lparen	"("	false	0	0.000000	(6, 5) -> (6, 6)
test/test_023.md:6:6:	ident	"x"	false	0	0.000000	(6, 6) -> (6, 7)
test/test_023.md:6:7:	rparen	")"	false	0	0.000000	(6, 7) -> (6, 8)
test/test_023.md:6:8:	semicolon	";"	false	0	0.000000	(6, 8) -> (6, 9)
test/test_023.md:7:2:	ident	"x"	false	0	0.000000	(7, 2) -> (7, 3)
test/test_023.md:7:3:	dot	"."	false	0	0.000000	(7, 3) -> (7, 4)
test/test_023.md:7:4:	ident	"y"	false	0	0.000000	(7, 4) -> (7, 5)
test/test_023.md:7:5:	lbracket	"["	false	0	0.000000	(7, 5) -> (7, 6)
test/test_023.md:7:6:	integer	"33"	false	33	0.000000	(7, 6) -> (7, 8)
test/test_023.md:7:8:	rbracket	"]"	false	0	0.000000	(7, 8) -> (7, 9)
test/test_023.md:7:10:	assign	":="	false	0	0.000000	(7, 10) -> (7, 12)
test/test_023.md:7:13:	integer	"1"	false	1	0.000000	(7, 13) -> (7, 14)
test/test_023.md:7:14:	semicolon	";"	false	0	0.000000	(7, 14) -> (7, 15)
test/test_023.md:8:0:	end	"END"	false	0	0.000000	(8, 0) -> (8, 3)
test/test_023.md:8:4:	ident	"PointerErrors"	false	0	0.000000	(8, 4) -> (8, 17)
test/test_023.md:8:17:	dot	"."	false	0	0.000000	(8, 17) -> (8, 18)
test/test_023.md:9:0:	eof	""	false	0	0.000000	(9, 0) -> (9, 0)
```
## Parser errors
```
test/test_023.md:3:15: ERROR: unexpected token: ;
test/test_023.md:3:15: INFO : expected: to
```
