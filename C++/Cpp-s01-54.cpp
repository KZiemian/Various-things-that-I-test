#include <iostream>

int main() {
  const int pojemnosc = 5;
  const int *staly_wsk = nullptr;
  // int *zwykly_wsk = nullptr;


  staly_wsk = &pojemnosc;
  // zwykly_wsk = &pojemnosc;

  std::cout << "*staly_wsk: " << *staly_wsk << '\n';



  return 0;
}
