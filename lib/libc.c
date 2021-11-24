#include <stdint.h>
#include <stddef.h>

int strcmp(const char* a, const char* b) {
    while (1) {
        unsigned char ac = *a, bc = *b;
        if (ac == 0 || bc == 0 || ac != bc) {
            return (ac > bc) - (ac < bc);
        }
        ++a, ++b;
    }
}

size_t strlen(const char* p) {
    size_t ret;
    for (ret = 0; p[ret]; ++ret)
        ;
    return ret;
}

void* memset(void* _p, int c, size_t n) {
    char *p = _p, *e = p + n;

    while (p < e)
        *p++ = c;
    return _p;
}

void* memmove(void* dst, const void* src, size_t count) {
    char* a = dst;
    const char* b = src;

    if (src == dst)
        return dst;

    if (src > dst) {
        while (count--)
            *a++ = *b++;
    } else {
        a += count - 1;
        b += count - 1;
        while (count--)
            *a-- = *b--;
    }

    return dst;
}

#define aligned(ptr, n) ((unsigned) ptr % n == 0)
#define aligned4(ptr) aligned(ptr, 4)

void* memcpy(void* dst, const void* src, size_t nbytes) {
    if (aligned4(dst) && aligned4(src) && aligned4(nbytes)) {
        unsigned n = nbytes / 4;
        unsigned* d = dst;
        const unsigned* s = src;

        for (unsigned i = 0; i < n; i++)
            d[i] = s[i];
    } else {
        unsigned char* d = dst;
        const unsigned char* s = src;
        for (unsigned i = 0; i < nbytes; i++)
            d[i] = s[i];
    }
    return dst;
}
