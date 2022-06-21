#include <iostream>

#include "Cpp_osoba_s01_01.h"
#include "Cpp_bilet_s01_01.h"

void prezentacja(osoba ktos) {
  std::cout << "Mam zaszczyt przedstawic panstwu," << '\n'
	    << "Oto we wlasnej osobie:";
  ktos.wypisz();
}

void funkcjaWInnymPliku();



int main() {
  osoba kompozytor, autor;


  kompozytor.zapamietaj("Fryderyk Chopin", 36);
  autor.zapamietaj("Marcel Proust", 34);

  prezentacja(kompozytor);
  prezentacja(autor);

  std::cout << '\n' << "Uzywamy w tym pliku tez klasy 'bilet'" << '\n';


  bilet zolty, niebieski;


  zolty.zapamietaj("Frankfurt", "Paris", bilet::ekspres, 1);
  zolty.wypisz();

  niebieski.zapamietaj("Zurich", "Genewa", bilet::przyspieszony, 2);
  niebieski.wypisz();

  niebieski.zmienRodzajPociagu(bilet::pospieszny);
  std::cout << '\n' << "Po zmianie tego biletu..." << '\n';
  niebieski.wypisz();

  funkcjaWInnymPliku();



  return 0;
}
