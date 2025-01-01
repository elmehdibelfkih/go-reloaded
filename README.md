# go-reloaded

## Description
**go-reloaded** is a text completion, editing, and auto-correction tool built in Go. The program takes an input file, applies specific text transformations, and writes the modified content to an output file. It demonstrates mastery of Go's file system (fs) API, string manipulation, and numeric conversions.

This project is part of a peer-audited process where students review and test each other's code. You are encouraged to create your own tests for this tool.

---

## Objectives
- Build a tool to process text files and apply various modifications.
- Follow Go best practices to write clean, maintainable code.
- Create unit tests to ensure correctness.
- Participate in the peer-auditing process.

---

## Features
1. **Hexadecimal to Decimal Conversion**:
   Replace words marked with `(hex)` with their decimal equivalents.  
   Example:  
"1E (hex) files were added" -> "30 files were added"

2. **Binary to Decimal Conversion**:
Replace words marked with `(bin)` with their decimal equivalents.  
Example:  
"10 (bin) years" -> "2 years"

3. **Case Transformations**:
- `(up)`: Convert the preceding word to uppercase.  
  Example: `"go (up) now!" -> "GO now!"`  
- `(low)`: Convert the preceding word to lowercase.  
  Example: `"SHOUTING (low)" -> "shouting"`  
- `(cap)`: Capitalize the preceding word.  
  Example: `"bridge (cap)" -> "Bridge"`  

**Multi-word transformations**:  
- `(up, n)`, `(low, n)`, `(cap, n)` modify the last `n` words.  
  Example: `"exciting (up, 2)" -> "SO EXCITING"`

4. **Punctuation Formatting**:
Ensure punctuation `.`, `,`, `!`, `?`, `:` and `;` is correctly formatted:
- Close to the preceding word.
- Spaced apart from the next word.  
Example:  
"boring ,what do you think ?" -> "boring, what do you think?"

5. **A/An Correction**:
Replace `a` with `an` when the next word starts with a vowel or `h`.  
Example:  
"a amazing rock" -> "an amazing rock"

6. **Quotation Marks**:
Ensure `'` is placed directly around words or groups of words.  
Examples:  
"' awesome '" -> "'awesome'" "As Elton John said: ' I am the one '" -> "As Elton John said: 'I am the one'"

---

## Prerequisites
- Install [Go](https://go.dev/).

---

## Usage
1. Create a text file with the content you want to modify (e.g., `sample.txt`).
2. Run the program:
```bash
go run . sample.txt result.txt

---

## Prerequisites
- Install [Go](https://go.dev/).

---

## Usage
1. Create a text file with the content you want to modify (e.g., `sample.txt`).
2. Run the program:
```bash
go run . sample.txt result.txt
