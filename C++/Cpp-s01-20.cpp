#include <iostream>
#include <string>

int main() {
  std::string sciezka("C:/katalogA/");
  std::string nazwaPliku("wspolczynniki.txt");
  std::string dlugaNazwa = "";


  std::cout << "sciezka: " << sciezka << '\n'
	    << "nazwaPliku: " << nazwaPliku << "\n\n";

  dlugaNazwa = sciezka + nazwaPliku;
  std::cout << "dlugaNazwa: " << dlugaNazwa << '\n';



  return 0;
}
