#include <iostream>
#include <string>

int main() {
  std::string poemat = "";
  std::string::size_type poprzPoj = poemat.capacity();
  std::string::size_type poj = poemat.capacity();


  std::cout << "Pusty string ma poj = " << poj << '\n';

  for (int i = 1; i < 500; i++) {
    poemat += 'A';

    poj = poemat.capacity();

    if (poj != poprzPoj) {
      std::cout << "Gdy mamy liczbę " << i << " liter, zwiekszylas sie "
		<< "poj = " << poj << ", (+" << (poj - poprzPoj) << ")"
		<< '\n';

      poprzPoj = poj;
    }
  }



  return 0;
}
