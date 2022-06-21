#include <iostream>

union skrytka {
  char c;
  double f;
};

int main() {
  skrytka moja  = {'z'};
  skrytka twoja = {'x'};


  std::cout << "moja.c: " << moja.c << '\n'
	    << "moja.f: " << moja.f << '\n';

  std::cout << "twoja.c: " << twoja.c << '\n'
	    << "twoja.f: " << twoja.f << '\n';

  // skrytka jego = {'3.14'};
  // std::cout << "jego.f: " << jego.f << std::endl
  // 	    << "jego.c: " << jego.c << std::endl;



  return 0;
}
