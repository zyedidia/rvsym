#include "librock.h"

int main() {
    while (1) {
        unsigned btn_val = gpi_read(BTN);
        gpo_write(LED_R, btn_val);
    }
}
