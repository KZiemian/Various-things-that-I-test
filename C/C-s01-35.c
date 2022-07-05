#include <stdio.h>

int main() {
    double numberVar = 0.0;
    double *ptrDouble = &numberVar;


    printf("ptrDouble = %d\n\n", ptrDouble);

    printf("Eneter some double number: ");
    scanf("%lf", ptrDouble);

    printf("You entered %f.\n", numberVar);
    printf("ptrDouble = %d.\n", ptrDouble);



    return 0;
}