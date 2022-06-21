#include <iostream>

int main() {
  int x = 10;

  int& ref = x;


  std::cout << "x = " << x << '\n';
  std::cout << "ref = " << ref << '\n';

  ref = 20;
  std::cout << "x = " << x << '\n';
  std::cout << "ref = " << ref << '\n';

  x = 30;
  std::cout << "x = " << x << '\n';
  std::cout << "ref = " << x << '\n';



  return 0;
}
