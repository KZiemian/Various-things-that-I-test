#include <iostream>

int main() {
  int tabInt[10] = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9};
  double tabDouble[10] = {0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0};

  int *wskInt = &tabInt[0];
  double *wskDouble = &tabDouble[0];


  for (int i = 0; i < 10; i++) {
    *wskDouble = double(i) / 10.0;
    wskDouble++;
  }

  std::cout << "Tresc tablic na poczatku" << '\n';
  wskInt = &tabInt[0];
  wskDouble = &tabDouble[0];

  for (int i = 0; i < 10; i++) {
    std::cout << i << ") " << *wskInt << "\t" << *wskDouble << '\n';
    wskInt++;
    wskDouble++;
  }

  wskInt = &tabInt[5];
  wskDouble = &tabDouble[0] + 2;

  for (int j = 0; j < 4; j++) {
    // std::cout << "Trzecia petla" << std::endl;
    // std::cout << "wskInt: " << wskInt << std::endl;
    // std::cout << "wskDouble: " << wskDouble << std::endl;
    *wskInt = -222;
    *wskDouble = -777.5;
    // std::cout << "*wskInt: " << *wskInt << std::endl;

    wskInt++;
    wskDouble++;
  }

  std::cout << "Tresc tablic po wstawieniu nowych wartosci" << '\n';
  wskInt = &tabInt[0];
  wskDouble = &tabDouble[0];

  for (int p = 0; p < 10; p++) {
    std::cout << "tabInt[" << p << "] = " << *wskInt << "    "
	      << "tabDouble[" << p << "] = " << *wskDouble << '\n';
    wskInt++;
    wskDouble++;
  }



  return 0;
}
