# test/test_019.md
## Source
```pascal
MODULE ParseErrors;

VAR set: SET;
    x  : INTEGER;

PROCEDURE Dummy
END

BEGIN
  x := (1+2;
  set := {1, 2;};
  set := {1..};
  print(3+
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
test/test_019.md:6:0:	procedure	"PROCEDURE"	false	0	0.000000	(6, 0) -> (6, 9)
test/test_019.md:6:10:	ident	"Dummy"	false	0	0.000000	(6, 10) -> (6, 15)
test/test_019.md:7:0:	end	"END"	false	0	0.000000	(7, 0) -> (7, 3)
test/test_019.md:9:0:	begin	"BEGIN"	false	0	0.000000	(9, 0) -> (9, 5)
test/test_019.md:10:2:	ident	"x"	false	0	0.000000	(10, 2) -> (10, 3)
test/test_019.md:10:4:	assign	":="	false	0	0.000000	(10, 4) -> (10, 6)
test/test_019.md:10:7:	lparen	"("	false	0	0.000000	(10, 7) -> (10, 8)
test/test_019.md:10:8:	integer	"1"	false	1	0.000000	(10, 8) -> (10, 9)
test/test_019.md:10:9:	plus	"+"	false	0	0.000000	(10, 9) -> (10, 10)
test/test_019.md:10:10:	integer	"2"	false	2	0.000000	(10, 10) -> (10, 11)
test/test_019.md:10:11:	semicolon	";"	false	0	0.000000	(10, 11) -> (10, 12)
test/test_019.md:11:2:	ident	"set"	false	0	0.000000	(11, 2) -> (11, 5)
test/test_019.md:11:6:	assign	":="	false	0	0.000000	(11, 6) -> (11, 8)
test/test_019.md:11:9:	lbrace	"{"	false	0	0.000000	(11, 9) -> (11, 10)
test/test_019.md:11:10:	integer	"1"	false	1	0.000000	(11, 10) -> (11, 11)
test/test_019.md:11:11:	comma	","	false	0	0.000000	(11, 11) -> (11, 12)
test/test_019.md:11:13:	integer	"2"	false	2	0.000000	(11, 13) -> (11, 14)
test/test_019.md:11:14:	semicolon	";"	false	0	0.000000	(11, 14) -> (11, 15)
test/test_019.md:11:15:	rbrace	"}"	false	0	0.000000	(11, 15) -> (11, 16)
test/test_019.md:11:16:	semicolon	";"	false	0	0.000000	(11, 16) -> (11, 17)
test/test_019.md:12:2:	ident	"set"	false	0	0.000000	(12, 2) -> (12, 5)
test/test_019.md:12:6:	assign	":="	false	0	0.000000	(12, 6) -> (12, 8)
test/test_019.md:12:9:	lbrace	"{"	false	0	0.000000	(12, 9) -> (12, 10)
test/test_019.md:12:10:	integer	"1"	false	1	0.000000	(12, 10) -> (12, 13)
test/test_019.md:12:13:	dotdot	".."	false	0	0.000000	(12, 13) -> (12, 15)
test/test_019.md:12:15:	rbrace	"}"	false	0	0.000000	(12, 15) -> (12, 16)
test/test_019.md:12:16:	semicolon	";"	false	0	0.000000	(12, 16) -> (12, 17)
test/test_019.md:13:2:	ident	"print"	false	0	0.000000	(13, 2) -> (13, 7)
test/test_019.md:13:7:	lparen	"("	false	0	0.000000	(13, 7) -> (13, 8)
test/test_019.md:13:8:	integer	"3"	false	3	0.000000	(13, 8) -> (13, 9)
test/test_019.md:13:9:	plus	"+"	false	0	0.000000	(13, 9) -> (13, 10)
test/test_019.md:14:0:	end	"END"	false	0	0.000000	(14, 0) -> (14, 3)
test/test_019.md:14:4:	ident	"Parse"	false	0	0.000000	(14, 4) -> (14, 9)
test/test_019.md:14:9:	dot	"."	false	0	0.000000	(14, 9) -> (14, 10)
test/test_019.md:15:0:	eof	""	false	0	0.000000	(15, 0) -> (15, 0)
```
## Parser errors
```
test/test_019.md:6:10: ERROR: PROCEDURE is unimplemented
test/test_019.md:10:11: ERROR: unexpected token: ;
test/test_019.md:10:11: INFO : expected: rparen
test/test_019.md:11:14: ERROR: unexpected token: ;
test/test_019.md:11:14: INFO : expected: rbrace
test/test_019.md:11:15: ERROR: unexpected token: }
test/test_019.md:11:15: INFO : expected a statement
test/test_019.md:12:15: ERROR: unexpected token: }
test/test_019.md:12:15: INFO : expected: integer
test/test_019.md:12:15: ERROR: unexpected token: }
test/test_019.md:12:15: INFO : expected: semicolon
test/test_019.md:12:16: ERROR: unexpected token: ;
test/test_019.md:12:16: INFO : expected: end
test/test_019.md:13:7: ERROR: unexpected identifier "print" at end of module
test/test_019.md:13:7: INFO : module name is "ParseErrors"
test/test_019.md:13:7: ERROR: unexpected token: (
test/test_019.md:13:7: INFO : expected: dot
```
