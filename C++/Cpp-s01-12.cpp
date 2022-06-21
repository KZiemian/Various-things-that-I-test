#include <iostream>


int main() {
  int tabInt[6] = {0, 1, 2, 3, 4, 5};
  double tabDouble[6] = {0.0, 1.0, 2.0, 3.0, 4.0, 5.0};

  int *wskInt = &tabInt[0];
  double *wskDouble = &tabDouble[0];


  std::cout << "Oto, jak przy inkrementacji wskaznikow" << '\n'
	    << "zmieniaja sie ukryte w nich adresy:" << '\n';

  for (int i = 0; i < 6; i++) {
    std::cout << "i = " << i << ") wskInt = " << wskInt
	      << ", wskDouble = " << wskDouble << '\n';
    wskInt++;
    wskDouble++;
  }



  return 0;
}
