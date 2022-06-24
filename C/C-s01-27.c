#include <stdio.h>
#include <math.h>

int main() {
    double exponentVarWhole = 1.0;
    double exponentVarFraction = 1.5;
    double powerVarWhole = 1.0;
    double powerVarFraction = 1.0;
    double rationLoc = 1.0;


    for (int i = 1; i < 20; i++) {
        exponentVarWhole = (double)(i);
        exponentVarFraction = exponentVarWhole + 0.5;
        powerVarWhole = pow(2.0, exponentVarWhole);
        powerVarFraction = pow(2.0, exponentVarFraction);

        rationLoc = powerVarFraction / powerVarWhole - 1;

        printf("(pow(2.0, %f) - pow(2.0, %f)) / pow(2.0, %f) = %f\n",
            exponentVarFraction, exponentVarWhole, exponentVarWhole, 
            rationLoc);
    }

    printf("pow(2.0, 0.5) - 1.0 = %f\n", pow(2.0, 0.5) - 1.0);



    return 0;
}