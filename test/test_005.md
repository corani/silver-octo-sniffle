# test/test_005.md
## Source
```pascal
MODULE Comparisons;

BEGIN
    Texts.WriteInt(ORD(1=2));               Texts.WriteLn;
    Texts.WriteInt(ORD(3.0/2.0=6.0/4.0));   Texts.WriteLn;
    Texts.WriteInt(ORD(1#2));               Texts.WriteLn;
    Texts.WriteInt(ORD(3 DIV 2#6 DIV 4));   Texts.WriteLn;
    Texts.WriteInt(ORD(1<2));               Texts.WriteLn;
    Texts.WriteInt(ORD(2.0/1.0<1.0/1.0));   Texts.WriteLn;
    Texts.WriteInt(ORD(1<=1));              Texts.WriteLn;
    Texts.WriteInt(ORD(1.0/1.0<=1.0/1.0));  Texts.WriteLn;
    Texts.WriteInt(ORD(1>2));               Texts.WriteLn;
    Texts.WriteInt(ORD(2.0/2.0>2.0/1.0));   Texts.WriteLn;
    Texts.WriteInt(ORD(2>=2));              Texts.WriteLn;
    Texts.WriteInt(ORD(2.0/1.0>=2.0/1.0));  Texts.WriteLn;
END Comparisons.
```
## Tokens
```tsv
test/test_005.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_005.md:1:8:	ident	"Comparisons"	false	0	0.000000	(1, 8) -> (1, 19)
test/test_005.md:1:19:	semicolon	";"	false	0	0.000000	(1, 19) -> (1, 20)
test/test_005.md:3:0:	begin	"BEGIN"	false	0	0.000000	(3, 0) -> (3, 5)
test/test_005.md:4:4:	ident	"Texts"	false	0	0.000000	(4, 4) -> (4, 9)
test/test_005.md:4:9:	dot	"."	false	0	0.000000	(4, 9) -> (4, 10)
test/test_005.md:4:10:	ident	"WriteInt"	false	0	0.000000	(4, 10) -> (4, 18)
test/test_005.md:4:18:	lparen	"("	false	0	0.000000	(4, 18) -> (4, 19)
test/test_005.md:4:19:	ident	"ORD"	false	0	0.000000	(4, 19) -> (4, 22)
test/test_005.md:4:22:	lparen	"("	false	0	0.000000	(4, 22) -> (4, 23)
test/test_005.md:4:23:	integer	"1"	false	1	0.000000	(4, 23) -> (4, 24)
test/test_005.md:4:24:	eq	"="	false	0	0.000000	(4, 24) -> (4, 25)
test/test_005.md:4:25:	integer	"2"	false	2	0.000000	(4, 25) -> (4, 26)
test/test_005.md:4:26:	rparen	")"	false	0	0.000000	(4, 26) -> (4, 27)
test/test_005.md:4:27:	rparen	")"	false	0	0.000000	(4, 27) -> (4, 28)
test/test_005.md:4:28:	semicolon	";"	false	0	0.000000	(4, 28) -> (4, 29)
test/test_005.md:4:44:	ident	"Texts"	false	0	0.000000	(4, 44) -> (4, 49)
test/test_005.md:4:49:	dot	"."	false	0	0.000000	(4, 49) -> (4, 50)
test/test_005.md:4:50:	ident	"WriteLn"	false	0	0.000000	(4, 50) -> (4, 57)
test/test_005.md:4:57:	semicolon	";"	false	0	0.000000	(4, 57) -> (4, 58)
test/test_005.md:5:4:	ident	"Texts"	false	0	0.000000	(5, 4) -> (5, 9)
test/test_005.md:5:9:	dot	"."	false	0	0.000000	(5, 9) -> (5, 10)
test/test_005.md:5:10:	ident	"WriteInt"	false	0	0.000000	(5, 10) -> (5, 18)
test/test_005.md:5:18:	lparen	"("	false	0	0.000000	(5, 18) -> (5, 19)
test/test_005.md:5:19:	ident	"ORD"	false	0	0.000000	(5, 19) -> (5, 22)
test/test_005.md:5:22:	lparen	"("	false	0	0.000000	(5, 22) -> (5, 23)
test/test_005.md:5:23:	real	"3.0"	false	0	3.000000	(5, 23) -> (5, 26)
test/test_005.md:5:26:	slash	"/"	false	0	0.000000	(5, 26) -> (5, 27)
test/test_005.md:5:27:	real	"2.0"	false	0	2.000000	(5, 27) -> (5, 30)
test/test_005.md:5:30:	eq	"="	false	0	0.000000	(5, 30) -> (5, 31)
test/test_005.md:5:31:	real	"6.0"	false	0	6.000000	(5, 31) -> (5, 34)
test/test_005.md:5:34:	slash	"/"	false	0	0.000000	(5, 34) -> (5, 35)
test/test_005.md:5:35:	real	"4.0"	false	0	4.000000	(5, 35) -> (5, 38)
test/test_005.md:5:38:	rparen	")"	false	0	0.000000	(5, 38) -> (5, 39)
test/test_005.md:5:39:	rparen	")"	false	0	0.000000	(5, 39) -> (5, 40)
test/test_005.md:5:40:	semicolon	";"	false	0	0.000000	(5, 40) -> (5, 41)
test/test_005.md:5:44:	ident	"Texts"	false	0	0.000000	(5, 44) -> (5, 49)
test/test_005.md:5:49:	dot	"."	false	0	0.000000	(5, 49) -> (5, 50)
test/test_005.md:5:50:	ident	"WriteLn"	false	0	0.000000	(5, 50) -> (5, 57)
test/test_005.md:5:57:	semicolon	";"	false	0	0.000000	(5, 57) -> (5, 58)
test/test_005.md:6:4:	ident	"Texts"	false	0	0.000000	(6, 4) -> (6, 9)
test/test_005.md:6:9:	dot	"."	false	0	0.000000	(6, 9) -> (6, 10)
test/test_005.md:6:10:	ident	"WriteInt"	false	0	0.000000	(6, 10) -> (6, 18)
test/test_005.md:6:18:	lparen	"("	false	0	0.000000	(6, 18) -> (6, 19)
test/test_005.md:6:19:	ident	"ORD"	false	0	0.000000	(6, 19) -> (6, 22)
test/test_005.md:6:22:	lparen	"("	false	0	0.000000	(6, 22) -> (6, 23)
test/test_005.md:6:23:	integer	"1"	false	1	0.000000	(6, 23) -> (6, 24)
test/test_005.md:6:24:	ne	"#"	false	0	0.000000	(6, 24) -> (6, 25)
test/test_005.md:6:25:	integer	"2"	false	2	0.000000	(6, 25) -> (6, 26)
test/test_005.md:6:26:	rparen	")"	false	0	0.000000	(6, 26) -> (6, 27)
test/test_005.md:6:27:	rparen	")"	false	0	0.000000	(6, 27) -> (6, 28)
test/test_005.md:6:28:	semicolon	";"	false	0	0.000000	(6, 28) -> (6, 29)
test/test_005.md:6:44:	ident	"Texts"	false	0	0.000000	(6, 44) -> (6, 49)
test/test_005.md:6:49:	dot	"."	false	0	0.000000	(6, 49) -> (6, 50)
test/test_005.md:6:50:	ident	"WriteLn"	false	0	0.000000	(6, 50) -> (6, 57)
test/test_005.md:6:57:	semicolon	";"	false	0	0.000000	(6, 57) -> (6, 58)
test/test_005.md:7:4:	ident	"Texts"	false	0	0.000000	(7, 4) -> (7, 9)
test/test_005.md:7:9:	dot	"."	false	0	0.000000	(7, 9) -> (7, 10)
test/test_005.md:7:10:	ident	"WriteInt"	false	0	0.000000	(7, 10) -> (7, 18)
test/test_005.md:7:18:	lparen	"("	false	0	0.000000	(7, 18) -> (7, 19)
test/test_005.md:7:19:	ident	"ORD"	false	0	0.000000	(7, 19) -> (7, 22)
test/test_005.md:7:22:	lparen	"("	false	0	0.000000	(7, 22) -> (7, 23)
test/test_005.md:7:23:	integer	"3"	false	3	0.000000	(7, 23) -> (7, 24)
test/test_005.md:7:25:	div	"DIV"	false	0	0.000000	(7, 25) -> (7, 28)
test/test_005.md:7:29:	integer	"2"	false	2	0.000000	(7, 29) -> (7, 30)
test/test_005.md:7:30:	ne	"#"	false	0	0.000000	(7, 30) -> (7, 31)
test/test_005.md:7:31:	integer	"6"	false	6	0.000000	(7, 31) -> (7, 32)
test/test_005.md:7:33:	div	"DIV"	false	0	0.000000	(7, 33) -> (7, 36)
test/test_005.md:7:37:	integer	"4"	false	4	0.000000	(7, 37) -> (7, 38)
test/test_005.md:7:38:	rparen	")"	false	0	0.000000	(7, 38) -> (7, 39)
test/test_005.md:7:39:	rparen	")"	false	0	0.000000	(7, 39) -> (7, 40)
test/test_005.md:7:40:	semicolon	";"	false	0	0.000000	(7, 40) -> (7, 41)
test/test_005.md:7:44:	ident	"Texts"	false	0	0.000000	(7, 44) -> (7, 49)
test/test_005.md:7:49:	dot	"."	false	0	0.000000	(7, 49) -> (7, 50)
test/test_005.md:7:50:	ident	"WriteLn"	false	0	0.000000	(7, 50) -> (7, 57)
test/test_005.md:7:57:	semicolon	";"	false	0	0.000000	(7, 57) -> (7, 58)
test/test_005.md:8:4:	ident	"Texts"	false	0	0.000000	(8, 4) -> (8, 9)
test/test_005.md:8:9:	dot	"."	false	0	0.000000	(8, 9) -> (8, 10)
test/test_005.md:8:10:	ident	"WriteInt"	false	0	0.000000	(8, 10) -> (8, 18)
test/test_005.md:8:18:	lparen	"("	false	0	0.000000	(8, 18) -> (8, 19)
test/test_005.md:8:19:	ident	"ORD"	false	0	0.000000	(8, 19) -> (8, 22)
test/test_005.md:8:22:	lparen	"("	false	0	0.000000	(8, 22) -> (8, 23)
test/test_005.md:8:23:	integer	"1"	false	1	0.000000	(8, 23) -> (8, 24)
test/test_005.md:8:24:	lt	"<"	false	0	0.000000	(8, 24) -> (8, 25)
test/test_005.md:8:25:	integer	"2"	false	2	0.000000	(8, 25) -> (8, 26)
test/test_005.md:8:26:	rparen	")"	false	0	0.000000	(8, 26) -> (8, 27)
test/test_005.md:8:27:	rparen	")"	false	0	0.000000	(8, 27) -> (8, 28)
test/test_005.md:8:28:	semicolon	";"	false	0	0.000000	(8, 28) -> (8, 29)
test/test_005.md:8:44:	ident	"Texts"	false	0	0.000000	(8, 44) -> (8, 49)
test/test_005.md:8:49:	dot	"."	false	0	0.000000	(8, 49) -> (8, 50)
test/test_005.md:8:50:	ident	"WriteLn"	false	0	0.000000	(8, 50) -> (8, 57)
test/test_005.md:8:57:	semicolon	";"	false	0	0.000000	(8, 57) -> (8, 58)
test/test_005.md:9:4:	ident	"Texts"	false	0	0.000000	(9, 4) -> (9, 9)
test/test_005.md:9:9:	dot	"."	false	0	0.000000	(9, 9) -> (9, 10)
test/test_005.md:9:10:	ident	"WriteInt"	false	0	0.000000	(9, 10) -> (9, 18)
test/test_005.md:9:18:	lparen	"("	false	0	0.000000	(9, 18) -> (9, 19)
test/test_005.md:9:19:	ident	"ORD"	false	0	0.000000	(9, 19) -> (9, 22)
test/test_005.md:9:22:	lparen	"("	false	0	0.000000	(9, 22) -> (9, 23)
test/test_005.md:9:23:	real	"2.0"	false	0	2.000000	(9, 23) -> (9, 26)
test/test_005.md:9:26:	slash	"/"	false	0	0.000000	(9, 26) -> (9, 27)
test/test_005.md:9:27:	real	"1.0"	false	0	1.000000	(9, 27) -> (9, 30)
test/test_005.md:9:30:	lt	"<"	false	0	0.000000	(9, 30) -> (9, 31)
test/test_005.md:9:31:	real	"1.0"	false	0	1.000000	(9, 31) -> (9, 34)
test/test_005.md:9:34:	slash	"/"	false	0	0.000000	(9, 34) -> (9, 35)
test/test_005.md:9:35:	real	"1.0"	false	0	1.000000	(9, 35) -> (9, 38)
test/test_005.md:9:38:	rparen	")"	false	0	0.000000	(9, 38) -> (9, 39)
test/test_005.md:9:39:	rparen	")"	false	0	0.000000	(9, 39) -> (9, 40)
test/test_005.md:9:40:	semicolon	";"	false	0	0.000000	(9, 40) -> (9, 41)
test/test_005.md:9:44:	ident	"Texts"	false	0	0.000000	(9, 44) -> (9, 49)
test/test_005.md:9:49:	dot	"."	false	0	0.000000	(9, 49) -> (9, 50)
test/test_005.md:9:50:	ident	"WriteLn"	false	0	0.000000	(9, 50) -> (9, 57)
test/test_005.md:9:57:	semicolon	";"	false	0	0.000000	(9, 57) -> (9, 58)
test/test_005.md:10:4:	ident	"Texts"	false	0	0.000000	(10, 4) -> (10, 9)
test/test_005.md:10:9:	dot	"."	false	0	0.000000	(10, 9) -> (10, 10)
test/test_005.md:10:10:	ident	"WriteInt"	false	0	0.000000	(10, 10) -> (10, 18)
test/test_005.md:10:18:	lparen	"("	false	0	0.000000	(10, 18) -> (10, 19)
test/test_005.md:10:19:	ident	"ORD"	false	0	0.000000	(10, 19) -> (10, 22)
test/test_005.md:10:22:	lparen	"("	false	0	0.000000	(10, 22) -> (10, 23)
test/test_005.md:10:23:	integer	"1"	false	1	0.000000	(10, 23) -> (10, 24)
test/test_005.md:10:24:	le	"<="	false	0	0.000000	(10, 24) -> (10, 26)
test/test_005.md:10:26:	integer	"1"	false	1	0.000000	(10, 26) -> (10, 27)
test/test_005.md:10:27:	rparen	")"	false	0	0.000000	(10, 27) -> (10, 28)
test/test_005.md:10:28:	rparen	")"	false	0	0.000000	(10, 28) -> (10, 29)
test/test_005.md:10:29:	semicolon	";"	false	0	0.000000	(10, 29) -> (10, 30)
test/test_005.md:10:44:	ident	"Texts"	false	0	0.000000	(10, 44) -> (10, 49)
test/test_005.md:10:49:	dot	"."	false	0	0.000000	(10, 49) -> (10, 50)
test/test_005.md:10:50:	ident	"WriteLn"	false	0	0.000000	(10, 50) -> (10, 57)
test/test_005.md:10:57:	semicolon	";"	false	0	0.000000	(10, 57) -> (10, 58)
test/test_005.md:11:4:	ident	"Texts"	false	0	0.000000	(11, 4) -> (11, 9)
test/test_005.md:11:9:	dot	"."	false	0	0.000000	(11, 9) -> (11, 10)
test/test_005.md:11:10:	ident	"WriteInt"	false	0	0.000000	(11, 10) -> (11, 18)
test/test_005.md:11:18:	lparen	"("	false	0	0.000000	(11, 18) -> (11, 19)
test/test_005.md:11:19:	ident	"ORD"	false	0	0.000000	(11, 19) -> (11, 22)
test/test_005.md:11:22:	lparen	"("	false	0	0.000000	(11, 22) -> (11, 23)
test/test_005.md:11:23:	real	"1.0"	false	0	1.000000	(11, 23) -> (11, 26)
test/test_005.md:11:26:	slash	"/"	false	0	0.000000	(11, 26) -> (11, 27)
test/test_005.md:11:27:	real	"1.0"	false	0	1.000000	(11, 27) -> (11, 30)
test/test_005.md:11:30:	le	"<="	false	0	0.000000	(11, 30) -> (11, 32)
test/test_005.md:11:32:	real	"1.0"	false	0	1.000000	(11, 32) -> (11, 35)
test/test_005.md:11:35:	slash	"/"	false	0	0.000000	(11, 35) -> (11, 36)
test/test_005.md:11:36:	real	"1.0"	false	0	1.000000	(11, 36) -> (11, 39)
test/test_005.md:11:39:	rparen	")"	false	0	0.000000	(11, 39) -> (11, 40)
test/test_005.md:11:40:	rparen	")"	false	0	0.000000	(11, 40) -> (11, 41)
test/test_005.md:11:41:	semicolon	";"	false	0	0.000000	(11, 41) -> (11, 42)
test/test_005.md:11:44:	ident	"Texts"	false	0	0.000000	(11, 44) -> (11, 49)
test/test_005.md:11:49:	dot	"."	false	0	0.000000	(11, 49) -> (11, 50)
test/test_005.md:11:50:	ident	"WriteLn"	false	0	0.000000	(11, 50) -> (11, 57)
test/test_005.md:11:57:	semicolon	";"	false	0	0.000000	(11, 57) -> (11, 58)
test/test_005.md:12:4:	ident	"Texts"	false	0	0.000000	(12, 4) -> (12, 9)
test/test_005.md:12:9:	dot	"."	false	0	0.000000	(12, 9) -> (12, 10)
test/test_005.md:12:10:	ident	"WriteInt"	false	0	0.000000	(12, 10) -> (12, 18)
test/test_005.md:12:18:	lparen	"("	false	0	0.000000	(12, 18) -> (12, 19)
test/test_005.md:12:19:	ident	"ORD"	false	0	0.000000	(12, 19) -> (12, 22)
test/test_005.md:12:22:	lparen	"("	false	0	0.000000	(12, 22) -> (12, 23)
test/test_005.md:12:23:	integer	"1"	false	1	0.000000	(12, 23) -> (12, 24)
test/test_005.md:12:24:	gt	">"	false	0	0.000000	(12, 24) -> (12, 25)
test/test_005.md:12:25:	integer	"2"	false	2	0.000000	(12, 25) -> (12, 26)
test/test_005.md:12:26:	rparen	")"	false	0	0.000000	(12, 26) -> (12, 27)
test/test_005.md:12:27:	rparen	")"	false	0	0.000000	(12, 27) -> (12, 28)
test/test_005.md:12:28:	semicolon	";"	false	0	0.000000	(12, 28) -> (12, 29)
test/test_005.md:12:44:	ident	"Texts"	false	0	0.000000	(12, 44) -> (12, 49)
test/test_005.md:12:49:	dot	"."	false	0	0.000000	(12, 49) -> (12, 50)
test/test_005.md:12:50:	ident	"WriteLn"	false	0	0.000000	(12, 50) -> (12, 57)
test/test_005.md:12:57:	semicolon	";"	false	0	0.000000	(12, 57) -> (12, 58)
test/test_005.md:13:4:	ident	"Texts"	false	0	0.000000	(13, 4) -> (13, 9)
test/test_005.md:13:9:	dot	"."	false	0	0.000000	(13, 9) -> (13, 10)
test/test_005.md:13:10:	ident	"WriteInt"	false	0	0.000000	(13, 10) -> (13, 18)
test/test_005.md:13:18:	lparen	"("	false	0	0.000000	(13, 18) -> (13, 19)
test/test_005.md:13:19:	ident	"ORD"	false	0	0.000000	(13, 19) -> (13, 22)
test/test_005.md:13:22:	lparen	"("	false	0	0.000000	(13, 22) -> (13, 23)
test/test_005.md:13:23:	real	"2.0"	false	0	2.000000	(13, 23) -> (13, 26)
test/test_005.md:13:26:	slash	"/"	false	0	0.000000	(13, 26) -> (13, 27)
test/test_005.md:13:27:	real	"2.0"	false	0	2.000000	(13, 27) -> (13, 30)
test/test_005.md:13:30:	gt	">"	false	0	0.000000	(13, 30) -> (13, 31)
test/test_005.md:13:31:	real	"2.0"	false	0	2.000000	(13, 31) -> (13, 34)
test/test_005.md:13:34:	slash	"/"	false	0	0.000000	(13, 34) -> (13, 35)
test/test_005.md:13:35:	real	"1.0"	false	0	1.000000	(13, 35) -> (13, 38)
test/test_005.md:13:38:	rparen	")"	false	0	0.000000	(13, 38) -> (13, 39)
test/test_005.md:13:39:	rparen	")"	false	0	0.000000	(13, 39) -> (13, 40)
test/test_005.md:13:40:	semicolon	";"	false	0	0.000000	(13, 40) -> (13, 41)
test/test_005.md:13:44:	ident	"Texts"	false	0	0.000000	(13, 44) -> (13, 49)
test/test_005.md:13:49:	dot	"."	false	0	0.000000	(13, 49) -> (13, 50)
test/test_005.md:13:50:	ident	"WriteLn"	false	0	0.000000	(13, 50) -> (13, 57)
test/test_005.md:13:57:	semicolon	";"	false	0	0.000000	(13, 57) -> (13, 58)
test/test_005.md:14:4:	ident	"Texts"	false	0	0.000000	(14, 4) -> (14, 9)
test/test_005.md:14:9:	dot	"."	false	0	0.000000	(14, 9) -> (14, 10)
test/test_005.md:14:10:	ident	"WriteInt"	false	0	0.000000	(14, 10) -> (14, 18)
test/test_005.md:14:18:	lparen	"("	false	0	0.000000	(14, 18) -> (14, 19)
test/test_005.md:14:19:	ident	"ORD"	false	0	0.000000	(14, 19) -> (14, 22)
test/test_005.md:14:22:	lparen	"("	false	0	0.000000	(14, 22) -> (14, 23)
test/test_005.md:14:23:	integer	"2"	false	2	0.000000	(14, 23) -> (14, 24)
test/test_005.md:14:24:	ge	">="	false	0	0.000000	(14, 24) -> (14, 26)
test/test_005.md:14:26:	integer	"2"	false	2	0.000000	(14, 26) -> (14, 27)
test/test_005.md:14:27:	rparen	")"	false	0	0.000000	(14, 27) -> (14, 28)
test/test_005.md:14:28:	rparen	")"	false	0	0.000000	(14, 28) -> (14, 29)
test/test_005.md:14:29:	semicolon	";"	false	0	0.000000	(14, 29) -> (14, 30)
test/test_005.md:14:44:	ident	"Texts"	false	0	0.000000	(14, 44) -> (14, 49)
test/test_005.md:14:49:	dot	"."	false	0	0.000000	(14, 49) -> (14, 50)
test/test_005.md:14:50:	ident	"WriteLn"	false	0	0.000000	(14, 50) -> (14, 57)
test/test_005.md:14:57:	semicolon	";"	false	0	0.000000	(14, 57) -> (14, 58)
test/test_005.md:15:4:	ident	"Texts"	false	0	0.000000	(15, 4) -> (15, 9)
test/test_005.md:15:9:	dot	"."	false	0	0.000000	(15, 9) -> (15, 10)
test/test_005.md:15:10:	ident	"WriteInt"	false	0	0.000000	(15, 10) -> (15, 18)
test/test_005.md:15:18:	lparen	"("	false	0	0.000000	(15, 18) -> (15, 19)
test/test_005.md:15:19:	ident	"ORD"	false	0	0.000000	(15, 19) -> (15, 22)
test/test_005.md:15:22:	lparen	"("	false	0	0.000000	(15, 22) -> (15, 23)
test/test_005.md:15:23:	real	"2.0"	false	0	2.000000	(15, 23) -> (15, 26)
test/test_005.md:15:26:	slash	"/"	false	0	0.000000	(15, 26) -> (15, 27)
test/test_005.md:15:27:	real	"1.0"	false	0	1.000000	(15, 27) -> (15, 30)
test/test_005.md:15:30:	ge	">="	false	0	0.000000	(15, 30) -> (15, 32)
test/test_005.md:15:32:	real	"2.0"	false	0	2.000000	(15, 32) -> (15, 35)
test/test_005.md:15:35:	slash	"/"	false	0	0.000000	(15, 35) -> (15, 36)
test/test_005.md:15:36:	real	"1.0"	false	0	1.000000	(15, 36) -> (15, 39)
test/test_005.md:15:39:	rparen	")"	false	0	0.000000	(15, 39) -> (15, 40)
test/test_005.md:15:40:	rparen	")"	false	0	0.000000	(15, 40) -> (15, 41)
test/test_005.md:15:41:	semicolon	";"	false	0	0.000000	(15, 41) -> (15, 42)
test/test_005.md:15:44:	ident	"Texts"	false	0	0.000000	(15, 44) -> (15, 49)
test/test_005.md:15:49:	dot	"."	false	0	0.000000	(15, 49) -> (15, 50)
test/test_005.md:15:50:	ident	"WriteLn"	false	0	0.000000	(15, 50) -> (15, 57)
test/test_005.md:15:57:	semicolon	";"	false	0	0.000000	(15, 57) -> (15, 58)
test/test_005.md:16:0:	end	"END"	false	0	0.000000	(16, 0) -> (16, 3)
test/test_005.md:16:4:	ident	"Comparisons"	false	0	0.000000	(16, 4) -> (16, 15)
test/test_005.md:16:15:	dot	"."	false	0	0.000000	(16, 15) -> (16, 16)
test/test_005.md:17:0:	eof	""	false	0	0.000000	(17, 0) -> (17, 0)
```
## AST
```scheme
(module "Comparisons"
  (stmts
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteInt")
        (call
          (procedure [i64] "ORD")
          (eq [boolean]
            (number [i64] 1)
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
        (procedure [void] "Texts.WriteInt")
        (call
          (procedure [i64] "ORD")
          (eq [boolean]
            (slash [f64]
              (number [f64] 3.000000)
              (number [f64] 2.000000)
            )
            (slash [f64]
              (number [f64] 6.000000)
              (number [f64] 4.000000)
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
        (procedure [void] "Texts.WriteInt")
        (call
          (procedure [i64] "ORD")
          (ne [boolean]
            (number [i64] 1)
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
        (procedure [void] "Texts.WriteInt")
        (call
          (procedure [i64] "ORD")
          (ne [boolean]
            (div [i64]
              (number [i64] 3)
              (number [i64] 2)
            )
            (div [i64]
              (number [i64] 6)
              (number [i64] 4)
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
        (procedure [void] "Texts.WriteInt")
        (call
          (procedure [i64] "ORD")
          (lt [boolean]
            (number [i64] 1)
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
        (procedure [void] "Texts.WriteInt")
        (call
          (procedure [i64] "ORD")
          (lt [boolean]
            (slash [f64]
              (number [f64] 2.000000)
              (number [f64] 1.000000)
            )
            (slash [f64]
              (number [f64] 1.000000)
              (number [f64] 1.000000)
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
        (procedure [void] "Texts.WriteInt")
        (call
          (procedure [i64] "ORD")
          (le [boolean]
            (number [i64] 1)
            (number [i64] 1)
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
        (call
          (procedure [i64] "ORD")
          (le [boolean]
            (slash [f64]
              (number [f64] 1.000000)
              (number [f64] 1.000000)
            )
            (slash [f64]
              (number [f64] 1.000000)
              (number [f64] 1.000000)
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
        (procedure [void] "Texts.WriteInt")
        (call
          (procedure [i64] "ORD")
          (gt [boolean]
            (number [i64] 1)
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
        (procedure [void] "Texts.WriteInt")
        (call
          (procedure [i64] "ORD")
          (gt [boolean]
            (slash [f64]
              (number [f64] 2.000000)
              (number [f64] 2.000000)
            )
            (slash [f64]
              (number [f64] 2.000000)
              (number [f64] 1.000000)
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
        (procedure [void] "Texts.WriteInt")
        (call
          (procedure [i64] "ORD")
          (ge [boolean]
            (number [i64] 2)
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
        (procedure [void] "Texts.WriteInt")
        (call
          (procedure [i64] "ORD")
          (ge [boolean]
            (slash [f64]
              (number [f64] 2.000000)
              (number [f64] 1.000000)
            )
            (slash [f64]
              (number [f64] 2.000000)
              (number [f64] 1.000000)
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
  )
)
```
## IR
```llvm
@0 = global [3 x i8] c"%d\00"
@1 = global [1 x i8] c"\00"

declare i64 @puts(i8* %str)

declare i64 @rand()

declare i64 @sprintf(i8* %buf, i8* %format, ...)

declare i64 @printf(i8* %format, ...)

declare i8* @malloc(i64 %size)

declare i8* @free(i8* %ptr)

define i64 @main() {
entry:
	%0 = icmp eq i64 1, 2
	%1 = zext i1 %0 to i64
	%2 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%3 = call i64 (i8*, ...) @printf(i8* %2, i64 %1)
	%4 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%5 = call i64 @puts(i8* %4)
	%6 = fdiv double 3.0, 2.0
	%7 = fdiv double 6.0, 4.0
	%8 = fcmp ueq double %6, %7
	%9 = zext i1 %8 to i64
	%10 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%11 = call i64 (i8*, ...) @printf(i8* %10, i64 %9)
	%12 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%13 = call i64 @puts(i8* %12)
	%14 = icmp ne i64 1, 2
	%15 = zext i1 %14 to i64
	%16 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%17 = call i64 (i8*, ...) @printf(i8* %16, i64 %15)
	%18 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%19 = call i64 @puts(i8* %18)
	%20 = sdiv i64 3, 2
	%21 = sdiv i64 6, 4
	%22 = icmp ne i64 %20, %21
	%23 = zext i1 %22 to i64
	%24 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%25 = call i64 (i8*, ...) @printf(i8* %24, i64 %23)
	%26 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%27 = call i64 @puts(i8* %26)
	%28 = icmp slt i64 1, 2
	%29 = zext i1 %28 to i64
	%30 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%31 = call i64 (i8*, ...) @printf(i8* %30, i64 %29)
	%32 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%33 = call i64 @puts(i8* %32)
	%34 = fdiv double 2.0, 1.0
	%35 = fdiv double 1.0, 1.0
	%36 = fcmp ult double %34, %35
	%37 = zext i1 %36 to i64
	%38 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%39 = call i64 (i8*, ...) @printf(i8* %38, i64 %37)
	%40 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%41 = call i64 @puts(i8* %40)
	%42 = icmp sle i64 1, 1
	%43 = zext i1 %42 to i64
	%44 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%45 = call i64 (i8*, ...) @printf(i8* %44, i64 %43)
	%46 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%47 = call i64 @puts(i8* %46)
	%48 = fdiv double 1.0, 1.0
	%49 = fdiv double 1.0, 1.0
	%50 = fcmp ule double %48, %49
	%51 = zext i1 %50 to i64
	%52 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%53 = call i64 (i8*, ...) @printf(i8* %52, i64 %51)
	%54 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%55 = call i64 @puts(i8* %54)
	%56 = icmp sgt i64 1, 2
	%57 = zext i1 %56 to i64
	%58 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%59 = call i64 (i8*, ...) @printf(i8* %58, i64 %57)
	%60 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%61 = call i64 @puts(i8* %60)
	%62 = fdiv double 2.0, 2.0
	%63 = fdiv double 2.0, 1.0
	%64 = fcmp ugt double %62, %63
	%65 = zext i1 %64 to i64
	%66 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%67 = call i64 (i8*, ...) @printf(i8* %66, i64 %65)
	%68 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%69 = call i64 @puts(i8* %68)
	%70 = icmp sge i64 2, 2
	%71 = zext i1 %70 to i64
	%72 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%73 = call i64 (i8*, ...) @printf(i8* %72, i64 %71)
	%74 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%75 = call i64 @puts(i8* %74)
	%76 = fdiv double 2.0, 1.0
	%77 = fdiv double 2.0, 1.0
	%78 = fcmp uge double %76, %77
	%79 = zext i1 %78 to i64
	%80 = getelementptr [3 x i8], [3 x i8]* @0, i64 0, i64 0
	%81 = call i64 (i8*, ...) @printf(i8* %80, i64 %79)
	%82 = getelementptr [1 x i8], [1 x i8]* @1, i64 0, i64 0
	%83 = call i64 @puts(i8* %82)
	ret i64 0
}

```
## Run
```bash
0
1
1
0
1
0
1
1
0
0
1
1
```
