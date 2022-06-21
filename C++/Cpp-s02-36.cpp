char *wsk_do_znakow;
double *wsk_do_doubleow;

X objX;
X *wskX = &objX;
Y *wskY;

// wskY = wskX;
wskY = reinterpret_cast<Y*>(wskX);

int numerycznyAdres = 0x0f6a21;
int *wskSygnal = reinterpret_cast<int*>(numerycznyAdres);
int schowek = 0;

wskDouble = (double*)wskVoid;
wskInt = (int*)wskVoid;
wskChar = (char*)wskVoid;

wskDouble = reinterpret_cast<double*>(wskVoid);
wskInt = reinterpret_cast<int*>(wskVoid);
wskChar = reinterpret_cast<char*>(wskChar);

class pralka {
public:
  int    nrProgramu;
  double temperaturPrania;
  char   nazwa[80];
};

pralka czerwona;
pralka *wskaz;
pralka &ruda = czerwona;

czerwona.temperaturaPrania = 60;
wskaz = &czerwona;
wskaz->temperaturaPrania = 60;
ruda.temperaturaPrania = 60;

class pralka {
public:
  void pierze(int program);
  void wiruj(int minuty);

  int    nrProgramu;
  double temperaturPrania;
  char   nazwa[80];

  int krochmalenie(void);
};

class pralka {
public:
  void pierze(int program);
  void wiruj(int minuty);
  int krochmalenie(void);

  int    nrProgramu;
  double temperaturaPrania;
  char   nazwa[80];
};

void funkcja() {
  int a = 0;
  int b = 0;
  int c = 15;


  a = 4 + c;


  int n = 0;


  n = a + 6;
  // ....
}

class naszTyp {
private:
  int    liczba;
  double temperatura;
  char   komunikat[80];

  int czy_gotowe();

public:
  double predkosc;
  int zrobPomiar();
};

class dzika {
  int    a;
  double b;

  void fun1(int);



protected:
  char m;
  void fun2(void);



public:
  int  x;
  void fun3(char*);



private:
  int d;



public:
  void fun4(void);
};
