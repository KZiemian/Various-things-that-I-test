#include <iostream>

union skarbiec {
  char c;
  int i;
};

int main() {
  skarbiec varUnion;


  varUnion.c = 'A';
  std::cout << "varUnion.c: " << varUnion.c << '\n';
  std::cout << "varUnion.i: " << varUnion.i << '\n';



  return 0;
}
