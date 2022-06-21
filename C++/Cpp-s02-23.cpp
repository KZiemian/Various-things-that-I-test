#include <iostream>

class Vehicle {
public:
  Vehicle() {
    std::cout << "This is a Vehicle\n";
  }
};

class Car : public Vehicle {};

class Bus : public Vehicle {};

int main() {
  Car obj1;
  Bus obj2;



  return 0;
}
