#include <stdio.h>

int main(int argc, char *argv[])
{
	char *input = argv[1];
	printf("Input string is: %s\n", input);

	int i = 0;
	while (input[i]) {
		input[i] += 32;

		i++;
	}
	printf("Output string is: %s\n", input);

	return 0;
}
