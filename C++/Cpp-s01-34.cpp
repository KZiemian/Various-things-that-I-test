#include <iostream>

#include "Cpp_osoba_s01_02.h"

void prezentacja(osoba ktos) {
  std::cout << "Mam zaszczyt przedstawic panstwu," << '\n'
	    << "Oto we wlasnej osobie: ";

  ktos.wypisz();
}

int main() {
  osoba kompozytor, autor;


  kompozytor.zapamietaj("Fryderyk Chopin", 36);
  autor.zapamietaj("Marcel Proust", 34);

  prezentacja(kompozytor);
  prezentacja(autor);



  return 0;
}
