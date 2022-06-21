#include <iostream>

class Vehicle {
public:
  Vehicle() {
    std::cout << "This is a Vehicle\n";
  }
};

class Fare {
public:
  Fare() {
    std::cout << "Fare of Vehicle\n";
  }
};

class Car : public Vehicle {};

class Bus : public Vehicle, public Fare {};

int main() {
  Bus obj2;



  return 0;
}
