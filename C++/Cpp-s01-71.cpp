#include <iostream>

int fun(const int& x) {
  return x;
}

int main() {
  std::cout << "fun(10): " << fun(10) << '\n';



  return 0;
}
