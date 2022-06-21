#include <iostream>

class base {
public:
  void fun_1() {
    std::cout << "base-1\n";
  }

  virtual void fun_2() {
    std::cout << "base-2\n";
  }

  virtual void fun_3() {
    std::cout << "base-3\n";
  }

  virtual void fun_4() {
    std::cout << "base-4\n";
  }
};

class derived : public base {
public:
  void fun_1() {
    std::cout << "derived-1\n";
  }

  virtual void fun_2() {
    std::cout << "derived-2\n";
  }

  virtual void fun_4(int x) {
    std::cout << "derived-4\n";
  }
};

int main() {
  base *p = nullptr;
  derived obj1;

  p = &obj1;

  p->fun_1();
  p->fun_2();
  p->fun_3();
  p->fun_4();



  return 0;
}
