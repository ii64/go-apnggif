CC = gcc
CC2 = g++
all: zlib.dll png.dll apng2gif.dll
zlib.dll:zlib/*.c
	$(CC) -shared -fPIC -o zlib.dll zlib/*.c

png.dll:libpng/*.c
	$(CC) -Izlib -L. -lzlib -shared -fPIC -o png.dll libpng/*.c
	
apng2gif.dll:
	$(CC2) -L. -lzlib -lpng -Ilibpng -shared -fPIC -o apng2gif.dll apng2gif.cpp