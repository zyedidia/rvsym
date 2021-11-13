#include "librock.h"

int main() {
    unsigned pin = 0;
    unsigned value = 1;

    while (1) {
        gpio_write(pin, value);
        for (int i = 0; i < 500000; i++) {
            stall();
        }
        value = !value;
    }
}
