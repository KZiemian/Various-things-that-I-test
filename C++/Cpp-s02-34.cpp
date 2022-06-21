#include <iostream>

class Shape {
public:
  Shape(int l, int w) {
    length = l;
    width = w;
  }

  int getArea() {
    std::cout << "This is call to parent class\n";

    return 1;
  }



protected:
  int length;
  int width;
};

class Square : public Shape {
public:
  Square(int l = 0, int w = 0) : Shape(l, w) {}

  int getArea() {
    std::cout << "Square area: " << length * width << '\n';

    return length * width;
  }
};

class Rectangle : public Shape {
public:
  Rectangle(int l = 0, int w = 0) : Shape(l, w) {}

  int getArea() {
    std::cout << "Rectangle area: " << length * width << '\n';

    return length * width;
  }
};

int main() {
  Shape *s;
  Square sq(5, 5);
  Rectangle rec(4, 5);


  s = &sq;
  s->getArea();
  s = &rec;
  s->getArea();



  return 0;
}
