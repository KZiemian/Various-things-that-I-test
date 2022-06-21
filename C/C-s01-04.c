#include <stdio.h>


#if __STDC_VERSION__ == 201710L
int compilerVersion = 18;
#elif __STDC_VERSION__ == 201112L
int compilerVersion = 11;
#elif __STDC_VERSION__ == 199901L
int compilerVersion = 99;
#elif __STDC_VERSION__ == 1
int compilerVersion = 90;
#else
int compilerVersion = -1;
#endif



int main() {
  if (compilerVersion == -1) {
    printf("Some werid compiler\n");
  } else {
    printf("C%d\n", compilerVersion);
  }



  return 0;
}
