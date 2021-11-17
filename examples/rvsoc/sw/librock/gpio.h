typedef enum gpio_pin {
    GPIO_0 = 0,
    GPIO_1 = 1,
    GPIO_5 = 2,
    GPIO_6 = 3,
    GPIO_9 = 4,
    GPIO_10 = 5,
    GPIO_11 = 6,
    GPIO_12 = 7,
    GPIO_13 = 8,
    GPIO_A0 = 9,
    GPIO_A1 = 10,
    GPIO_A2 = 11,
    GPIO_A3 = 12
} gpio_pin_t;

#define GPIO_BASE 0x3000

static volatile unsigned* gpio_set_reg   = (unsigned*) (GPIO_BASE + 0x00);
static volatile unsigned* gpio_clear_reg = (unsigned*) (GPIO_BASE + 0x04);
static volatile unsigned* gpio_level_reg = (unsigned*) (GPIO_BASE + 0x08);
static volatile unsigned* gpio_fsel_reg  = (unsigned*) (GPIO_BASE + 0x0c);

static inline void gpio_set(unsigned pin) {
    put32(gpio_set_reg, 1 << pin);
}

static inline void gpio_clear(unsigned pin) {
    put32(gpio_clear_reg, 1 << pin);
}

static inline void gpio_set_input(unsigned pin) {
    put32(gpio_fsel_reg, get32(gpio_fsel_reg) | (1 << pin));
}

static inline void gpio_set_output(unsigned pin) {
    put32(gpio_fsel_reg, get32(gpio_fsel_reg) & ~(1 << pin));
}

static inline void gpio_write(unsigned pin, unsigned value) {
    if (value) {
        gpio_set(pin);
    } else {
        gpio_clear(pin);
    }
}

static inline unsigned gpio_read(unsigned pin) {
    return (get32(gpio_level_reg) >> pin) & 0x1;
}
