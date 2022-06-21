#include <iostream>

void swap(int& first, int& second) {
  int temp = first;


  first = second;
  second = temp;
}

int main() {
  int a = 2;
  int b = 3;


  std::cout << "a = " << a << '\n'
	    << "b = " << b << '\n';

  swap(a, b);
  std::cout << "a = " << a << '\n'
	    << "b = " << b << '\n';
}
