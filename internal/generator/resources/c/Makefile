CC=gcc
CFLAGS=-Wall -Wextra -std=c99 -Iservices/src
SRC_DIR=services/src
BUILD_DIR=build
TESTS_DIR=tests
TARGET=$(BUILD_DIR)/main
TEST_TARGET=$(BUILD_DIR)/test_main

# Source files
SOURCES=$(wildcard $(SRC_DIR)/*.c)
OBJECTS=$(SOURCES:$(SRC_DIR)/%.c=$(BUILD_DIR)/%.o)

# Test files
TEST_SOURCES=$(wildcard $(TESTS_DIR)/*.c)
TEST_OBJECTS=$(TEST_SOURCES:$(TESTS_DIR)/%.c=$(BUILD_DIR)/%.o)

all: $(TARGET)

$(TARGET): $(OBJECTS) | $(BUILD_DIR)
	$(CC) $(CFLAGS) -o $(TARGET) $(OBJECTS)

$(BUILD_DIR)/%.o: $(SRC_DIR)/%.c | $(BUILD_DIR)
	$(CC) $(CFLAGS) -c $< -o $@

test: $(TEST_TARGET)
	./$(TEST_TARGET)

$(TEST_TARGET): $(TEST_OBJECTS) | $(BUILD_DIR)
	$(CC) $(CFLAGS) -o $(TEST_TARGET) $(TEST_OBJECTS)

$(BUILD_DIR)/%.o: $(TESTS_DIR)/%.c | $(BUILD_DIR)
	$(CC) $(CFLAGS) -c $< -o $@

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

clean:
	rm -rf $(BUILD_DIR)

run: $(TARGET)
	./$(TARGET)

.PHONY: all test clean run
