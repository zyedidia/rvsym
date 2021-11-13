#define TIMER_ADDR 0x4000

static inline unsigned timer_get_usec() {
    return get32((const volatile void*) TIMER_ADDR);
}

static inline void delay_us(unsigned us) {
    unsigned rb = timer_get_usec();
    while (1) {
        unsigned ra = timer_get_usec();
        if ((ra - rb) >= us) {
            break;
        }
    }
}

static inline void delay_ms(unsigned ms) {
    delay_us(1000 * ms);
}
