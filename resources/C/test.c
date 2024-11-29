

char ok() {
    return 1;
}




int main() {
    int a = 'a';
    float b = 0;
    double d = 0;
    char c = 'a';
    long l = 0;
    short s = 0;
    unsigned int u = 0;
    unsigned long ul = 0;
    unsigned short us = 0;
    unsigned char uc = 'a';
    signed int si = 0;
    signed long sl = 0;

    int * pa = &a;
    signed short ss = 0;
    signed char sc = 'a';
    char str[] = "Hello, World!";

    if (pa == a) {
        a = 1;
    }

    if (a == c) {
        a = 'a';
    }

    a = ok();

    //char str[] = "Hello, World!";
    a = 0.5234f;

}