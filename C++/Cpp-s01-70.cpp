#include <iostream>

int &fun() {
  static int x = 10;


  return x;
}

int main() {
  std::cout << "fun(): " << fun() << '\n';
  fun() = 30;
  std::cout << "fun(): " << fun() << '\n';



  return 0;
}
