#include <iostream>
#include <cassert>
#include "../services/src/utils.hpp"

// Test class for organizing tests
class TestSuite {
public:
    static void test_example() {
        assert(1 == 1);
        std::cout << "Test passed: example test" << std::endl;
    }

    static void test_math() {
        assert(2 + 2 == 4);
        std::cout << "Test passed: math test" << std::endl;
    }

    static void test_utils() {
        assert(add(3, 4) == 7);
        std::cout << "Test passed: utils test" << std::endl;
    }

    static void run_all_tests() {
        std::cout << "Running C++ tests..." << std::endl;
        test_example();
        test_math();
        test_utils();
        std::cout << "All tests passed!" << std::endl;
    }
};

int main() {
    TestSuite::run_all_tests();
    return 0;
}
