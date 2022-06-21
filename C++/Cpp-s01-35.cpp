#include <iostream>

union skarbiec {
  short int m;
  long l;
  char z;
};

int main() {
  skarbiec varUnion;


  varUnion.m = 1;

  std::cout << "varUnion: " << varUnion.m << '\n';

  varUnion.l = 2;

  std::cout << "varUnion: " << varUnion.l << '\n';

  varUnion.z = 'a';

  std::cout << "varUnion: " << varUnion.z << '\n';



  return 0;
}
