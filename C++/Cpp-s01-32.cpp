#include <iostream>

void spiew() {
  std::cout << "Zwykla funkcja spiew (nie majaca nic wspolnego z klasa)\n";
}

class opera {
public:
  int n;
  double balkon;


  void funkcja();
  void spiew() {
    std::cout << "Funckja spiew (z opery): tra-la-la!" << '\n';
  }
};

void opera::funkcja() {
  std::cout << "balkon (skladnik, klasy) = " << balkon << '\n';
  // std::cout << "balkon (zmienna globalna) = " << ::balkon << std::endl;

  char balkon = 'M';


  std::cout << '\n' << "Po definicji zmiennej lokalnej ---" << '\n';
  std::cout << "balkon (zmienna lokalna) = " << balkon << '\n';
  std::cout << "balkon (skladnik klasy) = " << opera::balkon << '\n';
  // std::cout << "balkon (zmienna globalna) = " << ::balkon << std::endl;

  spiew();
  int spiew = 7;

  std::cout << "Po zaslonieci da sie wywolac funkcje spiew tylko tak\n";
  opera::spiew();
}

int main() {
  opera Lohengrin;


  Lohengrin.balkon = 6;
  Lohengrin.funkcja();
  spiew();



  return 0;
}
