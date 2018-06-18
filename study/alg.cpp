typedef struct {
  int *qu_array;
  int count;
} QuickUnion;


static int root(QuickUnion qu, int i){
  while(qu->qu_array[i] != i){
    i = qu->qu_array[i];
  }
  return i;
}

int qu_union(QuickUnion qu, int p, int q){
  int proot, qroot;
  proot = root(qu,p);
  qroot = root(qu,q);
  qu->qu_array[proot] = qroot;
  return 0;
}



int qu_connected(QuickUnion qu, int p, int q){
  return root(qu,p) == root(qu,q);
}
