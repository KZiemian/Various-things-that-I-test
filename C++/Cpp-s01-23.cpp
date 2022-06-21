#include <iostream>
#include <string>

int main() {
  for (int i = 0; i < 15; i++) {
    std::string nazwaPliku("urzadzenie_");

    int liczba = i;


    liczba %= 100;
    char dziesiatki = '0' + (liczba / 10);


    liczba %= 10;
    // char jednostki = '0' + (liczba / 1);
    char jednostki = '0' + liczba;


    nazwaPliku += dziesiatki;
    nazwaPliku += jednostki;
    nazwaPliku += ".parametry";

    std::cout << "Dla i = " << i << " nazwaPliku = " << nazwaPliku << '\n';
  }



  return 0;
}
