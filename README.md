# Changus Screenus

[![Golang Gopher](https://godoc.org/github.com/disintegration/imaging?status.svg)](https://golang.org/)

Changus Screenus is a Go program that changes your screensaver to a random picture from a folder found in this project name `images`. It is currently only supported by Windows. Linux and MacOS are among future updates.

## Features

- Selects a random image from a specified folder.
- Sets the screensaver to the selected image.
- Supports Windows, macOS, and Linux.

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/kyleochata/changus_screenus.git
   cd changus_screenus/cmd/change_screensaver
   ```

2. Build the project:

```sh
go build
```

- This will create a `change_screensaver.exe` file in the `change_screensaver` directory.

3. Add desired images to the `images` directory found at the root level of the repository.

- Ensure that they are in a format that is accepted by Windows (I stick with jpegs)

4. Changus_Screenus!

```
./change_screensaver -folder ../../images/
```

## Usage

This project was built for learning purposes and is open to critique and usage under the chosen [liscence](#liscence).

<p align="right">(<a href="#back_to_top">back to top</a>)</p>

## Liscence

MIT License

Copyright (c) 2023 ultimated1228

Permission is hereby granted, free of charge, to any person or organization
obtaining a copy of the software and accompanying documentation covered by
this license (the "Software") to use, reproduce, display, distribute,
execute, and transmit the Software, and to prepare derivative works of the
Software, and to permit third-parties to whom the Software is furnished to
do so, all subject to the following:

The copyright notices in the Software and this entire statement, including
the above license grant, this restriction and the following disclaimer,
must be included in all copies of the Software, in whole or in part, and
all derivative works of the Software, unless such copies or derivative
works are solely in the form of machine-executable object code generated by
a source language processor.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE, TITLE AND NON-INFRINGEMENT. IN NO EVENT
SHALL THE COPYRIGHT HOLDERS OR ANYONE DISTRIBUTING THE SOFTWARE BE LIABLE
FOR ANY DAMAGES OR OTHER LIABILITY, WHETHER IN CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
DEALINGS IN THE SOFTWARE.

Fore more details on the [MIT License](https://opensource.org/licenses/MIT) please click the link, or check out the license file in the repo.

<p align="right">(<a href="#back_to_top">back to top</a>)</p>This project is licensed under the MIT License

## Contact

Kyle Etrata

- [Kyle Etrata Github](https://github.com/kyleochata)
- [Email Kyle](mailto:kyleochata@gmail.com)
