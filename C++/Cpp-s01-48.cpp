#include <iostream>
#include <typeinfo>

int main() {
  std::cout << "nullptr: " << nullptr << '\n'
	    << "typeid(std::nullptr).name(): " << typeid(nullptr).name()
	    << "\n\n";

  std::cout << "NULL: " << NULL << '\n'
	    << "std::typeid(NULL).name(): " << typeid(NULL).name() << '\n';



  return 0;
}
