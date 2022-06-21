#include <iostream>

int main() {
  int i = 3;

  int *ptr = &i;

  int &ref = i;


  std::cout << " i: " << i << '\n';
  std::cout << "*i: " << *ptr << '\n';
  std::cout << " ref: " << ref << '\n';
}
