#include <stdio.h>
#include <math.h>

int main() {
    double exponentVar = 0.5;


    for (int i = 0; i < 20; i++) {
        printf("pwo(2.0, %f) = %f\n", exponentVar, pow(2.0, exponentVar));

        exponentVar += 0.5;
    }



    return 0;
}