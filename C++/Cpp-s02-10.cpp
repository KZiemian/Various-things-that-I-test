#include <iostream>

class Pet {
public:
  virtual std::string GetDescription() const {
    return "This is Pet class";
  }
};

class Dog : public Pet {
public:
  virtual std::string GetDescription() const {
    return "This is Dog Class";
  }
};

void describe(const Pet& p) {
  std::cout << p.GetDescription() << '\n';
}

int main() {
  Dog d;


  describe(d);



  return 0;
}
