class osoba {
public:
  void zapamietaj(const char*, int);
  void wypisz();



private:
  char nazwisko[80];
  int wiek;
};

osoba student1, student2, profesor, pilot;

class paszport {
private:
  char nazwisko[40];
  char imie[40];
  int numer;
  // Błędna instrukcja.
  // int wzrost = 176;
};

profesor.zapamietaj("Albert Einsterin", 55);

student1.zapamietaj("Ruediger Schubart", 26);
student2.zapamietaj("Claudia Bach", 25);
pilot.zapamietaj("Neil Armstrong", 37);

osoba *wsk;
wsk = &profesor;

wsk->zapamietaj("Albert Einstein", 55);

osoba &belfer = profesor;

belfer.zapamietaj("Albert Einstein", 55);

class osoba {
public:
  void zapamietaj(const char *napis, int lata) {
    strcpy(nazwisko, (napis ? napis : "Anonim"));
    wiek = lata;
  }

  void wypisz() {
    std::cout << nazwisko << ", lat: " << wiek << '\n';
  }



private:
  char nazwisko[80];
  int  wiek;
};

class osoba {
public:
  void zapamietaj(const char *napis, int lata);
  void wypisz();



private:
  char nazwisko[80];
  int  wiek
};



void osoba::zapamietaj(const char *napis, int lat) {
  strcpy(nazwisko, (napis ? napis : "Anonim"));
  wiek = lata;
}

void osoba::wypisz() {
  std::cout << nazwisko << ", lat:" << wiek << '\n';
}

inline void osoba::wypisz() {
  // ...
}

for (auto& e : c) {
  // ...
  use(e);
  // ...
 }

for ( e : c ) {
  // ...
  use(e);
  // ...
 }

unique_ptr<widget> factory();
void caller() {
  auto w = factory();
  auto g = make_unique<gadget>();
  use(*w, *g);
}
