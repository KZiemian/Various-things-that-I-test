#include <iostream>

class base {
public:
  virtual void show() {
    std::cout << "In base\n";
  }
};

class derived : public base {
public:
  void show() {
    std::cout << "In derived\n";
  }
};

void print(base& b) {
  b.show();
}

int main() {
  base b;
  derived d;


  print(b);
  print(d);



  return 0;
}
