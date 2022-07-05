#include <stdio.h>

void zeroIntArray(int array[], int length) {
    for (int i = 0; i < length; i++) {
        array[i] = 0;
    }
}

int main() {
    int funnyNumbers[10];


    zeroIntArray(funnyNumbers, 10);

    for (int i = 0; i < 10; i++) {
        printf("funnyNumber[%d] = %d\n", i, funnyNumbers[i]);
    }
    printf("\n\n");

    funnyNumbers[0] = 90;
    funnyNumbers[1] = 80;

    for (int i = 0; i < 10; i++) {
        printf("funnyNumber[%d] = %d\n", i, funnyNumbers[i]);
    }



    return 0;
}