#include <iostream>
#include <vector>

int main() {
  std::vector<int> v1;


  for (int i = 1; i <= 10; i++) {
    v1.push_back(i * 10);
  }

  std::cout << "Reference operator [g]: v1[2] = " << v1[2] << '\n';
  std::cout << "at: v1.at(4) = " << v1.at(4) << '\n';
  std::cout << "front(): v1.front() = " << v1.front() << '\n';
  std::cout << "back(): v1.back() = " << v1.back() << '\n';


  int *pos = v1.data();


  std::cout << "The first element is " << *pos << '\n';



  return 0;
}
