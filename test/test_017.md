# test/test_017.md
## Source
```pascal
MODULE Sets;

CONST xc = {1, 2, 3, 10};
      yc = {10..19};

VAR x: SET;
    y: SET;
    z: SET;

BEGIN
    Texts.WriteInt(ORD(xc));    Texts.WriteLn;
    Texts.WriteInt(ORD(yc));    Texts.WriteLn;
    Texts.WriteInt(ORD(xc+yc)); Texts.WriteLn;
    Texts.WriteInt(ORD(xc-yc)); Texts.WriteLn;
    Texts.WriteInt(ORD(xc*yc)); Texts.WriteLn;
    Texts.WriteInt(ORD(xc/yc)); Texts.WriteLn;

    x := xc; y := yc;

    INCL(x, 4);
    EXCL(y, 30);

    z := x + y;

    Texts.WriteInt(ORD(z));     Texts.WriteLn;

    INCL(z, 30);
    EXCL(z, 4);

    Texts.WriteInt(ORD(z));     Texts.WriteLn;

    IF 30 IN z THEN
        Texts.WriteInt(30);     Texts.WriteLn;
    END;

    IF ~(31 IN z) THEN
        Texts.WriteInt(31);     Texts.WriteLn;
    END;

    Texts.WriteInt(ORD({1, 2})); Texts.WriteLn;
END Sets.
```
## Tokens
```tsv
test/test_017.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_017.md:1:8:	ident	"Sets"	false	0	0.000000	(1, 8) -> (1, 12)
test/test_017.md:1:12:	semicolon	";"	false	0	0.000000	(1, 12) -> (1, 13)
test/test_017.md:3:0:	const	"CONST"	false	0	0.000000	(3, 0) -> (3, 5)
test/test_017.md:3:6:	ident	"xc"	false	0	0.000000	(3, 6) -> (3, 8)
test/test_017.md:3:9:	eq	"="	false	0	0.000000	(3, 9) -> (3, 10)
test/test_017.md:3:11:	lbrace	"{"	false	0	0.000000	(3, 11) -> (3, 12)
test/test_017.md:3:12:	integer	"1"	false	1	0.000000	(3, 12) -> (3, 13)
test/test_017.md:3:13:	comma	","	false	0	0.000000	(3, 13) -> (3, 14)
test/test_017.md:3:15:	integer	"2"	false	2	0.000000	(3, 15) -> (3, 16)
test/test_017.md:3:16:	comma	","	false	0	0.000000	(3, 16) -> (3, 17)
test/test_017.md:3:18:	integer	"3"	false	3	0.000000	(3, 18) -> (3, 19)
test/test_017.md:3:19:	comma	","	false	0	0.000000	(3, 19) -> (3, 20)
test/test_017.md:3:21:	integer	"10"	false	10	0.000000	(3, 21) -> (3, 23)
test/test_017.md:3:23:	rbrace	"}"	false	0	0.000000	(3, 23) -> (3, 24)
test/test_017.md:3:24:	semicolon	";"	false	0	0.000000	(3, 24) -> (3, 25)
test/test_017.md:4:6:	ident	"yc"	false	0	0.000000	(4, 6) -> (4, 8)
test/test_017.md:4:9:	eq	"="	false	0	0.000000	(4, 9) -> (4, 10)
test/test_017.md:4:11:	lbrace	"{"	false	0	0.000000	(4, 11) -> (4, 12)
test/test_017.md:4:12:	integer	"10"	false	10	0.000000	(4, 12) -> (4, 16)
test/test_017.md:4:16:	dotdot	".."	false	0	0.000000	(4, 16) -> (4, 18)
test/test_017.md:4:18:	integer	"19"	false	19	0.000000	(4, 18) -> (4, 20)
test/test_017.md:4:20:	rbrace	"}"	false	0	0.000000	(4, 20) -> (4, 21)
test/test_017.md:4:21:	semicolon	";"	false	0	0.000000	(4, 21) -> (4, 22)
test/test_017.md:6:0:	var	"VAR"	false	0	0.000000	(6, 0) -> (6, 3)
test/test_017.md:6:4:	ident	"x"	false	0	0.000000	(6, 4) -> (6, 5)
test/test_017.md:6:5:	colon	":"	false	0	0.000000	(6, 5) -> (6, 6)
test/test_017.md:6:7:	ident	"SET"	false	0	0.000000	(6, 7) -> (6, 10)
test/test_017.md:6:10:	semicolon	";"	false	0	0.000000	(6, 10) -> (6, 11)
test/test_017.md:7:4:	ident	"y"	false	0	0.000000	(7, 4) -> (7, 5)
test/test_017.md:7:5:	colon	":"	false	0	0.000000	(7, 5) -> (7, 6)
test/test_017.md:7:7:	ident	"SET"	false	0	0.000000	(7, 7) -> (7, 10)
test/test_017.md:7:10:	semicolon	";"	false	0	0.000000	(7, 10) -> (7, 11)
test/test_017.md:8:4:	ident	"z"	false	0	0.000000	(8, 4) -> (8, 5)
test/test_017.md:8:5:	colon	":"	false	0	0.000000	(8, 5) -> (8, 6)
test/test_017.md:8:7:	ident	"SET"	false	0	0.000000	(8, 7) -> (8, 10)
test/test_017.md:8:10:	semicolon	";"	false	0	0.000000	(8, 10) -> (8, 11)
test/test_017.md:10:0:	begin	"BEGIN"	false	0	0.000000	(10, 0) -> (10, 5)
test/test_017.md:11:4:	ident	"Texts"	false	0	0.000000	(11, 4) -> (11, 9)
test/test_017.md:11:9:	dot	"."	false	0	0.000000	(11, 9) -> (11, 10)
test/test_017.md:11:10:	ident	"WriteInt"	false	0	0.000000	(11, 10) -> (11, 18)
test/test_017.md:11:18:	lparen	"("	false	0	0.000000	(11, 18) -> (11, 19)
test/test_017.md:11:19:	ident	"ORD"	false	0	0.000000	(11, 19) -> (11, 22)
test/test_017.md:11:22:	lparen	"("	false	0	0.000000	(11, 22) -> (11, 23)
test/test_017.md:11:23:	ident	"xc"	false	0	0.000000	(11, 23) -> (11, 25)
test/test_017.md:11:25:	rparen	")"	false	0	0.000000	(11, 25) -> (11, 26)
test/test_017.md:11:26:	rparen	")"	false	0	0.000000	(11, 26) -> (11, 27)
test/test_017.md:11:27:	semicolon	";"	false	0	0.000000	(11, 27) -> (11, 28)
test/test_017.md:11:32:	ident	"Texts"	false	0	0.000000	(11, 32) -> (11, 37)
test/test_017.md:11:37:	dot	"."	false	0	0.000000	(11, 37) -> (11, 38)
test/test_017.md:11:38:	ident	"WriteLn"	false	0	0.000000	(11, 38) -> (11, 45)
test/test_017.md:11:45:	semicolon	";"	false	0	0.000000	(11, 45) -> (11, 46)
test/test_017.md:12:4:	ident	"Texts"	false	0	0.000000	(12, 4) -> (12, 9)
test/test_017.md:12:9:	dot	"."	false	0	0.000000	(12, 9) -> (12, 10)
test/test_017.md:12:10:	ident	"WriteInt"	false	0	0.000000	(12, 10) -> (12, 18)
test/test_017.md:12:18:	lparen	"("	false	0	0.000000	(12, 18) -> (12, 19)
test/test_017.md:12:19:	ident	"ORD"	false	0	0.000000	(12, 19) -> (12, 22)
test/test_017.md:12:22:	lparen	"("	false	0	0.000000	(12, 22) -> (12, 23)
test/test_017.md:12:23:	ident	"yc"	false	0	0.000000	(12, 23) -> (12, 25)
test/test_017.md:12:25:	rparen	")"	false	0	0.000000	(12, 25) -> (12, 26)
test/test_017.md:12:26:	rparen	")"	false	0	0.000000	(12, 26) -> (12, 27)
test/test_017.md:12:27:	semicolon	";"	false	0	0.000000	(12, 27) -> (12, 28)
test/test_017.md:12:32:	ident	"Texts"	false	0	0.000000	(12, 32) -> (12, 37)
test/test_017.md:12:37:	dot	"."	false	0	0.000000	(12, 37) -> (12, 38)
test/test_017.md:12:38:	ident	"WriteLn"	false	0	0.000000	(12, 38) -> (12, 45)
test/test_017.md:12:45:	semicolon	";"	false	0	0.000000	(12, 45) -> (12, 46)
test/test_017.md:13:4:	ident	"Texts"	false	0	0.000000	(13, 4) -> (13, 9)
test/test_017.md:13:9:	dot	"."	false	0	0.000000	(13, 9) -> (13, 10)
test/test_017.md:13:10:	ident	"WriteInt"	false	0	0.000000	(13, 10) -> (13, 18)
test/test_017.md:13:18:	lparen	"("	false	0	0.000000	(13, 18) -> (13, 19)
test/test_017.md:13:19:	ident	"ORD"	false	0	0.000000	(13, 19) -> (13, 22)
test/test_017.md:13:22:	lparen	"("	false	0	0.000000	(13, 22) -> (13, 23)
test/test_017.md:13:23:	ident	"xc"	false	0	0.000000	(13, 23) -> (13, 25)
test/test_017.md:13:25:	plus	"+"	false	0	0.000000	(13, 25) -> (13, 26)
test/test_017.md:13:26:	ident	"yc"	false	0	0.000000	(13, 26) -> (13, 28)
test/test_017.md:13:28:	rparen	")"	false	0	0.000000	(13, 28) -> (13, 29)
test/test_017.md:13:29:	rparen	")"	false	0	0.000000	(13, 29) -> (13, 30)
test/test_017.md:13:30:	semicolon	";"	false	0	0.000000	(13, 30) -> (13, 31)
test/test_017.md:13:32:	ident	"Texts"	false	0	0.000000	(13, 32) -> (13, 37)
test/test_017.md:13:37:	dot	"."	false	0	0.000000	(13, 37) -> (13, 38)
test/test_017.md:13:38:	ident	"WriteLn"	false	0	0.000000	(13, 38) -> (13, 45)
test/test_017.md:13:45:	semicolon	";"	false	0	0.000000	(13, 45) -> (13, 46)
test/test_017.md:14:4:	ident	"Texts"	false	0	0.000000	(14, 4) -> (14, 9)
test/test_017.md:14:9:	dot	"."	false	0	0.000000	(14, 9) -> (14, 10)
test/test_017.md:14:10:	ident	"WriteInt"	false	0	0.000000	(14, 10) -> (14, 18)
test/test_017.md:14:18:	lparen	"("	false	0	0.000000	(14, 18) -> (14, 19)
test/test_017.md:14:19:	ident	"ORD"	false	0	0.000000	(14, 19) -> (14, 22)
test/test_017.md:14:22:	lparen	"("	false	0	0.000000	(14, 22) -> (14, 23)
test/test_017.md:14:23:	ident	"xc"	false	0	0.000000	(14, 23) -> (14, 25)
test/test_017.md:14:25:	minus	"-"	false	0	0.000000	(14, 25) -> (14, 26)
test/test_017.md:14:26:	ident	"yc"	false	0	0.000000	(14, 26) -> (14, 28)
test/test_017.md:14:28:	rparen	")"	false	0	0.000000	(14, 28) -> (14, 29)
test/test_017.md:14:29:	rparen	")"	false	0	0.000000	(14, 29) -> (14, 30)
test/test_017.md:14:30:	semicolon	";"	false	0	0.000000	(14, 30) -> (14, 31)
test/test_017.md:14:32:	ident	"Texts"	false	0	0.000000	(14, 32) -> (14, 37)
test/test_017.md:14:37:	dot	"."	false	0	0.000000	(14, 37) -> (14, 38)
test/test_017.md:14:38:	ident	"WriteLn"	false	0	0.000000	(14, 38) -> (14, 45)
test/test_017.md:14:45:	semicolon	";"	false	0	0.000000	(14, 45) -> (14, 46)
test/test_017.md:15:4:	ident	"Texts"	false	0	0.000000	(15, 4) -> (15, 9)
test/test_017.md:15:9:	dot	"."	false	0	0.000000	(15, 9) -> (15, 10)
test/test_017.md:15:10:	ident	"WriteInt"	false	0	0.000000	(15, 10) -> (15, 18)
test/test_017.md:15:18:	lparen	"("	false	0	0.000000	(15, 18) -> (15, 19)
test/test_017.md:15:19:	ident	"ORD"	false	0	0.000000	(15, 19) -> (15, 22)
test/test_017.md:15:22:	lparen	"("	false	0	0.000000	(15, 22) -> (15, 23)
test/test_017.md:15:23:	ident	"xc"	false	0	0.000000	(15, 23) -> (15, 25)
test/test_017.md:15:25:	asterisk	"*"	false	0	0.000000	(15, 25) -> (15, 26)
test/test_017.md:15:26:	ident	"yc"	false	0	0.000000	(15, 26) -> (15, 28)
test/test_017.md:15:28:	rparen	")"	false	0	0.000000	(15, 28) -> (15, 29)
test/test_017.md:15:29:	rparen	")"	false	0	0.000000	(15, 29) -> (15, 30)
test/test_017.md:15:30:	semicolon	";"	false	0	0.000000	(15, 30) -> (15, 31)
test/test_017.md:15:32:	ident	"Texts"	false	0	0.000000	(15, 32) -> (15, 37)
test/test_017.md:15:37:	dot	"."	false	0	0.000000	(15, 37) -> (15, 38)
test/test_017.md:15:38:	ident	"WriteLn"	false	0	0.000000	(15, 38) -> (15, 45)
test/test_017.md:15:45:	semicolon	";"	false	0	0.000000	(15, 45) -> (15, 46)
test/test_017.md:16:4:	ident	"Texts"	false	0	0.000000	(16, 4) -> (16, 9)
test/test_017.md:16:9:	dot	"."	false	0	0.000000	(16, 9) -> (16, 10)
test/test_017.md:16:10:	ident	"WriteInt"	false	0	0.000000	(16, 10) -> (16, 18)
test/test_017.md:16:18:	lparen	"("	false	0	0.000000	(16, 18) -> (16, 19)
test/test_017.md:16:19:	ident	"ORD"	false	0	0.000000	(16, 19) -> (16, 22)
test/test_017.md:16:22:	lparen	"("	false	0	0.000000	(16, 22) -> (16, 23)
test/test_017.md:16:23:	ident	"xc"	false	0	0.000000	(16, 23) -> (16, 25)
test/test_017.md:16:25:	slash	"/"	false	0	0.000000	(16, 25) -> (16, 26)
test/test_017.md:16:26:	ident	"yc"	false	0	0.000000	(16, 26) -> (16, 28)
test/test_017.md:16:28:	rparen	")"	false	0	0.000000	(16, 28) -> (16, 29)
test/test_017.md:16:29:	rparen	")"	false	0	0.000000	(16, 29) -> (16, 30)
test/test_017.md:16:30:	semicolon	";"	false	0	0.000000	(16, 30) -> (16, 31)
test/test_017.md:16:32:	ident	"Texts"	false	0	0.000000	(16, 32) -> (16, 37)
test/test_017.md:16:37:	dot	"."	false	0	0.000000	(16, 37) -> (16, 38)
test/test_017.md:16:38:	ident	"WriteLn"	false	0	0.000000	(16, 38) -> (16, 45)
test/test_017.md:16:45:	semicolon	";"	false	0	0.000000	(16, 45) -> (16, 46)
test/test_017.md:18:4:	ident	"x"	false	0	0.000000	(18, 4) -> (18, 5)
test/test_017.md:18:6:	assign	":="	false	0	0.000000	(18, 6) -> (18, 8)
test/test_017.md:18:9:	ident	"xc"	false	0	0.000000	(18, 9) -> (18, 11)
test/test_017.md:18:11:	semicolon	";"	false	0	0.000000	(18, 11) -> (18, 12)
test/test_017.md:18:13:	ident	"y"	false	0	0.000000	(18, 13) -> (18, 14)
test/test_017.md:18:15:	assign	":="	false	0	0.000000	(18, 15) -> (18, 17)
test/test_017.md:18:18:	ident	"yc"	false	0	0.000000	(18, 18) -> (18, 20)
test/test_017.md:18:20:	semicolon	";"	false	0	0.000000	(18, 20) -> (18, 21)
test/test_017.md:20:4:	ident	"INCL"	false	0	0.000000	(20, 4) -> (20, 8)
test/test_017.md:20:8:	lparen	"("	false	0	0.000000	(20, 8) -> (20, 9)
test/test_017.md:20:9:	ident	"x"	false	0	0.000000	(20, 9) -> (20, 10)
test/test_017.md:20:10:	comma	","	false	0	0.000000	(20, 10) -> (20, 11)
test/test_017.md:20:12:	integer	"4"	false	4	0.000000	(20, 12) -> (20, 13)
test/test_017.md:20:13:	rparen	")"	false	0	0.000000	(20, 13) -> (20, 14)
test/test_017.md:20:14:	semicolon	";"	false	0	0.000000	(20, 14) -> (20, 15)
test/test_017.md:21:4:	ident	"EXCL"	false	0	0.000000	(21, 4) -> (21, 8)
test/test_017.md:21:8:	lparen	"("	false	0	0.000000	(21, 8) -> (21, 9)
test/test_017.md:21:9:	ident	"y"	false	0	0.000000	(21, 9) -> (21, 10)
test/test_017.md:21:10:	comma	","	false	0	0.000000	(21, 10) -> (21, 11)
test/test_017.md:21:12:	integer	"30"	false	30	0.000000	(21, 12) -> (21, 14)
test/test_017.md:21:14:	rparen	")"	false	0	0.000000	(21, 14) -> (21, 15)
test/test_017.md:21:15:	semicolon	";"	false	0	0.000000	(21, 15) -> (21, 16)
test/test_017.md:23:4:	ident	"z"	false	0	0.000000	(23, 4) -> (23, 5)
test/test_017.md:23:6:	assign	":="	false	0	0.000000	(23, 6) -> (23, 8)
test/test_017.md:23:9:	ident	"x"	false	0	0.000000	(23, 9) -> (23, 10)
test/test_017.md:23:11:	plus	"+"	false	0	0.000000	(23, 11) -> (23, 12)
test/test_017.md:23:13:	ident	"y"	false	0	0.000000	(23, 13) -> (23, 14)
test/test_017.md:23:14:	semicolon	";"	false	0	0.000000	(23, 14) -> (23, 15)
test/test_017.md:25:4:	ident	"Texts"	false	0	0.000000	(25, 4) -> (25, 9)
test/test_017.md:25:9:	dot	"."	false	0	0.000000	(25, 9) -> (25, 10)
test/test_017.md:25:10:	ident	"WriteInt"	false	0	0.000000	(25, 10) -> (25, 18)
test/test_017.md:25:18:	lparen	"("	false	0	0.000000	(25, 18) -> (25, 19)
test/test_017.md:25:19:	ident	"ORD"	false	0	0.000000	(25, 19) -> (25, 22)
test/test_017.md:25:22:	lparen	"("	false	0	0.000000	(25, 22) -> (25, 23)
test/test_017.md:25:23:	ident	"z"	false	0	0.000000	(25, 23) -> (25, 24)
test/test_017.md:25:24:	rparen	")"	false	0	0.000000	(25, 24) -> (25, 25)
test/test_017.md:25:25:	rparen	")"	false	0	0.000000	(25, 25) -> (25, 26)
test/test_017.md:25:26:	semicolon	";"	false	0	0.000000	(25, 26) -> (25, 27)
test/test_017.md:25:32:	ident	"Texts"	false	0	0.000000	(25, 32) -> (25, 37)
test/test_017.md:25:37:	dot	"."	false	0	0.000000	(25, 37) -> (25, 38)
test/test_017.md:25:38:	ident	"WriteLn"	false	0	0.000000	(25, 38) -> (25, 45)
test/test_017.md:25:45:	semicolon	";"	false	0	0.000000	(25, 45) -> (25, 46)
test/test_017.md:27:4:	ident	"INCL"	false	0	0.000000	(27, 4) -> (27, 8)
test/test_017.md:27:8:	lparen	"("	false	0	0.000000	(27, 8) -> (27, 9)
test/test_017.md:27:9:	ident	"z"	false	0	0.000000	(27, 9) -> (27, 10)
test/test_017.md:27:10:	comma	","	false	0	0.000000	(27, 10) -> (27, 11)
test/test_017.md:27:12:	integer	"30"	false	30	0.000000	(27, 12) -> (27, 14)
test/test_017.md:27:14:	rparen	")"	false	0	0.000000	(27, 14) -> (27, 15)
test/test_017.md:27:15:	semicolon	";"	false	0	0.000000	(27, 15) -> (27, 16)
test/test_017.md:28:4:	ident	"EXCL"	false	0	0.000000	(28, 4) -> (28, 8)
test/test_017.md:28:8:	lparen	"("	false	0	0.000000	(28, 8) -> (28, 9)
test/test_017.md:28:9:	ident	"z"	false	0	0.000000	(28, 9) -> (28, 10)
test/test_017.md:28:10:	comma	","	false	0	0.000000	(28, 10) -> (28, 11)
test/test_017.md:28:12:	integer	"4"	false	4	0.000000	(28, 12) -> (28, 13)
test/test_017.md:28:13:	rparen	")"	false	0	0.000000	(28, 13) -> (28, 14)
test/test_017.md:28:14:	semicolon	";"	false	0	0.000000	(28, 14) -> (28, 15)
test/test_017.md:30:4:	ident	"Texts"	false	0	0.000000	(30, 4) -> (30, 9)
test/test_017.md:30:9:	dot	"."	false	0	0.000000	(30, 9) -> (30, 10)
test/test_017.md:30:10:	ident	"WriteInt"	false	0	0.000000	(30, 10) -> (30, 18)
test/test_017.md:30:18:	lparen	"("	false	0	0.000000	(30, 18) -> (30, 19)
test/test_017.md:30:19:	ident	"ORD"	false	0	0.000000	(30, 19) -> (30, 22)
test/test_017.md:30:22:	lparen	"("	false	0	0.000000	(30, 22) -> (30, 23)
test/test_017.md:30:23:	ident	"z"	false	0	0.000000	(30, 23) -> (30, 24)
test/test_017.md:30:24:	rparen	")"	false	0	0.000000	(30, 24) -> (30, 25)
test/test_017.md:30:25:	rparen	")"	false	0	0.000000	(30, 25) -> (30, 26)
test/test_017.md:30:26:	semicolon	";"	false	0	0.000000	(30, 26) -> (30, 27)
test/test_017.md:30:32:	ident	"Texts"	false	0	0.000000	(30, 32) -> (30, 37)
test/test_017.md:30:37:	dot	"."	false	0	0.000000	(30, 37) -> (30, 38)
test/test_017.md:30:38:	ident	"WriteLn"	false	0	0.000000	(30, 38) -> (30, 45)
test/test_017.md:30:45:	semicolon	";"	false	0	0.000000	(30, 45) -> (30, 46)
test/test_017.md:32:4:	if	"IF"	false	0	0.000000	(32, 4) -> (32, 6)
test/test_017.md:32:7:	integer	"30"	false	30	0.000000	(32, 7) -> (32, 9)
test/test_017.md:32:10:	in	"IN"	false	0	0.000000	(32, 10) -> (32, 12)
test/test_017.md:32:13:	ident	"z"	false	0	0.000000	(32, 13) -> (32, 14)
test/test_017.md:32:15:	then	"THEN"	false	0	0.000000	(32, 15) -> (32, 19)
test/test_017.md:33:8:	ident	"Texts"	false	0	0.000000	(33, 8) -> (33, 13)
test/test_017.md:33:13:	dot	"."	false	0	0.000000	(33, 13) -> (33, 14)
test/test_017.md:33:14:	ident	"WriteInt"	false	0	0.000000	(33, 14) -> (33, 22)
test/test_017.md:33:22:	lparen	"("	false	0	0.000000	(33, 22) -> (33, 23)
test/test_017.md:33:23:	integer	"30"	false	30	0.000000	(33, 23) -> (33, 25)
test/test_017.md:33:25:	rparen	")"	false	0	0.000000	(33, 25) -> (33, 26)
test/test_017.md:33:26:	semicolon	";"	false	0	0.000000	(33, 26) -> (33, 27)
test/test_017.md:33:32:	ident	"Texts"	false	0	0.000000	(33, 32) -> (33, 37)
test/test_017.md:33:37:	dot	"."	false	0	0.000000	(33, 37) -> (33, 38)
test/test_017.md:33:38:	ident	"WriteLn"	false	0	0.000000	(33, 38) -> (33, 45)
test/test_017.md:33:45:	semicolon	";"	false	0	0.000000	(33, 45) -> (33, 46)
test/test_017.md:34:4:	end	"END"	false	0	0.000000	(34, 4) -> (34, 7)
test/test_017.md:34:7:	semicolon	";"	false	0	0.000000	(34, 7) -> (34, 8)
test/test_017.md:36:4:	if	"IF"	false	0	0.000000	(36, 4) -> (36, 6)
test/test_017.md:36:7:	tilde	"~"	false	0	0.000000	(36, 7) -> (36, 8)
test/test_017.md:36:8:	lparen	"("	false	0	0.000000	(36, 8) -> (36, 9)
test/test_017.md:36:9:	integer	"31"	false	31	0.000000	(36, 9) -> (36, 11)
test/test_017.md:36:12:	in	"IN"	false	0	0.000000	(36, 12) -> (36, 14)
test/test_017.md:36:15:	ident	"z"	false	0	0.000000	(36, 15) -> (36, 16)
test/test_017.md:36:16:	rparen	")"	false	0	0.000000	(36, 16) -> (36, 17)
test/test_017.md:36:18:	then	"THEN"	false	0	0.000000	(36, 18) -> (36, 22)
test/test_017.md:37:8:	ident	"Texts"	false	0	0.000000	(37, 8) -> (37, 13)
test/test_017.md:37:13:	dot	"."	false	0	0.000000	(37, 13) -> (37, 14)
test/test_017.md:37:14:	ident	"WriteInt"	false	0	0.000000	(37, 14) -> (37, 22)
test/test_017.md:37:22:	lparen	"("	false	0	0.000000	(37, 22) -> (37, 23)
test/test_017.md:37:23:	integer	"31"	false	31	0.000000	(37, 23) -> (37, 25)
test/test_017.md:37:25:	rparen	")"	false	0	0.000000	(37, 25) -> (37, 26)
test/test_017.md:37:26:	semicolon	";"	false	0	0.000000	(37, 26) -> (37, 27)
test/test_017.md:37:32:	ident	"Texts"	false	0	0.000000	(37, 32) -> (37, 37)
test/test_017.md:37:37:	dot	"."	false	0	0.000000	(37, 37) -> (37, 38)
test/test_017.md:37:38:	ident	"WriteLn"	false	0	0.000000	(37, 38) -> (37, 45)
test/test_017.md:37:45:	semicolon	";"	false	0	0.000000	(37, 45) -> (37, 46)
test/test_017.md:38:4:	end	"END"	false	0	0.000000	(38, 4) -> (38, 7)
test/test_017.md:38:7:	semicolon	";"	false	0	0.000000	(38, 7) -> (38, 8)
test/test_017.md:40:4:	ident	"Texts"	false	0	0.000000	(40, 4) -> (40, 9)
test/test_017.md:40:9:	dot	"."	false	0	0.000000	(40, 9) -> (40, 10)
test/test_017.md:40:10:	ident	"WriteInt"	false	0	0.000000	(40, 10) -> (40, 18)
test/test_017.md:40:18:	lparen	"("	false	0	0.000000	(40, 18) -> (40, 19)
test/test_017.md:40:19:	ident	"ORD"	false	0	0.000000	(40, 19) -> (40, 22)
test/test_017.md:40:22:	lparen	"("	false	0	0.000000	(40, 22) -> (40, 23)
test/test_017.md:40:23:	lbrace	"{"	false	0	0.000000	(40, 23) -> (40, 24)
test/test_017.md:40:24:	integer	"1"	false	1	0.000000	(40, 24) -> (40, 25)
test/test_017.md:40:25:	comma	","	false	0	0.000000	(40, 25) -> (40, 26)
test/test_017.md:40:27:	integer	"2"	false	2	0.000000	(40, 27) -> (40, 28)
test/test_017.md:40:28:	rbrace	"}"	false	0	0.000000	(40, 28) -> (40, 29)
test/test_017.md:40:29:	rparen	")"	false	0	0.000000	(40, 29) -> (40, 30)
test/test_017.md:40:30:	rparen	")"	false	0	0.000000	(40, 30) -> (40, 31)
test/test_017.md:40:31:	semicolon	";"	false	0	0.000000	(40, 31) -> (40, 32)
test/test_017.md:40:33:	ident	"Texts"	false	0	0.000000	(40, 33) -> (40, 38)
test/test_017.md:40:38:	dot	"."	false	0	0.000000	(40, 38) -> (40, 39)
test/test_017.md:40:39:	ident	"WriteLn"	false	0	0.000000	(40, 39) -> (40, 46)
test/test_017.md:40:46:	semicolon	";"	false	0	0.000000	(40, 46) -> (40, 47)
test/test_017.md:41:0:	end	"END"	false	0	0.000000	(41, 0) -> (41, 3)
test/test_017.md:41:4:	ident	"Sets"	false	0	0.000000	(41, 4) -> (41, 8)
test/test_017.md:41:8:	dot	"."	false	0	0.000000	(41, 8) -> (41, 9)
test/test_017.md:42:0:	eof	""	false	0	0.000000	(42, 0) -> (42, 0)
```
## AST
```scheme
(module "Sets"
  (consts
    (xc [set]
      (set (1..3, 10))
    )
    (yc [set]
      (set (10..19))
    )
  )
  (vars
    (x 
      (SET [set])
    )
    (y 
      (SET [set])
    )
    (z 
      (SET [set])
    )
  )
  (stmts
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteInt")
        (call
          (procedure [i64] "ORD")
          (constant [set] "xc")
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
          (constant [set] "yc")
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
          (plus [set]
            (constant [set] "xc")
            (constant [set] "yc")
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
          (minus [set]
            (constant [set] "xc")
            (constant [set] "yc")
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
          (asterisk [set]
            (constant [set] "xc")
            (constant [set] "yc")
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
          (slash [set]
            (constant [set] "xc")
            (constant [set] "yc")
          )
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
    (assign
      (variable [set] "x")
      (constant [set] "xc")
    )
    (assign
      (variable [set] "y")
      (constant [set] "yc")
    )
    (expr2stmt
      (call
        (procedure [void] "INCL")
        (variable [set] "x")
        (number [i64] 4)
      )
    )
    (expr2stmt
      (call
        (procedure [void] "EXCL")
        (variable [set] "y")
        (number [i64] 30)
      )
    )
    (assign
      (variable [set] "z")
      (plus [set]
        (variable [set] "x")
        (variable [set] "y")
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteInt")
        (call
          (procedure [i64] "ORD")
          (variable [set] "z")
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
        (procedure [void] "INCL")
        (variable [set] "z")
        (number [i64] 30)
      )
    )
    (expr2stmt
      (call
        (procedure [void] "EXCL")
        (variable [set] "z")
        (number [i64] 4)
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteInt")
        (call
          (procedure [i64] "ORD")
          (variable [set] "z")
        )
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteLn")
      )
    )
    (if
      (in [boolean]
        (number [i64] 30)
        (variable [set] "z")
      )
      (stmts
        (expr2stmt
          (call
            (procedure [void] "Texts.WriteInt")
            (number [i64] 30)
          )
        )
        (expr2stmt
          (call
            (procedure [void] "Texts.WriteLn")
          )
        )
      )
      (stmts
      )
    )
    (if
      (not [boolean]
        (in [boolean]
          (number [i64] 31)
          (variable [set] "z")
        )
      )
      (stmts
        (expr2stmt
          (call
            (procedure [void] "Texts.WriteInt")
            (number [i64] 31)
          )
        )
        (expr2stmt
          (call
            (procedure [void] "Texts.WriteLn")
          )
        )
      )
      (stmts
      )
    )
    (expr2stmt
      (call
        (procedure [void] "Texts.WriteInt")
        (call
          (procedure [i64] "ORD")
          (set (1..2))
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
@0 = global i64 0
@1 = global i64 0
@2 = global i64 0
@3 = global [3 x i8] c"%d\00"
@4 = global [1 x i8] c"\00"

declare i64 @puts(i8* %str)

declare i64 @rand()

declare i64 @sprintf(i8* %buf, i8* %format, ...)

declare i64 @printf(i8* %format, ...)

declare i8* @malloc(i64 %size)

declare i8* @free(i8* %ptr)

define i64 @main() {
entry:
	%0 = getelementptr [3 x i8], [3 x i8]* @3, i64 0, i64 0
	%1 = call i64 (i8*, ...) @printf(i8* %0, i64 1038)
	%2 = getelementptr [1 x i8], [1 x i8]* @4, i64 0, i64 0
	%3 = call i64 @puts(i8* %2)
	%4 = getelementptr [3 x i8], [3 x i8]* @3, i64 0, i64 0
	%5 = call i64 (i8*, ...) @printf(i8* %4, i64 1047552)
	%6 = getelementptr [1 x i8], [1 x i8]* @4, i64 0, i64 0
	%7 = call i64 @puts(i8* %6)
	%8 = or i64 1038, 1047552
	%9 = getelementptr [3 x i8], [3 x i8]* @3, i64 0, i64 0
	%10 = call i64 (i8*, ...) @printf(i8* %9, i64 %8)
	%11 = getelementptr [1 x i8], [1 x i8]* @4, i64 0, i64 0
	%12 = call i64 @puts(i8* %11)
	%13 = xor i64 1047552, -1
	%14 = and i64 1038, %13
	%15 = getelementptr [3 x i8], [3 x i8]* @3, i64 0, i64 0
	%16 = call i64 (i8*, ...) @printf(i8* %15, i64 %14)
	%17 = getelementptr [1 x i8], [1 x i8]* @4, i64 0, i64 0
	%18 = call i64 @puts(i8* %17)
	%19 = and i64 1038, 1047552
	%20 = getelementptr [3 x i8], [3 x i8]* @3, i64 0, i64 0
	%21 = call i64 (i8*, ...) @printf(i8* %20, i64 %19)
	%22 = getelementptr [1 x i8], [1 x i8]* @4, i64 0, i64 0
	%23 = call i64 @puts(i8* %22)
	%24 = xor i64 1038, 1047552
	%25 = getelementptr [3 x i8], [3 x i8]* @3, i64 0, i64 0
	%26 = call i64 (i8*, ...) @printf(i8* %25, i64 %24)
	%27 = getelementptr [1 x i8], [1 x i8]* @4, i64 0, i64 0
	%28 = call i64 @puts(i8* %27)
	store i64 1038, i64* @0
	store i64 1047552, i64* @1
	%29 = load i64, i64* @0
	%30 = shl i64 1, 4
	%31 = or i64 %29, %30
	store i64 %31, i64* @0
	%32 = load i64, i64* @1
	%33 = shl i64 1, 30
	%34 = xor i64 %33, -1
	%35 = and i64 %32, %34
	store i64 %35, i64* @1
	%36 = load i64, i64* @0
	%37 = load i64, i64* @1
	%38 = or i64 %36, %37
	store i64 %38, i64* @2
	%39 = load i64, i64* @2
	%40 = getelementptr [3 x i8], [3 x i8]* @3, i64 0, i64 0
	%41 = call i64 (i8*, ...) @printf(i8* %40, i64 %39)
	%42 = getelementptr [1 x i8], [1 x i8]* @4, i64 0, i64 0
	%43 = call i64 @puts(i8* %42)
	%44 = load i64, i64* @2
	%45 = shl i64 1, 30
	%46 = or i64 %44, %45
	store i64 %46, i64* @2
	%47 = load i64, i64* @2
	%48 = shl i64 1, 4
	%49 = xor i64 %48, -1
	%50 = and i64 %47, %49
	store i64 %50, i64* @2
	%51 = load i64, i64* @2
	%52 = getelementptr [3 x i8], [3 x i8]* @3, i64 0, i64 0
	%53 = call i64 (i8*, ...) @printf(i8* %52, i64 %51)
	%54 = getelementptr [1 x i8], [1 x i8]* @4, i64 0, i64 0
	%55 = call i64 @puts(i8* %54)
	%56 = load i64, i64* @2
	%57 = shl i64 1, 30
	%58 = and i64 %56, %57
	%59 = icmp ne i64 %58, 0
	br i1 %59, label %60, label %65

60:
	%61 = getelementptr [3 x i8], [3 x i8]* @3, i64 0, i64 0
	%62 = call i64 (i8*, ...) @printf(i8* %61, i64 30)
	%63 = getelementptr [1 x i8], [1 x i8]* @4, i64 0, i64 0
	%64 = call i64 @puts(i8* %63)
	br label %66

65:
	br label %66

66:
	%67 = load i64, i64* @2
	%68 = shl i64 1, 31
	%69 = and i64 %67, %68
	%70 = icmp ne i64 %69, 0
	%71 = icmp eq i1 %70, false
	br i1 %71, label %72, label %77

72:
	%73 = getelementptr [3 x i8], [3 x i8]* @3, i64 0, i64 0
	%74 = call i64 (i8*, ...) @printf(i8* %73, i64 31)
	%75 = getelementptr [1 x i8], [1 x i8]* @4, i64 0, i64 0
	%76 = call i64 @puts(i8* %75)
	br label %78

77:
	br label %78

78:
	%79 = getelementptr [3 x i8], [3 x i8]* @3, i64 0, i64 0
	%80 = call i64 (i8*, ...) @printf(i8* %79, i64 6)
	%81 = getelementptr [1 x i8], [1 x i8]* @4, i64 0, i64 0
	%82 = call i64 @puts(i8* %81)
	ret i64 0
}

```
## Run
```bash
1038
1047552
1047566
14
1024
1046542
1047582
1074789390
30
31
6
```
