# test/test_004.md
## Source
```
(* boolean operators *)
print(~TRUE);
print(~FALSE);
print(TRUE & FALSE);
print(TRUE OR FALSE)
```
## Tokens
```tsv
test/test_004.md:2:0:	ident	"print"	false	0	0.000000	(2, 0) -> (2, 5)
test/test_004.md:2:5:	lparen	"("	false	0	0.000000	(2, 5) -> (2, 6)
test/test_004.md:2:6:	tilde	"~"	false	0	0.000000	(2, 6) -> (2, 7)
test/test_004.md:2:7:	boolean	"TRUE"	true	0	0.000000	(2, 7) -> (2, 11)
test/test_004.md:2:11:	rparen	")"	false	0	0.000000	(2, 11) -> (2, 12)
test/test_004.md:2:12:	semicolon	";"	false	0	0.000000	(2, 12) -> (2, 13)
test/test_004.md:3:0:	ident	"print"	false	0	0.000000	(3, 0) -> (3, 5)
test/test_004.md:3:5:	lparen	"("	false	0	0.000000	(3, 5) -> (3, 6)
test/test_004.md:3:6:	tilde	"~"	false	0	0.000000	(3, 6) -> (3, 7)
test/test_004.md:3:7:	boolean	"FALSE"	false	0	0.000000	(3, 7) -> (3, 12)
test/test_004.md:3:12:	rparen	")"	false	0	0.000000	(3, 12) -> (3, 13)
test/test_004.md:3:13:	semicolon	";"	false	0	0.000000	(3, 13) -> (3, 14)
test/test_004.md:4:0:	ident	"print"	false	0	0.000000	(4, 0) -> (4, 5)
test/test_004.md:4:5:	lparen	"("	false	0	0.000000	(4, 5) -> (4, 6)
test/test_004.md:4:6:	boolean	"TRUE"	true	0	0.000000	(4, 6) -> (4, 10)
test/test_004.md:4:11:	ampersand	"&"	false	0	0.000000	(4, 11) -> (4, 12)
test/test_004.md:4:13:	boolean	"FALSE"	false	0	0.000000	(4, 13) -> (4, 18)
test/test_004.md:4:18:	rparen	")"	false	0	0.000000	(4, 18) -> (4, 19)
test/test_004.md:4:19:	semicolon	";"	false	0	0.000000	(4, 19) -> (4, 20)
test/test_004.md:5:0:	ident	"print"	false	0	0.000000	(5, 0) -> (5, 5)
test/test_004.md:5:5:	lparen	"("	false	0	0.000000	(5, 5) -> (5, 6)
test/test_004.md:5:6:	boolean	"TRUE"	true	0	0.000000	(5, 6) -> (5, 10)
test/test_004.md:5:11:	or	"OR"	false	0	0.000000	(5, 11) -> (5, 13)
test/test_004.md:5:14:	boolean	"FALSE"	false	0	0.000000	(5, 14) -> (5, 19)
test/test_004.md:5:19:	rparen	")"	false	0	0.000000	(5, 19) -> (5, 20)
test/test_004.md:6:0:	eof	""	false	0	0.000000	(6, 0) -> (6, 0)
```
## AST
```scheme
(module
  (expr2stmt
    (print [void]
      (not [boolean]
        #true
      )
    )
  )
  (expr2stmt
    (print [void]
      (not [boolean]
        #false
      )
    )
  )
  (expr2stmt
    (print [void]
      (ampersand [boolean]
        #true
        #false
      )
    )
  )
  (expr2stmt
    (print [void]
      (or [boolean]
        #true
        #false
      )
    )
  )
)
```
## IR
```llvm
@0 = global [5 x i8] c"TRUE\00"
@1 = global [6 x i8] c"FALSE\00"

declare i32 @puts(i8* %str)

declare i32 @rand()

declare i32 @sprintf(i8* %buf, i8* %format, ...)

declare i32 @printf(i8* %format, ...)

define i32 @main() {
entry:
	%0 = icmp eq i1 true, false
	br i1 %0, label %1, label %3

1:
	%2 = getelementptr [5 x i8], [5 x i8]* @0, i32 0, i32 0
	br label %5

3:
	%4 = getelementptr [6 x i8], [6 x i8]* @1, i32 0, i32 0
	br label %5

5:
	%6 = phi i8* [ %2, %1 ], [ %4, %3 ]
	%7 = call i32 @puts(i8* %6)
	%8 = icmp eq i1 false, false
	br i1 %8, label %9, label %11

9:
	%10 = getelementptr [5 x i8], [5 x i8]* @0, i32 0, i32 0
	br label %13

11:
	%12 = getelementptr [6 x i8], [6 x i8]* @1, i32 0, i32 0
	br label %13

13:
	%14 = phi i8* [ %10, %9 ], [ %12, %11 ]
	%15 = call i32 @puts(i8* %14)
	%16 = and i1 true, false
	br i1 %16, label %17, label %19

17:
	%18 = getelementptr [5 x i8], [5 x i8]* @0, i32 0, i32 0
	br label %21

19:
	%20 = getelementptr [6 x i8], [6 x i8]* @1, i32 0, i32 0
	br label %21

21:
	%22 = phi i8* [ %18, %17 ], [ %20, %19 ]
	%23 = call i32 @puts(i8* %22)
	%24 = or i1 true, false
	br i1 %24, label %25, label %27

25:
	%26 = getelementptr [5 x i8], [5 x i8]* @0, i32 0, i32 0
	br label %29

27:
	%28 = getelementptr [6 x i8], [6 x i8]* @1, i32 0, i32 0
	br label %29

29:
	%30 = phi i8* [ %26, %25 ], [ %28, %27 ]
	%31 = call i32 @puts(i8* %30)
	ret i32 0
}

```
## Run
```bash
FALSE
TRUE
FALSE
TRUE
```
