# test/test_022.md
## Source
```
print(~TRUE);
print(~FALSE)
```
## Tokens
```tsv
test/test_022.md:1:1:	ident	"print"	false	0	0.000000	(1, 1) -> (1, 6)
test/test_022.md:1:6:	lparen	"("	false	0	0.000000	(1, 6) -> (1, 7)
test/test_022.md:1:7:	tilde	"~"	false	0	0.000000	(1, 7) -> (1, 8)
test/test_022.md:1:8:	boolean	"TRUE"	true	0	0.000000	(1, 8) -> (1, 12)
test/test_022.md:1:12:	rparen	")"	false	0	0.000000	(1, 12) -> (1, 13)
test/test_022.md:1:13:	semicolon	";"	false	0	0.000000	(1, 13) -> (1, 14)
test/test_022.md:2:0:	ident	"print"	false	0	0.000000	(2, 0) -> (2, 5)
test/test_022.md:2:5:	lparen	"("	false	0	0.000000	(2, 5) -> (2, 6)
test/test_022.md:2:6:	tilde	"~"	false	0	0.000000	(2, 6) -> (2, 7)
test/test_022.md:2:7:	boolean	"FALSE"	false	0	0.000000	(2, 7) -> (2, 12)
test/test_022.md:2:12:	rparen	")"	false	0	0.000000	(2, 12) -> (2, 13)
test/test_022.md:3:0:	eof	""	false	0	0.000000	(3, 0) -> (3, 0)
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
	ret i32 0
}

```
## Run
```bash
FALSE
TRUE
```
