#include <iostream>

class base {
public:
  virtual void Print() {
    std::cout << "Print base class\n";
  }

  void Show() {
    std::cout << "Show base class\n";
  }
};

class derived : public base {
public:
  virtual void Print() {
    std::cout << "Print derived class\n";
  }

  void Show() {
    std::cout << "Show derived class\n";
  }
};

int main() {
  base *bptr = nullptr;
  derived d;


  bptr = &d;

  bptr->Print();
  bptr->Show();



  return 0;
}
