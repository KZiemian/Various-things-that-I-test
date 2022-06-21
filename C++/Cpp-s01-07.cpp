#include <iostream>

int main() {
  int liczba = 0x0a;
  int *wskInt = &liczba;


  std::cout << " liczba: " <<  liczba << '\n';
  std::cout << "*wskInt: " << *wskInt << '\n';
  std::cout << " wskInt: " <<  wskInt << '\n';

  wskInt = reinterpret_cast<int*>(liczba);
  std::cout << "wskInt: " << wskInt << '\n';
  std::cout << "reinterpret_cast<int>(wskInt): "
    << reinterpret_cast<int>(wskInt) << '\n';



  return 0;
}
