#include <stdio.h>

void collatzSequence(int number) {
    printf("Element of sequence: %d\n", number);

    if (number < 1) {
        printf("Number %d is smaller than 1.\n", number);
    } else {
        while (number != 1) {
            if ((number % 2) == 0) {
                number /= 2;
            } else if ((number % 2) == 1) {
                number = 3*number + 1;
            } else {
                printf("Somethings goes wrong!");
            }

            printf("Element of sequence: %d\n", number);
        }
    }
}

int main() {
    collatzSequence(0);
    printf("\n\n");

    collatzSequence(1);
    printf("\n\n");

    collatzSequence(2);
    printf("\n\n");

    collatzSequence(3);
    printf("\n\n");

    collatzSequence(4);
    printf("\n\n");

    collatzSequence(5);
    printf("\n\n");

    collatzSequence(10);
    printf("\n\n");

    collatzSequence(100);



    return 0;
}