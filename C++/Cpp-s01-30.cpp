#include <iostream>
#include <cstring>

class osoba {
public:
  void zapamietaj(const char *napis, int lata);

  void wypisz() {
    std::cout << "\t" << nazwisko << ", lat: " << wiek << '\n';
  }



private:
  char nazwisko[80];
  int  wiek;
};

void osoba::zapamietaj(const char *napis, int lata) {
  strcpy(nazwisko, (napis ? napis : "Anonim"));
  wiek = lata;
}


int main() {
  osoba student1,
    student2,
    profesor,
    pilot;
  char magazynek[80];
  int  wiek;


  std::cout << "Dla informacji podaje, ze jeden obiekt klasy osoba\n"
	    << "ma rozmiar: "
	    << sizeof(osoba)
	    << " bajty. To samo inaczej: "
	    << sizeof(student1) << '\n';

  profesor.zapamietaj("Albert Einstein", 55);
  student1.zapamietaj("Ruediger Schubart", 26);
  student2.zapamietaj("Claudia Bach", 25);
  pilot.zapamietaj("Neil Armstrong", 37);

  std::cout << "Po wpisaniu informacji do obiektow. Sprawdzamy.\n"
	    << "Dane z obiektu profesor:\n";
  profesor.wypisz();

  std::cout << '\n'
	    << "Dane z obiektu student1:" << '\n';
  student1.wypisz();

  std::cout << '\n'
	    << "Dane z obiektu student2:" << '\n';
  student2.wypisz();

  std::cout << '\n'
	    << "Dane z obiektu pilot:" << '\n';
  pilot.wypisz();

  std::cout << '\n'
	    << "Podaj swoje nazwisko (tylko nazwisko): ";
  std::cin >> magazynek;

  std::cout << '\n'
	    << "Podaj swoj wiek: ";
  std::cin >> wiek;
  pilot.zapamietaj(magazynek, wiek);

  std::cout << '\n'
	    << "Oto dane ktore teraz sa zapamietane w obiektach "
	    << "w obiektach profesor i pilot" << '\n';
  profesor.wypisz();
  pilot.wypisz();



  return 0;
}
