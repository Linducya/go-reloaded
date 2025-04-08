# Go-Reloaded: A Text Modifier Tool

Go-Reloaded is a text modifier auto-correction tool written in Go. It processes an input text file, applies various modifications, and writes the modified content to an output file.

## Features

1. **Hexadecimal to Decimal Conversion**: Converts hexadecimal numbers followed by `(hex)` to their decimal equivalents.
   - Example: `"1E (hex) files were added"` → `"30 files were added"`

2. **Binary to Decimal Conversion**: Converts binary numbers followed by `(bin)` to their decimal equivalents.
   - Example: `"It has been 10 (bin) years"` → `"It has been 2 years"`

3. **Text Case Modifications**:
   - `(up)`: Converts the preceding word to uppercase.
     - Example: `"Ready, set, go (up)!"` → `"Ready, set, GO!"`
   - `(low)`: Converts the preceding word to lowercase.
     - Example: `"I should stop SHOUTING (low)"` → `"I should stop shouting"`
   - `(cap)`: Capitalizes the preceding word.
     - Example: `"Welcome to the Brooklyn bridge (cap)"` → `"Welcome to the Brooklyn Bridge"`

4. **Multi-Word Case Modifications**: Applies `(low, <number>)`, `(up, <number>)`, or `(cap, <number>)` to the specified number of preceding words.
   - Example: `"This is so exciting (up, 2)"` → `"This is SO EXCITING"`

5. **Punctuation Formatting**:
   - Ensures punctuation marks (e.g., `.`, `,`, `!`, `?`, `:`, `;`) are properly spaced.
     - Example: `"I was sitting over there ,and then BAMM !!"` → `"I was sitting over there, and then BAMM!!"`
   - Handles special cases such as groups of punctuation like `...` or `!?` correctly.
     - Example: `"I was thinking ... You were right"` → `"I was thinking... You were right"`

6. **Quotation Marks**: Ensures single quotes (`'`), which will always be found with another instance of it, are placed correctly around words or phrases without any spaces.
   - Example: `"I am exactly how they describe me: ' awesome '"` → `"I am exactly how they describe me: 'awesome'"`

7. **Article Correction**: Replaces `a` with `an` if the next word starts with a vowel (a, e, i, o, u).
   - Example: `"There it was. A amazing rock!"` → `"There it was. An amazing rock!"`

## Project Structure

The Go package for this project is called `textutils`. It contains utility functions for text processing, such as capitalization, punctuation formatting, and case modifications.

### Folder Structure

- **`textutils/`**: Contains the Go source files for text processing.
  - Example: `textutils/capitalizeword.go` includes the `CapitalizeWord` function, which capitalizes the first alphabetic character of a word while preserving punctuation.

### Example Usage of `textutils`

You can import and use the `textutils` package in your Go code as follows:

```go
package main

import (
    "fmt"
    "textutils"
)

func main() {
    word := "hello"
    capitalized := textutils.CapitalizeWord(word)
    fmt.Println(capitalized) // Output: Hello
}
```
## Usage

1. Create an input file (`sample.txt`) with the text to be modified.
2. Run the program with the input and output file names as arguments:
   ```bash
   $ go run main.go sample.txt result.txt
   ```
3. View the modified content in the output file (result.txt).
   ```bash
   $ cat result.txt
   ```

## Test Files
Example input test files can be found in the input_files folder. These files demonstrate various features of the tool, such as:

- Case modifications
- Hexadecimal and binary conversions
- Article corrections
- Punctuation formatting

To use these files, simply reference them when running the program. For example:
```bash
$ go run . input_files/sample.txt result.txt
```

## Example 1: Case Modifications
```bash
$ cat input_files/sample_case_modifications.txt
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.

$ go run . input_files/sample_case_modifications.txt result.txt

$ cat result.txt
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.
```

## Example  2: Hexadecimal and Binary Conversion 
```bash
$ cat input_files/sample_hex_bin_conversion.txt
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.

$ go run . input_files/sample_hex_bin_conversion.txt result.txt

$ cat result.txt
Simply add 66 and 2 and you will see the result is 68.
```

## Example 3: Article Correction
```bash
$ cat input_files/sample_article_correction.txt
There is no greater agony than bearing a untold story inside you.

$ go run . input_files/sample_article_correction.txt result.txt

$ cat result.txt
There is no greater agony than bearing an untold story inside you.
```

## Example 4: Punctuation Formatting
```bash
$ cat input_files/sample_punctuation_formatting.txt
Punctuation tests are ... kinda boring ,what do you think ?

$ go run . input_files/sample_punctuation_formatting.txt result.txt

$ cat result.txt
Punctuation tests are... kinda boring, what do you think?
```

## Learning Objectives
This project will help you learn about:

- The Go file system (os and io) API
- String and number manipulation
- Writing clean and maintainable Go code
- Unit testing in Go

## Allowed Packages
This project only uses standard Go packages.

## Contribution
Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Commit your changes with a descriptive message.
4. Push your branch and create a pull request.

## License
This project is licensed under the MIT License. See the LICENSE file for details.