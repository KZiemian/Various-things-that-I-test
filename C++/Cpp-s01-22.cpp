#include <iostream>
#include <string>

int main() {
  std::string napis = "string";


  std::cout << "napis: " << napis << "\n\n";

  napis += '0';
  napis += '1';
  napis += '2';
  std::cout << "napis: " << napis << '\n';



  return 0;
}
