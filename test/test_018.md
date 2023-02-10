# test/test_018.md
## Source
```pascal
MODULE LexErrors;

BEGIN
    print(1.1EF, @33AB, 9999999999999999999999999999999999999999999999999999999);
    print(123X)
END LexErrors.
```
## Lexer errors
```
test/test_018.md:4:10: ERROR: expected '+' or '-' after real exponent: "1.1E"
test/test_018.md:4:17: ERROR: unknown character '@'
test/test_018.md:4:18: ERROR: expected 'H' or 'X' after hex digits: "33AB"
test/test_018.md:4:24: ERROR: strconv.ParseInt: parsing "9999999999999999999999999999999999999999999999999999999": value out of range
test/test_018.md:5:10: ERROR: encoding/hex: odd length hex string
```
