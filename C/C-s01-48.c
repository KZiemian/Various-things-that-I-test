#include <stdio.h>

int main() {
    int funnyNumbers[6] = {4, 8, 15, 16, 23, 41};


    for (int i = 0; i < 6; i++) {
        printf("funnyNumbers[%d] = %d\n", i, funnyNumbers[i]);
    }
    printf("\n\n");
    

    funnyNumbers[1] = 200;

    for (int i = 0; i < 6; i++) {
        printf("funnyNumbers[%d] = %d\n", i, funnyNumbers[i]);
    }



    return 0;
}