#include <iostream>

class Geeks {
public:
  void func(int x) {
    std::cout << "Value of x is " << x << '\n';
  }

  void func(double x) {
    std::cout << "Value of x is " << x << '\n';
  }

  void func(int x, int y) {
    std::cout << "Value of x and y is " << x << ", " << y << '\n';
  }
};

int main() {
  Geeks obj1;


  obj1.func(7);
  obj1.func(9.132);
  obj1.func(85, 64);



  return 0;
}
