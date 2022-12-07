# json-converter

![image](asset/excel.png)

json-converter is CLI to convert JSON to Any other files like: 
- Excel
- CSV
- and more

## Installation

### Using go install
```
$ go install github.com/kecci/json-converter@latest
go: downloading github.com/kecci/json-converter v0.0.1

$ json-converter
json-converter - a simple CLI to conversion json and excel files.

Usage:
   [flags]
   [command]

Available Commands:
  excel       Convert json to excel file
  help        Help about any command

Flags:
  -h, --help      help for this command
  -v, --version   version for this command

Use " [command] --help" for more information about a command.
```

### Using the github repository
```sh
git clone https://github.com/kecci/json-converter.git
cd json-converter
go build -o json-converter .
./json-converter --help
```

## Usage
```
./json-converter excel
```

## Example

### Example JSON
`./example/example.json`
```json
[
    {
      "isbn": "123-456-222",
      "author": {
        "lastname": "Doe",
        "firstname": "Jane"
      },
      "editor": {
        "lastname": "Smith",
        "firstname": "Jane"
      },
      "title": "The Ultimate Database Study Guide",
      "category": [
        "Non-Fiction",
        "Technology"
      ]
    },
    {
      "isbn": "345-666-777",
      "author": {
        "lastname": "Han",
        "firstname": "Some"
      },
      "editor": {
        "lastname": "Marry",
        "firstname": "Jane"
      },
      "title": "The Story Behind Spiderman",
      "category": [
        "Fiction",
        "Comic"
      ]
    }
  ]
```

### Running Command json-converter excel
```sh
./json-converter excel ./example/example.json ./example/example.xlsx
```

### Result
| author.firstname | author.lastname | category.0  | category.1 | editor.firstname | editor.lastname | isbn        | title                             |
|------------------|-----------------|-------------|------------|------------------|-----------------|-------------|-----------------------------------|
| Jane             | Doe             | Non-Fiction | Technology | Jane             | Smith           | 123-456-222 | The Ultimate Database Study Guide |
| Some             | Han             | Fiction     | Comic      | Jane             | Marry           | 345-666-777 | The Story Behind Spiderman        |