#include <stdio.h>
#include "hashmap.h"

int main() {

  //unsigned int a, b, c, d, e, f, g, h, j, k, l, m;
  const char *aa = "Michael";
  const char *bb = "Yuhas";
  const char *cc = "John";
  const char *dd = "地土";
  const char *ee = "风吹";
  const char *ff = "烈火";
  const char *gg = "冰水";
  const char *hh = "马马虎虎";
  const char *jj = "万事如意";
  const char *kk = "百发百中";
  const char *ll = "万紫千红";
  const char *mm = "一路顺风";
  /*hash( aa, &a );
  hash( bb, &b );
  hash( cc, &c );
  hash( dd, &d );
  hash( ee, &e );
  hash( ff, &f );
  hash( gg, &g );
  hash( hh, &h );
  hash( jj, &j );
  hash( kk, &k );
  hash( ll, &l );
  hash( mm, &m );
  printf("%u\n",a);
  printf("%u\n",b);
  printf("%u\n",c);
  printf("%u\n",d);
  printf("%u\n",e);
  printf("%u\n",f);
  printf("%u\n",g);
  printf("%u\n",h);
  printf("%u\n",j);
  printf("%u\n",k);
  printf("%u\n",l);
  printf("%u\n",m);*/

  struct Hashmap hashmap1;
  init_hashmap( &hashmap1, 1024 );

  struct Var a;
  struct Var b;
  struct Var c;
  struct Var d;
  a.type=1;
  a.pointerToString="地土";
  b.type=2;
  b.pointerToFloat=1.1;
  c.type=3;
  c.pointerToInt=2;
  d.type=3;
  d.pointerToInt=3;

  add_key_value_pair( &hashmap1, "地土", a );
  add_key_value_pair( &hashmap1, "风吹", b );
  add_key_value_pair( &hashmap1, "烈火", c );
  add_key_value_pair( &hashmap1, "冰水", d );

  struct Var A;
  struct Var B;
  struct Var C;
  struct Var D;

  get_value_at_key( &hashmap1, "地土", &A );
  get_value_at_key( &hashmap1, "风吹", &B );
  get_value_at_key( &hashmap1, "烈火", &C );
  get_value_at_key( &hashmap1, "冰水", &D );

  printf("A: %s\n",A.pointerToString);
  printf("B: %f\n",B.pointerToFloat);
  printf("C: %d\n",C.pointerToInt);
  printf("D: %d\n",D.pointerToInt);

  destroy_hashmap( &hashmap1 );


  return 0;
}
