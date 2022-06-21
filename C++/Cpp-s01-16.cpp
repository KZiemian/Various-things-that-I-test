#include <iostream>
#include <string>

int main() {
  std::string hello = "Hello, World!";


  std::cout << "hello.length(): " << hello.length() << '\n';

  for (int i = 0; i < hello.length(); i++) {
    std::cout << "hello[" << i << "]: " << hello[i] << '\n';
  }



  return 0;
}
