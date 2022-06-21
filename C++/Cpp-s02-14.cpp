#include <iostream>

int main() {
  int i = 0;
  int *w = &i;


  std::cout << " i: " <<  i << '\n';
  std::cout << " w: " <<  w << '\n';
  std::cout << "*w: " << *w << '\n';

  *w = 1;
  std::cout << " i: " <<  i << '\n';
  std::cout << " w: " <<  w << '\n';
  std::cout << "*w: " << *w << '\n';



  return 0;
}
