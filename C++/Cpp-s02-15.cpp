#include <iostream>

class geeks {
public:
  std::string Geekname;
  int Id;


  void Printname();

  void PrintId() {
    std::cout << "Geek id is: " << Id << '\n';
  }
};

void geeks::Printname() {
  std::cout << "Geekname is: " << Geekname << '\n';
}

int main() {
  geeks obj1;


  obj1.Geekname = "xyz";
  obj1.Id = 15;

  obj1.Printname();
  obj1.PrintId();



  return 0;
}
