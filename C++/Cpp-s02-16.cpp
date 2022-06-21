#include <iostream>

class geeks {
public:
  int Id;


  geeks() {
    std::cout << "Default Constructor called\n";

    Id = -1;
  }

  geeks(int x) {
    std::cout << "Parameter Constructor called\n";

    Id = x;
  }
};

int main() {
  geeks obj1;


  std::cout << "Geek Id is: " << obj1.Id << '\n';


  geeks obj2(21);


  std::cout << "Geek Id is: " << obj2.Id << '\n';



  return 0;
}
