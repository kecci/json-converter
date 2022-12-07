# json-converter

![image](asset/excel.png)

json-converter is CLI to convert JSON to Any other files like: 
- Excel
- CSV
- and more

## Installation

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