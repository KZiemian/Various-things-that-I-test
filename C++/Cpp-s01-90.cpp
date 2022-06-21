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
  std::cout << '\n';


  for (auto i = v1.begin(); i != v1.end(); i++) {
    *i *= 2;
  }

  for (auto i = v1.begin(); i != v1.end(); i++) {
    std::cout << "*i: " << *i << '\n';
  }
  std::cout << '\n';

  // for (auto i = v1.begin(); i != v1.end(); i++) {
  //   std::cout << int(i) << std::endl;
  // }



  return 0;
}
