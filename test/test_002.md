# test/test_002.md
## Source
```pascal
(* print binary operations *)
print(-42);
print( 34 + 35 + 42000 );
print(66-33);
print( -34 + 35 );
print( (34 + 35) + 42000 );
print( (34 + 35) * 33 );
print( (34 + 35) / 33 );
print(11 DIV 3);
print(11 MOD 3);
print(11 / 3 / 2);
print((11 / 3) + (7 / 2));
print((11/3)-(7/2));
print(11/3*3);
print(69+0.420);
print(0.69E+2+420.E-3)
```
## Tokens
```tsv
test/test_002.md:2:0:	ident	"print"	false	0	0.000000	(2, 0) -> (2, 5)
test/test_002.md:2:5:	lparen	"("	false	0	0.000000	(2, 5) -> (2, 6)
test/test_002.md:2:6:	minus	"-"	false	0	0.000000	(2, 6) -> (2, 7)
test/test_002.md:2:7:	integer	"42"	false	42	0.000000	(2, 7) -> (2, 9)
test/test_002.md:2:9:	rparen	")"	false	0	0.000000	(2, 9) -> (2, 10)
test/test_002.md:2:10:	semicolon	";"	false	0	0.000000	(2, 10) -> (2, 11)
test/test_002.md:3:0:	ident	"print"	false	0	0.000000	(3, 0) -> (3, 5)
test/test_002.md:3:5:	lparen	"("	false	0	0.000000	(3, 5) -> (3, 6)
test/test_002.md:3:7:	integer	"34"	false	34	0.000000	(3, 7) -> (3, 9)
test/test_002.md:3:10:	plus	"+"	false	0	0.000000	(3, 10) -> (3, 11)
test/test_002.md:3:12:	integer	"35"	false	35	0.000000	(3, 12) -> (3, 14)
test/test_002.md:3:15:	plus	"+"	false	0	0.000000	(3, 15) -> (3, 16)
test/test_002.md:3:17:	integer	"42000"	false	42000	0.000000	(3, 17) -> (3, 22)
test/test_002.md:3:23:	rparen	")"	false	0	0.000000	(3, 23) -> (3, 24)
test/test_002.md:3:24:	semicolon	";"	false	0	0.000000	(3, 24) -> (3, 25)
test/test_002.md:4:0:	ident	"print"	false	0	0.000000	(4, 0) -> (4, 5)
test/test_002.md:4:5:	lparen	"("	false	0	0.000000	(4, 5) -> (4, 6)
test/test_002.md:4:6:	integer	"66"	false	66	0.000000	(4, 6) -> (4, 8)
test/test_002.md:4:8:	minus	"-"	false	0	0.000000	(4, 8) -> (4, 9)
test/test_002.md:4:9:	integer	"33"	false	33	0.000000	(4, 9) -> (4, 11)
test/test_002.md:4:11:	rparen	")"	false	0	0.000000	(4, 11) -> (4, 12)
test/test_002.md:4:12:	semicolon	";"	false	0	0.000000	(4, 12) -> (4, 13)
test/test_002.md:5:0:	ident	"print"	false	0	0.000000	(5, 0) -> (5, 5)
test/test_002.md:5:5:	lparen	"("	false	0	0.000000	(5, 5) -> (5, 6)
test/test_002.md:5:7:	minus	"-"	false	0	0.000000	(5, 7) -> (5, 8)
test/test_002.md:5:8:	integer	"34"	false	34	0.000000	(5, 8) -> (5, 10)
test/test_002.md:5:11:	plus	"+"	false	0	0.000000	(5, 11) -> (5, 12)
test/test_002.md:5:13:	integer	"35"	false	35	0.000000	(5, 13) -> (5, 15)
test/test_002.md:5:16:	rparen	")"	false	0	0.000000	(5, 16) -> (5, 17)
test/test_002.md:5:17:	semicolon	";"	false	0	0.000000	(5, 17) -> (5, 18)
test/test_002.md:6:0:	ident	"print"	false	0	0.000000	(6, 0) -> (6, 5)
test/test_002.md:6:5:	lparen	"("	false	0	0.000000	(6, 5) -> (6, 6)
test/test_002.md:6:7:	lparen	"("	false	0	0.000000	(6, 7) -> (6, 8)
test/test_002.md:6:8:	integer	"34"	false	34	0.000000	(6, 8) -> (6, 10)
test/test_002.md:6:11:	plus	"+"	false	0	0.000000	(6, 11) -> (6, 12)
test/test_002.md:6:13:	integer	"35"	false	35	0.000000	(6, 13) -> (6, 15)
test/test_002.md:6:15:	rparen	")"	false	0	0.000000	(6, 15) -> (6, 16)
test/test_002.md:6:17:	plus	"+"	false	0	0.000000	(6, 17) -> (6, 18)
test/test_002.md:6:19:	integer	"42000"	false	42000	0.000000	(6, 19) -> (6, 24)
test/test_002.md:6:25:	rparen	")"	false	0	0.000000	(6, 25) -> (6, 26)
test/test_002.md:6:26:	semicolon	";"	false	0	0.000000	(6, 26) -> (6, 27)
test/test_002.md:7:0:	ident	"print"	false	0	0.000000	(7, 0) -> (7, 5)
test/test_002.md:7:5:	lparen	"("	false	0	0.000000	(7, 5) -> (7, 6)
test/test_002.md:7:7:	lparen	"("	false	0	0.000000	(7, 7) -> (7, 8)
test/test_002.md:7:8:	integer	"34"	false	34	0.000000	(7, 8) -> (7, 10)
test/test_002.md:7:11:	plus	"+"	false	0	0.000000	(7, 11) -> (7, 12)
test/test_002.md:7:13:	integer	"35"	false	35	0.000000	(7, 13) -> (7, 15)
test/test_002.md:7:15:	rparen	")"	false	0	0.000000	(7, 15) -> (7, 16)
test/test_002.md:7:17:	asterisk	"*"	false	0	0.000000	(7, 17) -> (7, 18)
test/test_002.md:7:19:	integer	"33"	false	33	0.000000	(7, 19) -> (7, 21)
test/test_002.md:7:22:	rparen	")"	false	0	0.000000	(7, 22) -> (7, 23)
test/test_002.md:7:23:	semicolon	";"	false	0	0.000000	(7, 23) -> (7, 24)
test/test_002.md:8:0:	ident	"print"	false	0	0.000000	(8, 0) -> (8, 5)
test/test_002.md:8:5:	lparen	"("	false	0	0.000000	(8, 5) -> (8, 6)
test/test_002.md:8:7:	lparen	"("	false	0	0.000000	(8, 7) -> (8, 8)
test/test_002.md:8:8:	integer	"34"	false	34	0.000000	(8, 8) -> (8, 10)
test/test_002.md:8:11:	plus	"+"	false	0	0.000000	(8, 11) -> (8, 12)
test/test_002.md:8:13:	integer	"35"	false	35	0.000000	(8, 13) -> (8, 15)
test/test_002.md:8:15:	rparen	")"	false	0	0.000000	(8, 15) -> (8, 16)
test/test_002.md:8:17:	slash	"/"	false	0	0.000000	(8, 17) -> (8, 18)
test/test_002.md:8:19:	integer	"33"	false	33	0.000000	(8, 19) -> (8, 21)
test/test_002.md:8:22:	rparen	")"	false	0	0.000000	(8, 22) -> (8, 23)
test/test_002.md:8:23:	semicolon	";"	false	0	0.000000	(8, 23) -> (8, 24)
test/test_002.md:9:0:	ident	"print"	false	0	0.000000	(9, 0) -> (9, 5)
test/test_002.md:9:5:	lparen	"("	false	0	0.000000	(9, 5) -> (9, 6)
test/test_002.md:9:6:	integer	"11"	false	11	0.000000	(9, 6) -> (9, 8)
test/test_002.md:9:9:	div	"DIV"	false	0	0.000000	(9, 9) -> (9, 12)
test/test_002.md:9:13:	integer	"3"	false	3	0.000000	(9, 13) -> (9, 14)
test/test_002.md:9:14:	rparen	")"	false	0	0.000000	(9, 14) -> (9, 15)
test/test_002.md:9:15:	semicolon	";"	false	0	0.000000	(9, 15) -> (9, 16)
test/test_002.md:10:0:	ident	"print"	false	0	0.000000	(10, 0) -> (10, 5)
test/test_002.md:10:5:	lparen	"("	false	0	0.000000	(10, 5) -> (10, 6)
test/test_002.md:10:6:	integer	"11"	false	11	0.000000	(10, 6) -> (10, 8)
test/test_002.md:10:9:	mod	"MOD"	false	0	0.000000	(10, 9) -> (10, 12)
test/test_002.md:10:13:	integer	"3"	false	3	0.000000	(10, 13) -> (10, 14)
test/test_002.md:10:14:	rparen	")"	false	0	0.000000	(10, 14) -> (10, 15)
test/test_002.md:10:15:	semicolon	";"	false	0	0.000000	(10, 15) -> (10, 16)
test/test_002.md:11:0:	ident	"print"	false	0	0.000000	(11, 0) -> (11, 5)
test/test_002.md:11:5:	lparen	"("	false	0	0.000000	(11, 5) -> (11, 6)
test/test_002.md:11:6:	integer	"11"	false	11	0.000000	(11, 6) -> (11, 8)
test/test_002.md:11:9:	slash	"/"	false	0	0.000000	(11, 9) -> (11, 10)
test/test_002.md:11:11:	integer	"3"	false	3	0.000000	(11, 11) -> (11, 12)
test/test_002.md:11:13:	slash	"/"	false	0	0.000000	(11, 13) -> (11, 14)
test/test_002.md:11:15:	integer	"2"	false	2	0.000000	(11, 15) -> (11, 16)
test/test_002.md:11:16:	rparen	")"	false	0	0.000000	(11, 16) -> (11, 17)
test/test_002.md:11:17:	semicolon	";"	false	0	0.000000	(11, 17) -> (11, 18)
test/test_002.md:12:0:	ident	"print"	false	0	0.000000	(12, 0) -> (12, 5)
test/test_002.md:12:5:	lparen	"("	false	0	0.000000	(12, 5) -> (12, 6)
test/test_002.md:12:6:	lparen	"("	false	0	0.000000	(12, 6) -> (12, 7)
test/test_002.md:12:7:	integer	"11"	false	11	0.000000	(12, 7) -> (12, 9)
test/test_002.md:12:10:	slash	"/"	false	0	0.000000	(12, 10) -> (12, 11)
test/test_002.md:12:12:	integer	"3"	false	3	0.000000	(12, 12) -> (12, 13)
test/test_002.md:12:13:	rparen	")"	false	0	0.000000	(12, 13) -> (12, 14)
test/test_002.md:12:15:	plus	"+"	false	0	0.000000	(12, 15) -> (12, 16)
test/test_002.md:12:17:	lparen	"("	false	0	0.000000	(12, 17) -> (12, 18)
test/test_002.md:12:18:	integer	"7"	false	7	0.000000	(12, 18) -> (12, 19)
test/test_002.md:12:20:	slash	"/"	false	0	0.000000	(12, 20) -> (12, 21)
test/test_002.md:12:22:	integer	"2"	false	2	0.000000	(12, 22) -> (12, 23)
test/test_002.md:12:23:	rparen	")"	false	0	0.000000	(12, 23) -> (12, 24)
test/test_002.md:12:24:	rparen	")"	false	0	0.000000	(12, 24) -> (12, 25)
test/test_002.md:12:25:	semicolon	";"	false	0	0.000000	(12, 25) -> (12, 26)
test/test_002.md:13:0:	ident	"print"	false	0	0.000000	(13, 0) -> (13, 5)
test/test_002.md:13:5:	lparen	"("	false	0	0.000000	(13, 5) -> (13, 6)
test/test_002.md:13:6:	lparen	"("	false	0	0.000000	(13, 6) -> (13, 7)
test/test_002.md:13:7:	integer	"11"	false	11	0.000000	(13, 7) -> (13, 9)
test/test_002.md:13:9:	slash	"/"	false	0	0.000000	(13, 9) -> (13, 10)
test/test_002.md:13:10:	integer	"3"	false	3	0.000000	(13, 10) -> (13, 11)
test/test_002.md:13:11:	rparen	")"	false	0	0.000000	(13, 11) -> (13, 12)
test/test_002.md:13:12:	minus	"-"	false	0	0.000000	(13, 12) -> (13, 13)
test/test_002.md:13:13:	lparen	"("	false	0	0.000000	(13, 13) -> (13, 14)
test/test_002.md:13:14:	integer	"7"	false	7	0.000000	(13, 14) -> (13, 15)
test/test_002.md:13:15:	slash	"/"	false	0	0.000000	(13, 15) -> (13, 16)
test/test_002.md:13:16:	integer	"2"	false	2	0.000000	(13, 16) -> (13, 17)
test/test_002.md:13:17:	rparen	")"	false	0	0.000000	(13, 17) -> (13, 18)
test/test_002.md:13:18:	rparen	")"	false	0	0.000000	(13, 18) -> (13, 19)
test/test_002.md:13:19:	semicolon	";"	false	0	0.000000	(13, 19) -> (13, 20)
test/test_002.md:14:0:	ident	"print"	false	0	0.000000	(14, 0) -> (14, 5)
test/test_002.md:14:5:	lparen	"("	false	0	0.000000	(14, 5) -> (14, 6)
test/test_002.md:14:6:	integer	"11"	false	11	0.000000	(14, 6) -> (14, 8)
test/test_002.md:14:8:	slash	"/"	false	0	0.000000	(14, 8) -> (14, 9)
test/test_002.md:14:9:	integer	"3"	false	3	0.000000	(14, 9) -> (14, 10)
test/test_002.md:14:10:	asterisk	"*"	false	0	0.000000	(14, 10) -> (14, 11)
test/test_002.md:14:11:	integer	"3"	false	3	0.000000	(14, 11) -> (14, 12)
test/test_002.md:14:12:	rparen	")"	false	0	0.000000	(14, 12) -> (14, 13)
test/test_002.md:14:13:	semicolon	";"	false	0	0.000000	(14, 13) -> (14, 14)
test/test_002.md:15:0:	ident	"print"	false	0	0.000000	(15, 0) -> (15, 5)
test/test_002.md:15:5:	lparen	"("	false	0	0.000000	(15, 5) -> (15, 6)
test/test_002.md:15:6:	integer	"69"	false	69	0.000000	(15, 6) -> (15, 8)
test/test_002.md:15:8:	plus	"+"	false	0	0.000000	(15, 8) -> (15, 9)
test/test_002.md:15:9:	real	"0.420"	false	0	0.420000	(15, 9) -> (15, 14)
test/test_002.md:15:14:	rparen	")"	false	0	0.000000	(15, 14) -> (15, 15)
test/test_002.md:15:15:	semicolon	";"	false	0	0.000000	(15, 15) -> (15, 16)
test/test_002.md:16:0:	ident	"print"	false	0	0.000000	(16, 0) -> (16, 5)
test/test_002.md:16:5:	lparen	"("	false	0	0.000000	(16, 5) -> (16, 6)
test/test_002.md:16:6:	real	"0.69E+2"	false	0	69.000000	(16, 6) -> (16, 13)
test/test_002.md:16:13:	plus	"+"	false	0	0.000000	(16, 13) -> (16, 14)
test/test_002.md:16:14:	real	"420.E-3"	false	0	0.420000	(16, 14) -> (16, 21)
test/test_002.md:16:21:	rparen	")"	false	0	0.000000	(16, 21) -> (16, 22)
test/test_002.md:17:0:	eof	""	false	0	0.000000	(17, 0) -> (17, 0)
```
## AST
```scheme
(module ""
  (stmts
    (expr2stmt
      (print [void]
        (minus [i64]
          (number [i64] 0)
          (number [i64] 42)
        )
      )
    )
    (expr2stmt
      (print [void]
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
      (print [void]
        (minus [i64]
          (number [i64] 66)
          (number [i64] 33)
        )
      )
    )
    (expr2stmt
      (print [void]
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
      (print [void]
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
      (print [void]
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
      (print [void]
        (slash [f64]
          (plus [i64]
            (number [i64] 34)
            (number [i64] 35)
          )
          (number [i64] 33)
        )
      )
    )
    (expr2stmt
      (print [void]
        (div [i64]
          (number [i64] 11)
          (number [i64] 3)
        )
      )
    )
    (expr2stmt
      (print [void]
        (mod [i64]
          (number [i64] 11)
          (number [i64] 3)
        )
      )
    )
    (expr2stmt
      (print [void]
        (slash [f64]
          (slash [f64]
            (number [i64] 11)
            (number [i64] 3)
          )
          (number [i64] 2)
        )
      )
    )
    (expr2stmt
      (print [void]
        (plus [f64]
          (slash [f64]
            (number [i64] 11)
            (number [i64] 3)
          )
          (slash [f64]
            (number [i64] 7)
            (number [i64] 2)
          )
        )
      )
    )
    (expr2stmt
      (print [void]
        (minus [f64]
          (slash [f64]
            (number [i64] 11)
            (number [i64] 3)
          )
          (slash [f64]
            (number [i64] 7)
            (number [i64] 2)
          )
        )
      )
    )
    (expr2stmt
      (print [void]
        (asterisk [f64]
          (slash [f64]
            (number [i64] 11)
            (number [i64] 3)
          )
          (number [i64] 3)
        )
      )
    )
    (expr2stmt
      (print [void]
        (plus [f64]
          (number [i64] 69)
          (number [f64] 0.420000)
        )
      )
    )
    (expr2stmt
      (print [void]
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

define i64 @main() {
entry:
	%0 = sub i64 0, 42
	%1 = getelementptr [4 x i8], [4 x i8]* @0, i64 0, i64 0
	%2 = call i64 (i8*, ...) @printf(i8* %1, i64 %0)
	%3 = add i64 34, 35
	%4 = add i64 %3, 42000
	%5 = getelementptr [4 x i8], [4 x i8]* @0, i64 0, i64 0
	%6 = call i64 (i8*, ...) @printf(i8* %5, i64 %4)
	%7 = sub i64 66, 33
	%8 = getelementptr [4 x i8], [4 x i8]* @0, i64 0, i64 0
	%9 = call i64 (i8*, ...) @printf(i8* %8, i64 %7)
	%10 = sub i64 0, 34
	%11 = add i64 %10, 35
	%12 = getelementptr [4 x i8], [4 x i8]* @0, i64 0, i64 0
	%13 = call i64 (i8*, ...) @printf(i8* %12, i64 %11)
	%14 = add i64 34, 35
	%15 = add i64 %14, 42000
	%16 = getelementptr [4 x i8], [4 x i8]* @0, i64 0, i64 0
	%17 = call i64 (i8*, ...) @printf(i8* %16, i64 %15)
	%18 = add i64 34, 35
	%19 = mul i64 %18, 33
	%20 = getelementptr [4 x i8], [4 x i8]* @0, i64 0, i64 0
	%21 = call i64 (i8*, ...) @printf(i8* %20, i64 %19)
	%22 = add i64 34, 35
	%23 = sitofp i64 %22 to double
	%24 = sitofp i64 33 to double
	%25 = fdiv double %23, %24
	%26 = getelementptr [4 x i8], [4 x i8]* @1, i64 0, i64 0
	%27 = call i64 (i8*, ...) @printf(i8* %26, double %25)
	%28 = sdiv i64 11, 3
	%29 = getelementptr [4 x i8], [4 x i8]* @0, i64 0, i64 0
	%30 = call i64 (i8*, ...) @printf(i8* %29, i64 %28)
	%31 = srem i64 11, 3
	%32 = getelementptr [4 x i8], [4 x i8]* @0, i64 0, i64 0
	%33 = call i64 (i8*, ...) @printf(i8* %32, i64 %31)
	%34 = sitofp i64 11 to double
	%35 = sitofp i64 3 to double
	%36 = fdiv double %34, %35
	%37 = sitofp i64 2 to double
	%38 = fdiv double %36, %37
	%39 = getelementptr [4 x i8], [4 x i8]* @1, i64 0, i64 0
	%40 = call i64 (i8*, ...) @printf(i8* %39, double %38)
	%41 = sitofp i64 11 to double
	%42 = sitofp i64 3 to double
	%43 = fdiv double %41, %42
	%44 = sitofp i64 7 to double
	%45 = sitofp i64 2 to double
	%46 = fdiv double %44, %45
	%47 = fadd double %43, %46
	%48 = getelementptr [4 x i8], [4 x i8]* @1, i64 0, i64 0
	%49 = call i64 (i8*, ...) @printf(i8* %48, double %47)
	%50 = sitofp i64 11 to double
	%51 = sitofp i64 3 to double
	%52 = fdiv double %50, %51
	%53 = sitofp i64 7 to double
	%54 = sitofp i64 2 to double
	%55 = fdiv double %53, %54
	%56 = fsub double %52, %55
	%57 = getelementptr [4 x i8], [4 x i8]* @1, i64 0, i64 0
	%58 = call i64 (i8*, ...) @printf(i8* %57, double %56)
	%59 = sitofp i64 11 to double
	%60 = sitofp i64 3 to double
	%61 = fdiv double %59, %60
	%62 = sitofp i64 3 to double
	%63 = fmul double %61, %62
	%64 = getelementptr [4 x i8], [4 x i8]* @1, i64 0, i64 0
	%65 = call i64 (i8*, ...) @printf(i8* %64, double %63)
	%66 = sitofp i64 69 to double
	%67 = fadd double %66, 0x3FDAE147AE147AE1
	%68 = getelementptr [4 x i8], [4 x i8]* @1, i64 0, i64 0
	%69 = call i64 (i8*, ...) @printf(i8* %68, double %67)
	%70 = fadd double 69.0, 0x3FDAE147AE147AE1
	%71 = getelementptr [4 x i8], [4 x i8]* @1, i64 0, i64 0
	%72 = call i64 (i8*, ...) @printf(i8* %71, double %70)
	ret i64 0
}

```
## Run
```bash
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
