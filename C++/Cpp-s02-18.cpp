#include <iostream>

class parent {
public:
  int Id_p;


  parent() {
    Id_p = 0;
  }
};

class child : public parent {
public:
  int Id_c;


  child() {
    Id_c = 0;
  }
};

int main() {
  child obj1;


  std::cout << "obj1.Id_c: " << obj1.Id_c << '\n';

  obj1.Id_c = 7;
  obj1.Id_p = 91;

  std::cout << "Child Id is: " << obj1.Id_c << '\n';
  std::cout << "Parent Id is: " << obj1.Id_p << '\n';



  return 0;
}
