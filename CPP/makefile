# Makefile for cpl
# (C) 2016 Michael Yuhas

CC = g++
CFLAGS = -Wall -c
LFLAGS = -Wall
BUILD = build
SRC = src
INCLUDE = include
TARGET = bin/cpl

all: $(TARGET)
$(TARGET): $(BUILD)/main.o $(BUILD)/StringUtils.o $(BUILD)/UnicodeString.o $(BUILD)/InvalidUTF8Exception.o
	$(CC) $(LFLAGS) $(BUILD)/main.o $(BUILD)/StringUtils.o $(BUILD)/UnicodeString.o $(BUILD)/InvalidUTF8Exception.o -o $(TARGET)

$(BUILD)/main.o: $(SRC)/main.cpp $(INCLUDE)/StringUtils.hpp $(INCLUDE)/UnicodeString.hpp
	$(CC) $(CFLAGS) $(SRC)/main.cpp -o $(BUILD)/main.o

$(BUILD)/StringUtils.o: $(INCLUDE)/StringUtils.hpp $(SRC)/StringUtils.cpp
	$(CC) $(CFLAGS) $(SRC)/StringUtils.cpp -o $(BUILD)/StringUtils.o

$(BUILD)/UnicodeString.o: $(INCLUDE)/UnicodeString.hpp $(SRC)/UnicodeString.cpp $(INCLUDE)/InvalidUTF8Exception.hpp
	$(CC) $(CFLAGS) $(SRC)/UnicodeString.cpp -o $(BUILD)/UnicodeString.o

$(BUILD)/InvalidUTF8Exception.o: $(INCLUDE)/InvalidUTF8Exception.hpp $(SRC)/InvalidUTF8Exception.cpp
	$(CC) $(CFLAGS) $(SRC)/InvalidUTF8Exception.cpp -o $(BUILD)/InvalidUTF8Exception.o

clean:
	rm $(BUILD)/*.o; rm $(TARGET)

test: bin/test
bin/test: $(BUILD)/StringUtils.o $(BUILD)/tests.o
	$(CC) $(LFLAGS) $(BUILD)/tests.o $(BUILD)/StringUtils.o -o bin/test -lgtest -lgtest_main -googleTest -lpthread

$(BUILD)/tests.o: $(SRC)/tests.cpp $(INCLUDE)/StringUtils.hpp
	$(CC) -lpthread $(CFLAGS) $(SRC)/tests.cpp -o $(BUILD)/tests.o -lgtest -lgtest_main -googleTest
