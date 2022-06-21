#include <iostream>

int main() {
  int a = 10;
  void *ptrA = &a;
  // void& aref = a;


  std::cout << "a = " << a << '\n'
	    << "*(int*)(ptrA) = " << *(int*)(ptrA) << '\n'
	    << "ptrA = " << ptrA << '\n';



  return 0;
}
