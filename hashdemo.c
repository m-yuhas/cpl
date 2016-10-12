#include <stdio.h>
#include "hashmap2.c"

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
  //init_hashmap( &hashmap1, 1024 );
  int a = 1;
  int b = 2;
  int c = 10;
  int d = 50;
  printf("HERE\n");

  add_key_value_pair_int( &hashmap1, "地土", a );
  add_key_value_pair_int( &hashmap1, "风吹", b );
  add_key_value_pair_int( &hashmap1, "烈火", c );
  add_key_value_pair_int( &hashmap1, "冰水", d );

  printf("HERE2\n");

  int A,B,C,D;
  printf("HERE3\n");

  get_int_at_key( &hashmap1, "地土", &A );
  printf("HERE4\n");
  get_int_at_key( &hashmap1, "风吹", &B );
  get_int_at_key( &hashmap1, "烈火", &C );
  get_int_at_key( &hashmap1, "冰水", &D );

  printf("A: %d\n",A);
  printf("B: %d\n",B);
  printf("C: %d\n",C);
  printf("D: %d\n",D);


  return 0;
}
