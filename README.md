# Word Frequency Counter

A simple Go program that analyzes text files and displays the most frequently used words.

## Usage
```bash
go run main.go <filepath> <N>
```

Where:
- `<filepath>` - Path to the text file to analyze
- `<N>` - Number of top words to display

## Example
```bash
go run main.go sample.txt 10
```

Output:
```
The top 10 words are:
the: 42
and: 31
go: 28
...
```

## Features

- Cleans text (removes punctuation, converts to lowercase)
- Sorts by frequency (most common first)
- Handles edge cases (file not found, invalid input)