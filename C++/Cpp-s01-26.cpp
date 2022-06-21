#include <iostream>
#include <string>

// 9 223 372 036 854 775 807

int main() {
  std::string pusty = "";
  std::string napis = "string";


  std::cout << "pusty.max_size(): " << pusty.max_size() << '\n'
	    << "napis.max_size(): " << napis.max_size() << '\n';



  return 0;
}
