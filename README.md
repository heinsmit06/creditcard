The project MUST be compiled by the following command in the project's root directory:
```$ go build -o creditcard .```

## Validate
The validate feature checks if a credit card number is valid using Luhn's Algorithm.

### Requirements:

- The number must be at least 13 digits long.
- If valid, print OK to stdout and exit with status 0.
- If invalid, print INCORRECT to stderr and exit with status 1.
- Support passing multiple entries.
- Support --stdin flag to pass number from stdin.

## Generate
The generate feature creates possible credit card numbers by replacing asterisks (*) with digits.

### Requirements:

- Replace up to 4 asterisks (*) with digits. If more - it's an error. Asterisks should be at the end of the given credit card number.
- Print the generated numbers to stdout.
- Numbers must be printed in ascending order.
- Exit with status 1 if there is any error.
- Support --pick flag to randomly pick a single entry.

## Information
The information feature provides details about the card based on data in brands.txt and issuers.txt.

### Requirements:

- Output the card number, validity, brand, and issuer.
- Support --stdin flag to pass number from stdin.
- Support passing multiple entries.

## Issue
The issue feature generates a random valid credit card number for a specified brand and issuer.

### Requirements:

- Pick a random number for the specified brand and issuer.
- Exit with status 1 if there is any error.

Example:
<br/>```$ ./creditcard issue --brands=brands.txt --issuers=issuers.txt --brand=VISA --issuer="Kaspi Gold"```
<br/>```4400430180300003```



