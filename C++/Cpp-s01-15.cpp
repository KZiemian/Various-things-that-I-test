#include <iostream>
#include <string>

void plakat(std::string tresc) {
  std::string gwiazdki(tresc.length() + 6, '*');

  std::cout << gwiazdki << '\n'
	    << gwiazdki << "\r"
	    << "** " << tresc << " **\n"
	    << gwiazdki << '\n';
}

int main() {
  std::string klub1("Wisla");
  std::string klub2("Unia");

  std::string miasto1 = "Krakow";
  std::string miasto2 = "Tarnow";

  std::string polaczenie = "";


  polaczenie = klub1 + klub2;
  std::cout << "Dodanie dwoch stringow[" << polaczenie << "]\n\n";

  std::string zapowiedz = klub1 + "-" + miasto1;
  zapowiedz = zapowiedz + " vs. ";

  zapowiedz += klub2;
  zapowiedz += ('-' + miasto2);

  std::cout << "!!! Zapraszamy na mecz !!!\n";
  plakat(zapowiedz);

  zapowiedz += " Rezultat meczu ";
  std::string wynik = "";
  std::cin >> wynik;

  std::cout << "\n\nUWAGA: Wczoraj odbyl sie mecz\n";
  plakat(zapowiedz + wynik);



  return 0;
}
