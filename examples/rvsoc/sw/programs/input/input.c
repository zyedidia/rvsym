#include "librock.h"

int main() {
    gpio_set_input(GPIO_1);
    gpio_set_output(GPIO_0);

    while (1) {
        gpio_write(GPIO_0, gpio_read(GPIO_1));
    }
}
