#include <iostream>

int main() {
  int varInt = 3;
  int anotherVarInt = 5;
  int& refInt = varInt;


  std::cout << "varInt = " << varInt << '\n'
	    << "refInt = " << refInt << '\n'
	    << "anotherVarInt = " << anotherVarInt << "\n\n";


  refInt = anotherVarInt;
  std::cout << "refInt = " << refInt << "\n\n";

  refInt = 10;
  std::cout << "varInt = " << varInt << '\n'
	    << "anotherVarInt = " << anotherVarInt << '\n'
	    << "refInt = " << refInt << '\n';



  return 0;
}
