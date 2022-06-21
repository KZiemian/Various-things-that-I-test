#include <iostream>
#include <string>

int main() {
  std::string text = "sample text";


  for (const char& c : text) {
    std::cout << c << '\n';
  }

  return 0;
}
