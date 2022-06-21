#include <iostream>

void swap(std::string& string1, std::string& string2) {
  std::string temp = string1;


  string1 = string2;
  string2 = temp;
}

int main() {
  std::string string1 = "GEEKS";
  std::string string2 = "FOR GEEKS";


  swap(string1, string2);

  std::cout << "string1 is " << string1 << '\n';
  std::cout << "string2 is " << string2 << '\n';
}
