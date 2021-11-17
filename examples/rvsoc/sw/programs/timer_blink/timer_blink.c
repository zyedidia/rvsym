#include "librock.h"

int main() {
    while (1) {
        for (gpo_pin_t led = LED_R; led <= LED_B; led++) {
            // set all colors to 0
            gpo_write(LED_R, 0);
            gpo_write(LED_G, 0);
            gpo_write(LED_B, 0);

            // turn on the color we want
            gpo_write(led, 1);

            delay_ms(1000);
        }
    }
}
