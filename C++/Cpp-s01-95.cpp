#include <iostream>

struct struct1 {
  unsigned int Bitfield : 3;
};

class class1 {
public:
  unsigned int Bitfield : 2;
};

int main() {
  struct1 structVar;
  class1 classVar;


  structVar.Bitfield = 7;
  classVar.Bitfield = 1;

  std::cout << "structVar.bitField: " << structVar.Bitfield << '\n';
  std::cout << "classVar.bitField: " << classVar.Bitfield << '\n';



  return 0;
}
