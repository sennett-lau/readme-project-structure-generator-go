# ReadMe Project Structure Generator

## Description
This is a command-line application that dynamically generates the project structure into a `structure.md` file for user's `README.md`.

## Table of Contents
- [File Structure](#file-structure)
- [Installation](#installation)
- [Implementation](#implementation)
  - [Basic Usage](#basic-usage)
  - [Advanced Usage](#advanced-usage)
- [Contributing](#contributing)
- [License](#license)

## File Structure
```
.
├── test
│   ├── utils_args_check_test.go
│   └── utils_func_test.go
├── types
│   └── directory.go
├── utils
│   ├── args_check.go
│   ├── config.go
│   └── func.go
├── README.md
├── go.mod
├── main.go
└── structure.md
```

## Installation
To install this application:
```
git clone https://github.com/sennett-lau/readme-project-structure-generator-go.git
cd readme-project-structure-generator-go
make install
```

## Implementation

### Basic Usage
To use this application, simply run the following command at the root of your project and a `structure.md` file will be generated.

```bash
rpsg
```

There are some default ignored files which located at `config/GetDefaultIgnoreList` including:
- `.git`
- `.gitignore`
- `.idea`
- `.vscode`
- `node_modules`
- `vendor`
- `*.exe`
- `*.dll`
- `*.so`
- `*.dylib`
- `*.min.js`
- `*.min.css`
- `*.min.html`
- `*.min.json`
- `*.min.xml`
- `venv`
- `__pycache__`
- *.pyc

### Advanced Usage
There are more available options for users to customize the output of the `structure.md` file.
- [`--extend-ignore-list`](#--extend-ignore-list) / `-e`
- [`--max-depth`](#--max-depth) / `-d`
- [`--show-ignore-list`](#--show-ignore-list) / `-s`

#### `--extend-ignore-list`
The `--extend-ignore-list` option allows users to add more ignored files to the default ignored list.
<br/>
For example, if you want to add `.DS_Store` and `.gitkeep` to the ignored list, you can run the following command:

```bash
rpsg --extend-ignore-list=.DS_Store,.gitkeep
```

#### `--max-depth`
The `--max-depth` option allows users to set the maximum depth of the file structure, while the default value is `6` and the allowed range is from `1` to `10`

```txt
// the default maximum depth of 6, as shown below:
.
├── 1
│   └── 2
│       └── 3
│           └── 4
│               └── 5
│                   └── 6
│                       └── ...
└── ...
```

For example, if you want to set the maximum depth to `3`, you can run the following command:

```bash
rpsg --max-depth=3
```

output:

```txt
.
├── 1
│   └── 2
│       └── 3
│           └── ...
└── ...
```

#### `--show-ignore-list`
The `--show-ignore-list` option allows users to show the default ignored files in the `structure.md` file and will NOT generate the file structure.

For example, if you want to show the default ignored files, you can run the following command:

```bash
rpsg --show-ignore-list
```

output:

```txt
Ignore list:
.git
.gitignore
...
```

## Contributing
Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

### General Steps

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a pull request

## License
This library is licensed under the MIT License. See the [LICENSE](LICENSE.txt) file
