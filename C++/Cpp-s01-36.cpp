#include <iostream>

union skarbiec {
  short int i;
  long l;
  char c;
};

int main() {
  skarbiec varUnion;


  varUnion.i = 65;

  std::cout << "varUnion.i: " << varUnion.i << '\n';
  std::cout << "varUnion.l: " << varUnion.l << '\n';
  std::cout << "varUnion.c: " << varUnion.c << '\n';

  varUnion.c = 'x';
  std::cout << "varUnion.i: " << varUnion.i << '\n';
  std::cout << "varUnion.l: " << varUnion.l << '\n';
  std::cout << "varUnion.c: " << varUnion.c << '\n';



  return 0;
}
