std::vector<int> sampleVector;


for (std::vector<int>::const_iterator it sampleVector.cbegin();
     it != sampleVector.cend(); it++) {
  std::cout << *it << '\n';
 }

for (auto it = sampleVector.cbegin(); it != sampleVector.cend(); it++) {
  std::cout << *it << '\n';
 }

void function(int a_Number);
void function(const char* a_Pointer);

function(NULL);
function(nullptr);

#include <vector>

int main() {
  std::vector<std::vector<int>> myVector;



  return 0;
}

#include <vector>

int main() {
  std::vector<std::vector<int> > myVector;


  return 0;
}

enum class Level {
  First,
  Second,
  Third
};

Level myLevel = Level::Second;

int *wsk_temperatury = nullptr;

wsk_temperatury = reinterpret_cast<int*>(93952);

char *wsk;
wsk = new char;

delete wsk;

double *wsk;

w = new double[15];

delete [] w;

std::unique_ptr
std::shared_ptr
std::weak_ptr

void MyFunction() {
  std::unique_ptr<int> IntPointer;


  IntPointer.reset(new int(4));

  if (IntPointer != nullptr) {
    std::cout << "Value: " << *IntPointer << '\n';
  }
}

std::shared_ptr<MyClass> sp1, sp2;

sp1 = sp2;

int *ptr = new int;

std::shared_ptr<int> sp1(ptr);
std::shared_ptr<int> sp = std::make_share<int>(4);

std::shared_ptr<int> sp1 = std::make_share<int>(10);

std::shared_ptr<int> sp2(sp1);

std::shared_ptr<int> sp1 = std::make_shared<int>(4);
std::weak_ptr<int> wp(sp1);

std::shared_ptr<int> sp2 = wp.lock();

if (sp2) {
  //
 }

class port {
  unsigned int odczyt : 3;
};

class portA {
  unsigned int odczyt : 1;
  unsigned int tryp : 3;
  unsigned int gotow : 1;
  unsigned int dana : 4;
};
