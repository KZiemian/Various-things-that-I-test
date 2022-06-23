#include <iostream>

int main() {
  int cale = 0;
  double centymetry, przelicznik = 2.54;


  std::cout << "Podaj dlugosc w calach: ";
  std::cin >> cale;
  centymetry = cale * przelicznik;

  std::cout << cale << " cali, to " << centymetry
	    << " centymetrow\n";



  return 0;
}
