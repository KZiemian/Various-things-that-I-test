#include <iostream>

int main() {
  int *ptr = nullptr;
  int *&ref = ptr;


  std::cout << "ptr: " << ptr << '\n'
	    << "ref: " << ref << '\n';



  return 0;
}
