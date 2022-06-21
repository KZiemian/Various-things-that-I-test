#include <iostream>

int main() {
  int liczba = 0;
  int *wskInt = &liczba;
  char znak = 'a';
  char *wskZnak = &znak;
  void *wskVoid = wskInt;
  void *wskVoid2 = wskInt;


  std::cout << " liczba: " <<  liczba << '\n'
	    << "*wskInt: " << *wskInt << '\n'
	    << " wskInt: " <<  wskInt << '\n'
	    << " znak: " << znak << '\n'
	    << "*wskZnak: " << *wskZnak << '\n'
	    << " wskZnak: " <<  wskZnak << "\n\n";

  znak = 'b';
  std::cout << " znak: " << znak << '\n'
	    << "*wskZnak: " << *wskZnak << '\n'
	    << " wskZnak: " <<  wskZnak << "\n\n";

  // std::cout << "*wskVoid: " << *wskVoid << "\n\n";

  std::cout << " wskVoid: " <<  wskVoid << '\n'
	    << " wskVoid2: " << wskVoid2 << "\n\n";

  wskVoid = wskZnak;
  std::cout << " wskVoid: " << wskVoid << '\n';

  wskVoid2 = wskVoid;
  std::cout << " wskVoid2: " << wskVoid2 << '\n';



  return 0;
}
