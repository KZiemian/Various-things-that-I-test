#include <iostream>
#include <memory>

void MyFunction1() {
  std::unique_ptr<int> intPointer;


  intPointer.reset(new int(4));

  if (intPointer != nullptr) {
    std::cout << "Value: " << *intPointer << '\n';
  }
}

void MyFunction2() {
  std::unique_ptr<int> intPointer;


  intPointer.reset(new int(7));

  if (intPointer) {
    std::cout << "Value: " << *intPointer << '\n';
  }
}

int main() {
  MyFunction1();
  MyFunction2();



  return 0;
}
