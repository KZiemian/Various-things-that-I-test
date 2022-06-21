#include <iostream>

int main() {
  int a = 5;
  int b = 6;
  int& p = a;
  int& p = b;


  std::cout << "a: " << a << '\n';
  std::cout << "b: " << b << '\n';
  std::cout << "p: " << p << '\n';



  return 0;
}
