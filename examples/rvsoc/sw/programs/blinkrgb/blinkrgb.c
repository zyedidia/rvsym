#include "librock.h"

int main() {
    // rgb_led_r: 0
    // rgb_led_g: 1
    // rgb_led_b: 2

    while (1) {
        for (gpo_pin_t led = LED_R; led <= LED_B; led++) {
            // set all colors to 0
            gpo_write(LED_R, 0);
            gpo_write(LED_G, 0);
            gpo_write(LED_B, 0);

            // turn on the color we want
            gpo_write(led, 1);

            // stall
            for (int i = 0; i < 500000; i++) {
                stall();
            }
        }
    }
}
