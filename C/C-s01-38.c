#include <stdio.h>

int main() {
    char nameAndSurname[20];


    printf("Enter your name and surname: ");
    fgets(nameAndSurname, 20, stdin);
    printf("Your name is %s.\n", nameAndSurname);



    return 0;
}