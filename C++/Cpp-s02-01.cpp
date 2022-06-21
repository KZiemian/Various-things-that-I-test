#include <iostream>

int main() {
  int a = 10;
  int& p;


  p = a;

  std::cout << "a: " << a << '\n';
  std::cout << "p: " << p << '\n';



  return 0;
}
