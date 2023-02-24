# test/test_018.md
## Source
```pascal
MODULE LexErrors;

BEGIN
    Texts.WriteReal(1.1EF);
    Texts.WriteString(@33AB);
    Texts.WriteInt(9999999999999999999999999999999999999999999999999999999);
    Texts.WriteString(123X)
END LexErrors.
```
## Lexer errors
```
test/test_018.md:4:20: ERROR: expected '+' or '-' after real exponent: "1.1E"
test/test_018.md:5:22: ERROR: unknown character '@'
test/test_018.md:5:23: ERROR: expected 'H' or 'X' after hex digits: "33AB"
test/test_018.md:6:19: ERROR: strconv.ParseInt: parsing "9999999999999999999999999999999999999999999999999999999": value out of range
test/test_018.md:7:22: ERROR: encoding/hex: odd length hex string
```
