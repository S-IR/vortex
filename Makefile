# Makefile for my_raylib_app

# Compiler
CC = clang

# Source files - Use wildcard to include all .c files in src directory
SRC = $(wildcard src/*.c)

# Output executable
TARGET = a.out

# Include and library directories
INCLUDE_DIR = /usr/local/include    # Adjust this to the correct include path
LIB_DIR = /usr/local/lib           # Adjust this to the correct library path

# Static library
LIBRARY = -lraylib  # Link against raylib (you might not need -l:libraylib.a)

# Compiler flags
CFLAGS = -Wall -Wextra -I $(INCLUDE_DIR) -L $(LIB_DIR) $(LIBRARY) -lm -lpthread -ldl -lrt -lX11

# Build target
all: $(TARGET)

$(TARGET): $(SRC)
	$(CC) $(SRC) -o $(TARGET) $(CFLAGS)

# Clean up
clean:
	rm -f $(TARGET)

.PHONY: all clean
