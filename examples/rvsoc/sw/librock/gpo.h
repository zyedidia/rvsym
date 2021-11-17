#define GPO_BASE 0x2000

static volatile unsigned* gpo_set_reg   = (unsigned*) (GPO_BASE + 0x00);
static volatile unsigned* gpo_clear_reg = (unsigned*) (GPO_BASE + 0x04);

typedef enum gpo_pin {
    LED_R = 0,
    LED_G = 1,
    LED_B = 2
} gpo_pin_t;

static inline void gpo_set(unsigned pin) {
    put32(gpo_set_reg, 1 << pin);
}

static inline void gpo_clear(unsigned pin) {
    put32(gpo_clear_reg, 1 << pin);
}

static inline void gpo_write(unsigned pin, unsigned value) {
    if (value) {
        gpo_set(pin);
    } else {
        gpo_clear(pin);
    }
}
