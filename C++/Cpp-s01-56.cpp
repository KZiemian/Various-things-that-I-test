#include <iostream>

int main() {
  int *wsk = nullptr;


  wsk = new int(32);

  std::cout << "*wsk = " << *wsk << '\n'
	    << " wsk = " <<  wsk << '\n';



  return 0;
}
