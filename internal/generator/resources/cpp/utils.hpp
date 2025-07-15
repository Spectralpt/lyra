#ifndef UTILS_HPP
#define UTILS_HPP

#include <string>

class Utils {
private:
    std::string version;

public:
    Utils();
    void print_version() const;
    std::string get_version() const;
    void set_version(const std::string& new_version);
};

// Function overloading examples
int add(int a, int b);
double add(double a, double b);

#endif // UTILS_HPP
