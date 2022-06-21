#include <iostream>

int main() {
  int a = 10;
  int& p = a;
  int& q = p;


  std::cout << "a: " << a << '\n';
  std::cout << "p: " << p << '\n';
  std::cout << "q: " << q << '\n';

  p = 9;
  std::cout << "a: " << a << '\n';
  std::cout << "p: " << p << '\n';
  std::cout << "q: " << q << '\n';

  q = 8;
  std::cout << "a: " << a << '\n';
  std::cout << "p: " << p << '\n';
  std::cout << "q: " << q << '\n';



  return 0;
}
