#include <iostream>

class Vehicle {
public:
  Vehicle() {
    std::cout << "This is a Vehicle\n";
  }
};

class Car : public Vehicle {};

int main() {
  Car obj;



  return 0;
}
