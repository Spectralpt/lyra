cmake_minimum_required(VERSION 3.10)
project(CppProject)

# Set C++ standard
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# Set compiler flags
set(CMAKE_CXX_FLAGS "-Wall -Wextra")
set(CMAKE_CXX_FLAGS_DEBUG "-g -DDEBUG")
set(CMAKE_CXX_FLAGS_RELEASE "-O3 -DNDEBUG")

# Include directories
include_directories(services/src)

# Source files
file(GLOB SOURCES "services/src/*.cpp")

# Main executable
add_executable(main ${SOURCES})

# Test executable
file(GLOB TEST_SOURCES "tests/*.cpp")
add_executable(test_main ${TEST_SOURCES})

# Enable testing
enable_testing()
add_test(NAME cpp_tests COMMAND test_main)
