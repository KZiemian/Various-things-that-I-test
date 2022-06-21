#include <iostream>

int main() {
  int tablica[10] = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9};


  for (int i = 0; i < 10; i++) {
    std::cout << "tablica[" << i << "] = " << tablica[i] << '\n';
  }

  std::cout << std::endl;
  std::cout << "tablica[0]: " << tablica[0] << '\n'
	    << "tablica: " << tablica << '\n';



  return 0;
}
