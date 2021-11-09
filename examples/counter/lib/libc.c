int strcmp(const char* a, const char* b) {
    while (1) {
        unsigned char ac = *a, bc = *b;
        if (ac == 0 || bc == 0 || ac != bc) {
            return (ac > bc) - (ac < bc);
        }
        ++a, ++b;
    }
}

void free(void* ptr) {
    return;
}
