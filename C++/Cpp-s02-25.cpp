#include <iostream>

class ClassA {
public:
  int A;
};

class ClassB : public ClassA {
public:
  int B;
};

class ClassC : public ClassA {
public:
  int C;
};

class ClassD : public ClassB, public ClassC {
public:
  int D;
};

int main() {
  ClassD obj;


  // obj.a = 10;

  obj.ClassB::A = 10;
  obj.ClassC::A = 100;

  obj.B = 20;
  obj.C = 30;
  obj.D = 40;


  std::cout << " a from ClassB : " << obj.ClassB::A;
  std::cout << "\n a from ClassC : " << obj.ClassC::A;

  std::cout << "\n b : " << obj.B;
  std::cout << "\n c : " << obj.C;
  std::cout << "\n d : " << obj.D << '\n';

  // std::cout << "obj.a: " << obj.a << '\n';



  return 0;
}
