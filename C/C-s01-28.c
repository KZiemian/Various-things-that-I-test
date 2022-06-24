#include <stdio.h>
#include <math.h>

int main() {
    double argumentDouble = 0.0;


    printf("sqrt(2.0) = %f\n", sqrt(2.0));

    for (int i = 1; i <= 10; i++) {
        argumentDouble = (double)(i);

        printf("sqrt(%f) = %f\n", argumentDouble, sqrt(argumentDouble));
    }



    return 0;
}