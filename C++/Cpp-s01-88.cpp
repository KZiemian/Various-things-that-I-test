#include <iostream>
#include <vector>

int main() {
  std::vector<int> g1;


  for (int i = 1; i <= 5; i++) {
    g1.push_back(i);
  }

  for (auto i = g1.begin(); i != g1.end(); i++) {
    std::cout << "*i: " << *i << '\n';
  }
  std::cout << '\n';

  for (int i = 0; i < 5; i++) {
    std::cout << "g1[" << i << "]: " << g1[i] << '\n';
  }



  return 0;
}
