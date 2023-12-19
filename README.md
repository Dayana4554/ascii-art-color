**Ascii-art-color**

Objectives

The ascii-art-color project expands on the basic ASCII art concept by introducing color manipulation. This project requires you to implement a feature that allows users to color specific letters or sets of letters in the ASCII art output. The color customization is achieved using the --color=<color> <letters to be colored> flag, where <color> is the desired color (supporting various color code systems) and <letters to be colored> are the specific letters to be colored.
Key Features:

    Color Individual Letters or Sets: Choose to color a single letter or a group of letters.
    Flexible Color Notation: Supports different color code systems (RGB, HSL, ANSI, etc.).
    Default Coloring: If no specific letter is mentioned, the entire string will be colored.
    Strict Flag Formatting: The flag must follow the format --color=<color> <letters to be colored>. Any deviation leads to a usage message.

Usage

```shell

Usage: go run . [OPTION] [STRING]

Example: go run . --color=<color> <letters to be colored> "your text"
```

If implementing alongside other ASCII art projects, the program should support other [OPTION] and/or [BANNER] formats and must run with a single [STRING] argument.
Instructions

    Language: The project must be written in Go.
    Coding Practices: Adherence to good coding practices is expected.
    Testing: Unit tests are recommended.
