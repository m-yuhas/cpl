# Makefile for cpl
# (C) 2016 Michael Yuhas

CC = g++
CFLAGS = -Wall -c
LFLAGS = -Wall
BUILDDIR = build
SRCDIR = src
INCLUDEDIR = include
TARGET = bin/cpl

$(TARGET): $(BUILDDIR)/main.o $(BUILDDIR)/StringUtils.o
	$(CC) $(LFLAGS) $(BUILDDIR)/main.o $(BUILDDIR)/StringUtils.o -o $(TARGET)

$(BUILDDIR)/main.o: $(SRCDIR)/main.cpp $(INCLUDEDIR)/StringUtils.hpp
	$(CC) $(CFLAGS) $(SRCDIR)/main.cpp -o $(BUILDDIR)/main.o

$(BUILDDIR)/StringUtils.o: $(INCLUDEDIR)/StringUtils.hpp $(SRCDIR)/StringUtils.cpp
	$(CC) $(CFLAGS) $(SRCDIR)/StringUtils.cpp -o $(BUILDDIR)/StringUtils.o

clean:
	rm $(BUILDDIR)/*.o; rm $(TARGET)
