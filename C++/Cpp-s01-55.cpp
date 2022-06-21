#include <iostream>

char *producent(void) {
  char *ptr = nullptr;


  std::cout << "Wlasnie produkuje obiekt\n";

  ptr = new char;

  return ptr;
}

int main() {
  char *ptr1 = producent();
  char *ptr2 = producent();
  char *ptr3 = producent();
  char *ptr4 = producent();


  *ptr1 = 'H';
  *ptr2 = 'M';
  *ptr3 = 'I';

  std::cout << "Oto 3 znaki: " << *ptr1 << *ptr2 << *ptr3 << '\n'
	    << "Oraz smieci w czwartym: " << *ptr4 << '\n';

  // std::cout << "ptr4: " << ptr4 << std::endl;
  // std::cout << "reinterpret_cast<int>(ptr4): " << reinterpret_cast<int>(ptr4)
  // 	    << std::endl;

  delete ptr1;
  delete ptr2;
  delete ptr3;
  delete ptr4;

  // *ptr1 = 'O';

  // std::cout << "*ptr1: " << *ptr1 << std::endl;

  std::cout << "Koniec\n";



  return 0;
}
