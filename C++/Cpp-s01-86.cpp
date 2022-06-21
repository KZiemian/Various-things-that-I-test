#include <iostream>
#include <vector>

int main() {
  std::vector<int> g1;


  for (int i = 1; i <= 5; i++) {
    g1.push_back(i);
  }

  // std::cout << "g1" << g1 << std::endl;
  std::cout << "Output of begin and end: ";

  for (auto i = g1.begin(); i != g1.end(); i++) {
    std::cout << *i << " ";
  }
  std::cout << std::endl << std::endl;

  // for (auto i = g1.begin(); i != g1.end(); i++) {
  //   std::cout << i << " ";
  // }

  std::cout << "Output of begin and end: ";
  for (auto i = g1.begin(); i != g1.end(); i++) {
    std::cout << *i << " ";
  }
  std::cout << '\n';



  return 0;
}
