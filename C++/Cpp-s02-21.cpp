#include <iostream>

class Vehicle {
public:
  Vehicle() {
    std::cout << "This is a Vehicle\n";
  }
};

class FourWheeler {
public:
  FourWheeler() {
    std::cout << "This is a 4 wheeler Vehicle\n";
  }
};

class Car : public Vehicle, public FourWheeler {};

int main() {
  Car obj;



  return 0;
}
