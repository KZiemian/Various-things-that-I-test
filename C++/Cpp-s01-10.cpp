#include <iostream>

int main() {
  int tablica[11] = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
  int *wskInt = &tablica[4];


  for (int i = 0; i < 11; i++) {
    std::cout << "tablica[" << i << "] = " << tablica[i] << '\n';
  }

  std::cout << '\n';
  std::cout << "*wskInt = " << *wskInt << '\n'
	    << " wskInt = " <<  wskInt << "\n\n";

  wskInt = wskInt + 1;
  std::cout << "*wskInt = " << *wskInt << '\n'
	    << " wskInt = " <<  wskInt << "\n\n";

  wskInt++;
  std::cout << "*wskInt = " << *wskInt << '\n'
	    << " wskInt = " <<  wskInt << "\n\n";

  wskInt += 3;
  std::cout << "*wskInt = " << *wskInt << '\n'
	    << " wskInt = " <<  wskInt << '\n';



  return 0;
}
