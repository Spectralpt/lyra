#include <iostream>
#include "utils.hpp"

Utils::Utils() {
    version = "1.0.0";
}

void Utils::print_version() const {
    std::cout << "Version: " << version << std::endl;
}

std::string Utils::get_version() const {
    return version;
}

void Utils::set_version(const std::string& new_version) {
    version = new_version;
}

int add(int a, int b) {
    return a + b;
}

double add(double a, double b) {
    return a + b;
}
