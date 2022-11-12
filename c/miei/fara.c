#include <stdio.h>
#define N 6
#define L 0.01

int main(){
    float v[N];
	for (int i=0; i<N; i++){
		float in;
		scanf("%f", &in);
		v[i] = in;
	}

	for (int i=0; i<N; i++){
		if (v[i]>L){
			printf("%.1f ", v[i]);
		}
	}
	printf("\n");
	return 0;
}
