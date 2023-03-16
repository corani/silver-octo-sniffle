# test/test_024.md
## Source
```pascal
MODULE ParseProcedures;

PROCEDURE NoArgs;
BEGIN END NoArgs;

PROCEDURE OneArg(x : INTEGER);
BEGIN END OneArg;

PROCEDURE Return(x : INTEGER) : REAL;
BEGIN 
    RETURN 0
END Return;

PROCEDURE MoreArg(x, y : INTEGER; z : REAL);
BEGIN END MoreArg;

PROCEDURE VarArg(VAR x : INTEGER);
BEGIN END VarArg;

PROCEDURE WithDecls;
VAR x : INTEGER;
BEGIN END WithDecls;

BEGIN
END ParseProcedures.
```
## Tokens
```tsv
test/test_024.md:1:1:	module	"MODULE"	false	0	0.000000	(1, 1) -> (1, 7)
test/test_024.md:1:8:	ident	"ParseProcedures"	false	0	0.000000	(1, 8) -> (1, 23)
test/test_024.md:1:23:	semicolon	";"	false	0	0.000000	(1, 23) -> (1, 24)
test/test_024.md:3:0:	procedure	"PROCEDURE"	false	0	0.000000	(3, 0) -> (3, 9)
test/test_024.md:3:10:	ident	"NoArgs"	false	0	0.000000	(3, 10) -> (3, 16)
test/test_024.md:3:16:	semicolon	";"	false	0	0.000000	(3, 16) -> (3, 17)
test/test_024.md:4:0:	begin	"BEGIN"	false	0	0.000000	(4, 0) -> (4, 5)
test/test_024.md:4:6:	end	"END"	false	0	0.000000	(4, 6) -> (4, 9)
test/test_024.md:4:10:	ident	"NoArgs"	false	0	0.000000	(4, 10) -> (4, 16)
test/test_024.md:4:16:	semicolon	";"	false	0	0.000000	(4, 16) -> (4, 17)
test/test_024.md:6:0:	procedure	"PROCEDURE"	false	0	0.000000	(6, 0) -> (6, 9)
test/test_024.md:6:10:	ident	"OneArg"	false	0	0.000000	(6, 10) -> (6, 16)
test/test_024.md:6:16:	lparen	"("	false	0	0.000000	(6, 16) -> (6, 17)
test/test_024.md:6:17:	ident	"x"	false	0	0.000000	(6, 17) -> (6, 18)
test/test_024.md:6:19:	colon	":"	false	0	0.000000	(6, 19) -> (6, 20)
test/test_024.md:6:21:	ident	"INTEGER"	false	0	0.000000	(6, 21) -> (6, 28)
test/test_024.md:6:28:	rparen	")"	false	0	0.000000	(6, 28) -> (6, 29)
test/test_024.md:6:29:	semicolon	";"	false	0	0.000000	(6, 29) -> (6, 30)
test/test_024.md:7:0:	begin	"BEGIN"	false	0	0.000000	(7, 0) -> (7, 5)
test/test_024.md:7:6:	end	"END"	false	0	0.000000	(7, 6) -> (7, 9)
test/test_024.md:7:10:	ident	"OneArg"	false	0	0.000000	(7, 10) -> (7, 16)
test/test_024.md:7:16:	semicolon	";"	false	0	0.000000	(7, 16) -> (7, 17)
test/test_024.md:9:0:	procedure	"PROCEDURE"	false	0	0.000000	(9, 0) -> (9, 9)
test/test_024.md:9:10:	ident	"Return"	false	0	0.000000	(9, 10) -> (9, 16)
test/test_024.md:9:16:	lparen	"("	false	0	0.000000	(9, 16) -> (9, 17)
test/test_024.md:9:17:	ident	"x"	false	0	0.000000	(9, 17) -> (9, 18)
test/test_024.md:9:19:	colon	":"	false	0	0.000000	(9, 19) -> (9, 20)
test/test_024.md:9:21:	ident	"INTEGER"	false	0	0.000000	(9, 21) -> (9, 28)
test/test_024.md:9:28:	rparen	")"	false	0	0.000000	(9, 28) -> (9, 29)
test/test_024.md:9:30:	colon	":"	false	0	0.000000	(9, 30) -> (9, 31)
test/test_024.md:9:32:	ident	"REAL"	false	0	0.000000	(9, 32) -> (9, 36)
test/test_024.md:9:36:	semicolon	";"	false	0	0.000000	(9, 36) -> (9, 37)
test/test_024.md:10:0:	begin	"BEGIN"	false	0	0.000000	(10, 0) -> (10, 5)
test/test_024.md:11:4:	return	"RETURN"	false	0	0.000000	(11, 4) -> (11, 10)
test/test_024.md:11:11:	integer	"0"	false	0	0.000000	(11, 11) -> (11, 12)
test/test_024.md:12:0:	end	"END"	false	0	0.000000	(12, 0) -> (12, 3)
test/test_024.md:12:4:	ident	"Return"	false	0	0.000000	(12, 4) -> (12, 10)
test/test_024.md:12:10:	semicolon	";"	false	0	0.000000	(12, 10) -> (12, 11)
test/test_024.md:14:0:	procedure	"PROCEDURE"	false	0	0.000000	(14, 0) -> (14, 9)
test/test_024.md:14:10:	ident	"MoreArg"	false	0	0.000000	(14, 10) -> (14, 17)
test/test_024.md:14:17:	lparen	"("	false	0	0.000000	(14, 17) -> (14, 18)
test/test_024.md:14:18:	ident	"x"	false	0	0.000000	(14, 18) -> (14, 19)
test/test_024.md:14:19:	comma	","	false	0	0.000000	(14, 19) -> (14, 20)
test/test_024.md:14:21:	ident	"y"	false	0	0.000000	(14, 21) -> (14, 22)
test/test_024.md:14:23:	colon	":"	false	0	0.000000	(14, 23) -> (14, 24)
test/test_024.md:14:25:	ident	"INTEGER"	false	0	0.000000	(14, 25) -> (14, 32)
test/test_024.md:14:32:	semicolon	";"	false	0	0.000000	(14, 32) -> (14, 33)
test/test_024.md:14:34:	ident	"z"	false	0	0.000000	(14, 34) -> (14, 35)
test/test_024.md:14:36:	colon	":"	false	0	0.000000	(14, 36) -> (14, 37)
test/test_024.md:14:38:	ident	"REAL"	false	0	0.000000	(14, 38) -> (14, 42)
test/test_024.md:14:42:	rparen	")"	false	0	0.000000	(14, 42) -> (14, 43)
test/test_024.md:14:43:	semicolon	";"	false	0	0.000000	(14, 43) -> (14, 44)
test/test_024.md:15:0:	begin	"BEGIN"	false	0	0.000000	(15, 0) -> (15, 5)
test/test_024.md:15:6:	end	"END"	false	0	0.000000	(15, 6) -> (15, 9)
test/test_024.md:15:10:	ident	"MoreArg"	false	0	0.000000	(15, 10) -> (15, 17)
test/test_024.md:15:17:	semicolon	";"	false	0	0.000000	(15, 17) -> (15, 18)
test/test_024.md:17:0:	procedure	"PROCEDURE"	false	0	0.000000	(17, 0) -> (17, 9)
test/test_024.md:17:10:	ident	"VarArg"	false	0	0.000000	(17, 10) -> (17, 16)
test/test_024.md:17:16:	lparen	"("	false	0	0.000000	(17, 16) -> (17, 17)
test/test_024.md:17:17:	var	"VAR"	false	0	0.000000	(17, 17) -> (17, 20)
test/test_024.md:17:21:	ident	"x"	false	0	0.000000	(17, 21) -> (17, 22)
test/test_024.md:17:23:	colon	":"	false	0	0.000000	(17, 23) -> (17, 24)
test/test_024.md:17:25:	ident	"INTEGER"	false	0	0.000000	(17, 25) -> (17, 32)
test/test_024.md:17:32:	rparen	")"	false	0	0.000000	(17, 32) -> (17, 33)
test/test_024.md:17:33:	semicolon	";"	false	0	0.000000	(17, 33) -> (17, 34)
test/test_024.md:18:0:	begin	"BEGIN"	false	0	0.000000	(18, 0) -> (18, 5)
test/test_024.md:18:6:	end	"END"	false	0	0.000000	(18, 6) -> (18, 9)
test/test_024.md:18:10:	ident	"VarArg"	false	0	0.000000	(18, 10) -> (18, 16)
test/test_024.md:18:16:	semicolon	";"	false	0	0.000000	(18, 16) -> (18, 17)
test/test_024.md:20:0:	procedure	"PROCEDURE"	false	0	0.000000	(20, 0) -> (20, 9)
test/test_024.md:20:10:	ident	"WithDecls"	false	0	0.000000	(20, 10) -> (20, 19)
test/test_024.md:20:19:	semicolon	";"	false	0	0.000000	(20, 19) -> (20, 20)
test/test_024.md:21:0:	var	"VAR"	false	0	0.000000	(21, 0) -> (21, 3)
test/test_024.md:21:4:	ident	"x"	false	0	0.000000	(21, 4) -> (21, 5)
test/test_024.md:21:6:	colon	":"	false	0	0.000000	(21, 6) -> (21, 7)
test/test_024.md:21:8:	ident	"INTEGER"	false	0	0.000000	(21, 8) -> (21, 15)
test/test_024.md:21:15:	semicolon	";"	false	0	0.000000	(21, 15) -> (21, 16)
test/test_024.md:22:0:	begin	"BEGIN"	false	0	0.000000	(22, 0) -> (22, 5)
test/test_024.md:22:6:	end	"END"	false	0	0.000000	(22, 6) -> (22, 9)
test/test_024.md:22:10:	ident	"WithDecls"	false	0	0.000000	(22, 10) -> (22, 19)
test/test_024.md:22:19:	semicolon	";"	false	0	0.000000	(22, 19) -> (22, 20)
test/test_024.md:24:0:	begin	"BEGIN"	false	0	0.000000	(24, 0) -> (24, 5)
test/test_024.md:25:0:	end	"END"	false	0	0.000000	(25, 0) -> (25, 3)
test/test_024.md:25:4:	ident	"ParseProcedures"	false	0	0.000000	(25, 4) -> (25, 19)
test/test_024.md:25:19:	dot	"."	false	0	0.000000	(25, 19) -> (25, 20)
test/test_024.md:26:0:	eof	""	false	0	0.000000	(26, 0) -> (26, 0)
```
## AST
```scheme
(module "ParseProcedures"
  (procs
    (NoArgs [
        VOID
      ]
      (stmts
      )
    )
    (OneArg [
        VOID
      ]
      (stmts
      )
    )
    (Return [
        (typeref REAL)
      ]
      (stmts
        (return
          (number [i64] 0)
        )
      )
    )
    (MoreArg [
        VOID
      ]
      (stmts
      )
    )
    (VarArg [
        VOID
      ]
      (stmts
      )
    )
    (WithDecls [
        VOID
      ]
      (vars
        (x 
          (typeref INTEGER)
        )
      )
      (stmts
      )
    )
  )
  (stmts
  )
)
```
## IR
```llvm
@__argc = global i64 0
@__argv = global i8** inttoptr (i8 0 to i8**)

declare i64 @puts(i8* %str)

declare i64 @rand()

declare i64 @sprintf(i8* %buf, i8* %format, ...)

declare i64 @printf(i8* %format, ...)

declare i8* @malloc(i64 %size)

declare i8* @free(i8* %ptr)

define void @oberonMain() {
entry:
	ret void
}

define i64 @main(i64 %argc, i8** %argv) {
entry:
	store i64 %argc, i64* @__argc
	store i8** %argv, i8*** @__argv
	call void @oberonMain()
	ret i64 0
}

```
## Run
```bash
```
