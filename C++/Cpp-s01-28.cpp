#include <iostream>
#include <string>

int main() {
  std::string poemat = "";
  std::string::size_type poprzPoj = poemat.capacity();
  std::string::size_type poj = poemat.capacity();


  std::cout << "Pusty string ma poj = " << poj << '\n';

  for (int i = 0; i < 2000; i++) {
    poemat += 'A';

    poj = poemat.capacity();

    if (poj != poprzPoj) {
      std::cout << "Gdy mamy liczbę " << i << " liter, zwiekszyla sie "
		<< "poj = " << poj << ", (+" << (poj - poprzPoj) << ")"
		<< '\n';

      poprzPoj = poj;
    }
  }



  return 0;
}
