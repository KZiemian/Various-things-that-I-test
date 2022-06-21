#include <iostream>

int main() {
  int zmienna = 8;
  int drugi = 4;
  int *wskaznik = &zmienna;


  std::cout << "zmienna = " << zmienna << '\n'
	    << "a odczytana przez wskaznik = " << *wskaznik << '\n';

  zmienna = 10;
  std::cout << "zmienna = " << zmienna << '\n'
	    << "\na odczytana przez wskaznik = " << *wskaznik << '\n';

  *wskaznik = 200;
  std::cout << "zmienna = " << '\n'
	    << "a odczytana przez wskaznik = " << *wskaznik << '\n';

  wskaznik = &drugi;
  std::cout << "zmienna = " << zmienna << '\n'
	    << "a odczytana przez wskaznik = " << *wskaznik << '\n';



  return 0;
}
