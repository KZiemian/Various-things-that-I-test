#include <iostream>
#include <vector>

int main() {
  std::vector<int> v;


  v.assign(5, 10);

  std::cout << "The vector elements are: ";
  for (int i = 0; i < v.size(); i++) {
    std::cout << v[i] << " ";
  }
  std::cout << '\n';

  v.push_back(15);

  std::cout << "The vector elements are: ";
  for (int i = 0; i < v.size(); i++) {
    std::cout << v[i] << " ";
  }
  std::cout << '\n';


  int n = v.size();


  std::cout << "The last element is: " << v[n - 1] << '\n';

  v.pop_back();

  std::cout << "The vector elements are: ";
  for (int i = 0; i < v.size(); i++) {
    std::cout << v[i] << " ";
  }
  std::cout << '\n';

  v.insert(v.begin(), 5);
  std::cout << "The first element is: " << v[0] << '\n';

  v.erase(v.begin());
  std::cout << "The first element is: " << v[0] << '\n';

  v.emplace(v.begin(), 3);
  std::cout << "The first element is: " << v[0] << '\n';

  v.emplace_back(20);
  n = v.size();
  std::cout << "The last element is: " << v[n - 1] << '\n';

  v.clear();
  std::cout << "Vector size after erase(): " << v.size() << '\n';


  std::vector<int> v1, v2;


  v1.push_back(1);
  v1.push_back(2);
  v2.push_back(3);
  v2.push_back(4);

  std::cout << "\nVector 1: ";
  for (int i = 0; i < v1.size(); i++) {
    std::cout << v1[i] << " ";
  }
  std::cout << '\n';

  std::cout << "Vector 2: ";
  for (int i = 0; i < v2.size(); i++) {
    std::cout << v2[i] << " ";
  }
  std::cout << '\n';

  v1.swap(v2);

  std::cout << "After swap\n" << "Vector 1: ";
  for (int i = 0; i < v1.size(); i++) {
    std::cout << v1[i] << " ";
  }
  std::cout << '\n';

  std::cout << "Vector 2: ";
  for (int i = 0; i < v2.size(); i++) {
    std::cout << v2[i] << " ";
  }
  std::cout << '\n';



  return 0;
}
