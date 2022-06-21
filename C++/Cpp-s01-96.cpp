 #include<iostream>

struct struct1 {
  unsigned int Bitfield : 4;
};

int main() {
  struct1 structVar;


  structVar.Bitfield = 2;

  std::cout << "structVar.Bitfield: " << structVar.Bitfield << '\n';



  return 0;
}
