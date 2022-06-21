#include <iostream>

void pokazywacz(const int *wsk, int ile) {
  std::cout << "Dziala pokazywasz\n";

  for (int i = 0; i < ile; i++) {
    // *wsk += 22;
    std::cout << "Element nr. " << i << " ma wartosc " << *wsk << '\n';
    wsk++;
  }
}

void zmieniacz(int *wsk, int ile) {
  std::cout << "Dziala zmieniacz\n";

  for (int i = 0; i < ile; i++) {
    *wsk += 500;
    std::cout << "Element nr. " << i << " ma wartosc " << *wsk << '\n';
    wsk++;
  }
}

int main() {
  int tablica[4] = {110, 120, 130, 140};


  pokazywacz(tablica, 4);
  zmieniacz(tablica, 4);
  pokazywacz(tablica, 4);

  std::cout << "Dla potwierdzenia tablica[3] = " << tablica[3] << '\n';



  return 0;
}
