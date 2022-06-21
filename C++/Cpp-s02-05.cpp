#include <iostream>

int main() {
  int a = 10;
  int &p = a;


  std::cout << " a: " << a  << '\n'
	    << " p: " << p  << '\n'
	    << "&a: " << &a << '\n'
	    << "&p: " << &p << '\n';



  return 0;
}
