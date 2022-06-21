#include <iostream>

class geeks {
public:
  int Id;


  geeks() {
    Id = 0;
  }

  ~geeks() {
    std::cout << "Destruction called for Id: " << Id << '\n';
  }
};

int main() {
  geeks obj1;


  obj1.Id = 7;


  int i = 0;


  while (i < 5) {
    geeks obj2;


    obj2.Id = i;
    i++;
  }



  return 0;
}
