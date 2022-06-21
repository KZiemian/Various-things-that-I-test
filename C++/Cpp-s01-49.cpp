#include <iostream>

int main() {
  int *ptr1 = nullptr;
  int *ptr2;


  if (ptr1 == nullptr) {
    std::cout << "Wskaznik ptr1: nullptr" << '\n';
  }
  if (ptr2 != nullptr) {
    std::cout << "Wskaznik ptr2 ustawiony na jakieś smieci" << '\n'
	      << "Smiertelnie niebezpieczna sytuacja" << '\n';
  } else if (ptr2 == nullptr) {
    std::cout << "Wskaznik ptr2: nullptr" << '\n';
  }

  std::cout << "ptr1: " << ptr1 << '\n'
	    << "ptr2: " << ptr2 << '\n';



  return 0;
}
