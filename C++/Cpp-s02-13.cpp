#include <iostream>

void swap(int& x, int& y) {
  int z = x;


  x = y;
  y = z;
}

int main() {
  int a = 45;
  int b = 35;


  std::cout << "Before swap\n";
  std::cout << "a = " << a << ", b = " << b << "\n";

  swap(a, b);

  std::cout << "After swap with pass by reference\n";
  std::cout << "a = " << a << ", b = " << b << "\n";



  return 0;
}
