void f(widget& w) {
  use(w);
}

void g(widget *w) {
  if (w) {
    use(*w);
  }
}

void osoba::mojeWakacje() {
  zapisOsobyNaWycieczke(this);
}

void zapisOsobyNaWycieczke(const osoba *jejAdres);

class owoc {
public:
  int  pestka;
  void funkcja1();



private:
  int scisleTajne;
};

owoc cytryna, gruszka;

cytryna.pestka = 16;
gruszka.pestka = 12;

std::cout << "Moja cytryna ma " << cytryna.pestka << " pestek.";
std::cout << "Moja gruszka ma " << gruszka.pestka << " pestek.";

cytryna.scisleTajne = 6;

void spiew();

void opera::spiew();
void opera::spiew(int);
void opera::spiew(double*);
void opera::spiew(char*);

union skarbier {
  short int m;
  long l;
  char z;
};
