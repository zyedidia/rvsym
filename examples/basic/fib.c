#include <stdio.h>
int fib(int n) {
    if (n <= 1)
        return n;
    return fib(n - 1) + fib(n - 2);
}

int main() {
    int a = fib(32);
    return a;
}
