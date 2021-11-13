static inline void put32(volatile void* addr, uint32_t v) {
    *((volatile uint32_t*) addr) = v;
}

static inline uint32_t get32(const volatile void* addr) {
    return *((volatile uint32_t*) addr);
}
