#include <iostream>

class Vehicle {
public:
  Vehicle() {
    std::cout << "This is a Vehicle\n";
  }
};

class fourWheeler : public Vehicle {
public:
  fourWheeler() {
    std::cout << "Objects with 4 wheels are vehicles\n";
  }
};

class Car : public fourWheeler {
public:
  Car() {
    std::cout << "Car has 4 wheels\n";
  }
};

int main() {
  Car obj;



  return 0;
}
