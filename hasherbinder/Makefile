# export DYLD_LIBRARY_PATH=/usr/local/lib:$DYLD_LIBRARY_PATH

build-binding:
gcc -shared -o hasherbinding.so \
	-fPIC src/libhasher.c \
	-I/opt/homebrew/Cellar/lua/5.4.7/include/lua5.4 \
	-L/opt/homebrew/Cellar/lua/5.4.7/lib \
	-llua -lhasher