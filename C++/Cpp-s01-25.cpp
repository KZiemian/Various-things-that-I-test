#include <iostream>
#include <string>

int main() {
  std::string napis1 = "string";
  std::string napis2 = "";


  std::cout << "napis1: " << napis1 << '\n'
	    << "napis1.empty(): " << napis1.empty() << "\n\n";

  std::cout << "napis2: " << napis2 << '\n'
	    << "napis2.empty(): " << napis2.empty() << "\n\n";

  if (napis2.empty() == true) {
    std::cout << "string napis1 jest pusty" << '\n';
  }



  return 0;
}
