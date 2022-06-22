#include <stdio.h>

int main() {
    char characterName[] = "John";
    int characterAge = 67;


    printf("There once was a man named %s\n", characterName);
    printf("he was %d years old.\n", characterAge);

    characterAge = 99;

    printf("He really liked the name %s.\n", characterName);
    printf("but did not like being %d.\n", characterAge);



    return 0;
}