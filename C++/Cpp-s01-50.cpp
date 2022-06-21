#include <iostream>

void hydraulik(int *wsk_do_kranu) {
  *wsk_do_kranu = 100;
}

int main() {
  int kran = -1;


  std::cout << "Stan techniczny kranu = " << kran << '\n';

  hydraulik(&kran);

  std::cout << "Po wezwaniu hydraulika stan techniczny kranu = " << kran
	    << '\n';



  return 0;
}
