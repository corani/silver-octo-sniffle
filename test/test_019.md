# test/test_019.md
## Source
```pascal
MODULE ParseErrors;

VAR set: SET;
    x  : INTEGER;

BEGIN
  x := (1+2;
  set := {1, 2;};
  set := {1..};
  Texts.WriteInt(3+
END Parse.
```
## Tokens
```tsv
test/test_019.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_019.md:1:8:	ident	"ParseErrors"	false	0	0.000000	(1, 8) -> (1, 19)
test/test_019.md:1:19:	semicolon	";"	false	0	0.000000	(1, 19) -> (1, 20)
test/test_019.md:3:0:	var	"VAR"	false	0	0.000000	(3, 0) -> (3, 3)
test/test_019.md:3:4:	ident	"set"	false	0	0.000000	(3, 4) -> (3, 7)
test/test_019.md:3:7:	colon	":"	false	0	0.000000	(3, 7) -> (3, 8)
test/test_019.md:3:9:	ident	"SET"	false	0	0.000000	(3, 9) -> (3, 12)
test/test_019.md:3:12:	semicolon	";"	false	0	0.000000	(3, 12) -> (3, 13)
test/test_019.md:4:4:	ident	"x"	false	0	0.000000	(4, 4) -> (4, 5)
test/test_019.md:4:7:	colon	":"	false	0	0.000000	(4, 7) -> (4, 8)
test/test_019.md:4:9:	ident	"INTEGER"	false	0	0.000000	(4, 9) -> (4, 16)
test/test_019.md:4:16:	semicolon	";"	false	0	0.000000	(4, 16) -> (4, 17)
test/test_019.md:6:0:	begin	"BEGIN"	false	0	0.000000	(6, 0) -> (6, 5)
test/test_019.md:7:2:	ident	"x"	false	0	0.000000	(7, 2) -> (7, 3)
test/test_019.md:7:4:	assign	":="	false	0	0.000000	(7, 4) -> (7, 6)
test/test_019.md:7:7:	lparen	"("	false	0	0.000000	(7, 7) -> (7, 8)
test/test_019.md:7:8:	integer	"1"	false	1	0.000000	(7, 8) -> (7, 9)
test/test_019.md:7:9:	plus	"+"	false	0	0.000000	(7, 9) -> (7, 10)
test/test_019.md:7:10:	integer	"2"	false	2	0.000000	(7, 10) -> (7, 11)
test/test_019.md:7:11:	semicolon	";"	false	0	0.000000	(7, 11) -> (7, 12)
test/test_019.md:8:2:	ident	"set"	false	0	0.000000	(8, 2) -> (8, 5)
test/test_019.md:8:6:	assign	":="	false	0	0.000000	(8, 6) -> (8, 8)
test/test_019.md:8:9:	lbrace	"{"	false	0	0.000000	(8, 9) -> (8, 10)
test/test_019.md:8:10:	integer	"1"	false	1	0.000000	(8, 10) -> (8, 11)
test/test_019.md:8:11:	comma	","	false	0	0.000000	(8, 11) -> (8, 12)
test/test_019.md:8:13:	integer	"2"	false	2	0.000000	(8, 13) -> (8, 14)
test/test_019.md:8:14:	semicolon	";"	false	0	0.000000	(8, 14) -> (8, 15)
test/test_019.md:8:15:	rbrace	"}"	false	0	0.000000	(8, 15) -> (8, 16)
test/test_019.md:8:16:	semicolon	";"	false	0	0.000000	(8, 16) -> (8, 17)
test/test_019.md:9:2:	ident	"set"	false	0	0.000000	(9, 2) -> (9, 5)
test/test_019.md:9:6:	assign	":="	false	0	0.000000	(9, 6) -> (9, 8)
test/test_019.md:9:9:	lbrace	"{"	false	0	0.000000	(9, 9) -> (9, 10)
test/test_019.md:9:10:	integer	"1"	false	1	0.000000	(9, 10) -> (9, 13)
test/test_019.md:9:13:	dotdot	".."	false	0	0.000000	(9, 13) -> (9, 15)
test/test_019.md:9:15:	rbrace	"}"	false	0	0.000000	(9, 15) -> (9, 16)
test/test_019.md:9:16:	semicolon	";"	false	0	0.000000	(9, 16) -> (9, 17)
test/test_019.md:10:2:	ident	"Texts"	false	0	0.000000	(10, 2) -> (10, 7)
test/test_019.md:10:7:	dot	"."	false	0	0.000000	(10, 7) -> (10, 8)
test/test_019.md:10:8:	ident	"WriteInt"	false	0	0.000000	(10, 8) -> (10, 16)
test/test_019.md:10:16:	lparen	"("	false	0	0.000000	(10, 16) -> (10, 17)
test/test_019.md:10:17:	integer	"3"	false	3	0.000000	(10, 17) -> (10, 18)
test/test_019.md:10:18:	plus	"+"	false	0	0.000000	(10, 18) -> (10, 19)
test/test_019.md:11:0:	end	"END"	false	0	0.000000	(11, 0) -> (11, 3)
test/test_019.md:11:4:	ident	"Parse"	false	0	0.000000	(11, 4) -> (11, 9)
test/test_019.md:11:9:	dot	"."	false	0	0.000000	(11, 9) -> (11, 10)
test/test_019.md:12:0:	eof	""	false	0	0.000000	(12, 0) -> (12, 0)
```
## Parser errors
```
test/test_019.md:7:11: ERROR: unexpected token: ;
test/test_019.md:7:11: INFO : expected: rparen
test/test_019.md:8:14: ERROR: unexpected token: ;
test/test_019.md:8:14: INFO : expected: rbrace
test/test_019.md:8:15: ERROR: unexpected token: }
test/test_019.md:8:15: INFO : expected a statement
test/test_019.md:9:15: ERROR: unexpected token: }
test/test_019.md:9:15: INFO : expected: integer
test/test_019.md:9:15: ERROR: unexpected token: }
test/test_019.md:9:15: INFO : expected: semicolon
test/test_019.md:9:16: ERROR: unexpected token: ;
test/test_019.md:9:16: INFO : expected: end
test/test_019.md:10:7: ERROR: unexpected identifier "Texts" at end of module
test/test_019.md:10:7: INFO : module name is "ParseErrors"
test/test_019.md:10:8: ERROR: unexpected token: WriteInt
test/test_019.md:10:8: INFO : expected: eof
```
