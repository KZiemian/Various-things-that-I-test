#include <iostream>
#include <string>

int main() {
  std::string stary("Mieszko");
  std::string nowy = "";


  std::cout << "stary: " << stary << '\n'
	    << "nowy: " << nowy << "<--" << '\n';

  nowy = stary;
  std::cout << "stary: " << stary << '\n'
	    << "nowy: " << nowy << "<--" << '\n';



  return 0;
}
