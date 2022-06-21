#include <iostream>

class ClassA {
public:
  int a;
};

class ClassB : virtual public ClassA {
public:
  int b;
};

class ClassC : virtual public ClassA {
public:
  int c;
};

class ClassD : public ClassB, public ClassC {
public:
  int d;
};

int main() {
  ClassD obj;


  obj.a = 10;
  obj.a = 100;

  obj.b = 20;
  obj.c = 30;
  obj.d = 40;

  std::cout << " a : " << obj.a;
  std::cout << "\n b : " << obj.b;
  std::cout << "\n c : " << obj.c;
  std::cout << "\n d : " << obj.d << '\n';



  return 0;
}
