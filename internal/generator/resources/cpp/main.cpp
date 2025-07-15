#include <iostream>
#include "utils.hpp"

int main() {
    std::cout << "Hello, C++ World!" << std::endl;

    Utils utils;
    utils.print_version();

    std::cout << "Addition result: " << add(5, 3) << std::endl;

    return 0;
}
