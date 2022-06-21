#include <iostream>
#include <string>

struct Student {
  std::string Name;
  std::string Address;
  int RollNo;
};

void print(const Student &student) {
  std::cout << student.Name << "  " << student.Address << "  "
	    << student.RollNo << '\n';
}

int main() {
  Student Kamil = {"Kamil", "Gdzies", 1};


  print(Kamil);



  return 0;
}
