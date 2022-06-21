#include <iostream>

int& func() {
  int x = 10;


  return x;
}

int main() {
  func() = 30;
  std::cout << "func(): " << func() << '\n';



  return 0;
}
