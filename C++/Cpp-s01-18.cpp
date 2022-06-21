#include <iostream>
#include <string>

int main() {
  // std::string napis("0123456789ABCDEF", 11);
  std::string napis("0123456789ABCDEF", -3); // Niestety to też działa,
  // dostajemy tym błąd czasu działania.


  std::cout << "napis: " << napis << '\n';



  return 0;
}
