#include <iostream>

int main() {
  int a = 10;
  int *p = &a;
  int *ptr = nullptr;
  int& ref = a;



  ptr = &a;

  std::cout << "a: " << a << '\n';
  std::cout << "*p: " << *p << '\n';
  std::cout << "*ptr: " << *ptr << '\n';
  std::cout << "ref: " << ref << '\n';



  return 0;
}
