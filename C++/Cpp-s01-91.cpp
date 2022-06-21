#include <iostream>
#include <vector>

int main() {
  std::vector<int> v1;


  for (int i = 0; i < 5; i++) {
    v1.push_back(i);
  }

  for (int i = 0; i < 5; i++) {
    std::cout << "v1[" << i << "]: " << v1[i] << '\n';
  }
  // std::cout << std::endl;

  // for (auto i = v1.cbegin(), i != v1.cend(), i++) {
  //   *i -= 3;
  // }
  // std::cout << std::endl;

  // for (int i = 0; i < 5; i++) {
  //   std::cout << "v1[" << i << "]: " << v1[i] << std::endl;
  // }



  return 0;
}
