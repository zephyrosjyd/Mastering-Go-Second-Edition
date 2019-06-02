#include <stdio.h>
#include "usedByC.h"

int main(int argc, char **argv) {
	GoInt x = 12;
	GoInt y = 23;
	
	printf("About to call a Go function!\n");
	PrintMessage();
	GoInt p = Multiply(x,y);
	printf("Product: %d\n", (int)p);
	GoInt q = Add(x,y);
	printf("Plus: %d\n", (int)q);
	GoInt r = Subtract(x,y);
	printf("Minus: %d\n", (int)r);
	GoFloat32 s = divide(x,y);
	printf("Division: %f\n", (float)s);
	
	printf("It worked!\n");
	return 0;
}