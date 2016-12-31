#include <stdio.h>
#include <stdlib.h>

#define SIZE 3

struct result {
	int aresult, bresult;
};

struct result *compareTriplet(int *a, int *b, int size)
{
	struct result *res = (struct result *) malloc(sizeof(struct result));
	int i = 0;
	for (i = 0; i < size; i++) {
		if (a[i] == b[i]) {
			continue;
		} else if (a[i] < b[i]) {
			res->bresult++;
		} else {
			res->aresult++;
		}
	}
	return res;
}

int main()
{
	int *as = (int *) malloc(sizeof(int) * SIZE);
	int *bs = (int *) malloc(sizeof(int) * SIZE);
	int i = 0;
	for (i = 0; i < SIZE; i++) {
		scanf("%d", &as[i]);
	}
	int j = 0;
	for (j = 0; j < SIZE; j++) {
		scanf("%d", &bs[j]);
	}
	struct result *res = compareTriplet(as, bs, SIZE);
	printf("%d %d", res->aresult, res->bresult);
	free(as);
	free(bs);
	free(res);
	return 0;
}
