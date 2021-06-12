#include <stdio.h>
#include "zlib/zlib.h"
#include "apng2gif.h"

int
main() {
    printf("running apng2gif...\n");
    int s = apng2gif("anim.png", "anim.gif", 0, "");
    printf("%d\n", s);
    return 0;
}