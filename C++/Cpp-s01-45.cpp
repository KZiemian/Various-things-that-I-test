#include <iostream>
#include <string>
#include <typeinfo>

int main() {
  int varInt = 0;
  double varDouble = 0.0;
  float varFloat = 0.0;
  std::string varString = "string";
  char varChar = 'a';


  std::cout << "Type varInt: " << typeid(varInt).name() << '\n'
	    << "Type varDouble: " << typeid(varDouble).name() << '\n'
	    << "Type varFloat: " << typeid(varFloat).name() << '\n'
	    << "Type varString: " << typeid(varString).name() << '\n'
	    << "Type varChar: " << typeid(varChar).name() << '\n';



  return 0;
}
