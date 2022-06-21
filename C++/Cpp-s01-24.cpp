#include <iostream>
#include <string>

int main() {
  std::string spiewak("Jan Kiepura");
  char napis[] = "Jan Kiepur";


  std::cout << "spiewak.length(): " << spiewak.length() << '\n'
	    << "spiewak.size(): " << spiewak.size() << "\n\n";

  std::cout << "sizeof(napis) = " << sizeof(napis) << '\n';



  return 0;
}
