# test/test_002.md
## Source
```pascal
MODULE BinaryOperations;

BEGIN
    print(+33);
    print(-42);
    print( 34 + 35 + 42000 );
    print(66-33);
    print( -34 + 35 );
    print( (34 + 35) + 42000 );
    print( (34 + 35) * 33 );
    print( FLT(34 + 35) / FLT(33) );
    print(11 DIV 3);
    print(11 MOD 3);
    print(FLT(11) / FLT(3) / FLT(2));
    print((FLT(11) / FLT(3)) + (FLT(7) / FLT(2)));
    print((11.0/3.0)-(7.0/2.0));
    print(FLT(11)/FLT(3)*FLT(3));
    print(FLT(69)+0.420);
    print(0.69E+2+420.E-3)
END BinaryOperations.
```
## Tokens
```tsv
test/test_002.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_002.md:1:8:	ident	"BinaryOperations"	false	0	0.000000	(1, 8) -> (1, 24)
test/test_002.md:1:24:	semicolon	";"	false	0	0.000000	(1, 24) -> (1, 25)
test/test_002.md:3:0:	begin	"BEGIN"	false	0	0.000000	(3, 0) -> (3, 5)
test/test_002.md:4:4:	ident	"print"	false	0	0.000000	(4, 4) -> (4, 9)
test/test_002.md:4:9:	lparen	"("	false	0	0.000000	(4, 9) -> (4, 10)
test/test_002.md:4:10:	plus	"+"	false	0	0.000000	(4, 10) -> (4, 11)
test/test_002.md:4:11:	integer	"33"	false	33	0.000000	(4, 11) -> (4, 13)
test/test_002.md:4:13:	rparen	")"	false	0	0.000000	(4, 13) -> (4, 14)
test/test_002.md:4:14:	semicolon	";"	false	0	0.000000	(4, 14) -> (4, 15)
test/test_002.md:5:4:	ident	"print"	false	0	0.000000	(5, 4) -> (5, 9)
test/test_002.md:5:9:	lparen	"("	false	0	0.000000	(5, 9) -> (5, 10)
test/test_002.md:5:10:	minus	"-"	false	0	0.000000	(5, 10) -> (5, 11)
test/test_002.md:5:11:	integer	"42"	false	42	0.000000	(5, 11) -> (5, 13)
test/test_002.md:5:13:	rparen	")"	false	0	0.000000	(5, 13) -> (5, 14)
test/test_002.md:5:14:	semicolon	";"	false	0	0.000000	(5, 14) -> (5, 15)
test/test_002.md:6:4:	ident	"print"	false	0	0.000000	(6, 4) -> (6, 9)
test/test_002.md:6:9:	lparen	"("	false	0	0.000000	(6, 9) -> (6, 10)
test/test_002.md:6:11:	integer	"34"	false	34	0.000000	(6, 11) -> (6, 13)
test/test_002.md:6:14:	plus	"+"	false	0	0.000000	(6, 14) -> (6, 15)
test/test_002.md:6:16:	integer	"35"	false	35	0.000000	(6, 16) -> (6, 18)
test/test_002.md:6:19:	plus	"+"	false	0	0.000000	(6, 19) -> (6, 20)
test/test_002.md:6:21:	integer	"42000"	false	42000	0.000000	(6, 21) -> (6, 26)
test/test_002.md:6:27:	rparen	")"	false	0	0.000000	(6, 27) -> (6, 28)
test/test_002.md:6:28:	semicolon	";"	false	0	0.000000	(6, 28) -> (6, 29)
test/test_002.md:7:4:	ident	"print"	false	0	0.000000	(7, 4) -> (7, 9)
test/test_002.md:7:9:	lparen	"("	false	0	0.000000	(7, 9) -> (7, 10)
test/test_002.md:7:10:	integer	"66"	false	66	0.000000	(7, 10) -> (7, 12)
test/test_002.md:7:12:	minus	"-"	false	0	0.000000	(7, 12) -> (7, 13)
test/test_002.md:7:13:	integer	"33"	false	33	0.000000	(7, 13) -> (7, 15)
test/test_002.md:7:15:	rparen	")"	false	0	0.000000	(7, 15) -> (7, 16)
test/test_002.md:7:16:	semicolon	";"	false	0	0.000000	(7, 16) -> (7, 17)
test/test_002.md:8:4:	ident	"print"	false	0	0.000000	(8, 4) -> (8, 9)
test/test_002.md:8:9:	lparen	"("	false	0	0.000000	(8, 9) -> (8, 10)
test/test_002.md:8:11:	minus	"-"	false	0	0.000000	(8, 11) -> (8, 12)
test/test_002.md:8:12:	integer	"34"	false	34	0.000000	(8, 12) -> (8, 14)
test/test_002.md:8:15:	plus	"+"	false	0	0.000000	(8, 15) -> (8, 16)
test/test_002.md:8:17:	integer	"35"	false	35	0.000000	(8, 17) -> (8, 19)
test/test_002.md:8:20:	rparen	")"	false	0	0.000000	(8, 20) -> (8, 21)
test/test_002.md:8:21:	semicolon	";"	false	0	0.000000	(8, 21) -> (8, 22)
test/test_002.md:9:4:	ident	"print"	false	0	0.000000	(9, 4) -> (9, 9)
test/test_002.md:9:9:	lparen	"("	false	0	0.000000	(9, 9) -> (9, 10)
test/test_002.md:9:11:	lparen	"("	false	0	0.000000	(9, 11) -> (9, 12)
test/test_002.md:9:12:	integer	"34"	false	34	0.000000	(9, 12) -> (9, 14)
test/test_002.md:9:15:	plus	"+"	false	0	0.000000	(9, 15) -> (9, 16)
test/test_002.md:9:17:	integer	"35"	false	35	0.000000	(9, 17) -> (9, 19)
test/test_002.md:9:19:	rparen	")"	false	0	0.000000	(9, 19) -> (9, 20)
test/test_002.md:9:21:	plus	"+"	false	0	0.000000	(9, 21) -> (9, 22)
test/test_002.md:9:23:	integer	"42000"	false	42000	0.000000	(9, 23) -> (9, 28)
test/test_002.md:9:29:	rparen	")"	false	0	0.000000	(9, 29) -> (9, 30)
test/test_002.md:9:30:	semicolon	";"	false	0	0.000000	(9, 30) -> (9, 31)
test/test_002.md:10:4:	ident	"print"	false	0	0.000000	(10, 4) -> (10, 9)
test/test_002.md:10:9:	lparen	"("	false	0	0.000000	(10, 9) -> (10, 10)
test/test_002.md:10:11:	lparen	"("	false	0	0.000000	(10, 11) -> (10, 12)
test/test_002.md:10:12:	integer	"34"	false	34	0.000000	(10, 12) -> (10, 14)
test/test_002.md:10:15:	plus	"+"	false	0	0.000000	(10, 15) -> (10, 16)
test/test_002.md:10:17:	integer	"35"	false	35	0.000000	(10, 17) -> (10, 19)
test/test_002.md:10:19:	rparen	")"	false	0	0.000000	(10, 19) -> (10, 20)
test/test_002.md:10:21:	asterisk	"*"	false	0	0.000000	(10, 21) -> (10, 22)
test/test_002.md:10:23:	integer	"33"	false	33	0.000000	(10, 23) -> (10, 25)
test/test_002.md:10:26:	rparen	")"	false	0	0.000000	(10, 26) -> (10, 27)
test/test_002.md:10:27:	semicolon	";"	false	0	0.000000	(10, 27) -> (10, 28)
test/test_002.md:11:4:	ident	"print"	false	0	0.000000	(11, 4) -> (11, 9)
test/test_002.md:11:9:	lparen	"("	false	0	0.000000	(11, 9) -> (11, 10)
test/test_002.md:11:11:	ident	"FLT"	false	0	0.000000	(11, 11) -> (11, 14)
test/test_002.md:11:14:	lparen	"("	false	0	0.000000	(11, 14) -> (11, 15)
test/test_002.md:11:15:	integer	"34"	false	34	0.000000	(11, 15) -> (11, 17)
test/test_002.md:11:18:	plus	"+"	false	0	0.000000	(11, 18) -> (11, 19)
test/test_002.md:11:20:	integer	"35"	false	35	0.000000	(11, 20) -> (11, 22)
test/test_002.md:11:22:	rparen	")"	false	0	0.000000	(11, 22) -> (11, 23)
test/test_002.md:11:24:	slash	"/"	false	0	0.000000	(11, 24) -> (11, 25)
test/test_002.md:11:26:	ident	"FLT"	false	0	0.000000	(11, 26) -> (11, 29)
test/test_002.md:11:29:	lparen	"("	false	0	0.000000	(11, 29) -> (11, 30)
test/test_002.md:11:30:	integer	"33"	false	33	0.000000	(11, 30) -> (11, 32)
test/test_002.md:11:32:	rparen	")"	false	0	0.000000	(11, 32) -> (11, 33)
test/test_002.md:11:34:	rparen	")"	false	0	0.000000	(11, 34) -> (11, 35)
test/test_002.md:11:35:	semicolon	";"	false	0	0.000000	(11, 35) -> (11, 36)
test/test_002.md:12:4:	ident	"print"	false	0	0.000000	(12, 4) -> (12, 9)
test/test_002.md:12:9:	lparen	"("	false	0	0.000000	(12, 9) -> (12, 10)
test/test_002.md:12:10:	integer	"11"	false	11	0.000000	(12, 10) -> (12, 12)
test/test_002.md:12:13:	div	"DIV"	false	0	0.000000	(12, 13) -> (12, 16)
test/test_002.md:12:17:	integer	"3"	false	3	0.000000	(12, 17) -> (12, 18)
test/test_002.md:12:18:	rparen	")"	false	0	0.000000	(12, 18) -> (12, 19)
test/test_002.md:12:19:	semicolon	";"	false	0	0.000000	(12, 19) -> (12, 20)
test/test_002.md:13:4:	ident	"print"	false	0	0.000000	(13, 4) -> (13, 9)
test/test_002.md:13:9:	lparen	"("	false	0	0.000000	(13, 9) -> (13, 10)
test/test_002.md:13:10:	integer	"11"	false	11	0.000000	(13, 10) -> (13, 12)
test/test_002.md:13:13:	mod	"MOD"	false	0	0.000000	(13, 13) -> (13, 16)
test/test_002.md:13:17:	integer	"3"	false	3	0.000000	(13, 17) -> (13, 18)
test/test_002.md:13:18:	rparen	")"	false	0	0.000000	(13, 18) -> (13, 19)
test/test_002.md:13:19:	semicolon	";"	false	0	0.000000	(13, 19) -> (13, 20)
test/test_002.md:14:4:	ident	"print"	false	0	0.000000	(14, 4) -> (14, 9)
test/test_002.md:14:9:	lparen	"("	false	0	0.000000	(14, 9) -> (14, 10)
test/test_002.md:14:10:	ident	"FLT"	false	0	0.000000	(14, 10) -> (14, 13)
test/test_002.md:14:13:	lparen	"("	false	0	0.000000	(14, 13) -> (14, 14)
test/test_002.md:14:14:	integer	"11"	false	11	0.000000	(14, 14) -> (14, 16)
test/test_002.md:14:16:	rparen	")"	false	0	0.000000	(14, 16) -> (14, 17)
test/test_002.md:14:18:	slash	"/"	false	0	0.000000	(14, 18) -> (14, 19)
test/test_002.md:14:20:	ident	"FLT"	false	0	0.000000	(14, 20) -> (14, 23)
test/test_002.md:14:23:	lparen	"("	false	0	0.000000	(14, 23) -> (14, 24)
test/test_002.md:14:24:	integer	"3"	false	3	0.000000	(14, 24) -> (14, 25)
test/test_002.md:14:25:	rparen	")"	false	0	0.000000	(14, 25) -> (14, 26)
test/test_002.md:14:27:	slash	"/"	false	0	0.000000	(14, 27) -> (14, 28)
test/test_002.md:14:29:	ident	"FLT"	false	0	0.000000	(14, 29) -> (14, 32)
test/test_002.md:14:32:	lparen	"("	false	0	0.000000	(14, 32) -> (14, 33)
test/test_002.md:14:33:	integer	"2"	false	2	0.000000	(14, 33) -> (14, 34)
test/test_002.md:14:34:	rparen	")"	false	0	0.000000	(14, 34) -> (14, 35)
test/test_002.md:14:35:	rparen	")"	false	0	0.000000	(14, 35) -> (14, 36)
test/test_002.md:14:36:	semicolon	";"	false	0	0.000000	(14, 36) -> (14, 37)
test/test_002.md:15:4:	ident	"print"	false	0	0.000000	(15, 4) -> (15, 9)
test/test_002.md:15:9:	lparen	"("	false	0	0.000000	(15, 9) -> (15, 10)
test/test_002.md:15:10:	lparen	"("	false	0	0.000000	(15, 10) -> (15, 11)
test/test_002.md:15:11:	ident	"FLT"	false	0	0.000000	(15, 11) -> (15, 14)
test/test_002.md:15:14:	lparen	"("	false	0	0.000000	(15, 14) -> (15, 15)
test/test_002.md:15:15:	integer	"11"	false	11	0.000000	(15, 15) -> (15, 17)
test/test_002.md:15:17:	rparen	")"	false	0	0.000000	(15, 17) -> (15, 18)
test/test_002.md:15:19:	slash	"/"	false	0	0.000000	(15, 19) -> (15, 20)
test/test_002.md:15:21:	ident	"FLT"	false	0	0.000000	(15, 21) -> (15, 24)
test/test_002.md:15:24:	lparen	"("	false	0	0.000000	(15, 24) -> (15, 25)
test/test_002.md:15:25:	integer	"3"	false	3	0.000000	(15, 25) -> (15, 26)
test/test_002.md:15:26:	rparen	")"	false	0	0.000000	(15, 26) -> (15, 27)
test/test_002.md:15:27:	rparen	")"	false	0	0.000000	(15, 27) -> (15, 28)
test/test_002.md:15:29:	plus	"+"	false	0	0.000000	(15, 29) -> (15, 30)
test/test_002.md:15:31:	lparen	"("	false	0	0.000000	(15, 31) -> (15, 32)
test/test_002.md:15:32:	ident	"FLT"	false	0	0.000000	(15, 32) -> (15, 35)
test/test_002.md:15:35:	lparen	"("	false	0	0.000000	(15, 35) -> (15, 36)
test/test_002.md:15:36:	integer	"7"	false	7	0.000000	(15, 36) -> (15, 37)
test/test_002.md:15:37:	rparen	")"	false	0	0.000000	(15, 37) -> (15, 38)
test/test_002.md:15:39:	slash	"/"	false	0	0.000000	(15, 39) -> (15, 40)
test/test_002.md:15:41:	ident	"FLT"	false	0	0.000000	(15, 41) -> (15, 44)
test/test_002.md:15:44:	lparen	"("	false	0	0.000000	(15, 44) -> (15, 45)
test/test_002.md:15:45:	integer	"2"	false	2	0.000000	(15, 45) -> (15, 46)
test/test_002.md:15:46:	rparen	")"	false	0	0.000000	(15, 46) -> (15, 47)
test/test_002.md:15:47:	rparen	")"	false	0	0.000000	(15, 47) -> (15, 48)
test/test_002.md:15:48:	rparen	")"	false	0	0.000000	(15, 48) -> (15, 49)
test/test_002.md:15:49:	semicolon	";"	false	0	0.000000	(15, 49) -> (15, 50)
test/test_002.md:16:4:	ident	"print"	false	0	0.000000	(16, 4) -> (16, 9)
test/test_002.md:16:9:	lparen	"("	false	0	0.000000	(16, 9) -> (16, 10)
test/test_002.md:16:10:	lparen	"("	false	0	0.000000	(16, 10) -> (16, 11)
test/test_002.md:16:11:	real	"11.0"	false	0	11.000000	(16, 11) -> (16, 15)
test/test_002.md:16:15:	slash	"/"	false	0	0.000000	(16, 15) -> (16, 16)
test/test_002.md:16:16:	real	"3.0"	false	0	3.000000	(16, 16) -> (16, 19)
test/test_002.md:16:19:	rparen	")"	false	0	0.000000	(16, 19) -> (16, 20)
test/test_002.md:16:20:	minus	"-"	false	0	0.000000	(16, 20) -> (16, 21)
test/test_002.md:16:21:	lparen	"("	false	0	0.000000	(16, 21) -> (16, 22)
test/test_002.md:16:22:	real	"7.0"	false	0	7.000000	(16, 22) -> (16, 25)
test/test_002.md:16:25:	slash	"/"	false	0	0.000000	(16, 25) -> (16, 26)
test/test_002.md:16:26:	real	"2.0"	false	0	2.000000	(16, 26) -> (16, 29)
test/test_002.md:16:29:	rparen	")"	false	0	0.000000	(16, 29) -> (16, 30)
test/test_002.md:16:30:	rparen	")"	false	0	0.000000	(16, 30) -> (16, 31)
test/test_002.md:16:31:	semicolon	";"	false	0	0.000000	(16, 31) -> (16, 32)
test/test_002.md:17:4:	ident	"print"	false	0	0.000000	(17, 4) -> (17, 9)
test/test_002.md:17:9:	lparen	"("	false	0	0.000000	(17, 9) -> (17, 10)
test/test_002.md:17:10:	ident	"FLT"	false	0	0.000000	(17, 10) -> (17, 13)
test/test_002.md:17:13:	lparen	"("	false	0	0.000000	(17, 13) -> (17, 14)
test/test_002.md:17:14:	integer	"11"	false	11	0.000000	(17, 14) -> (17, 16)
test/test_002.md:17:16:	rparen	")"	false	0	0.000000	(17, 16) -> (17, 17)
test/test_002.md:17:17:	slash	"/"	false	0	0.000000	(17, 17) -> (17, 18)
test/test_002.md:17:18:	ident	"FLT"	false	0	0.000000	(17, 18) -> (17, 21)
test/test_002.md:17:21:	lparen	"("	false	0	0.000000	(17, 21) -> (17, 22)
test/test_002.md:17:22:	integer	"3"	false	3	0.000000	(17, 22) -> (17, 23)
test/test_002.md:17:23:	rparen	")"	false	0	0.000000	(17, 23) -> (17, 24)
test/test_002.md:17:24:	asterisk	"*"	false	0	0.000000	(17, 24) -> (17, 25)
test/test_002.md:17:25:	ident	"FLT"	false	0	0.000000	(17, 25) -> (17, 28)
test/test_002.md:17:28:	lparen	"("	false	0	0.000000	(17, 28) -> (17, 29)
test/test_002.md:17:29:	integer	"3"	false	3	0.000000	(17, 29) -> (17, 30)
test/test_002.md:17:30:	rparen	")"	false	0	0.000000	(17, 30) -> (17, 31)
test/test_002.md:17:31:	rparen	")"	false	0	0.000000	(17, 31) -> (17, 32)
test/test_002.md:17:32:	semicolon	";"	false	0	0.000000	(17, 32) -> (17, 33)
test/test_002.md:18:4:	ident	"print"	false	0	0.000000	(18, 4) -> (18, 9)
test/test_002.md:18:9:	lparen	"("	false	0	0.000000	(18, 9) -> (18, 10)
test/test_002.md:18:10:	ident	"FLT"	false	0	0.000000	(18, 10) -> (18, 13)
test/test_002.md:18:13:	lparen	"("	false	0	0.000000	(18, 13) -> (18, 14)
test/test_002.md:18:14:	integer	"69"	false	69	0.000000	(18, 14) -> (18, 16)
test/test_002.md:18:16:	rparen	")"	false	0	0.000000	(18, 16) -> (18, 17)
test/test_002.md:18:17:	plus	"+"	false	0	0.000000	(18, 17) -> (18, 18)
test/test_002.md:18:18:	real	"0.420"	false	0	0.420000	(18, 18) -> (18, 23)
test/test_002.md:18:23:	rparen	")"	false	0	0.000000	(18, 23) -> (18, 24)
test/test_002.md:18:24:	semicolon	";"	false	0	0.000000	(18, 24) -> (18, 25)
test/test_002.md:19:4:	ident	"print"	false	0	0.000000	(19, 4) -> (19, 9)
test/test_002.md:19:9:	lparen	"("	false	0	0.000000	(19, 9) -> (19, 10)
test/test_002.md:19:10:	real	"0.69E+2"	false	0	69.000000	(19, 10) -> (19, 17)
test/test_002.md:19:17:	plus	"+"	false	0	0.000000	(19, 17) -> (19, 18)
test/test_002.md:19:18:	real	"420.E-3"	false	0	0.420000	(19, 18) -> (19, 25)
test/test_002.md:19:25:	rparen	")"	false	0	0.000000	(19, 25) -> (19, 26)
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
        (procedure [void] "print")
        (number [i64] 33)
      )
    )
    (expr2stmt
      (call
        (procedure [void] "print")
        (minus [i64]
          (number [i64] 0)
          (number [i64] 42)
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "print")
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
        (procedure [void] "print")
        (minus [i64]
          (number [i64] 66)
          (number [i64] 33)
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "print")
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
        (procedure [void] "print")
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
        (procedure [void] "print")
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
        (procedure [void] "print")
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
        (procedure [void] "print")
        (div [i64]
          (number [i64] 11)
          (number [i64] 3)
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "print")
        (mod [i64]
          (number [i64] 11)
          (number [i64] 3)
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "print")
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
        (procedure [void] "print")
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
        (procedure [void] "print")
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
        (procedure [void] "print")
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
        (procedure [void] "print")
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
        (procedure [void] "print")
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
