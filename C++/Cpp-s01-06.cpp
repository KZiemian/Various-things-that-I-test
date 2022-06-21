#include <iostream>

int main() {
  int varInt1 = 0;
  int varInt2 = 1;
  double varDouble = 0.5;

  int *wskInt1 = &varInt1;
  int *wskInt2 = &varInt2;
  double *wskDouble = &varDouble;


  std::cout << " varInt1: " << varInt1 << '\n'
	    << "*wskInt1: " << *wskInt1 << '\n'
	    << " wskInt1: " << wskInt1 << "\n\n";

  std::cout << " varInt2: " << varInt2 << '\n'
	    << "*wskInt2: " << *wskInt2 << '\n'
	    << " wskInt2: " << wskInt2 << "\n\n";

  std::cout << " varDouble:" << varDouble << '\n'
	    << "*wskDouble: " << varDouble << '\n'
	    << " wskDouble: " << wskDouble << "\n\n";


  wskInt1 = wskInt2;

  std::cout << "*wskInt1: " << *wskInt1 << '\n'
	    << " wskInt1: " << wskInt1 << "\n\n";

  std::cout << "*wskInt2: " << *wskInt2 << '\n'
	    << " wskInt2: " << wskInt2 << "\n\n";

  // wskDouble = wskInt1;
  wskDouble = reinterpret_cast<double*>(wskInt1);
  std::cout << "*wskInt1: " << *wskInt1 << '\n'
	    << " wskInt1: " << wskInt1 << "\n\n";

  std::cout << "*wskDouble: " << *wskDouble << '\n'
	    << " wskDouble: " << wskDouble << '\n';



  return 0;
}
