#include <iostream>
#include <vector>

int main() {
  std::vector<int> v1;


  for (int i = 1; i <= 5; i++) {
    v1.push_back(i);
  }

  std::cout << "Size: " << v1.size() << '\n';
  std::cout << "Capacity: " << v1.capacity() << '\n';
  std::cout << "Max_size: " << v1.max_size() << '\n';

  v1.resize(4);

  std::cout << "\nSize: " << v1.size() << '\n';

  if (v1.empty() == false) {
    std::cout << "Vector is not empty\n";
  } else {
    std::cout << "Vector is empty\n";
  }

  v1.shrink_to_fit();
  std::cout << "Vector elements are: \n";
  for (auto it = v1.cbegin(); it != v1.cend(); it++) {
    std::cout << "*it: " << *it << '\n';
  }



  return 0;
}
