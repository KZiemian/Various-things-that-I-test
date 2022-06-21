#include <iostream>
#include <string>

int main() {
  std::string zlepek("nazwa");
  std::string b(".ext");


  std::cout << "zlepek: " << zlepek << '\n'
	    << "b: " << b << "\n\n";

  zlepek += b;
  std::cout << "zlepek: " << zlepek << '\n';



  return 0;
}
