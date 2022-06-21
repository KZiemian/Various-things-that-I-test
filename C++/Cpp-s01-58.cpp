#include <iostream>

int main() {
  const double *wsk_pi;


  wsk_pi = new const double(3.1416);

  // *wsk_pi = 888.1;

  std::cout << "*wsk_pi: " << *wsk_pi << '\n'
	    << " wsk_pi: " <<  wsk_pi << '\n';


  delete wsk_pi;


  // std::cout << "*wsk_pi: " << *wsk_pi << std::endl
  // 	    << " wsk_pi: " <<  wsk_pi << std::endl;

  return 0;
}
