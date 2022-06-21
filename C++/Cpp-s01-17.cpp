#include <iostream>
#include <iomanip>
#include <string>

void pokaz(std::string opis, std::string wlasciwy) {
  std::cout << "Tresc obietku " << std::setw(15)
	    << opis << ": -->" << wlasciwy << "<--" << '\n';
}

int main() {
  std::string napisA = "";
  std::string napisB1("Jakis tekst");

  char tablica[20] = { "Natenczas Wojski" };

  std::string napisB2(&tablica[0]);
  std::string wiadomosc(&tablica[5]);
  std::string ostrzezenie("Awaria studni", 8);
  std::string gwiazdki(25, '*');
  std::string inny = "ABCDEFGH";
  std::string nowy(inny, 4, 2);


  pokaz("napisA", napisA);
  pokaz("napisB1", napisB1);
  pokaz("napisB2", napisB2);
  pokaz("wiadomosc", wiadomosc);
  pokaz("ostrzezenie", ostrzezenie);
  pokaz("gwiazdki", gwiazdki);
  pokaz("nowy", nowy);



  return 0;
}
