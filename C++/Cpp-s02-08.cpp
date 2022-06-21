#include <iostream>

void func(int& x) {
  x = 20;
}

int main() {
  int x = 10;


  std::cout << "x: " << x << '\n';
  func(x);
  std::cout << "New value of x is: " << x << '\n';



  return 0;
}
