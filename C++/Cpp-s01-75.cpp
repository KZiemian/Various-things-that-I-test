#include <iostream>

int main() {
  int x = 10;
  int *ptr = &x;
  int *&ptr1 = ptr;


  std::cout << "x: " << x << '\n'
	    << "*ptr: " << *ptr << '\n'
	    << "ptr: " << ptr << '\n'
	    << "ptr1: " << ptr1 << '\n';



  return 0;
}
