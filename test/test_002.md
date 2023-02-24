# test/test_002.md
## Source
```pascal
MODULE BinaryOperations;

BEGIN
    C.print(+33);
    C.print(-42);
    C.print( 34 + 35 + 42000 );
    C.print(66-33);
    C.print( -34 + 35 );
    C.print( (34 + 35) + 42000 );
    C.print( (34 + 35) * 33 );
    C.print( FLT(34 + 35) / FLT(33) );
    C.print(11 DIV 3);
    C.print(11 MOD 3);
    C.print(FLT(11) / FLT(3) / FLT(2));
    C.print((FLT(11) / FLT(3)) + (FLT(7) / FLT(2)));
    C.print((11.0/3.0)-(7.0/2.0));
    C.print(FLT(11)/FLT(3)*FLT(3));
    C.print(FLT(69)+0.420);
    C.print(0.69E+2+420.E-3)
END BinaryOperations.
```
## Tokens
```tsv
test/test_002.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_002.md:1:8:	ident	"BinaryOperations"	false	0	0.000000	(1, 8) -> (1, 24)
test/test_002.md:1:24:	semicolon	";"	false	0	0.000000	(1, 24) -> (1, 25)
test/test_002.md:3:0:	begin	"BEGIN"	false	0	0.000000	(3, 0) -> (3, 5)
test/test_002.md:4:4:	ident	"C"	false	0	0.000000	(4, 4) -> (4, 5)
test/test_002.md:4:5:	dot	"."	false	0	0.000000	(4, 5) -> (4, 6)
test/test_002.md:4:6:	ident	"print"	false	0	0.000000	(4, 6) -> (4, 11)
test/test_002.md:4:11:	lparen	"("	false	0	0.000000	(4, 11) -> (4, 12)
test/test_002.md:4:12:	plus	"+"	false	0	0.000000	(4, 12) -> (4, 13)
test/test_002.md:4:13:	integer	"33"	false	33	0.000000	(4, 13) -> (4, 15)
test/test_002.md:4:15:	rparen	")"	false	0	0.000000	(4, 15) -> (4, 16)
test/test_002.md:4:16:	semicolon	";"	false	0	0.000000	(4, 16) -> (4, 17)
test/test_002.md:5:4:	ident	"C"	false	0	0.000000	(5, 4) -> (5, 5)
test/test_002.md:5:5:	dot	"."	false	0	0.000000	(5, 5) -> (5, 6)
test/test_002.md:5:6:	ident	"print"	false	0	0.000000	(5, 6) -> (5, 11)
test/test_002.md:5:11:	lparen	"("	false	0	0.000000	(5, 11) -> (5, 12)
test/test_002.md:5:12:	minus	"-"	false	0	0.000000	(5, 12) -> (5, 13)
test/test_002.md:5:13:	integer	"42"	false	42	0.000000	(5, 13) -> (5, 15)
test/test_002.md:5:15:	rparen	")"	false	0	0.000000	(5, 15) -> (5, 16)
test/test_002.md:5:16:	semicolon	";"	false	0	0.000000	(5, 16) -> (5, 17)
test/test_002.md:6:4:	ident	"C"	false	0	0.000000	(6, 4) -> (6, 5)
test/test_002.md:6:5:	dot	"."	false	0	0.000000	(6, 5) -> (6, 6)
test/test_002.md:6:6:	ident	"print"	false	0	0.000000	(6, 6) -> (6, 11)
test/test_002.md:6:11:	lparen	"("	false	0	0.000000	(6, 11) -> (6, 12)
test/test_002.md:6:13:	integer	"34"	false	34	0.000000	(6, 13) -> (6, 15)
test/test_002.md:6:16:	plus	"+"	false	0	0.000000	(6, 16) -> (6, 17)
test/test_002.md:6:18:	integer	"35"	false	35	0.000000	(6, 18) -> (6, 20)
test/test_002.md:6:21:	plus	"+"	false	0	0.000000	(6, 21) -> (6, 22)
test/test_002.md:6:23:	integer	"42000"	false	42000	0.000000	(6, 23) -> (6, 28)
test/test_002.md:6:29:	rparen	")"	false	0	0.000000	(6, 29) -> (6, 30)
test/test_002.md:6:30:	semicolon	";"	false	0	0.000000	(6, 30) -> (6, 31)
test/test_002.md:7:4:	ident	"C"	false	0	0.000000	(7, 4) -> (7, 5)
test/test_002.md:7:5:	dot	"."	false	0	0.000000	(7, 5) -> (7, 6)
test/test_002.md:7:6:	ident	"print"	false	0	0.000000	(7, 6) -> (7, 11)
test/test_002.md:7:11:	lparen	"("	false	0	0.000000	(7, 11) -> (7, 12)
test/test_002.md:7:12:	integer	"66"	false	66	0.000000	(7, 12) -> (7, 14)
test/test_002.md:7:14:	minus	"-"	false	0	0.000000	(7, 14) -> (7, 15)
test/test_002.md:7:15:	integer	"33"	false	33	0.000000	(7, 15) -> (7, 17)
test/test_002.md:7:17:	rparen	")"	false	0	0.000000	(7, 17) -> (7, 18)
test/test_002.md:7:18:	semicolon	";"	false	0	0.000000	(7, 18) -> (7, 19)
test/test_002.md:8:4:	ident	"C"	false	0	0.000000	(8, 4) -> (8, 5)
test/test_002.md:8:5:	dot	"."	false	0	0.000000	(8, 5) -> (8, 6)
test/test_002.md:8:6:	ident	"print"	false	0	0.000000	(8, 6) -> (8, 11)
test/test_002.md:8:11:	lparen	"("	false	0	0.000000	(8, 11) -> (8, 12)
test/test_002.md:8:13:	minus	"-"	false	0	0.000000	(8, 13) -> (8, 14)
test/test_002.md:8:14:	integer	"34"	false	34	0.000000	(8, 14) -> (8, 16)
test/test_002.md:8:17:	plus	"+"	false	0	0.000000	(8, 17) -> (8, 18)
test/test_002.md:8:19:	integer	"35"	false	35	0.000000	(8, 19) -> (8, 21)
test/test_002.md:8:22:	rparen	")"	false	0	0.000000	(8, 22) -> (8, 23)
test/test_002.md:8:23:	semicolon	";"	false	0	0.000000	(8, 23) -> (8, 24)
test/test_002.md:9:4:	ident	"C"	false	0	0.000000	(9, 4) -> (9, 5)
test/test_002.md:9:5:	dot	"."	false	0	0.000000	(9, 5) -> (9, 6)
test/test_002.md:9:6:	ident	"print"	false	0	0.000000	(9, 6) -> (9, 11)
test/test_002.md:9:11:	lparen	"("	false	0	0.000000	(9, 11) -> (9, 12)
test/test_002.md:9:13:	lparen	"("	false	0	0.000000	(9, 13) -> (9, 14)
test/test_002.md:9:14:	integer	"34"	false	34	0.000000	(9, 14) -> (9, 16)
test/test_002.md:9:17:	plus	"+"	false	0	0.000000	(9, 17) -> (9, 18)
test/test_002.md:9:19:	integer	"35"	false	35	0.000000	(9, 19) -> (9, 21)
test/test_002.md:9:21:	rparen	")"	false	0	0.000000	(9, 21) -> (9, 22)
test/test_002.md:9:23:	plus	"+"	false	0	0.000000	(9, 23) -> (9, 24)
test/test_002.md:9:25:	integer	"42000"	false	42000	0.000000	(9, 25) -> (9, 30)
test/test_002.md:9:31:	rparen	")"	false	0	0.000000	(9, 31) -> (9, 32)
test/test_002.md:9:32:	semicolon	";"	false	0	0.000000	(9, 32) -> (9, 33)
test/test_002.md:10:4:	ident	"C"	false	0	0.000000	(10, 4) -> (10, 5)
test/test_002.md:10:5:	dot	"."	false	0	0.000000	(10, 5) -> (10, 6)
test/test_002.md:10:6:	ident	"print"	false	0	0.000000	(10, 6) -> (10, 11)
test/test_002.md:10:11:	lparen	"("	false	0	0.000000	(10, 11) -> (10, 12)
test/test_002.md:10:13:	lparen	"("	false	0	0.000000	(10, 13) -> (10, 14)
test/test_002.md:10:14:	integer	"34"	false	34	0.000000	(10, 14) -> (10, 16)
test/test_002.md:10:17:	plus	"+"	false	0	0.000000	(10, 17) -> (10, 18)
test/test_002.md:10:19:	integer	"35"	false	35	0.000000	(10, 19) -> (10, 21)
test/test_002.md:10:21:	rparen	")"	false	0	0.000000	(10, 21) -> (10, 22)
test/test_002.md:10:23:	asterisk	"*"	false	0	0.000000	(10, 23) -> (10, 24)
test/test_002.md:10:25:	integer	"33"	false	33	0.000000	(10, 25) -> (10, 27)
test/test_002.md:10:28:	rparen	")"	false	0	0.000000	(10, 28) -> (10, 29)
test/test_002.md:10:29:	semicolon	";"	false	0	0.000000	(10, 29) -> (10, 30)
test/test_002.md:11:4:	ident	"C"	false	0	0.000000	(11, 4) -> (11, 5)
test/test_002.md:11:5:	dot	"."	false	0	0.000000	(11, 5) -> (11, 6)
test/test_002.md:11:6:	ident	"print"	false	0	0.000000	(11, 6) -> (11, 11)
test/test_002.md:11:11:	lparen	"("	false	0	0.000000	(11, 11) -> (11, 12)
test/test_002.md:11:13:	ident	"FLT"	false	0	0.000000	(11, 13) -> (11, 16)
test/test_002.md:11:16:	lparen	"("	false	0	0.000000	(11, 16) -> (11, 17)
test/test_002.md:11:17:	integer	"34"	false	34	0.000000	(11, 17) -> (11, 19)
test/test_002.md:11:20:	plus	"+"	false	0	0.000000	(11, 20) -> (11, 21)
test/test_002.md:11:22:	integer	"35"	false	35	0.000000	(11, 22) -> (11, 24)
test/test_002.md:11:24:	rparen	")"	false	0	0.000000	(11, 24) -> (11, 25)
test/test_002.md:11:26:	slash	"/"	false	0	0.000000	(11, 26) -> (11, 27)
test/test_002.md:11:28:	ident	"FLT"	false	0	0.000000	(11, 28) -> (11, 31)
test/test_002.md:11:31:	lparen	"("	false	0	0.000000	(11, 31) -> (11, 32)
test/test_002.md:11:32:	integer	"33"	false	33	0.000000	(11, 32) -> (11, 34)
test/test_002.md:11:34:	rparen	")"	false	0	0.000000	(11, 34) -> (11, 35)
test/test_002.md:11:36:	rparen	")"	false	0	0.000000	(11, 36) -> (11, 37)
test/test_002.md:11:37:	semicolon	";"	false	0	0.000000	(11, 37) -> (11, 38)
test/test_002.md:12:4:	ident	"C"	false	0	0.000000	(12, 4) -> (12, 5)
test/test_002.md:12:5:	dot	"."	false	0	0.000000	(12, 5) -> (12, 6)
test/test_002.md:12:6:	ident	"print"	false	0	0.000000	(12, 6) -> (12, 11)
test/test_002.md:12:11:	lparen	"("	false	0	0.000000	(12, 11) -> (12, 12)
test/test_002.md:12:12:	integer	"11"	false	11	0.000000	(12, 12) -> (12, 14)
test/test_002.md:12:15:	div	"DIV"	false	0	0.000000	(12, 15) -> (12, 18)
test/test_002.md:12:19:	integer	"3"	false	3	0.000000	(12, 19) -> (12, 20)
test/test_002.md:12:20:	rparen	")"	false	0	0.000000	(12, 20) -> (12, 21)
test/test_002.md:12:21:	semicolon	";"	false	0	0.000000	(12, 21) -> (12, 22)
test/test_002.md:13:4:	ident	"C"	false	0	0.000000	(13, 4) -> (13, 5)
test/test_002.md:13:5:	dot	"."	false	0	0.000000	(13, 5) -> (13, 6)
test/test_002.md:13:6:	ident	"print"	false	0	0.000000	(13, 6) -> (13, 11)
test/test_002.md:13:11:	lparen	"("	false	0	0.000000	(13, 11) -> (13, 12)
test/test_002.md:13:12:	integer	"11"	false	11	0.000000	(13, 12) -> (13, 14)
test/test_002.md:13:15:	mod	"MOD"	false	0	0.000000	(13, 15) -> (13, 18)
test/test_002.md:13:19:	integer	"3"	false	3	0.000000	(13, 19) -> (13, 20)
test/test_002.md:13:20:	rparen	")"	false	0	0.000000	(13, 20) -> (13, 21)
test/test_002.md:13:21:	semicolon	";"	false	0	0.000000	(13, 21) -> (13, 22)
test/test_002.md:14:4:	ident	"C"	false	0	0.000000	(14, 4) -> (14, 5)
test/test_002.md:14:5:	dot	"."	false	0	0.000000	(14, 5) -> (14, 6)
test/test_002.md:14:6:	ident	"print"	false	0	0.000000	(14, 6) -> (14, 11)
test/test_002.md:14:11:	lparen	"("	false	0	0.000000	(14, 11) -> (14, 12)
test/test_002.md:14:12:	ident	"FLT"	false	0	0.000000	(14, 12) -> (14, 15)
test/test_002.md:14:15:	lparen	"("	false	0	0.000000	(14, 15) -> (14, 16)
test/test_002.md:14:16:	integer	"11"	false	11	0.000000	(14, 16) -> (14, 18)
test/test_002.md:14:18:	rparen	")"	false	0	0.000000	(14, 18) -> (14, 19)
test/test_002.md:14:20:	slash	"/"	false	0	0.000000	(14, 20) -> (14, 21)
test/test_002.md:14:22:	ident	"FLT"	false	0	0.000000	(14, 22) -> (14, 25)
test/test_002.md:14:25:	lparen	"("	false	0	0.000000	(14, 25) -> (14, 26)
test/test_002.md:14:26:	integer	"3"	false	3	0.000000	(14, 26) -> (14, 27)
test/test_002.md:14:27:	rparen	")"	false	0	0.000000	(14, 27) -> (14, 28)
test/test_002.md:14:29:	slash	"/"	false	0	0.000000	(14, 29) -> (14, 30)
test/test_002.md:14:31:	ident	"FLT"	false	0	0.000000	(14, 31) -> (14, 34)
test/test_002.md:14:34:	lparen	"("	false	0	0.000000	(14, 34) -> (14, 35)
test/test_002.md:14:35:	integer	"2"	false	2	0.000000	(14, 35) -> (14, 36)
test/test_002.md:14:36:	rparen	")"	false	0	0.000000	(14, 36) -> (14, 37)
test/test_002.md:14:37:	rparen	")"	false	0	0.000000	(14, 37) -> (14, 38)
test/test_002.md:14:38:	semicolon	";"	false	0	0.000000	(14, 38) -> (14, 39)
test/test_002.md:15:4:	ident	"C"	false	0	0.000000	(15, 4) -> (15, 5)
test/test_002.md:15:5:	dot	"."	false	0	0.000000	(15, 5) -> (15, 6)
test/test_002.md:15:6:	ident	"print"	false	0	0.000000	(15, 6) -> (15, 11)
test/test_002.md:15:11:	lparen	"("	false	0	0.000000	(15, 11) -> (15, 12)
test/test_002.md:15:12:	lparen	"("	false	0	0.000000	(15, 12) -> (15, 13)
test/test_002.md:15:13:	ident	"FLT"	false	0	0.000000	(15, 13) -> (15, 16)
test/test_002.md:15:16:	lparen	"("	false	0	0.000000	(15, 16) -> (15, 17)
test/test_002.md:15:17:	integer	"11"	false	11	0.000000	(15, 17) -> (15, 19)
test/test_002.md:15:19:	rparen	")"	false	0	0.000000	(15, 19) -> (15, 20)
test/test_002.md:15:21:	slash	"/"	false	0	0.000000	(15, 21) -> (15, 22)
test/test_002.md:15:23:	ident	"FLT"	false	0	0.000000	(15, 23) -> (15, 26)
test/test_002.md:15:26:	lparen	"("	false	0	0.000000	(15, 26) -> (15, 27)
test/test_002.md:15:27:	integer	"3"	false	3	0.000000	(15, 27) -> (15, 28)
test/test_002.md:15:28:	rparen	")"	false	0	0.000000	(15, 28) -> (15, 29)
test/test_002.md:15:29:	rparen	")"	false	0	0.000000	(15, 29) -> (15, 30)
test/test_002.md:15:31:	plus	"+"	false	0	0.000000	(15, 31) -> (15, 32)
test/test_002.md:15:33:	lparen	"("	false	0	0.000000	(15, 33) -> (15, 34)
test/test_002.md:15:34:	ident	"FLT"	false	0	0.000000	(15, 34) -> (15, 37)
test/test_002.md:15:37:	lparen	"("	false	0	0.000000	(15, 37) -> (15, 38)
test/test_002.md:15:38:	integer	"7"	false	7	0.000000	(15, 38) -> (15, 39)
test/test_002.md:15:39:	rparen	")"	false	0	0.000000	(15, 39) -> (15, 40)
test/test_002.md:15:41:	slash	"/"	false	0	0.000000	(15, 41) -> (15, 42)
test/test_002.md:15:43:	ident	"FLT"	false	0	0.000000	(15, 43) -> (15, 46)
test/test_002.md:15:46:	lparen	"("	false	0	0.000000	(15, 46) -> (15, 47)
test/test_002.md:15:47:	integer	"2"	false	2	0.000000	(15, 47) -> (15, 48)
test/test_002.md:15:48:	rparen	")"	false	0	0.000000	(15, 48) -> (15, 49)
test/test_002.md:15:49:	rparen	")"	false	0	0.000000	(15, 49) -> (15, 50)
test/test_002.md:15:50:	rparen	")"	false	0	0.000000	(15, 50) -> (15, 51)
test/test_002.md:15:51:	semicolon	";"	false	0	0.000000	(15, 51) -> (15, 52)
test/test_002.md:16:4:	ident	"C"	false	0	0.000000	(16, 4) -> (16, 5)
test/test_002.md:16:5:	dot	"."	false	0	0.000000	(16, 5) -> (16, 6)
test/test_002.md:16:6:	ident	"print"	false	0	0.000000	(16, 6) -> (16, 11)
test/test_002.md:16:11:	lparen	"("	false	0	0.000000	(16, 11) -> (16, 12)
test/test_002.md:16:12:	lparen	"("	false	0	0.000000	(16, 12) -> (16, 13)
test/test_002.md:16:13:	real	"11.0"	false	0	11.000000	(16, 13) -> (16, 17)
test/test_002.md:16:17:	slash	"/"	false	0	0.000000	(16, 17) -> (16, 18)
test/test_002.md:16:18:	real	"3.0"	false	0	3.000000	(16, 18) -> (16, 21)
test/test_002.md:16:21:	rparen	")"	false	0	0.000000	(16, 21) -> (16, 22)
test/test_002.md:16:22:	minus	"-"	false	0	0.000000	(16, 22) -> (16, 23)
test/test_002.md:16:23:	lparen	"("	false	0	0.000000	(16, 23) -> (16, 24)
test/test_002.md:16:24:	real	"7.0"	false	0	7.000000	(16, 24) -> (16, 27)
test/test_002.md:16:27:	slash	"/"	false	0	0.000000	(16, 27) -> (16, 28)
test/test_002.md:16:28:	real	"2.0"	false	0	2.000000	(16, 28) -> (16, 31)
test/test_002.md:16:31:	rparen	")"	false	0	0.000000	(16, 31) -> (16, 32)
test/test_002.md:16:32:	rparen	")"	false	0	0.000000	(16, 32) -> (16, 33)
test/test_002.md:16:33:	semicolon	";"	false	0	0.000000	(16, 33) -> (16, 34)
test/test_002.md:17:4:	ident	"C"	false	0	0.000000	(17, 4) -> (17, 5)
test/test_002.md:17:5:	dot	"."	false	0	0.000000	(17, 5) -> (17, 6)
test/test_002.md:17:6:	ident	"print"	false	0	0.000000	(17, 6) -> (17, 11)
test/test_002.md:17:11:	lparen	"("	false	0	0.000000	(17, 11) -> (17, 12)
test/test_002.md:17:12:	ident	"FLT"	false	0	0.000000	(17, 12) -> (17, 15)
test/test_002.md:17:15:	lparen	"("	false	0	0.000000	(17, 15) -> (17, 16)
test/test_002.md:17:16:	integer	"11"	false	11	0.000000	(17, 16) -> (17, 18)
test/test_002.md:17:18:	rparen	")"	false	0	0.000000	(17, 18) -> (17, 19)
test/test_002.md:17:19:	slash	"/"	false	0	0.000000	(17, 19) -> (17, 20)
test/test_002.md:17:20:	ident	"FLT"	false	0	0.000000	(17, 20) -> (17, 23)
test/test_002.md:17:23:	lparen	"("	false	0	0.000000	(17, 23) -> (17, 24)
test/test_002.md:17:24:	integer	"3"	false	3	0.000000	(17, 24) -> (17, 25)
test/test_002.md:17:25:	rparen	")"	false	0	0.000000	(17, 25) -> (17, 26)
test/test_002.md:17:26:	asterisk	"*"	false	0	0.000000	(17, 26) -> (17, 27)
test/test_002.md:17:27:	ident	"FLT"	false	0	0.000000	(17, 27) -> (17, 30)
test/test_002.md:17:30:	lparen	"("	false	0	0.000000	(17, 30) -> (17, 31)
test/test_002.md:17:31:	integer	"3"	false	3	0.000000	(17, 31) -> (17, 32)
test/test_002.md:17:32:	rparen	")"	false	0	0.000000	(17, 32) -> (17, 33)
test/test_002.md:17:33:	rparen	")"	false	0	0.000000	(17, 33) -> (17, 34)
test/test_002.md:17:34:	semicolon	";"	false	0	0.000000	(17, 34) -> (17, 35)
test/test_002.md:18:4:	ident	"C"	false	0	0.000000	(18, 4) -> (18, 5)
test/test_002.md:18:5:	dot	"."	false	0	0.000000	(18, 5) -> (18, 6)
test/test_002.md:18:6:	ident	"print"	false	0	0.000000	(18, 6) -> (18, 11)
test/test_002.md:18:11:	lparen	"("	false	0	0.000000	(18, 11) -> (18, 12)
test/test_002.md:18:12:	ident	"FLT"	false	0	0.000000	(18, 12) -> (18, 15)
test/test_002.md:18:15:	lparen	"("	false	0	0.000000	(18, 15) -> (18, 16)
test/test_002.md:18:16:	integer	"69"	false	69	0.000000	(18, 16) -> (18, 18)
test/test_002.md:18:18:	rparen	")"	false	0	0.000000	(18, 18) -> (18, 19)
test/test_002.md:18:19:	plus	"+"	false	0	0.000000	(18, 19) -> (18, 20)
test/test_002.md:18:20:	real	"0.420"	false	0	0.420000	(18, 20) -> (18, 25)
test/test_002.md:18:25:	rparen	")"	false	0	0.000000	(18, 25) -> (18, 26)
test/test_002.md:18:26:	semicolon	";"	false	0	0.000000	(18, 26) -> (18, 27)
test/test_002.md:19:4:	ident	"C"	false	0	0.000000	(19, 4) -> (19, 5)
test/test_002.md:19:5:	dot	"."	false	0	0.000000	(19, 5) -> (19, 6)
test/test_002.md:19:6:	ident	"print"	false	0	0.000000	(19, 6) -> (19, 11)
test/test_002.md:19:11:	lparen	"("	false	0	0.000000	(19, 11) -> (19, 12)
test/test_002.md:19:12:	real	"0.69E+2"	false	0	69.000000	(19, 12) -> (19, 19)
test/test_002.md:19:19:	plus	"+"	false	0	0.000000	(19, 19) -> (19, 20)
test/test_002.md:19:20:	real	"420.E-3"	false	0	0.420000	(19, 20) -> (19, 27)
test/test_002.md:19:27:	rparen	")"	false	0	0.000000	(19, 27) -> (19, 28)
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
        (procedure [void] "C.print")
        (number [i64] 33)
      )
    )
    (expr2stmt
      (call
        (procedure [void] "C.print")
        (minus [i64]
          (number [i64] 0)
          (number [i64] 42)
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "C.print")
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
        (procedure [void] "C.print")
        (minus [i64]
          (number [i64] 66)
          (number [i64] 33)
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "C.print")
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
        (procedure [void] "C.print")
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
        (procedure [void] "C.print")
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
        (procedure [void] "C.print")
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
        (procedure [void] "C.print")
        (div [i64]
          (number [i64] 11)
          (number [i64] 3)
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "C.print")
        (mod [i64]
          (number [i64] 11)
          (number [i64] 3)
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "C.print")
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
        (procedure [void] "C.print")
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
        (procedure [void] "C.print")
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
        (procedure [void] "C.print")
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
        (procedure [void] "C.print")
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
        (procedure [void] "C.print")
        (plus [f64]
          (number [f64] 69.000000)
          (number [f64] 0.420000)
        )
      )
    )
  )
)
```
## IR
```llvm
@0 = global [4 x i8] c"%d\0A\00"
@1 = global [4 x i8] c"%f\0A\00"

declare i64 @puts(i8* %str)

declare i64 @rand()

declare i64 @sprintf(i8* %buf, i8* %format, ...)

declare i64 @printf(i8* %format, ...)

declare i8* @malloc(i64 %size)

declare i8* @free(i8* %ptr)

define i64 @main() {
entry:
	%0 = getelementptr [4 x i8], [4 x i8]* @0, i64 0, i64 0
	%1 = call i64 (i8*, ...) @printf(i8* %0, i64 33)
	%2 = sub i64 0, 42
	%3 = getelementptr [4 x i8], [4 x i8]* @0, i64 0, i64 0
	%4 = call i64 (i8*, ...) @printf(i8* %3, i64 %2)
	%5 = add i64 34, 35
	%6 = add i64 %5, 42000
	%7 = getelementptr [4 x i8], [4 x i8]* @0, i64 0, i64 0
	%8 = call i64 (i8*, ...) @printf(i8* %7, i64 %6)
	%9 = sub i64 66, 33
	%10 = getelementptr [4 x i8], [4 x i8]* @0, i64 0, i64 0
	%11 = call i64 (i8*, ...) @printf(i8* %10, i64 %9)
	%12 = sub i64 0, 34
	%13 = add i64 %12, 35
	%14 = getelementptr [4 x i8], [4 x i8]* @0, i64 0, i64 0
	%15 = call i64 (i8*, ...) @printf(i8* %14, i64 %13)
	%16 = add i64 34, 35
	%17 = add i64 %16, 42000
	%18 = getelementptr [4 x i8], [4 x i8]* @0, i64 0, i64 0
	%19 = call i64 (i8*, ...) @printf(i8* %18, i64 %17)
	%20 = add i64 34, 35
	%21 = mul i64 %20, 33
	%22 = getelementptr [4 x i8], [4 x i8]* @0, i64 0, i64 0
	%23 = call i64 (i8*, ...) @printf(i8* %22, i64 %21)
	%24 = add i64 34, 35
	%25 = sitofp i64 %24 to double
	%26 = sitofp i64 33 to double
	%27 = fdiv double %25, %26
	%28 = getelementptr [4 x i8], [4 x i8]* @1, i64 0, i64 0
	%29 = call i64 (i8*, ...) @printf(i8* %28, double %27)
	%30 = sdiv i64 11, 3
	%31 = getelementptr [4 x i8], [4 x i8]* @0, i64 0, i64 0
	%32 = call i64 (i8*, ...) @printf(i8* %31, i64 %30)
	%33 = srem i64 11, 3
	%34 = getelementptr [4 x i8], [4 x i8]* @0, i64 0, i64 0
	%35 = call i64 (i8*, ...) @printf(i8* %34, i64 %33)
	%36 = sitofp i64 11 to double
	%37 = sitofp i64 3 to double
	%38 = fdiv double %36, %37
	%39 = sitofp i64 2 to double
	%40 = fdiv double %38, %39
	%41 = getelementptr [4 x i8], [4 x i8]* @1, i64 0, i64 0
	%42 = call i64 (i8*, ...) @printf(i8* %41, double %40)
	%43 = sitofp i64 11 to double
	%44 = sitofp i64 3 to double
	%45 = fdiv double %43, %44
	%46 = sitofp i64 7 to double
	%47 = sitofp i64 2 to double
	%48 = fdiv double %46, %47
	%49 = fadd double %45, %48
	%50 = getelementptr [4 x i8], [4 x i8]* @1, i64 0, i64 0
	%51 = call i64 (i8*, ...) @printf(i8* %50, double %49)
	%52 = fdiv double 11.0, 3.0
	%53 = fdiv double 7.0, 2.0
	%54 = fsub double %52, %53
	%55 = getelementptr [4 x i8], [4 x i8]* @1, i64 0, i64 0
	%56 = call i64 (i8*, ...) @printf(i8* %55, double %54)
	%57 = sitofp i64 11 to double
	%58 = sitofp i64 3 to double
	%59 = fdiv double %57, %58
	%60 = sitofp i64 3 to double
	%61 = fmul double %59, %60
	%62 = getelementptr [4 x i8], [4 x i8]* @1, i64 0, i64 0
	%63 = call i64 (i8*, ...) @printf(i8* %62, double %61)
	%64 = sitofp i64 69 to double
	%65 = fadd double %64, 0x3FDAE147AE147AE1
	%66 = getelementptr [4 x i8], [4 x i8]* @1, i64 0, i64 0
	%67 = call i64 (i8*, ...) @printf(i8* %66, double %65)
	%68 = fadd double 69.0, 0x3FDAE147AE147AE1
	%69 = getelementptr [4 x i8], [4 x i8]* @1, i64 0, i64 0
	%70 = call i64 (i8*, ...) @printf(i8* %69, double %68)
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
