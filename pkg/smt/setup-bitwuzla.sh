git clone https://github.com/bitwuzla/bitwuzla
cd bitwuzla
./contrib/setup-cadical.sh
./contrib/setup-btor2tools.sh
./contrib/setup-symfpu.sh
./configure.sh && cd build && make

cd ..

mkdir -p ../deps/include
mkdir -p ../deps/lib
cp -r deps/install/include/* ../deps/include
cp -r deps/install/lib/* ../deps/lib
cp -r build/lib/* ../deps/lib
cp src/api/c/bitwuzla.h ../deps/include
