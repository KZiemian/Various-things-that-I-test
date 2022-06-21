#include <iostream>

void funkcjaWska(int *wsk, int rozmiar) {
  std::cout << "\nWewnatrz funkcji funkcja_wska\n";

  for (int i = 0; i < rozmiar; i++) {
    std::cout << *wsk << "\t";
    wsk++;
  }
}

void funkcjaTabl(int tab[], int rozmiar) {
  std::cout << "\nWewnatrz funkcja funkcja_tabl\n";

  for (int i = 0; i < rozmiar; i++) {
    std::cout << tab[i] << "\t";
  }
}

void funkcjaWsk2(int *wsk, int rozmiar) {
  std::cout << "\nWewnatrz funkcji funkcja_wsk2\n";

  for (int i = 0; i < rozmiar; i++) {
    std::cout << wsk[i] << "\t";
  }
}

int main() {
  int tafla[4] = {5, 10, 15, 20};


  funkcjaTabl(tafla, 4);
  funkcjaWska(tafla, 4);
  funkcjaWsk2(tafla, 4);

  std::cout << '\n';



  return 0;
}
