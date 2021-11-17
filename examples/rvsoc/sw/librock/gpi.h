#define GPI_BASE 0x2008

static volatile unsigned* gpi_level_reg   = (unsigned*) (GPI_BASE + 0x00);

typedef enum gpi_pin {
    BTN = 0,
} gpi_pin_t;

static inline unsigned gpi_read(unsigned pin) {
    return (get32(gpi_level_reg) >> pin) & 0x1;
}
