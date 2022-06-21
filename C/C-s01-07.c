#include <stdio.h>



int main(int argc, char *argv[]) {
  int n = 0;
  int optarg = 0;


  printf("\nargc = %d\n", argc);

  for (n = 0; n < argc; ++n) {
    printf("*argc[%d]: %s\n", n, argv[n]);
  }
  printf("\n");


  optarg = getopt(argc, argv, "a:b:c");

  while (optarg != 1) {
    switch (optarg) {
    case 'a':
      printf("case \'a\'\n");
      printf("    optarg = %s =\n", optarg);

      break;


    case 'b':
      printf("case \'b\'\n");
      printf("    optarg = %s = \n", optarg);

      break;


    case 'c':
      printf("case \'c\'\n");
      printf("    optarg = %s = \n", optarg);

      break;
    }


    optarg = getopt(argc, argv, "a:b:c");
  }



  return 0;
}
