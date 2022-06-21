#include <iostream>

int main() {
  int a = 10;
  int *p = &a;
  int **q = &p;


  std::cout << " a: " <<  a << '\n'
	    << "*p: " << *p << '\n'
	    << " p: " <<  p << '\n'
	    << "*q: " << *q << '\n'
	    << " q: " <<  q << '\n';



  return 0;
}
