#include <stdio.h>
#include <assert.h>

// Example test function
void test_example() {
    assert(1 == 1);
    printf("Test passed: example test\n");
}

void test_math() {
    assert(2 + 2 == 4);
    printf("Test passed: math test\n");
}

int main() {
    printf("Running tests...\n");
    test_example();
    test_math();
    printf("All tests passed!\n");
    return 0;
}
