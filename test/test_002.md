# test/test_002.md
## Source
```pascal
MODULE BinaryOperations;

BEGIN
    Texts.WriteInt(+33);                                        Texts.WriteLn;
    Texts.WriteInt(-42);                                        Texts.WriteLn;
    Texts.WriteInt( 34 + 35 + 42000 );                          Texts.WriteLn;
    Texts.WriteInt(66-33);                                      Texts.WriteLn;
    Texts.WriteInt( -34 + 35 );                                 Texts.WriteLn;
    Texts.WriteInt( (34 + 35) + 42000 );                        Texts.WriteLn;
    Texts.WriteInt( (34 + 35) * 33 );                           Texts.WriteLn;
    Texts.WriteReal( FLT(34 + 35) / FLT(33) );                  Texts.WriteLn;
    Texts.WriteInt(11 DIV 3);                                   Texts.WriteLn;
    Texts.WriteInt(11 MOD 3);                                   Texts.WriteLn;
    Texts.WriteReal(FLT(11) / FLT(3) / FLT(2));                 Texts.WriteLn;
    Texts.WriteReal((FLT(11) / FLT(3)) + (FLT(7) / FLT(2)));    Texts.WriteLn;
    Texts.WriteReal((11.0/3.0)-(7.0/2.0));                      Texts.WriteLn;
    Texts.WriteReal(FLT(11)/FLT(3)*FLT(3));                     Texts.WriteLn;
    Texts.WriteReal(FLT(69)+0.420);                             Texts.WriteLn;
    Texts.WriteReal(0.69E+2+420.E-3);                           Texts.WriteLn;
END BinaryOperations.
```
## Tokens
```tsv
test/test_002.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_002.md:1:8:	ident	"BinaryOperations"	false	0	0.000000	(1, 8) -> (1, 24)
test/test_002.md:1:24:	semicolon	";"	false	0	0.000000	(1, 24) -> (1, 25)
test/test_002.md:3:0:	begin	"BEGIN"	false	0	0.000000	(3, 0) -> (3, 5)
test/test_002.md:4:4:	ident	"Texts"	false	0	0.000000	(4, 4) -> (4, 9)
test/test_002.md:4:9:	dot	"."	false	0	0.000000	(4, 9) -> (4, 10)
test/test_002.md:4:10:	ident	"WriteInt"	false	0	0.000000	(4, 10) -> (4, 18)
test/test_002.md:4:18:	lparen	"("	false	0	0.000000	(4, 18) -> (4, 19)
test/test_002.md:4:19:	plus	"+"	false	0	0.000000	(4, 19) -> (4, 20)
test/test_002.md:4:20:	integer	"33"	false	33	0.000000	(4, 20) -> (4, 22)
test/test_002.md:4:22:	rparen	")"	false	0	0.000000	(4, 22) -> (4, 23)
test/test_002.md:4:23:	semicolon	";"	false	0	0.000000	(4, 23) -> (4, 24)
test/test_002.md:4:64:	ident	"Texts"	false	0	0.000000	(4, 64) -> (4, 69)
test/test_002.md:4:69:	dot	"."	false	0	0.000000	(4, 69) -> (4, 70)
test/test_002.md:4:70:	ident	"WriteLn"	false	0	0.000000	(4, 70) -> (4, 77)
test/test_002.md:4:77:	semicolon	";"	false	0	0.000000	(4, 77) -> (4, 78)
test/test_002.md:5:4:	ident	"Texts"	false	0	0.000000	(5, 4) -> (5, 9)
test/test_002.md:5:9:	dot	"."	false	0	0.000000	(5, 9) -> (5, 10)
test/test_002.md:5:10:	ident	"WriteInt"	false	0	0.000000	(5, 10) -> (5, 18)
test/test_002.md:5:18:	lparen	"("	false	0	0.000000	(5, 18) -> (5, 19)
test/test_002.md:5:19:	minus	"-"	false	0	0.000000	(5, 19) -> (5, 20)
test/test_002.md:5:20:	integer	"42"	false	42	0.000000	(5, 20) -> (5, 22)
test/test_002.md:5:22:	rparen	")"	false	0	0.000000	(5, 22) -> (5, 23)
test/test_002.md:5:23:	semicolon	";"	false	0	0.000000	(5, 23) -> (5, 24)
test/test_002.md:5:64:	ident	"Texts"	false	0	0.000000	(5, 64) -> (5, 69)
test/test_002.md:5:69:	dot	"."	false	0	0.000000	(5, 69) -> (5, 70)
test/test_002.md:5:70:	ident	"WriteLn"	false	0	0.000000	(5, 70) -> (5, 77)
test/test_002.md:5:77:	semicolon	";"	false	0	0.000000	(5, 77) -> (5, 78)
test/test_002.md:6:4:	ident	"Texts"	false	0	0.000000	(6, 4) -> (6, 9)
test/test_002.md:6:9:	dot	"."	false	0	0.000000	(6, 9) -> (6, 10)
test/test_002.md:6:10:	ident	"WriteInt"	false	0	0.000000	(6, 10) -> (6, 18)
test/test_002.md:6:18:	lparen	"("	false	0	0.000000	(6, 18) -> (6, 19)
test/test_002.md:6:20:	integer	"34"	false	34	0.000000	(6, 20) -> (6, 22)
test/test_002.md:6:23:	plus	"+"	false	0	0.000000	(6, 23) -> (6, 24)
test/test_002.md:6:25:	integer	"35"	false	35	0.000000	(6, 25) -> (6, 27)
test/test_002.md:6:28:	plus	"+"	false	0	0.000000	(6, 28) -> (6, 29)
test/test_002.md:6:30:	integer	"42000"	false	42000	0.000000	(6, 30) -> (6, 35)
test/test_002.md:6:36:	rparen	")"	false	0	0.000000	(6, 36) -> (6, 37)
test/test_002.md:6:37:	semicolon	";"	false	0	0.000000	(6, 37) -> (6, 38)
test/test_002.md:6:64:	ident	"Texts"	false	0	0.000000	(6, 64) -> (6, 69)
test/test_002.md:6:69:	dot	"."	false	0	0.000000	(6, 69) -> (6, 70)
test/test_002.md:6:70:	ident	"WriteLn"	false	0	0.000000	(6, 70) -> (6, 77)
test/test_002.md:6:77:	semicolon	";"	false	0	0.000000	(6, 77) -> (6, 78)
test/test_002.md:7:4:	ident	"Texts"	false	0	0.000000	(7, 4) -> (7, 9)
test/test_002.md:7:9:	dot	"."	false	0	0.000000	(7, 9) -> (7, 10)
test/test_002.md:7:10:	ident	"WriteInt"	false	0	0.000000	(7, 10) -> (7, 18)
test/test_002.md:7:18:	lparen	"("	false	0	0.000000	(7, 18) -> (7, 19)
test/test_002.md:7:19:	integer	"66"	false	66	0.000000	(7, 19) -> (7, 21)
test/test_002.md:7:21:	minus	"-"	false	0	0.000000	(7, 21) -> (7, 22)
test/test_002.md:7:22:	integer	"33"	false	33	0.000000	(7, 22) -> (7, 24)
test/test_002.md:7:24:	rparen	")"	false	0	0.000000	(7, 24) -> (7, 25)
test/test_002.md:7:25:	semicolon	";"	false	0	0.000000	(7, 25) -> (7, 26)
test/test_002.md:7:64:	ident	"Texts"	false	0	0.000000	(7, 64) -> (7, 69)
test/test_002.md:7:69:	dot	"."	false	0	0.000000	(7, 69) -> (7, 70)
test/test_002.md:7:70:	ident	"WriteLn"	false	0	0.000000	(7, 70) -> (7, 77)
test/test_002.md:7:77:	semicolon	";"	false	0	0.000000	(7, 77) -> (7, 78)
test/test_002.md:8:4:	ident	"Texts"	false	0	0.000000	(8, 4) -> (8, 9)
test/test_002.md:8:9:	dot	"."	false	0	0.000000	(8, 9) -> (8, 10)
test/test_002.md:8:10:	ident	"WriteInt"	false	0	0.000000	(8, 10) -> (8, 18)
test/test_002.md:8:18:	lparen	"("	false	0	0.000000	(8, 18) -> (8, 19)
test/test_002.md:8:20:	minus	"-"	false	0	0.000000	(8, 20) -> (8, 21)
test/test_002.md:8:21:	integer	"34"	false	34	0.000000	(8, 21) -> (8, 23)
test/test_002.md:8:24:	plus	"+"	false	0	0.000000	(8, 24) -> (8, 25)
test/test_002.md:8:26:	integer	"35"	false	35	0.000000	(8, 26) -> (8, 28)
test/test_002.md:8:29:	rparen	")"	false	0	0.000000	(8, 29) -> (8, 30)
test/test_002.md:8:30:	semicolon	";"	false	0	0.000000	(8, 30) -> (8, 31)
test/test_002.md:8:64:	ident	"Texts"	false	0	0.000000	(8, 64) -> (8, 69)
test/test_002.md:8:69:	dot	"."	false	0	0.000000	(8, 69) -> (8, 70)
test/test_002.md:8:70:	ident	"WriteLn"	false	0	0.000000	(8, 70) -> (8, 77)
test/test_002.md:8:77:	semicolon	";"	false	0	0.000000	(8, 77) -> (8, 78)
test/test_002.md:9:4:	ident	"Texts"	false	0	0.000000	(9, 4) -> (9, 9)
test/test_002.md:9:9:	dot	"."	false	0	0.000000	(9, 9) -> (9, 10)
test/test_002.md:9:10:	ident	"WriteInt"	false	0	0.000000	(9, 10) -> (9, 18)
test/test_002.md:9:18:	lparen	"("	false	0	0.000000	(9, 18) -> (9, 19)
test/test_002.md:9:20:	lparen	"("	false	0	0.000000	(9, 20) -> (9, 21)
test/test_002.md:9:21:	integer	"34"	false	34	0.000000	(9, 21) -> (9, 23)
test/test_002.md:9:24:	plus	"+"	false	0	0.000000	(9, 24) -> (9, 25)
test/test_002.md:9:26:	integer	"35"	false	35	0.000000	(9, 26) -> (9, 28)
test/test_002.md:9:28:	rparen	")"	false	0	0.000000	(9, 28) -> (9, 29)
test/test_002.md:9:30:	plus	"+"	false	0	0.000000	(9, 30) -> (9, 31)
test/test_002.md:9:32:	integer	"42000"	false	42000	0.000000	(9, 32) -> (9, 37)
test/test_002.md:9:38:	rparen	")"	false	0	0.000000	(9, 38) -> (9, 39)
test/test_002.md:9:39:	semicolon	";"	false	0	0.000000	(9, 39) -> (9, 40)
test/test_002.md:9:64:	ident	"Texts"	false	0	0.000000	(9, 64) -> (9, 69)
test/test_002.md:9:69:	dot	"."	false	0	0.000000	(9, 69) -> (9, 70)
test/test_002.md:9:70:	ident	"WriteLn"	false	0	0.000000	(9, 70) -> (9, 77)
test/test_002.md:9:77:	semicolon	";"	false	0	0.000000	(9, 77) -> (9, 78)
test/test_002.md:10:4:	ident	"Texts"	false	0	0.000000	(10, 4) -> (10, 9)
test/test_002.md:10:9:	dot	"."	false	0	0.000000	(10, 9) -> (10, 10)
test/test_002.md:10:10:	ident	"WriteInt"	false	0	0.000000	(10, 10) -> (10, 18)
test/test_002.md:10:18:	lparen	"("	false	0	0.000000	(10, 18) -> (10, 19)
test/test_002.md:10:20:	lparen	"("	false	0	0.000000	(10, 20) -> (10, 21)
test/test_002.md:10:21:	integer	"34"	false	34	0.000000	(10, 21) -> (10, 23)
test/test_002.md:10:24:	plus	"+"	false	0	0.000000	(10, 24) -> (10, 25)
test/test_002.md:10:26:	integer	"35"	false	35	0.000000	(10, 26) -> (10, 28)
test/test_002.md:10:28:	rparen	")"	false	0	0.000000	(10, 28) -> (10, 29)
test/test_002.md:10:30:	asterisk	"*"	false	0	0.000000	(10, 30) -> (10, 31)
test/test_002.md:10:32:	integer	"33"	false	33	0.000000	(10, 32) -> (10, 34)
test/test_002.md:10:35:	rparen	")"	false	0	0.000000	(10, 35) -> (10, 36)
test/test_002.md:10:36:	semicolon	";"	false	0	0.000000	(10, 36) -> (10, 37)
test/test_002.md:10:64:	ident	"Texts"	false	0	0.000000	(10, 64) -> (10, 69)
test/test_002.md:10:69:	dot	"."	false	0	0.000000	(10, 69) -> (10, 70)
test/test_002.md:10:70:	ident	"WriteLn"	false	0	0.000000	(10, 70) -> (10, 77)
test/test_002.md:10:77:	semicolon	";"	false	0	0.000000	(10, 77) -> (10, 78)
test/test_002.md:11:4:	ident	"Texts"	false	0	0.000000	(11, 4) -> (11, 9)
test/test_002.md:11:9:	dot	"."	false	0	0.000000	(11, 9) -> (11, 10)
test/test_002.md:11:10:	ident	"WriteReal"	false	0	0.000000	(11, 10) -> (11, 19)
test/test_002.md:11:19:	lparen	"("	false	0	0.000000	(11, 19) -> (11, 20)
test/test_002.md:11:21:	ident	"FLT"	false	0	0.000000	(11, 21) -> (11, 24)
test/test_002.md:11:24:	lparen	"("	false	0	0.000000	(11, 24) -> (11, 25)
test/test_002.md:11:25:	integer	"34"	false	34	0.000000	(11, 25) -> (11, 27)
test/test_002.md:11:28:	plus	"+"	false	0	0.000000	(11, 28) -> (11, 29)
test/test_002.md:11:30:	integer	"35"	false	35	0.000000	(11, 30) -> (11, 32)
test/test_002.md:11:32:	rparen	")"	false	0	0.000000	(11, 32) -> (11, 33)
test/test_002.md:11:34:	slash	"/"	false	0	0.000000	(11, 34) -> (11, 35)
test/test_002.md:11:36:	ident	"FLT"	false	0	0.000000	(11, 36) -> (11, 39)
test/test_002.md:11:39:	lparen	"("	false	0	0.000000	(11, 39) -> (11, 40)
test/test_002.md:11:40:	integer	"33"	false	33	0.000000	(11, 40) -> (11, 42)
test/test_002.md:11:42:	rparen	")"	false	0	0.000000	(11, 42) -> (11, 43)
test/test_002.md:11:44:	rparen	")"	false	0	0.000000	(11, 44) -> (11, 45)
test/test_002.md:11:45:	semicolon	";"	false	0	0.000000	(11, 45) -> (11, 46)
test/test_002.md:11:64:	ident	"Texts"	false	0	0.000000	(11, 64) -> (11, 69)
test/test_002.md:11:69:	dot	"."	false	0	0.000000	(11, 69) -> (11, 70)
test/test_002.md:11:70:	ident	"WriteLn"	false	0	0.000000	(11, 70) -> (11, 77)
test/test_002.md:11:77:	semicolon	";"	false	0	0.000000	(11, 77) -> (11, 78)
test/test_002.md:12:4:	ident	"Texts"	false	0	0.000000	(12, 4) -> (12, 9)
test/test_002.md:12:9:	dot	"."	false	0	0.000000	(12, 9) -> (12, 10)
test/test_002.md:12:10:	ident	"WriteInt"	false	0	0.000000	(12, 10) -> (12, 18)
test/test_002.md:12:18:	lparen	"("	false	0	0.000000	(12, 18) -> (12, 19)
test/test_002.md:12:19:	integer	"11"	false	11	0.000000	(12, 19) -> (12, 21)
test/test_002.md:12:22:	div	"DIV"	false	0	0.000000	(12, 22) -> (12, 25)
test/test_002.md:12:26:	integer	"3"	false	3	0.000000	(12, 26) -> (12, 27)
test/test_002.md:12:27:	rparen	")"	false	0	0.000000	(12, 27) -> (12, 28)
test/test_002.md:12:28:	semicolon	";"	false	0	0.000000	(12, 28) -> (12, 29)
test/test_002.md:12:64:	ident	"Texts"	false	0	0.000000	(12, 64) -> (12, 69)
test/test_002.md:12:69:	dot	"."	false	0	0.000000	(12, 69) -> (12, 70)
test/test_002.md:12:70:	ident	"WriteLn"	false	0	0.000000	(12, 70) -> (12, 77)
test/test_002.md:12:77:	semicolon	";"	false	0	0.000000	(12, 77) -> (12, 78)
test/test_002.md:13:4:	ident	"Texts"	false	0	0.000000	(13, 4) -> (13, 9)
test/test_002.md:13:9:	dot	"."	false	0	0.000000	(13, 9) -> (13, 10)
test/test_002.md:13:10:	ident	"WriteInt"	false	0	0.000000	(13, 10) -> (13, 18)
test/test_002.md:13:18:	lparen	"("	false	0	0.000000	(13, 18) -> (13, 19)
test/test_002.md:13:19:	integer	"11"	false	11	0.000000	(13, 19) -> (13, 21)
test/test_002.md:13:22:	mod	"MOD"	false	0	0.000000	(13, 22) -> (13, 25)
test/test_002.md:13:26:	integer	"3"	false	3	0.000000	(13, 26) -> (13, 27)
test/test_002.md:13:27:	rparen	")"	false	0	0.000000	(13, 27) -> (13, 28)
test/test_002.md:13:28:	semicolon	";"	false	0	0.000000	(13, 28) -> (13, 29)
test/test_002.md:13:64:	ident	"Texts"	false	0	0.000000	(13, 64) -> (13, 69)
test/test_002.md:13:69:	dot	"."	false	0	0.000000	(13, 69) -> (13, 70)
test/test_002.md:13:70:	ident	"WriteLn"	false	0	0.000000	(13, 70) -> (13, 77)
test/test_002.md:13:77:	semicolon	";"	false	0	0.000000	(13, 77) -> (13, 78)
test/test_002.md:14:4:	ident	"Texts"	false	0	0.000000	(14, 4) -> (14, 9)
test/test_002.md:14:9:	dot	"."	false	0	0.000000	(14, 9) -> (14, 10)
test/test_002.md:14:10:	ident	"WriteReal"	false	0	0.000000	(14, 10) -> (14, 19)
test/test_002.md:14:19:	lparen	"("	false	0	0.000000	(14, 19) -> (14, 20)
test/test_002.md:14:20:	ident	"FLT"	false	0	0.000000	(14, 20) -> (14, 23)
test/test_002.md:14:23:	lparen	"("	false	0	0.000000	(14, 23) -> (14, 24)
test/test_002.md:14:24:	integer	"11"	false	11	0.000000	(14, 24) -> (14, 26)
test/test_002.md:14:26:	rparen	")"	false	0	0.000000	(14, 26) -> (14, 27)
test/test_002.md:14:28:	slash	"/"	false	0	0.000000	(14, 28) -> (14, 29)
test/test_002.md:14:30:	ident	"FLT"	false	0	0.000000	(14, 30) -> (14, 33)
test/test_002.md:14:33:	lparen	"("	false	0	0.000000	(14, 33) -> (14, 34)
test/test_002.md:14:34:	integer	"3"	false	3	0.000000	(14, 34) -> (14, 35)
test/test_002.md:14:35:	rparen	")"	false	0	0.000000	(14, 35) -> (14, 36)
test/test_002.md:14:37:	slash	"/"	false	0	0.000000	(14, 37) -> (14, 38)
test/test_002.md:14:39:	ident	"FLT"	false	0	0.000000	(14, 39) -> (14, 42)
test/test_002.md:14:42:	lparen	"("	false	0	0.000000	(14, 42) -> (14, 43)
test/test_002.md:14:43:	integer	"2"	false	2	0.000000	(14, 43) -> (14, 44)
test/test_002.md:14:44:	rparen	")"	false	0	0.000000	(14, 44) -> (14, 45)
test/test_002.md:14:45:	rparen	")"	false	0	0.000000	(14, 45) -> (14, 46)
test/test_002.md:14:46:	semicolon	";"	false	0	0.000000	(14, 46) -> (14, 47)
test/test_002.md:14:64:	ident	"Texts"	false	0	0.000000	(14, 64) -> (14, 69)
test/test_002.md:14:69:	dot	"."	false	0	0.000000	(14, 69) -> (14, 70)
test/test_002.md:14:70:	ident	"WriteLn"	false	0	0.000000	(14, 70) -> (14, 77)
test/test_002.md:14:77:	semicolon	";"	false	0	0.000000	(14, 77) -> (14, 78)
test/test_002.md:15:4:	ident	"Texts"	false	0	0.000000	(15, 4) -> (15, 9)
test/test_002.md:15:9:	dot	"."	false	0	0.000000	(15, 9) -> (15, 10)
test/test_002.md:15:10:	ident	"WriteReal"	false	0	0.000000	(15, 10) -> (15, 19)
test/test_002.md:15:19:	lparen	"("	false	0	0.000000	(15, 19) -> (15, 20)
test/test_002.md:15:20:	lparen	"("	false	0	0.000000	(15, 20) -> (15, 21)
test/test_002.md:15:21:	ident	"FLT"	false	0	0.000000	(15, 21) -> (15, 24)
test/test_002.md:15:24:	lparen	"("	false	0	0.000000	(15, 24) -> (15, 25)
test/test_002.md:15:25:	integer	"11"	false	11	0.000000	(15, 25) -> (15, 27)
test/test_002.md:15:27:	rparen	")"	false	0	0.000000	(15, 27) -> (15, 28)
test/test_002.md:15:29:	slash	"/"	false	0	0.000000	(15, 29) -> (15, 30)
test/test_002.md:15:31:	ident	"FLT"	false	0	0.000000	(15, 31) -> (15, 34)
test/test_002.md:15:34:	lparen	"("	false	0	0.000000	(15, 34) -> (15, 35)
test/test_002.md:15:35:	integer	"3"	false	3	0.000000	(15, 35) -> (15, 36)
test/test_002.md:15:36:	rparen	")"	false	0	0.000000	(15, 36) -> (15, 37)
test/test_002.md:15:37:	rparen	")"	false	0	0.000000	(15, 37) -> (15, 38)
test/test_002.md:15:39:	plus	"+"	false	0	0.000000	(15, 39) -> (15, 40)
test/test_002.md:15:41:	lparen	"("	false	0	0.000000	(15, 41) -> (15, 42)
test/test_002.md:15:42:	ident	"FLT"	false	0	0.000000	(15, 42) -> (15, 45)
test/test_002.md:15:45:	lparen	"("	false	0	0.000000	(15, 45) -> (15, 46)
test/test_002.md:15:46:	integer	"7"	false	7	0.000000	(15, 46) -> (15, 47)
test/test_002.md:15:47:	rparen	")"	false	0	0.000000	(15, 47) -> (15, 48)
test/test_002.md:15:49:	slash	"/"	false	0	0.000000	(15, 49) -> (15, 50)
test/test_002.md:15:51:	ident	"FLT"	false	0	0.000000	(15, 51) -> (15, 54)
test/test_002.md:15:54:	lparen	"("	false	0	0.000000	(15, 54) -> (15, 55)
test/test_002.md:15:55:	integer	"2"	false	2	0.000000	(15, 55) -> (15, 56)
test/test_002.md:15:56:	rparen	")"	false	0	0.000000	(15, 56) -> (15, 57)
test/test_002.md:15:57:	rparen	")"	false	0	0.000000	(15, 57) -> (15, 58)
test/test_002.md:15:58:	rparen	")"	false	0	0.000000	(15, 58) -> (15, 59)
test/test_002.md:15:59:	semicolon	";"	false	0	0.000000	(15, 59) -> (15, 60)
test/test_002.md:15:64:	ident	"Texts"	false	0	0.000000	(15, 64) -> (15, 69)
test/test_002.md:15:69:	dot	"."	false	0	0.000000	(15, 69) -> (15, 70)
test/test_002.md:15:70:	ident	"WriteLn"	false	0	0.000000	(15, 70) -> (15, 77)
test/test_002.md:15:77:	semicolon	";"	false	0	0.000000	(15, 77) -> (15, 78)
test/test_002.md:16:4:	ident	"Texts"	false	0	0.000000	(16, 4) -> (16, 9)
test/test_002.md:16:9:	dot	"."	false	0	0.000000	(16, 9) -> (16, 10)
test/test_002.md:16:10:	ident	"WriteReal"	false	0	0.000000	(16, 10) -> (16, 19)
test/test_002.md:16:19:	lparen	"("	false	0	0.000000	(16, 19) -> (16, 20)
test/test_002.md:16:20:	lparen	"("	false	0	0.000000	(16, 20) -> (16, 21)
test/test_002.md:16:21:	real	"11.0"	false	0	11.000000	(16, 21) -> (16, 25)
test/test_002.md:16:25:	slash	"/"	false	0	0.000000	(16, 25) -> (16, 26)
test/test_002.md:16:26:	real	"3.0"	false	0	3.000000	(16, 26) -> (16, 29)
test/test_002.md:16:29:	rparen	")"	false	0	0.000000	(16, 29) -> (16, 30)
test/test_002.md:16:30:	minus	"-"	false	0	0.000000	(16, 30) -> (16, 31)
test/test_002.md:16:31:	lparen	"("	false	0	0.000000	(16, 31) -> (16, 32)
test/test_002.md:16:32:	real	"7.0"	false	0	7.000000	(16, 32) -> (16, 35)
test/test_002.md:16:35:	slash	"/"	false	0	0.000000	(16, 35) -> (16, 36)
test/test_002.md:16:36:	real	"2.0"	false	0	2.000000	(16, 36) -> (16, 39)
test/test_002.md:16:39:	rparen	")"	false	0	0.000000	(16, 39) -> (16, 40)
test/test_002.md:16:40:	rparen	")"	false	0	0.000000	(16, 40) -> (16, 41)
test/test_002.md:16:41:	semicolon	";"	false	0	0.000000	(16, 41) -> (16, 42)
test/test_002.md:16:64:	ident	"Texts"	false	0	0.000000	(16, 64) -> (16, 69)
test/test_002.md:16:69:	dot	"."	false	0	0.000000	(16, 69) -> (16, 70)
test/test_002.md:16:70:	ident	"WriteLn"	false	0	0.000000	(16, 70) -> (16, 77)
test/test_002.md:16:77:	semicolon	";"	false	0	0.000000	(16, 77) -> (16, 78)
test/test_002.md:17:4:	ident	"Texts"	false	0	0.000000	(17, 4) -> (17, 9)
test/test_002.md:17:9:	dot	"."	false	0	0.000000	(17, 9) -> (17, 10)
test/test_002.md:17:10:	ident	"WriteReal"	false	0	0.000000	(17, 10) -> (17, 19)
test/test_002.md:17:19:	lparen	"("	false	0	0.000000	(17, 19) -> (17, 20)
test/test_002.md:17:20:	ident	"FLT"	false	0	0.000000	(17, 20) -> (17, 23)
test/test_002.md:17:23:	lparen	"("	false	0	0.000000	(17, 23) -> (17, 24)
test/test_002.md:17:24:	integer	"11"	false	11	0.000000	(17, 24) -> (17, 26)
test/test_002.md:17:26:	rparen	")"	false	0	0.000000	(17, 26) -> (17, 27)
test/test_002.md:17:27:	slash	"/"	false	0	0.000000	(17, 27) -> (17, 28)
test/test_002.md:17:28:	ident	"FLT"	false	0	0.000000	(17, 28) -> (17, 31)
test/test_002.md:17:31:	lparen	"("	false	0	0.000000	(17, 31) -> (17, 32)
test/test_002.md:17:32:	integer	"3"	false	3	0.000000	(17, 32) -> (17, 33)
test/test_002.md:17:33:	rparen	")"	false	0	0.000000	(17, 33) -> (17, 34)
test/test_002.md:17:34:	asterisk	"*"	false	0	0.000000	(17, 34) -> (17, 35)
test/test_002.md:17:35:	ident	"FLT"	false	0	0.000000	(17, 35) -> (17, 38)
test/test_002.md:17:38:	lparen	"("	false	0	0.000000	(17, 38) -> (17, 39)
test/test_002.md:17:39:	integer	"3"	false	3	0.000000	(17, 39) -> (17, 40)
test/test_002.md:17:40:	rparen	")"	false	0	0.000000	(17, 40) -> (17, 41)
test/test_002.md:17:41:	rparen	")"	false	0	0.000000	(17, 41) -> (17, 42)
test/test_002.md:17:42:	semicolon	";"	false	0	0.000000	(17, 42) -> (17, 43)
test/test_002.md:17:64:	ident	"Texts"	false	0	0.000000	(17, 64) -> (17, 69)
test/test_002.md:17:69:	dot	"."	false	0	0.000000	(17, 69) -> (17, 70)
test/test_002.md:17:70:	ident	"WriteLn"	false	0	0.000000	(17, 70) -> (17, 77)
test/test_002.md:17:77:	semicolon	";"	false	0	0.000000	(17, 77) -> (17, 78)
test/test_002.md:18:4:	ident	"Texts"	false	0	0.000000	(18, 4) -> (18, 9)
test/test_002.md:18:9:	dot	"."	false	0	0.000000	(18, 9) -> (18, 10)
test/test_002.md:18:10:	ident	"WriteReal"	false	0	0.000000	(18, 10) -> (18, 19)
test/test_002.md:18:19:	lparen	"("	false	0	0.000000	(18, 19) -> (18, 20)
test/test_002.md:18:20:	ident	"FLT"	false	0	0.000000	(18, 20) -> (18, 23)
test/test_002.md:18:23:	lparen	"("	false	0	0.000000	(18, 23) -> (18, 24)
test/test_002.md:18:24:	integer	"69"	false	69	0.000000	(18, 24) -> (18, 26)
test/test_002.md:18:26:	rparen	")"	false	0	0.000000	(18, 26) -> (18, 27)
test/test_002.md:18:27:	plus	"+"	false	0	0.000000	(18, 27) -> (18, 28)
test/test_002.md:18:28:	real	"0.420"	false	0	0.420000	(18, 28) -> (18, 33)
test/test_002.md:18:33:	rparen	")"	false	0	0.000000	(18, 33) -> (18, 34)
test/test_002.md:18:34:	semicolon	";"	false	0	0.000000	(18, 34) -> (18, 35)
test/test_002.md:18:64:	ident	"Texts"	false	0	0.000000	(18, 64) -> (18, 69)
test/test_002.md:18:69:	dot	"."	false	0	0.000000	(18, 69) -> (18, 70)
test/test_002.md:18:70:	ident	"WriteLn"	false	0	0.000000	(18, 70) -> (18, 77)
test/test_002.md:18:77:	semicolon	";"	false	0	0.000000	(18, 77) -> (18, 78)
test/test_002.md:19:4:	ident	"Texts"	false	0	0.000000	(19, 4) -> (19, 9)
test/test_002.md:19:9:	dot	"."	false	0	0.000000	(19, 9) -> (19, 10)
test/test_002.md:19:10:	ident	"WriteReal"	false	0	0.000000	(19, 10) -> (19, 19)
test/test_002.md:19:19:	lparen	"("	false	0	0.000000	(19, 19) -> (19, 20)
test/test_002.md:19:20:	real	"0.69E+2"	false	0	69.000000	(19, 20) -> (19, 27)
test/test_002.md:19:27:	plus	"+"	false	0	0.000000	(19, 27) -> (19, 28)
test/test_002.md:19:28:	real	"420.E-3"	false	0	0.420000	(19, 28) -> (19, 35)
test/test_002.md:19:35:	rparen	")"	false	0	0.000000	(19, 35) -> (19, 36)
test/test_002.md:19:36:	semicolon	";"	false	0	0.000000	(19, 36) -> (19, 37)
test/test_002.md:19:64:	ident	"Texts"	false	0	0.000000	(19, 64) -> (19, 69)
test/test_002.md:19:69:	dot	"."	false	0	0.000000	(19, 69) -> (19, 70)
test/test_002.md:19:70:	ident	"WriteLn"	false	0	0.000000	(19, 70) -> (19, 77)
test/test_002.md:19:77:	semicolon	";"	false	0	0.000000	(19, 77) -> (19, 78)
test/test_002.md:20:0:	end	"END"	false	0	0.000000	(20, 0) -> (20, 3)
test/test_002.md:20:4:	ident	"BinaryOperations"	false	0	0.000000	(20, 4) -> (20, 20)
test/test_002.md:20:20:	dot	"."	false	0	0.000000	(20, 20) -> (20, 21)
test/test_002.md:21:0:	eof	""	false	0	0.000000	(21, 0) -> (21, 0)
```
## AST
```scheme
(module "BinaryOperations"
  (stmts
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteInt")
        (number [i64] 33)
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteInt")
        (minus [i64]
          (number [i64] 0)
          (number [i64] 42)
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteInt")
        (plus [i64]
          (plus [i64]
            (number [i64] 34)
            (number [i64] 35)
          )
          (number [i64] 42000)
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteInt")
        (minus [i64]
          (number [i64] 66)
          (number [i64] 33)
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteInt")
        (plus [i64]
          (minus [i64]
            (number [i64] 0)
            (number [i64] 34)
          )
          (number [i64] 35)
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteInt")
        (plus [i64]
          (plus [i64]
            (number [i64] 34)
            (number [i64] 35)
          )
          (number [i64] 42000)
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteInt")
        (asterisk [i64]
          (plus [i64]
            (number [i64] 34)
            (number [i64] 35)
          )
          (number [i64] 33)
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteReal")
        (slash [f64]
          (call
            (procedure [f64] "FLT")
            (plus [i64]
              (number [i64] 34)
              (number [i64] 35)
            )
          )
          (call
            (procedure [f64] "FLT")
            (number [i64] 33)
          )
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteInt")
        (div [i64]
          (number [i64] 11)
          (number [i64] 3)
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteInt")
        (mod [i64]
          (number [i64] 11)
          (number [i64] 3)
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteReal")
        (slash [f64]
          (slash [f64]
            (call
              (procedure [f64] "FLT")
              (number [i64] 11)
            )
            (call
              (procedure [f64] "FLT")
              (number [i64] 3)
            )
          )
          (call
            (procedure [f64] "FLT")
            (number [i64] 2)
          )
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteReal")
        (plus [f64]
          (slash [f64]
            (call
              (procedure [f64] "FLT")
              (number [i64] 11)
            )
            (call
              (procedure [f64] "FLT")
              (number [i64] 3)
            )
          )
          (slash [f64]
            (call
              (procedure [f64] "FLT")
              (number [i64] 7)
            )
            (call
              (procedure [f64] "FLT")
              (number [i64] 2)
            )
          )
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteReal")
        (minus [f64]
          (slash [f64]
            (number [f64] 11.000000)
            (number [f64] 3.000000)
          )
          (slash [f64]
            (number [f64] 7.000000)
            (number [f64] 2.000000)
          )
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteReal")
        (asterisk [f64]
          (slash [f64]
            (call
              (procedure [f64] "FLT")
              (number [i64] 11)
            )
            (call
              (procedure [f64] "FLT")
              (number [i64] 3)
            )
          )
          (call
            (procedure [f64] "FLT")
            (number [i64] 3)
          )
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteReal")
        (plus [f64]
          (call
            (procedure [f64] "FLT")
            (number [i64] 69)
          )
          (number [f64] 0.420000)
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteReal")
        (plus [f64]
          (number [f64] 69.000000)
          (number [f64] 0.420000)
        )
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
@2 = global [3 x i8] c"%f\00"

declare i64 @puts(i8* %str)

declare i64 @rand()

declare i64 @sprintf(i8* %buf, i8* %format, ...)

declare i64 @printf(i8* %format, ...)

declare i8* @malloc(i64 %size)

declare i8* @free(i8* %ptr)

define i64 @main() {
entry:
	%0 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%1 = call i64 (i8*, ...) @printf(i8* %0, i64 33)
	%2 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%3 = call i64 @puts(i8* %2)
	%4 = sub i64 0, 42
	%5 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%6 = call i64 (i8*, ...) @printf(i8* %5, i64 %4)
	%7 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%8 = call i64 @puts(i8* %7)
	%9 = add i64 34, 35
	%10 = add i64 %9, 42000
	%11 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%12 = call i64 (i8*, ...) @printf(i8* %11, i64 %10)
	%13 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%14 = call i64 @puts(i8* %13)
	%15 = sub i64 66, 33
	%16 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%17 = call i64 (i8*, ...) @printf(i8* %16, i64 %15)
	%18 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%19 = call i64 @puts(i8* %18)
	%20 = sub i64 0, 34
	%21 = add i64 %20, 35
	%22 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%23 = call i64 (i8*, ...) @printf(i8* %22, i64 %21)
	%24 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%25 = call i64 @puts(i8* %24)
	%26 = add i64 34, 35
	%27 = add i64 %26, 42000
	%28 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%29 = call i64 (i8*, ...) @printf(i8* %28, i64 %27)
	%30 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%31 = call i64 @puts(i8* %30)
	%32 = add i64 34, 35
	%33 = mul i64 %32, 33
	%34 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%35 = call i64 (i8*, ...) @printf(i8* %34, i64 %33)
	%36 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%37 = call i64 @puts(i8* %36)
	%38 = add i64 34, 35
	%39 = sitofp i64 %38 to double
	%40 = sitofp i64 33 to double
	%41 = fdiv double %39, %40
	%42 = getelementptr [3 x i8], [3 x i8]* @2, i64 0, i64 0
	%43 = call i64 (i8*, ...) @printf(i8* %42, double %41)
	%44 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%45 = call i64 @puts(i8* %44)
	%46 = sdiv i64 11, 3
	%47 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%48 = call i64 (i8*, ...) @printf(i8* %47, i64 %46)
	%49 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%50 = call i64 @puts(i8* %49)
	%51 = srem i64 11, 3
	%52 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%53 = call i64 (i8*, ...) @printf(i8* %52, i64 %51)
	%54 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%55 = call i64 @puts(i8* %54)
	%56 = sitofp i64 11 to double
	%57 = sitofp i64 3 to double
	%58 = fdiv double %56, %57
	%59 = sitofp i64 2 to double
	%60 = fdiv double %58, %59
	%61 = getelementptr [3 x i8], [3 x i8]* @2, i64 0, i64 0
	%62 = call i64 (i8*, ...) @printf(i8* %61, double %60)
	%63 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%64 = call i64 @puts(i8* %63)
	%65 = sitofp i64 11 to double
	%66 = sitofp i64 3 to double
	%67 = fdiv double %65, %66
	%68 = sitofp i64 7 to double
	%69 = sitofp i64 2 to double
	%70 = fdiv double %68, %69
	%71 = fadd double %67, %70
	%72 = getelementptr [3 x i8], [3 x i8]* @2, i64 0, i64 0
	%73 = call i64 (i8*, ...) @printf(i8* %72, double %71)
	%74 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%75 = call i64 @puts(i8* %74)
	%76 = fdiv double 11.0, 3.0
	%77 = fdiv double 7.0, 2.0
	%78 = fsub double %76, %77
	%79 = getelementptr [3 x i8], [3 x i8]* @2, i64 0, i64 0
	%80 = call i64 (i8*, ...) @printf(i8* %79, double %78)
	%81 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%82 = call i64 @puts(i8* %81)
	%83 = sitofp i64 11 to double
	%84 = sitofp i64 3 to double
	%85 = fdiv double %83, %84
	%86 = sitofp i64 3 to double
	%87 = fmul double %85, %86
	%88 = getelementptr [3 x i8], [3 x i8]* @2, i64 0, i64 0
	%89 = call i64 (i8*, ...) @printf(i8* %88, double %87)
	%90 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%91 = call i64 @puts(i8* %90)
	%92 = sitofp i64 69 to double
	%93 = fadd double %92, 0x3FDAE147AE147AE1
	%94 = getelementptr [3 x i8], [3 x i8]* @2, i64 0, i64 0
	%95 = call i64 (i8*, ...) @printf(i8* %94, double %93)
	%96 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%97 = call i64 @puts(i8* %96)
	%98 = fadd double 69.0, 0x3FDAE147AE147AE1
	%99 = getelementptr [3 x i8], [3 x i8]* @2, i64 0, i64 0
	%100 = call i64 (i8*, ...) @printf(i8* %99, double %98)
	%101 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%102 = call i64 @puts(i8* %101)
	ret i64 0
}

```
## Run
```bash
33
-42
42069
33
1
42069
2277
2.090909
3
2
1.833333
7.166667
0.166667
11.000000
69.420000
69.420000
```
