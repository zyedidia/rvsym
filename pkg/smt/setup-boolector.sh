cd ./pkg/smt

git clone https://github.com/zyedidia/boolector
cd boolector
./contrib/setup-lingeling.sh
./contrib/setup-btor2tools.sh
./configure.sh && cd build && make

cd ..

mkdir -p ../deps/include
mkdir -p ../deps/lib
cp -r deps/install/include/* ../deps/include
cp -r deps/install/lib/* ../deps/lib
cp -r build/lib/* ../deps/lib
cp src/boolector.h ../deps/include
cp src/btortypes.h ../deps/include
