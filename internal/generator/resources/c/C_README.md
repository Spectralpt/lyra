# C Project Structure

This project follows a standard C project layout:

## Directory Structure

- `services/src/` - Source code files (.c and .h files)
- `tests/` - Test files
- `build/` - Compiled binaries and object files (created during build)

## Building

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

To clean build artifacts:
```bash
make clean
```

## Files

- `main.c` - Main program entry point
- `utils.c` - Utility functions
- `utils.h` - Header file with function declarations
