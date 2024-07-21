# Password Checker

This Password Checker is a simple Go program designed to evaluate the strength of a given password based on its length and character composition. It estimates the time required to brute-force the password, providing a detailed breakdown in terms of years, months, weeks, days, hours, minutes, and seconds.

## Features

- Checks if the password contains lowercase letters, uppercase letters, numbers, and special characters.
- Calculates the character set size based on the types of characters used.
- Estimates the time to brute-force the password using a specified hash rate.

## Installation

1. Ensure you have Go installed on your machine.
2. Clone this repository:
3. Copy code
```git clone https://github.com/MartinDeBeer/passCheck```
4. Navigate to the project directory:
```cd password-checker```

## Usage

1. Build the program:
```go build -o password-checker```
2. Run the program:
```./password-checker```
3. Enter your password when prompted:
```Enter password:```
4. The program will then display the character set size, the number of possibilities, and the estimated time to crack the password.

## Example

```
Enter password:
P@ssw0rd123
Character set size: 94
Number of possibilities: 7.537181102180347e+17
Time to crack: 216118940 years, 2593427280 months, 11287205205 weeks, 79107182040 days, 1898572368960 hours, 113914342137600 minutes, 6834860528256000 seconds
```

## Explanation

The password checker performs the following steps:

1. Read the Password: Prompts the user to enter a password and trims any whitespace.
2. Pattern Matching: Uses regular expressions to check for the presence of lowercase letters, uppercase letters, numbers, and special characters.
3. Character Set Size: Calculates the total character set size based on the character types found in the password.
4. Number of Possibilities: Computes the total number of possible combinations based on the character set size and password length.
5. Time to Crack: Estimates the time to brute-force the password based on a hash rate of 1,102,200,000 hashes per second.
6. Duration Conversion: Converts the estimated time in seconds to a more human-readable format.

## License

This project is licensed under the MIT License

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## Acknowledgements

Special thanks to the Go community for their support and resources.
