#include <iostream>

class base {
public:
  void Fun1() {
    std::cout << "base-1\n";
  }

  virtual void Fun2() {
    std::cout << "base-2\n";
  }

  virtual void Fun3() {
    std::cout << "base-3\n";
  }

  virtual void Fun4() {
    std::cout << "base-4\n";
  }
};

class derived : public base {
public:
  void Fun1() {
    std::cout << "derived-1\n";
  }

  virtual void Fun2() {
    std::cout << "derived-2\n";
  }

  virtual void Fun4(int x) {
    std::cout << "derived-4\n";
  }
};

int main() {
  base *p = nullptr;
  derived obj1;


  p = &obj1;

  p->Fun1();
  p->Fun2();
  p->Fun3();
  p->Fun4();

  // p->Fun4(5);



  return 0;
}
