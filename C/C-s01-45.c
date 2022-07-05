#include <stdio.h>

int main() {
    char color[20] = "";
    char pluralNoun[20] = "";
    char celebrityFirstName[20] = "";
    char celebrityLastName[20] = "";


    printf("Enter a color: ");
    scanf("%s", &color);

    printf("Enter a plural noun: ");
    scanf("%s", &pluralNoun);

    printf("Enter a celebrity name: ");
    scanf("%s %s", &celebrityFirstName, &celebrityLastName);


    printf("Roses are %s\n", color);
    printf("%s are blue\n", pluralNoun);
    printf("I love %s %s\n", celebrityFirstName, celebrityLastName);



    return 0;
}