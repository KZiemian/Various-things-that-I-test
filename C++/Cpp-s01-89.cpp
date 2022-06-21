#include <iostream>
#include <vector>

int main() {
  std::vector<int> g1;


  for (int i = 0; i < 5; i++) {
    g1.push_back(i);
  }

  for (int i = 0; i < 5; i++) {
    std::cout << "g1[" << i << "]: " << g1[i] << '\n';
  }

  for (int i = 0; i < 5; i++) {
    g1[i] += 2;
  }

  std::cout << "\nDruga petla:\n";
  for (int i = 0; i < 5; i++) {
    std::cout << "g1[" << i << "]: " << g1[i] << '\n';
  }

  for (int i = 0; i < 5; i++) {
    g1[i] = i;
  }

  std::cout << "\nTrzecia petla:\n";
  for (int i = 0; i < 5; i++) {
    std::cout << "g1[" << i << "]: " << g1[i] << '\n';
  }

  for (int i = 0; i < 5; i++) {
    g1[i] *= 2;
  }

  std::cout << "\nCzwarta petla:\n";
  for (int i = 0; i < 5; i++) {
    std::cout << "g1[" << i << "]: " << g1[i] << '\n';
  }



  return 0;
}
