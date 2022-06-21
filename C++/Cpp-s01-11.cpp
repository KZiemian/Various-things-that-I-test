#include <iostream>

int main() {
  int tablica[11] = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
  int *wskInt = &tablica[0];


  for (int i = 0; i < 11; i++) {
    std::cout << " tablica[" << i << "] = " << tablica[i] << '\n';
  }

  std::cout << std::endl;

  std::cout << " tablica: " << tablica << '\n'
	    << "*wskInt: " << *wskInt << '\n'
	    << " wskInt: " <<  wskInt << '\n';

  tablica[0] = 36;
  std::cout << " tablica[0]: " << tablica[0] << '\n'
	    << "*wskInt: " << *wskInt << "\n\n";

  for (int i = 0; i < 11; i++) {
    std::cout << " tablica[" << i << "] = " << tablica[i] << '\n';
  }



  return 0;
}
