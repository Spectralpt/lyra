# C++ Project Structure

This project follows modern C++ project layout with support for both Make and CMake build systems:

## Directory Structure

- `services/src/` - Source code files (.cpp and .hpp files)
- `tests/` - Test files
- `build/` - Compiled binaries and object files (created during build)

## Building with Make

To build the project:
```bash
make
```

To run the program:
```bash
make run
```

To run tests:
```bash
make test
```

To build debug version:
```bash
make debug
```

To build release version:
```bash
make release
```

To clean build artifacts:
```bash
make clean
```

## Building with CMake

To build with CMake:
```bash
mkdir build
cd build
cmake ..
make
```

To run tests with CMake:
```bash
cd build
ctest
```

## Files

- `main.cpp` - Main program entry point
- `utils.cpp` - Utility class implementation
- `utils.hpp` - Header file with class declarations

## Features

- C++17 standard
- Object-oriented design with classes
- Function overloading examples
- Modern C++ practices
- Both Make and CMake build support
