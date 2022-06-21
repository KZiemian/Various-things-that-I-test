#include <iostream>

int fun(int& x) {
  return x;
}

int main() {
  int x = 0;


  std::cout << "fun(x): " << fun(x) << '\n';



  return 0;
}
