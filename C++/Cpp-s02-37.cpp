#include <iostream>

class geeks {
public:
  std::string Geekname;

  void Printname() {
    std::cout << "Geekname is " << Geekname << '\n';
  }
};

int main() {
  geeks obj1;

  obj1.Geekname = "Abhi";
  obj1.Printname();



  return 0;
}

class Employee {
public:
  virtual void raiseSalary() {}

  virtual void promote() {}
};

class Manager : public Employee {
  virtual void raiseSalary() {}
};

void globalRaiseSalary(Employee *emp[], int n) {
  for (int i = 0; i < n, i++) {
    emp[i]->raiseSalary();
  }
}
