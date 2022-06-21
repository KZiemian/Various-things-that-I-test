#include <stdio.h>

struct _iobuf {
  int m;
  int n;
};

int main() {
  struct _iobuf structVar = {0, 1};

  printf("structVar.m = %d\nstructVar.n = %d\n", structVar.m, structVar.n);



  return 0;
}
