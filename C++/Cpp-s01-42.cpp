#include <iostream>
#include <string>

int solution(int A, int B) {
  int result = -1;
  std::string Astr = std::to_string(A);
  std::string Bstr = std::to_string(B);

  int lengthAstr = Astr.length();
  int lengthBstr = Bstr.length();

  int correctDigits = 0;
  int AcurrentIndex = 0;
  int resultCandIndex = -1;


  for (int i = 0; i < lengthBstr; i++) {
    if ((Bstr[i] == Astr[AcurrentIndex]) && (AcurrentIndex == 0)) {
	resultCandIndex = i;
	correctDigits += 1;
	AcurrentIndex += 1;
      } else if (Bstr[i] == Astr[AcurrentIndex]) {
      correctDigits += 1;
      AcurrentIndex += 1;
    } else {
      resultCandIndex = -1;
      correctDigits = 0;
      AcurrentIndex = 0;
    }

    if (correctDigits == lengthAstr) {
      result = resultCandIndex;
      break;
    }
  }

  return result;
}

int main() {
  std::cout << "solution(33, 13334): " << solution(0, 101113) << '\n';



  return 0;
}
