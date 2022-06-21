#include <iostream>

void swap(char * &string1, char * &string2) {
  char *temp = string1;


  string1 = string2;
  string2 = temp;
}

int main() {
  char *string1 = "GEEKS";
  char *string2 = "FOR GEEKS";


  swap(string1, string2);

  std::cout << "string1 is " << string1 << '\n';
  std::cout << "string2 is " << string2 << '\n';



  return 0;
}
