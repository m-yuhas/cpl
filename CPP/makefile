# Makefile for cpl
# (C) 2016 Michael Yuhas

CC = g++
CFLAGS = -Wall -c
LFLAGS = -Wall
TFLAGS = -lgtest -lgtest_main -lpthread
BUILD = build
SRC = src
INCLUDE = include
TEST = test
BIN = bin

all: $(BIN)/cpl
$(BIN)/cpl: $(BUILD)/main.o $(BUILD)/StringUtils.o $(BUILD)/UnicodeString.o $(BUILD)/InvalidUTF8Exception.o
	$(CC) $(LFLAGS) $(BUILD)/main.o $(BUILD)/StringUtils.o $(BUILD)/UnicodeString.o $(BUILD)/InvalidUTF8Exception.o -o $(BIN)/cpl

$(BUILD)/main.o: $(SRC)/main.cpp $(INCLUDE)/StringUtils.hpp $(INCLUDE)/UnicodeString.hpp
	$(CC) $(CFLAGS) $(SRC)/main.cpp -o $(BUILD)/main.o

$(BUILD)/StringUtils.o: $(INCLUDE)/StringUtils.hpp $(SRC)/StringUtils.cpp
	$(CC) $(CFLAGS) $(SRC)/StringUtils.cpp -o $(BUILD)/StringUtils.o

$(BUILD)/UnicodeString.o: $(INCLUDE)/UnicodeString.hpp $(SRC)/UnicodeString.cpp $(INCLUDE)/InvalidUTF8Exception.hpp
	$(CC) $(CFLAGS) $(SRC)/UnicodeString.cpp -o $(BUILD)/UnicodeString.o

$(BUILD)/InvalidUTF8Exception.o: $(INCLUDE)/InvalidUTF8Exception.hpp $(SRC)/InvalidUTF8Exception.cpp
	$(CC) $(CFLAGS) $(SRC)/InvalidUTF8Exception.cpp -o $(BUILD)/InvalidUTF8Exception.o

clean:
	rm $(BUILD)/*.o; rm $(BIN)/*

test: $(BIN)/test
$(BIN)/test: $(BUILD)/StringUtils.o $(BUILD)/test_main.o
	$(CC) $(LFLAGS) $(BUILD)/test_main.o $(BUILD)/StringUtils.o -o $(BIN)/test $(TFLAGS)

$(BUILD)/test_main.o: $(TEST)/test_main.cpp $(INCLUDE)/StringUtils.hpp
	$(CC) $(CFLAGS) $(TEST)/test_main.cpp -o $(BUILD)/test_main.o $(TFLAGS)
