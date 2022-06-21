#include <stdio.h>

int main() {
  FILE *fp1 = NULL;
  char c = 'a';


  fp1 = fopen("test.dat", "r");

  while (!feof(fp1)) {
    c = getc(fp1);

    printf("%c %d feof = %d\n", c, c, feof(fp1));
  }



  return 0;
}
