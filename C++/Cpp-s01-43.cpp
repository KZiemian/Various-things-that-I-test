#include <iostream>

int main() {
  int arrayVar[15] = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14};
  int *wsk_a = &arrayVar[5];
  int *wsk_b = &arrayVar[10];
  int *wsk_c = &arrayVar[11];


  std::cout << "(wsk_b - wsk_a) = " << (wsk_b - wsk_a) << '\n'
	    << "(wsk_c - wsk_b) = " << (wsk_c - wsk_b) << '\n'
	    << "(wsk_a - wsk_c) = " << (wsk_a - wsk_c) << '\n'
	    << "(wsk_c - wsk_a) = " << (wsk_c - wsk_a) << '\n';



  return 0;
}
