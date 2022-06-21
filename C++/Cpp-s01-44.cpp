#include <iostream>

int main() {
  int arrayVar[5] = {0, 1, 2, 3, 4};
  int *ptrRed = &arrayVar[3];
  int *ptrGreen = &arrayVar[0];
  int i = 0;


  std::cout << "Mamy piecioelementowa tablice arrayVar" << '\n'
	    << "Wskaznik ptrRed pokazuje na element o indeksie 3"
	    << '\n'
	    << "Na ktory element ma pokazywac wskaznik zielony? (0-4)"
	    << '\n';
  i = 4;

  if (i < 0 || i > 4) {
    std::cout << "Nie ma takiego elementu w tej tablicy!" << '\n';
  } else {
    ptrGreen = &arrayVar[i];

    std::cout << '\n'
	      << "Z przeprowadzonego porownania wskaznikow" << '\n'
	      << "ptrRed z ptrGreen wynika, że:" << '\n';

    if (ptrRed > ptrGreen) {
      std::cout << "ptrGreen pokazuje na element blizej poczatku tablicy"
		<< '\n';
    } else if (ptrRed < ptrGreen) {
      std::cout << "ptrGreen pokazuje na element o wyzszym indeksie"
		<< '\n';
    } else {
      std::cout << "ptrGreen i ptrRed pokazuje na to samo" << '\n';
    }
  }



  return 0;
}
