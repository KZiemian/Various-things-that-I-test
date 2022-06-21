#include <stdio.h>

int main() {
  FILE *fp = fopen("./test.txt", "r");
  char buff[255];


  fscanf(fp, "%s", buff);
  printf("1: %s\n", buff);


  fgets(buff, 255, (FILE *) fp);
  printf("2: %s\n", buff);


  fgets(buff, 255, (FILE *) fp);
  printf("3: %s\n", buff);


  fclose(fp);



  return 0;
}
