#include <iostream>

int main() {
  int i = 0;
  int *ptr = nullptr;


  std::cout << "i: " << i << '\n'
	    << "ptr: " << ptr << "\n\n";

  ptr = &i;
  *ptr = 1;
  std::cout << " i: " << i << '\n'
	    << "*ptr: " << *ptr << '\n';



  return 0;
}
