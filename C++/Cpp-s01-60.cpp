#include <iostream>
#include <vector>

int main() {
  std::vector<int> values = {1, 3, 5, 7};


  for (auto& v : values) {
    v *= 2;
  }

  for (int v : values) {
    std::cout << v << " ";
  }

  std::cout << '\n';



  return 0;
}
